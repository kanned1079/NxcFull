package config

import (
	"context"
	"errors"
	"log"
	"orderServices/internal/dao"
	"strconv"
	"time"
)

//var (
//	AlipayConfCache   Alipay
//	WechatConfCache   Wechat
//	AppleConfCache    Apple
//	PaymentMethodsArr []string // 存储有效支付方式的名称
//)

type InviteConf struct {
	InviteRebateEnable bool  `json:"invite_rebate_enable"`
	InviteRebateRate   int32 `json:"invite_rebate_rate"`
}

var InviteConfCache InviteConf

func (i *InviteConf) FetchConfFromRedis() error {
	// 定义 Redis 键
	redisKey := "settings:invite"

	// 使用 context 管理超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 从 Redis 获取所有字段值
	entries, err := dao.Rdb.HGetAll(ctx, redisKey).Result()
	if err != nil {
		log.Printf("Failed to fetch Redis key '%s': %v\n", redisKey, err)
		return err
	}

	// 检查是否有数据返回
	if len(entries) == 0 {
		return errors.New("no data found for Redis key" + redisKey)
	}

	// 解析字段 invite_rebate_enable
	if enableStr, ok := entries["invite_rebate_enable"]; ok {
		parsedVal, err := strconv.ParseBool(enableStr)
		if err != nil {
			log.Printf("Failed to parse 'invite_rebate_enable' as bool: %v\n", err)
			return err
		} else {
			i.InviteRebateEnable = parsedVal
		}
	} else {
		log.Println("Key 'invite_rebate_enable' not found in Redis hash")
	}

	// 解析字段 invite_rebate_rate
	if rateStr, ok := entries["invite_rebate_rate"]; ok {
		parsedVal, err := strconv.Atoi(rateStr)
		if err != nil {
			log.Printf("Failed to parse 'invite_rebate_rate' as int: %v\n", err)
			return err
		} else {
			i.InviteRebateRate = int32(parsedVal)
		}
	} else {
		log.Println("Key 'invite_rebate_rate' not found in Redis hash")
		return err
	}
	return nil
}
