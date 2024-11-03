package utils

import (
	"context"
	"fmt"
	"orderHandleServices/internal/dao"
	"time"
)

func FreshUserPropertyInfoInRedis(userId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 1 删除密钥列表缓存
	// 构建匹配该用户的键的模式
	pattern := fmt.Sprintf("user_property:user_keys:%d:*", userId)
	// 使用 Scan 命令查找匹配的键，以避免使用 KEYS 可能带来的性能问题
	var cursor uint64
	for {
		// 扫描符合模式的键
		keys, nextCursor, err := dao.Rdb.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return fmt.Errorf("failed to scan keys with pattern %s: %v", pattern, err)
		}

		// 删除找到的键
		if len(keys) > 0 {
			if err := dao.Rdb.Del(ctx, keys...).Err(); err != nil {
				return fmt.Errorf("failed to delete keys %v: %v", keys, err)
			}
			fmt.Printf("Deleted keys: %v\n", keys)
		}

		// 更新游标
		cursor = nextCursor
		if cursor == 0 {
			break // 扫描完成
		}
	}

	// 2.删除用户首页的"我的订阅"中的缓存
	// 构建 Redis 键
	redisKey := fmt.Sprintf("user_property:active_order_summary:%d", userId)

	// 删除 Redis 中的缓存键
	if err := dao.Rdb.Del(ctx, redisKey).Err(); err != nil {
		return fmt.Errorf("failed to clear cache in Redis: %w", err)
	}

	fmt.Printf("All keys for user %d deleted successfully\n", userId)
	return nil
}
