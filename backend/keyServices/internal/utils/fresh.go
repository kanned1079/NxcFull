package utils

import (
	"context"
	"errors"
	"fmt"
	"keyServices/internal/dao"
	"keyServices/internal/model"
	"log"
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

// CloseActiveActivationCheckCacheByKeyId 清除redis中缓存的激活记录 如果它不成功则需要回滚事务
func CloseActiveActivationCheckCacheByKeyId(ctx context.Context, keyId int64) error {
	var clientId string = ""
	if err := dao.Db.Model(&model.Keys{}).Where("`id` = ?", keyId).Select("client_id").Scan(&clientId).Error; err != nil {
		log.Println("failed to get client_id:", err)
		return err
	}
	if clientId == "" || len(clientId) != 64 {
		return errors.New("client_id of key is not valid")
	} else {
		if err := dao.Rdb.Del(ctx, "user_property:active_activation_record_cache:"+clientId).Err(); err != nil {
			log.Println("failed to delete cache:", err)
			return err
		}
	}
	return nil
}
