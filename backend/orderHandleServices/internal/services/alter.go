package services

import (
	sysContext "context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	pb "orderHandleServices/api/proto"
	"orderHandleServices/internal/dao"
	"orderHandleServices/internal/model"
	"orderHandleServices/internal/utils"
	"time"
)

func (s *OrderHandleServices) CancelOrder(context context.Context, request *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	//var err error
	//log.Println("取消订单信息: ", request.UserId, request.OrderId)
	//
	//var pendingOrder model.Orders
	//redisKey := fmt.Sprintf("pending_order:%d:%s", request.UserId, request.OrderId)
	//
	//// 从 Redis 中取出数据并反序列化
	//redisCtx, cancel := sysContext.WithTimeout(sysContext.Background(), 5*time.Second)
	//defer cancel()
	//val, err := dao.Rdb.Get(redisCtx, redisKey).Result()
	//if errors.Is(err, redis.Nil) {
	//	log.Println("订单不存在于 Redis 中")
	//	return &pb.CancelOrderResponse{
	//		Code:      http.StatusNotFound,
	//		Cancelled: false,
	//		Msg:       "订单未找到",
	//	}, nil
	//} else if err != nil {
	//	log.Println("从 Redis 获取订单失败:", err)
	//	return &pb.CancelOrderResponse{
	//		Code:      http.StatusInternalServerError,
	//		Cancelled: false,
	//		Msg:       "获取订单信息失败" + err.Error(),
	//	}, nil
	//}
	//
	//// 反序列化 Redis 中的订单数据
	//err = json.Unmarshal([]byte(val), &pendingOrder)
	//if err != nil {
	//	log.Println("订单反序列化失败:", err)
	//	return &pb.CancelOrderResponse{
	//		Code:      http.StatusInternalServerError,
	//		Cancelled: false,
	//		Msg:       "订单反序列化失败",
	//	}, nil
	//}
	//
	//// 启动数据库事务
	//err = dao.Db.Transaction(func(tx *gorm.DB) error {
	//	// 将 pendingOrder 插入订单表
	//	pendingOrder.IsFinished = true       // 标记订单为已取消
	//	pendingOrder.FailureReason = "订单已取消" // 取消原因
	//
	//	if err := tx.Create(&pendingOrder).Error; err != nil {
	//		log.Println("订单插入数据库失败:", err)
	//		return err
	//	}
	//
	//	// 删除 Redis 中的订单
	//	if err := dao.Rdb.Del(redisCtx, redisKey).Err(); err != nil {
	//		log.Println("从 Redis 删除订单失败:", err)
	//		return err
	//	}
	//
	//	return nil
	//})
	//
	//// 事务操作失败处理
	//if err != nil {
	//	log.Println("取消订单失败:", err)
	//	return &pb.CancelOrderResponse{
	//		Code:      http.StatusInternalServerError,
	//		Cancelled: false,
	//		Msg:       "取消订单失败",
	//	}, nil
	//}
	//
	//// 事务成功，返回成功响应
	//return &pb.CancelOrderResponse{
	//	Code:      http.StatusOK,
	//	Cancelled: true,
	//	Msg:       "订单已取消",
	//}, nil
	code, cancelled, msg, err := CancelOrderUniversal(request.UserId, request.OrderId, "user")
	return &pb.CancelOrderResponse{
		Code:      code,
		Cancelled: cancelled,
		Msg:       msg,
	}, err

}

