package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"orderServices/internal/dao"
	"time"
)

// ReadSettingFromRedisCache 从redis中读取一个配置项目
func ReadSettingFromRedisCache(category, key string) (json.RawMessage, error) {
	redisKey := fmt.Sprintf("settings:%s", category)

	// 从 Redis 中获取指定 Key 的值
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	valueStr, err := dao.Rdb.HGet(ctx, redisKey, key).Result()
	if errors.Is(err, redis.Nil) {
		log.Printf("Setting not found in Redis cache for category '%s' and key '%s'.", category, key)
		return nil, nil // 未找到记录
	} else if err != nil {
		return nil, err // 其他错误
	}

	// 将字符串转换为 json.RawMessage
	// 直接使用 `json.RawMessage(valueStr)` 将 string 转换为 RawMessage
	return json.RawMessage(valueStr), nil
}
