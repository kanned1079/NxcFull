package consumer

import (
	sysContext "context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
	"math"
	"orderHandleServices/internal/dao"
	"orderHandleServices/internal/model"
	"time"
)

const (
	PendingOrderKeyPrefix = "pending_order"
)

func ProcessOrder(jsonData []byte) error {

	var postOrderData = &struct {
		OrderId  string `json:"order_id"`
		UserId   int64  `json:"user_id"`
		PlanId   int64  `json:"plan_id"`
		Period   string `json:"period"`
		CouponId int64  `json:"coupon_id"`
	}{}
	err := json.Unmarshal(jsonData, &postOrderData)
	if err != nil {
		return err
	}
	log.Println(postOrderData)

	// 开始数据库事务
	tx := dao.Db.Begin()
	if tx.Error != nil {
		log.Println("Failed to start transaction:", tx.Error)
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Recovered from panic, transaction rolled back.")
		}
	}()

	// 查询用户邮箱
	var userMail string
	if result := tx.Model(&model.User{}).Where("id = ?", postOrderData.UserId).Select("email").Limit(1).Find(&userMail); result.Error != nil {
		log.Println("获取邮箱地址失败")
		tx.Rollback()
		return result.Error
	}
	log.Println(userMail)

	var couponInfo model.Coupon
	if postOrderData.CouponId != 0 {
		if result := tx.Model(&model.Coupon{}).Where("id = ?", postOrderData.CouponId).First(&couponInfo); result.Error != nil {
			log.Println(result.Error.Error())
			tx.Rollback()
			return result.Error
		}
	}

	var plan model.Plan
	if result := tx.Where("id = ?", postOrderData.PlanId).First(&plan); result.Error != nil {
		log.Println(result.Error.Error())
		tx.Rollback()
		return result.Error
	}

	if plan.Residue <= 0 {
		log.Println("订阅剩余不足")
		tx.Rollback()
		return errors.New("订阅剩余不足")
	}

	// 构建订单信息
	order := model.Orders{
		OrderId:        postOrderData.OrderId,
		UserId:         postOrderData.UserId,
		PlanId:         postOrderData.PlanId,
		Period:         postOrderData.Period,
		Email:          userMail,
		DiscountAmount: 0.0,
		Amount:         0.0,
		IsFinished:     false,
		Status:         0,
		CouponId:       postOrderData.CouponId,
		CreatedAt:      time.Now(),
		CouponName:     couponInfo.Name,
		GroupId:        plan.GroupId,
		PlanName:       plan.Name,
	}

	var originalPrice float64
	switch postOrderData.Period {
	case "month":
		originalPrice = plan.MonthPrice
	case "quarter":
		originalPrice = plan.QuarterPrice
	case "half-year":
		originalPrice = plan.HalfYearPrice
	case "year":
		originalPrice = plan.YearPrice
	}

	var discountPercent float64 = couponInfo.PercentOff / 100
	var disCountAmount float64 = originalPrice * discountPercent
	order.Price = originalPrice
	order.DiscountAmount = math.Round((disCountAmount)*100) / 100
	order.Amount = math.Round((originalPrice-disCountAmount)*100) / 100

	// 插入订单记录
	// 先存入redis 超时后再存入数据库
	//if err := tx.Create(&order).Error; err != nil {
	//	log.Println("订单插入失败:", err)
	//	tx.Rollback()
	//	return err
	//}

	// 添加优惠券使用记录
	if postOrderData.CouponId != 0 {
		couponUsed := model.CouponUsage{
			CouponId: couponInfo.Id,
			UserId:   postOrderData.UserId,
			UsedAt:   time.Now(),
		}
		if err := tx.Create(&couponUsed).Error; err != nil {
			log.Println("添加优惠券使用记录失败:", err)
			tx.Rollback()
			return err
		}
		if err := tx.Model(&model.Coupon{}).Where("id = ?", postOrderData.CouponId).UpdateColumn("residue", gorm.Expr("residue - 1")).Error; err != nil {
			log.Println("更新优惠券剩余失败:", err)
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Println("事务提交失败:", err)
		return err
	}

	// 存储到 Redis 并设置 5 分钟超时
	orderKey := fmt.Sprintf("%s:%d:%s", PendingOrderKeyPrefix, postOrderData.UserId, postOrderData.OrderId)
	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Println("订单序列化失败")
		return err
	}

	ctx, cancel := sysContext.WithTimeout(sysContext.Background(), 5*time.Second)
	defer cancel()

	if err := dao.Rdb.Set(ctx, orderKey, orderJSON, 6*time.Minute).Err(); err != nil {
		log.Println("Redis 存储失败")
		return err
	}

	go func(orderId string) {
		log.Println("开始计时5分钟 订单号:", orderId)
		time.Sleep(5 * time.Minute)
		log.Println("订单确认失败")

		var order model.Orders
		orderKey := fmt.Sprintf("%s:%d:%s", PendingOrderKeyPrefix, postOrderData.UserId, orderId)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// 查询 Redis 中的订单
		val, err := dao.Rdb.Get(ctx, orderKey).Result()
		if errors.Is(err, redis.Nil) {
			// Redis 中没有此订单，说明已处理，直接返回
			log.Println("Order not found in Redis, possibly already processed.")
			return
		} else if err != nil {
			// 处理 Redis 错误
			log.Println("Failed to fetch from Redis:", err)
			return
		}

		// 反序列化订单
		err = json.Unmarshal([]byte(val), &order)
		if err != nil {
			log.Println("Failed to unmarshal order:", err)
			return
		}

		// 检查订单是否已完成
		if order.IsFinished {
			log.Println("Order is already finished, no action needed.")
			return
		}

		// 开始数据库事务
		tx := dao.Db.Begin()
		if tx.Error != nil {
			log.Println("Failed to start transaction:", tx.Error)
			return
		}

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				log.Println("Recovered from panic, transaction rolled back.")
			}
		}()

		// 更新订单状态
		order.IsFinished = true
		order.FailureReason = "Not confirmed within the validity period"

		if err := tx.Model(&model.Orders{}).Create(&order).Error; err != nil {
			tx.Rollback()
			log.Println("Failed to insert order into database:", err)
			return
		}

		// 处理优惠券恢复
		if order.CouponId != 0 {
			log.Println("恢复优惠券容量", order.CouponId)
			//if err := tx.Model(&model.Coupon{}).Where("id = ?", postOrderData.CouponId).UpdateColumn("residue", gorm.Expr("residue + 1")).Error; err != nil {
			if err := tx.Model(&model.Coupon{}).Where("id = ?", order.CouponId).UpdateColumn("residue", gorm.Expr("residue + 1")).Error; err != nil {

				tx.Rollback()
				log.Println("Failed to update coupon residue:", err)
				return
			}
		}

		// 提交事务
		if err := tx.Commit().Error; err != nil {
			log.Println("Failed to commit transaction:", err)
			return
		}

		// 删除 Redis 中的订单键
		if err := dao.Rdb.Del(ctx, orderKey).Err(); err != nil {
			log.Println("Failed to delete order from Redis:", err)
		}

		log.Println("Order not finished, inserted into database, and deleted from Redis.")
	}(postOrderData.OrderId)

	return nil

}