func (s *OrderHandleServices) PlaceOrder(ctx context.Context, request *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {
	log.Println("PlaceOrder", request.UserId, request.OrderId)
	redisKey := fmt.Sprintf("pending_order:%d:%s", request.UserId, request.OrderId)

	// 从 Redis 中获取订单信息
	orderJson, err := dao.Rdb.Get(ctx, redisKey).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get order from redis: %v", err)
	}

	// 将订单信息反序列化
	var order model.Orders
	if err := json.Unmarshal([]byte(orderJson), &order); err != nil {
		return nil, fmt.Errorf("failed to unmarshal order: %v", err)
	}

	// 检查用户余额
	var user model.User
	if err := dao.Db.Model(&model.User{}).Where("id = ?", request.UserId).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	// 如果余额不足，返回响应
	if user.Balance < float32(order.Amount) {
		return &pb.PlaceOrderResponse{
			Code:        http.StatusPaymentRequired,
			Msg:         "余额不足请充值后再次购买，该订单保留时间为5分钟。",
			Placed:      false,
			UserBalance: user.Balance,
		}, nil
	}

	// 删除 Redis 中的订单键
	if err := dao.Rdb.Del(ctx, redisKey).Err(); err != nil {
		log.Println("Failed to delete order from Redis:", err)
	}

	// 开启事务
	tx := dao.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新用户余额
	//newBalance := user.Balance - float32(order.Amount)
	//if err := tx.Model(&model.User{}).Where("id = ?", request.UserId).Update("balance", newBalance).Error; err != nil {
	//	tx.Rollback()
	//	return nil, fmt.Errorf("failed to update user balance: %v", err)
	//}

	// 通过 WHERE balance >= ? 限制，避免高并发下可能的余额不足问题。
	newBalance := user.Balance - float32(order.Amount)
	if err := tx.Model(&model.User{}).
		Where("id = ? AND balance >= ?", request.UserId, order.Amount).
		Update("balance", newBalance).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update user balance or insufficient funds: %v", err)
	}

	// 生成密钥
	newKey := model.Keys{
		//UserId:  order.UserId,
		//IsValid: true,  // 密钥有效
		//IsUsed:  false, // 还没有被使用
		//Expired: false, // 还未到期
		Key: generateKeyCode(order.Period),
	}
	log.Println("Key: ", newKey.Key)
	if err := tx.Model(&model.Keys{}).Create(&newKey).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create key: %v", err)
	}

	// 设置订单到期时间
	var expirationDate time.Time
	switch order.Period {
	case "month":
		expirationDate = time.Now().AddDate(0, 1, 0)
	case "quarter":
		expirationDate = time.Now().AddDate(0, 3, 0)
	case "half-year":
		expirationDate = time.Now().AddDate(0, 6, 0)
	case "year":
		expirationDate = time.Now().AddDate(1, 0, 0)
	}

	// 创建激活订单
	log.Println("KeyId: ", newKey.Id)
	activeOrder := model.ActiveOrders{
		OrderId:        order.OrderId,
		UserId:         request.UserId,
		PlanId:         order.PlanId,
		Email:          user.Email,
		GroupId:        order.GroupId,
		IsActive:       true,
		IsUsed:         false,
		IsValid:        true,
		KeyId:          newKey.Id,
		ExpirationDate: &expirationDate,
	}

	if err := tx.Model(&model.ActiveOrders{}).Create(&activeOrder).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create active order: %v", err)
	}

	// 更新订单状态
	order.IsFinished = true // 订单结束
	order.IsSuccess = true  // 订单成功
	var paidTime = time.Now()
	order.PaidAt = &paidTime
	//order.FailureReason = "Not confirmed within the validity period"

	if err := tx.Model(&model.Orders{}).Create(&order).Error; err != nil {
		tx.Rollback()
		log.Println("Failed to insert order into database:", err)
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	// 更新 Redis 中的用户信息
	userKey := fmt.Sprintf("user:%s", user.Email)
	userJson, err := json.Marshal(user)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to marshal user data: %v", err)
	}
	if err := dao.Rdb.Set(ctx, userKey, userJson, 0).Err(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update user info in redis: %v", err)
	}

	if err = utils.FreshUserPropertyInfoInRedis(request.UserId); err != nil {
		log.Println("Failed to update user property info in redis:", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	// 返回成功响应
	log.Println("执行成功")
	return &pb.PlaceOrderResponse{
		Code:         http.StatusOK,
		Msg:          "下单成功",
		Placed:       true,
		UserBalance:  newBalance,
		KeyGenerated: true,
	}, nil
}

func (s *OrderHandleServices) ManualPassOrderPayment(ctx context.Context, request *pb.ManualPassOrderPaymentRequest) (*pb.ManualPassOrderPaymentResponse, error) {
	log.Println("ManualPassOrderPayment", request.UserId, request.OrderId)

	// 从 Redis 中获取订单信息
	redisKey := fmt.Sprintf("pending_order:%d:%s", request.UserId, request.OrderId)
	orderJson, err := dao.Rdb.Get(ctx, redisKey).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get order from redis: %v", err)
	}

	// 反序列化订单信息
	var order model.Orders
	if err := json.Unmarshal([]byte(orderJson), &order); err != nil {
		return nil, fmt.Errorf("failed to unmarshal order: %v", err)
	}

	// 删除 Redis 中的订单键
	//if err := dao.Rdb.Del(ctx, redisKey).Err(); err != nil {
	//	log.Println("Failed to delete order from Redis:", err)
	//}

	// 获取用户信息
	var user model.User
	if err := dao.Db.Model(&model.User{}).Where("id = ?", request.UserId).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	// 开启事务
	tx := dao.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 生成密钥
	newKey := model.Keys{
		Key: generateKeyCode(order.Period),
	}
	log.Println("Key: ", newKey.Key)
	if err := tx.Model(&model.Keys{}).Create(&newKey).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create key: %v", err)
	}

	// 设置订单到期时间
	var expirationDate time.Time
	switch order.Period {
	case "month":
		expirationDate = time.Now().AddDate(0, 1, 0)
	case "quarter":
		expirationDate = time.Now().AddDate(0, 3, 0)
	case "half-year":
		expirationDate = time.Now().AddDate(0, 6, 0)
	case "year":
		expirationDate = time.Now().AddDate(1, 0, 0)
	}

	// 创建激活订单
	activeOrder := model.ActiveOrders{
		OrderId:        order.OrderId,
		UserId:         request.UserId,
		PlanId:         order.PlanId,
		Email:          user.Email,
		GroupId:        order.GroupId,
		IsActive:       true,
		IsUsed:         false,
		IsValid:        true,
		KeyId:          newKey.Id,
		ExpirationDate: &expirationDate,
	}
	if err := tx.Model(&model.ActiveOrders{}).Create(&activeOrder).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create active order: %v", err)
	}

	checkTime := time.Now()

	// 更新订单状态
	order.IsFinished = true
	order.IsSuccess = true
	order.PaidAt = &checkTime
	if err := tx.Model(&model.Orders{}).Create(&order).Error; err != nil {
		tx.Rollback()
		log.Println("Failed to insert order into database:", err)
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	// 更新 Redis 中的用户信息
	//userKey := fmt.Sprintf("user:%s", user.Email)
	//userJson, err := json.Marshal(user)
	//if err != nil {
	//	tx.Rollback()
	//	return nil, fmt.Errorf("failed to marshal user data: %v", err)
	//}
	//if err := dao.Rdb.Set(ctx, userKey, userJson, 0).Err(); err != nil {
	//	tx.Rollback()
	//	return nil, fmt.Errorf("failed to update user info in redis: %v", err)
	//}

	if err = utils.FreshUserPropertyInfoInRedis(request.UserId); err != nil {
		log.Println("Failed to update user property info in redis:", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	// 删除redis订单信息应该在最后进行删除 防止刷新不及时返回错误
	if err := dao.Rdb.Del(ctx, redisKey).Err(); err != nil {
		log.Println("Failed to delete order from Redis:", err)
	}

	// 返回成功响应
	log.Println("Order manually passed by admin successfully")
	return &pb.ManualPassOrderPaymentResponse{
		Code:   http.StatusOK,
		Passed: true,
		Msg:    "Order manually passed by admin",
	}, nil
}

func (s *OrderHandleServices) ManualCancelOrderPayment(ctx context.Context, request *pb.ManualCancelOrderPaymentRequest) (*pb.ManualCancelOrderPaymentResponse, error) {
	code, cancelled, msg, err := CancelOrderUniversal(request.UserId, request.OrderId, "admin")
	return &pb.ManualCancelOrderPaymentResponse{
		Code:      code,
		Cancelled: cancelled,
		Msg:       msg,
	}, err
}

func generateKeyCode(period string) string {
	// 根据不同订阅时长选择代码
	var periodCode string
	switch period {
	case "month":
		periodCode = "M"
	case "quarter":
		periodCode = "Q"
	case "half-year":
		periodCode = "H"
	case "year":
		periodCode = "Y"
	default:
		periodCode = "X" // 默认代码以防意外输入
	}

	// 获取当前月份代码（占两位）
	now := time.Now()
	monthCode := fmt.Sprintf("%02d", now.Month())
	dayCode := fmt.Sprintf("%02d", now.Day())
	// 生成一个随机字符串，包含大写英文字母和数字
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	generateRandomSegment := func() string {
		segment := make([]byte, 5)
		for i := range segment {
			segment[i] = charset[rand.Intn(len(charset))]
		}
		return string(segment)
	}

	// 拼接最终的密钥格式
	key := fmt.Sprintf("%s%s%s-%s-%s-%s-%s-%s",
		periodCode,
		monthCode,
		dayCode,
		generateRandomSegment(),
		generateRandomSegment(),
		generateRandomSegment(),
		generateRandomSegment(),
		generateRandomSegment(),
	)

	return key
}

func CancelOrderUniversal(userId int64, orderId string, role string) (code int32, cancelled bool, msg string, err error) {
	log.Println("取消订单信息: ", userId, userId)

	var pendingOrder model.Orders
	redisKey := fmt.Sprintf("pending_order:%d:%s", userId, orderId)

	// 从 Redis 中取出数据并反序列化
	redisCtx, cancel := sysContext.WithTimeout(sysContext.Background(), 5*time.Second)
	defer cancel()
	val, err := dao.Rdb.Get(redisCtx, redisKey).Result()
	if errors.Is(err, redis.Nil) {
		log.Println("订单不存在于 Redis 中")
		return http.StatusNotFound, false, "订单未找到", nil
	} else if err != nil {
		log.Println("从 Redis 获取订单失败:", err)
		return http.StatusInternalServerError, false, "获取订单信息失败" + err.Error(), nil
	}

	// 反序列化 Redis 中的订单数据
	err = json.Unmarshal([]byte(val), &pendingOrder)
	if err != nil {
		log.Println("订单反序列化失败:", err)
		return http.StatusInternalServerError, false, "订单反序列化失败", nil
	}

	// 启动数据库事务
	err = dao.Db.Transaction(func(tx *gorm.DB) error {
		// 将 pendingOrder 插入订单表
		pendingOrder.IsFinished = true // 标记订单为已取消
		if role == "admin" {
			pendingOrder.FailureReason = "The order has been cancelled by the administrator" // 由管理员取消
		} else if role == "user" {
			pendingOrder.FailureReason = "The order has been cancelled by the user"
		} else {
			pendingOrder.FailureReason = "The order has been cancelled"
		}

		if err := tx.Create(&pendingOrder).Error; err != nil {
			log.Println("订单插入数据库失败:", err)
			return err
		}

		// 处理优惠券恢复
		if pendingOrder.CouponId != 0 {

			// 将 model.Coupon{}.Residue + 1
			err := tx.Model(&model.Coupon{}).
				Where("id = ?", pendingOrder.CouponId).
				Update("residue", gorm.Expr("residue + ?", 1)).
				Error
			if err != nil {
				return err
			}
		}

		// 删除 Redis 中的订单
		if err := dao.Rdb.Del(redisCtx, redisKey).Err(); err != nil {
			log.Println("从 Redis 删除订单失败:", err)
			return err
		}

		return nil
	})

	// 事务操作失败处理
	if err != nil {
		log.Println("取消订单失败:", err)
		return http.StatusInternalServerError, false, "取消订单失败", nil
	}

	// 事务成功，返回成功响应
	return http.StatusOK, true, "订单已取消", nil
}
