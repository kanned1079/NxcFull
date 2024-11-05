package utils

import (
	"context"
	"documentServices/internal/dao"
	"fmt"
	"time"
)

//func FreshDocumentsRedisCache() error {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	redisKey := fmt.Sprintf("documents:*")
//
//	// 删除 Redis 中的缓存键
//	if err := dao.Rdb.Del(ctx, redisKey).Err(); err != nil {
//		return fmt.Errorf("failed to clear cache in Redis: %w", err)
//	}
//	return nil
//}

func FreshDocumentsRedisCache() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pattern := "documents:*"

	// 查找匹配 pattern 的所有键
	keys, err := dao.Rdb.Keys(ctx, pattern).Result()
	if err != nil {
		return fmt.Errorf("failed to find keys: %w", err)
	}

	// 如果存在匹配的键，则删除
	if len(keys) > 0 {
		if err := dao.Rdb.Del(ctx, keys...).Err(); err != nil {
			return fmt.Errorf("failed to clear cache in Redis: %w", err)
		}
	}
	return nil
}