/*

// 检查五分钟后订单有没有完成(order.isFinished为true则为完成)
	//如果未完成则将数据插入数据库order表(dao.Db.Model(&model.Orders).Create(...))
	// 并删除redis中此键值
	//go func(orderId string) {
	//	// 休眠五分钟
	//	log.Println("开始计时5分钟 订单号:", orderId)
	//	time.Sleep(5 * time.Minute)
	//	log.Println("订单确认失败")
	//	var order model.Orders
	//	// 查询订单，假设你在 Redis 中使用 orderId 作为键存储订单信息
	//	//redisKey := "pending_order:" + orderId
	//	orderKey := fmt.Sprintf("%s:%d:%s", PendingOrderKeyPrefix, postOrderData.UserId, orderId)
	//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//	defer cancel()
	//	val, err := dao.Rdb.Get(ctx, orderKey).Result()
	//	if errors.Is(err, redis.Nil) {
	//		// 如果 Redis 中没有此订单，说明可能已处理，直接返回 关闭此协程
	//		return
	//	} else if err != nil {
	//		// 处理 Redis 错误
	//		log.Println("Failed to fetch from Redis:", err)
	//		return
	//	}
	//
	//	// 反序列化订单信息，假设你存储的是 JSON 格式的订单数据
	//	err = json.Unmarshal([]byte(val), &order)
	//	if err != nil {
	//		log.Println("Failed to unmarshal order:", err)
	//		return
	//	}
	//
	//	// 检查订单是否已完成
	//	if order.IsFinished {
	//		log.Println("Order is already finished, no action needed.")
	//		return
	//	}
	//
	//	// 如果订单未完成，插入数据库
	//	order.IsFinished = true // 订单已经完成但是没有成功
	//	order.FailureReason = "Not confirmed within the validity period"
	//	if err := dao.Db.Model(&model.Orders{}).Create(&order).Error; err != nil {
	//		log.Println("Failed to insert order into database:", err)
	//		return
	//	}
	//	if order.CouponId != 0 {
	//		if err := dao.Db.Model(&model.Coupon{}).Where("id = ?", postOrderData.CouponId).First(&order).Error; err != nil {
	//		}
	//		// 将该优惠券的residue值 +1
	//	}
	//
	//	// 删除 Redis 中的订单键
	//	if err := dao.Rdb.Del(ctx, orderKey).Err(); err != nil {
	//		log.Println("Failed to delete order from Redis:", err)
	//	}
	//
	//	log.Println("Order not finished, inserted into database and deleted from Redis.")
	//}(postOrderData.OrderId)

*/

