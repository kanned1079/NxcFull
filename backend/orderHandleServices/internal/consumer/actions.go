package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
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

	// ------------------

	order := model.Orders{
		// 注意这里的修改后 密钥的处理查询OrderId而不是Id
		OrderId:        postOrderData.OrderId,
		UserId:         postOrderData.UserId, // 用户id
		PlanId:         postOrderData.PlanId, // 订阅计划id
		Period:         postOrderData.Period, // 付款周期
		Email:          "userMail",           // 用户邮箱
		DiscountAmount: 0.0,                  // 抵折金额
		Amount:         0.0,                  // 最终应该支付金额
		IsFinished:     false,                // 还未完成购买
		Status:         0,                    // 新购买
		CouponId:       postOrderData.CouponId,
		CreatedAt:      time.Now(),
	}

	// -----------

	var userMail string
	if result := dao.Db.Model(&model.User{}).Where("id = ?", postOrderData.UserId).Select("email").Limit(1).Find(&userMail); result.Error != nil {
		log.Println("获取邮箱地址失败")
		order.IsSuccess = false
		order.IsFinished = false
		order.FailureReason = "用户信息查询失败"
		ProcessErrorOrder(&order) // 直接处理失败的订单 跳过后面的步骤
		return result.Error
	}
	log.Println(userMail)

	var couponInfo model.Coupon
	if postOrderData.CouponId != 0 {
		if result := dao.Db.Model(&model.Coupon{}).Where("id = ?", postOrderData.CouponId).First(&couponInfo); result.Error != nil {
			log.Println(result.Error.Error())
			order.IsSuccess = false
			order.IsFinished = false
			order.FailureReason = "优惠券验证未通过"
			ProcessErrorOrder(&order) // 直接处理失败的订单 跳过后面的步骤
			return result.Error
		}
	}

	order.CouponName = couponInfo.Name

	var plan model.Plan
	if result := dao.Db.Where("id = ?", postOrderData.PlanId).First(&plan); result.Error != nil {
		log.Println(result.Error.Error())
		order.IsSuccess = false
		order.IsFinished = false
		order.FailureReason = "订阅计划非法"
		ProcessErrorOrder(&order) // 直接处理失败的订单 跳过后面的步骤
		return result.Error
	}

	if plan.Residue <= 0 {
		log.Println("订阅剩余不足")
		order.IsSuccess = false
		order.IsFinished = false
		order.FailureReason = "订阅剩余不足"
		ProcessErrorOrder(&order) // 直接处理失败的订单 跳过后面的步骤
		return errors.New("订阅剩余不足")
	}

	order.PlanName = plan.Name

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
	log.Println("原始价格", originalPrice)
	order.Price = originalPrice

	// 抵折百分比
	var discountPercent float64 = couponInfo.PercentOff / 100
	// 抵折的金额
	var disCountAmount float64 = originalPrice * discountPercent

	log.Println(discountPercent, disCountAmount, originalPrice-disCountAmount)

	order.DiscountAmount = disCountAmount // 优惠了的价格
	//order.Amount = originalPrice - disCountAmount // 订单价格
	order.Amount = math.Round((originalPrice-disCountAmount)*100) / 100 // 取整

	log.Println(order)

	couponUsed := model.CouponUsage{
		CouponId: couponInfo.Id,
		UserId:   postOrderData.UserId,
		UsedAt:   time.Now(),
	}

	if postOrderData.CouponId != 0 {
		if result := dao.Db.Model(&model.CouponUsage{}).Create(&couponUsed); result.Error != nil {
			// 添加优惠券使用记录出错
			order.IsSuccess = false
			order.IsFinished = false
			order.FailureReason = "添加优惠券使用记录出错"
			ProcessErrorOrder(&order) // 直接处理失败的订单 跳过后面的步骤
			return errors.New("添加优惠券使用记录出错" + result.Error.Error())
		}
	}

	// 存储到 Redis 并设置 5 分钟超时
	log.Println(postOrderData.OrderId)
	orderKey := fmt.Sprintf("%s:%d:%s", PendingOrderKeyPrefix, postOrderData.UserId, postOrderData.OrderId)
	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Println("订单序列化失败")
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := dao.Rdb.Set(ctx, orderKey, orderJSON, 6*time.Minute).Err(); err != nil {
		log.Println("Redis 存储失败")
		return err
	}

	// 检查五分钟后订单有没有完成(order.isFinished为true则为完成)
	//如果未完成则将数据插入数据库order表(dao.Db.Model(&model.Orders).Create(...))
	// 并删除redis中此键值
	go func(orderId string) {
		// 休眠五分钟
		log.Println("开始计时5分钟 订单号:", orderId)
		time.Sleep(5 * time.Minute)
		log.Println("订单确认失败")
		var order model.Orders
		// 查询订单，假设你在 Redis 中使用 orderId 作为键存储订单信息
		//redisKey := "pending_order:" + orderId
		orderKey := fmt.Sprintf("%s:%d:%s", PendingOrderKeyPrefix, postOrderData.UserId, orderId)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		val, err := dao.Rdb.Get(ctx, orderKey).Result()
		if errors.Is(err, redis.Nil) {
			// 如果 Redis 中没有此订单，说明可能已处理，直接返回 关闭此协程
			return
		} else if err != nil {
			// 处理 Redis 错误
			log.Println("Failed to fetch from Redis:", err)
			return
		}

		// 反序列化订单信息，假设你存储的是 JSON 格式的订单数据
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

		// 如果订单未完成，插入数据库
		order.IsFinished = true // 订单已经完成但是没有成功
		order.FailureReason = "Not confirmed within the validity period"
		if err := dao.Db.Model(&model.Orders{}).Create(&order).Error; err != nil {
			log.Println("Failed to insert order into database:", err)
			return
		}

		// 删除 Redis 中的订单键
		if err := dao.Rdb.Del(ctx, orderKey).Err(); err != nil {
			log.Println("Failed to delete order from Redis:", err)
		}

		log.Println("Order not finished, inserted into database and deleted from Redis.")
	}(postOrderData.OrderId)

	return nil

}

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
