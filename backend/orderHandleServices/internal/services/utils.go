package services

import (
	"context"
	"fmt"
	"log"
	"orderHandleServices/internal/dao"
)

func ClearUserKeysCache(ctx context.Context, userId int64) error {
	// 构造 Redis 缓存的键
	cacheSetKey := fmt.Sprintf("user_keys_set:%d", userId)
	cacheDataPrefix := fmt.Sprintf("user_key_data:%d:", userId)

	// 开始事务，先删除 SET，再删除每个订单详细数据
	pipe := dao.Rdb.TxPipeline()

	// 删除 SET 键
	pipe.Del(ctx, cacheSetKey)

	// 查找与用户相关的所有 key 数据
	keys, err := dao.Rdb.Keys(ctx, cacheDataPrefix+"*").Result()
	if err != nil {
		log.Printf("查询用户 %d 的缓存键失败: %v", userId, err)
		return err
	}

	// 删除每个详细数据缓存
	for _, key := range keys {
		pipe.Del(ctx, key)
	}

	// 提交事务
	if _, err := pipe.Exec(ctx); err != nil {
		log.Printf("删除用户 %d 的缓存失败: %v", userId, err)
		return err
	}

	log.Printf("用户 %d 的缓存删除成功", userId)
	return nil
}