//func ProcessOrder(jsonData []byte) error {
//	var postOrderData = &struct {
//		OrderId  string `json:"order_id"`
//		UserId   int64  `json:"user_id"`
//		PlanId   int64  `json:"plan_id"`
//		Period   string `json:"period"`
//		CouponId int64  `json:"coupon_id"`
//	}{}
//	err := json.Unmarshal(jsonData, &postOrderData)
//	if err != nil {
//		return err
//	}
//	log.Println(postOrderData)
//	var order model.Orders
//
//	// 开启事务
//	err = dao.Db.Transaction(func(tx *gorm.DB) error {
//		order = model.Orders{
//			OrderId:        postOrderData.OrderId,
//			UserId:         postOrderData.UserId,
//			PlanId:         postOrderData.PlanId,
//			Period:         postOrderData.Period,
//			Email:          "userMail",
//			DiscountAmount: 0.0,
//			Amount:         0.0,
//			IsFinished:     false,
//			Status:         0,
//			CouponId:       postOrderData.CouponId,
//			CreatedAt:      time.Now(),
//		}
//
//		// 查询用户邮箱
//		var userMail string
//		if result := tx.Model(&model.User{}).Where("id = ?", postOrderData.UserId).Select("email").Limit(1).Find(&userMail); result.Error != nil {
//			log.Println("获取邮箱地址失败")
//			order.IsSuccess = false
//			order.IsFinished = false
//			order.FailureReason = "用户信息查询失败"
//			ProcessErrorOrder(&order)
//			return result.Error
//		}
//		log.Println(userMail)
//		order.Email = userMail
//
//		// 查询优惠券
//		var couponInfo model.Coupon
//		if postOrderData.CouponId != 0 {
//			if result := tx.Model(&model.Coupon{}).Where("id = ?", postOrderData.CouponId).First(&couponInfo); result.Error != nil {
//				log.Println(result.Error.Error())
//				order.IsSuccess = false
//				order.IsFinished = false
//				order.FailureReason = "优惠券验证未通过"
//				ProcessErrorOrder(&order)
//				return result.Error
//			}
//		}
//		order.CouponName = couponInfo.Name
//
//		// 查询订阅计划
//		var plan model.Plan
//		if result := tx.Where("id = ?", postOrderData.PlanId).First(&plan); result.Error != nil {
//			log.Println(result.Error.Error())
//			order.IsSuccess = false
//			order.IsFinished = false
//			order.FailureReason = "订阅计划非法"
//			ProcessErrorOrder(&order)
//			return result.Error
//		}
//
//		if plan.Residue <= 0 {
//			log.Println("订阅剩余不足")
//			order.IsSuccess = false
//			order.IsFinished = false
//			order.FailureReason = "订阅剩余不足"
//			ProcessErrorOrder(&order)
//			return errors.New("订阅剩余不足")
//		}
//
//		order.PlanName = plan.Name
//
//		// 计算订单价格
//		var originalPrice float64
//		switch postOrderData.Period {
//		case "month":
//			originalPrice = plan.MonthPrice
//		case "quarter":
//			originalPrice = plan.QuarterPrice
//		case "half-year":
//			originalPrice = plan.HalfYearPrice
//		case "year":
//			originalPrice = plan.YearPrice
//		}
//		log.Println("原始价格", originalPrice)
//		order.Price = originalPrice
//
//		// 计算折扣
//		var discountPercent float64 = couponInfo.PercentOff / 100
//		var disCountAmount float64 = originalPrice * discountPercent
//		log.Println(discountPercent, disCountAmount, originalPrice-disCountAmount)
//		order.DiscountAmount = disCountAmount
//		order.Amount = math.Round((originalPrice-disCountAmount)*100) / 100
//
//		// 记录优惠券使用
//		if postOrderData.CouponId != 0 {
//			couponUsed := model.CouponUsage{
//				CouponId: couponInfo.Id,
//				UserId:   postOrderData.UserId,
//				UsedAt:   time.Now(),
//			}
//			if result := tx.Model(&model.CouponUsage{}).Create(&couponUsed); result.Error != nil {
//				order.IsSuccess = false
//				order.IsFinished = false
//				order.FailureReason = "添加优惠券使用记录出错"
//				ProcessErrorOrder(&order)
//				return errors.New("添加优惠券使用记录出错" + result.Error.Error())
//			}
//		}
//
//		// 存储订单
//		if err := tx.Create(&order).Error; err != nil {
//			log.Println("订单创建失败:", err)
//			return err
//		}
//
//		// 提交事务
//		return nil
//	})
//
//	if err != nil {
//		return err
//	}
//
//	// Redis 操作和订单处理逻辑可以放在事务外部
//	orderKey := fmt.Sprintf("%s:%d:%s", PendingOrderKeyPrefix, postOrderData.UserId, postOrderData.OrderId)
//	orderJSON, err := json.Marshal(order)
//	if err != nil {
//		log.Println("订单序列化失败")
//		return err
//	}
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	if err := dao.Rdb.Set(ctx, orderKey, orderJSON, 6*time.Minute).Err(); err != nil {
//		log.Println("Redis 存储失败")
//		return err
//	}
//
//	go handlePendingOrder(postOrderData.OrderId, postOrderData.UserId)
//
//	return nil
//}

//func handlePendingOrder(orderId string, userId int64) {
//	// 休眠五分钟
//	log.Println("开始计时5分钟 订单号:", orderId)
//	time.Sleep(1 * time.Minute)
//	log.Println("订单确认失败")
//
//	orderKey := fmt.Sprintf("%s:%d:%s", PendingOrderKeyPrefix, userId, orderId)
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	val, err := dao.Rdb.Get(ctx, orderKey).Result()
//	if errors.Is(err, redis.Nil) {
//		return
//	} else if err != nil {
//		log.Println("Failed to fetch from Redis:", err)
//		return
//	}
//
//	var order model.Orders
//	err = json.Unmarshal([]byte(val), &order)
//	if err != nil {
//		log.Println("Failed to unmarshal order:", err)
//		return
//	}
//
//	if order.IsFinished {
//		log.Println("Order is already finished, no action needed.")
//		return
//	}
//
//	order.IsFinished = true
//	order.FailureReason = "Not confirmed within the validity period"
//	if err := dao.Db.Model(&model.Orders{}).Create(&order).Error; err != nil {
//		log.Println("Failed to insert order into database:", err)
//		return
//	}
//
//	if err := dao.Rdb.Del(ctx, orderKey).Err(); err != nil {
//		log.Println("Failed to delete order from Redis:", err)
//	}
//
//	log.Println("Order not finished, inserted into database and deleted from Redis.")
//}

//return &pb.CommitNewOrderResponse{
//	Code:           http.StatusOK,
//	OrderId:        orderIdGenerated,
//	PlanName:       plan.Name,
//	OriginalPrice:  float32(originalPrice),
//	DiscountAmount: float32(disCountAmount),
//	PayPrice:       float32(originalPrice - discountPercent),
//	Period:         request.Period,
//	CreatedAt:      time.Now().Format("2006-01-02 15:04:05"),
//}, nil

func ProcessErrorOrder(badOrder *model.Orders) {
	if result := dao.Db.Model(&model.Orders{}).Create(badOrder); result.Error != nil {
		log.Println(result.Error.Error())
	}
}
