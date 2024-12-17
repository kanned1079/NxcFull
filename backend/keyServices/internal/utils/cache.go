package utils

import (
	"context"
	"fmt"
	"keyServices/internal/dao"
	"log"
	"time"
)

// ClearUserKeyCache 清除指定用户的所有缓存数据
func ClearUserKeyCache(userId int64) {
	log.Println("flush user key cache, userId: ", userId)
	// 构建缓存的 key 模式
	cacheKeyPattern := fmt.Sprintf("user_property:user_keys:%d:page:*:size:*", userId)

	// 获取所有匹配的缓存键
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	keys, err := dao.Rdb.Keys(ctx, cacheKeyPattern).Result()
	if err != nil {
		log.Printf("获取匹配缓存键失败: %v", err)
	}

	if len(keys) == 0 {
		log.Printf("没有找到用户(%d)的缓存", userId)
	}

	// 删除所有匹配的缓存键
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = dao.Rdb.Del(ctx, keys...).Err()
	if err != nil {
		log.Printf("清除用户(%d)缓存失败: %v", userId, err)
	}

	log.Printf("成功清除用户(%d)的缓存", userId)
}
