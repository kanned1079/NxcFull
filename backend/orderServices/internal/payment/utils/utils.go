package utils

import (
	"context"
	"fmt"
	"log"
	"orderServices/internal/dao"
	"strconv"
	"time"
)

//func SavePreCreateOrder2Redis(outTradeNo string, amount float32, finalAmount float32, discount float32) error {
//	// 定义 Redis 键
//	redisKey := fmt.Sprintf("topup:%s", outTradeNo)
//
//	// 定义要存储的订单数据
//	orderPreSave := map[string]any{
//		"out_trade_no": outTradeNo,  // 商家订单号
//		"amount":       amount,      // 充值金额
//		"final_amount": finalAmount, // 最终支付金额
//		"discount":     discount,    // 优惠金额
//	}
//
//	// 将数据存入 Redis
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	// 使用 HSet 存储哈希值
//	err := dao.Rdb.HSet(ctx, redisKey, orderPreSave).Err()
//	if err != nil {
//		log.Printf("Failed to save pre-create order to Redis: %v\n", err)
//		return err
//	}
//
//	// 设置过期时间 2h
//	// 预订单保留时间为2h
//	// 文档 https://opendocs.alipay.com/open/8ad49e4a_alipay.trade.precreate?pathHash=d18bff53&scene=2ae8516856f24a5592194d237f3f235d
//	expiration := 2 * time.Hour
//	err = dao.Rdb.Expire(ctx, redisKey, expiration).Err()
//	if err != nil {
//		log.Printf("Failed to set expiration for Redis key '%s': %v\n", redisKey, err)
//		return err
//	}
//
//	log.Printf("Pre-create order saved to Redis successfully: %s\n", redisKey)
//	return nil
//}

func SavePreCreateOrder2Redis(outTradeNo string, amount float32, finalAmount float32, discount float32) error {
	// 定义 Redis 键
	redisKey := fmt.Sprintf("topup:%s", outTradeNo)

	// 格式化浮点数为字符串（两位小数）
	orderPreSave := map[string]any{
		"out_trade_no": outTradeNo,                       // 商家订单号
		"amount":       fmt.Sprintf("%.2f", amount),      // 充值金额
		"final_amount": fmt.Sprintf("%.2f", finalAmount), // 最终支付金额
		"discount":     fmt.Sprintf("%.2f", discount),    // 优惠金额
	}

	// 将数据存入 Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := dao.Rdb.HSet(ctx, redisKey, orderPreSave).Err()
	if err != nil {
		log.Printf("Failed to save pre-create order to Redis: %v\n", err)
		return err
	}

	// 设置过期时间（假设订单预创建信息有效期为30分钟）
	expiration := 2 * time.Hour
	err = dao.Rdb.Expire(ctx, redisKey, expiration).Err()
	if err != nil {
		log.Printf("Failed to set expiration for Redis key '%s': %v\n", redisKey, err)
		return err
	}

	log.Printf("Pre-create order saved to Redis successfully: %s\n", redisKey)
	return nil
}

func GetPreCreateOrderFromRedis(outTradeNo string) (string, float32, float32, float32, error) {
	redisKey := fmt.Sprintf("topup:%s", outTradeNo)
	log.Println("查询redisKey：", redisKey)

	// 从 Redis 获取哈希数据
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entries, err := dao.Rdb.HGetAll(ctx, redisKey).Result()
	if err != nil {
		log.Printf("Failed to fetch pre-create order from Redis: %v\n", err)
		return "", 0, 0, 0, err
	}

	// 提取并解析数据
	var outTradeNoVal string
	var amount, finalAmount, discount float32

	// 从哈希表中获取并解析字段
	if val, exists := entries["out_trade_no"]; exists {
		outTradeNoVal = val
	}

	// 解析浮动金额字段
	if val, exists := entries["amount"]; exists {
		parsedValue, err := strconv.ParseFloat(val, 32)
		if err != nil {
			log.Printf("Error parsing float for amount: %v\n", err)
		} else {
			amount = float32(parsedValue)
		}
	}

	if val, exists := entries["final_amount"]; exists {
		parsedValue, err := strconv.ParseFloat(val, 32)
		if err != nil {
			log.Printf("Error parsing float for final_amount: %v\n", err)
		} else {
			finalAmount = float32(parsedValue)
		}
	}

	if val, exists := entries["discount"]; exists {
		parsedValue, err := strconv.ParseFloat(val, 32)
		if err != nil {
			log.Printf("Error parsing float for discount: %v\n", err)
		} else {
			discount = float32(parsedValue)
		}
	}

	log.Printf("Pre-create order fetched from Redis: out_trade_no=%s, amount=%.2f, final_amount=%.2f, discount=%.2f\n", outTradeNoVal, amount, finalAmount, discount)

	// 返回查询的值
	return outTradeNoVal, amount, finalAmount, discount, nil
}
