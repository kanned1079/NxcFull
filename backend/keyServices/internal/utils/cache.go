package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"keyServices/internal/dao"
	"log"
	"net/http"
	"strconv"
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

// SetActivationRecordCache2Redis 激活后放到redis缓存中 过期时间为到期时间-当前时间+1天
// SetActivationRecordCache2Redis 激活后放到redis缓存中 过期时间为到期时间-当前时间+1天
func SetActivationRecordCache2Redis(ctx context.Context, clientId, keyContent string, expirationDate *time.Time) error {
	log.Println("client_id: ", clientId)
	log.Println("key_content: ", keyContent)

	// 确保 expirationDate 不为空
	if expirationDate == nil {
		log.Println("expirationDate is nil")
		return fmt.Errorf("expirationDate cannot be nil")
	}

	// 将 expirationDate 转换为 Unix 时间戳（秒级别）
	expirationTimestamp := expirationDate.Unix()

	// 生成 Redis 键
	var redisKeys string = "user_property:active_activation_record_cache:" + clientId

	// 设置 Redis 键值对
	res := dao.Rdb.HSet(ctx, redisKeys, map[string]any{
		"key_content":     keyContent,
		"client_id":       clientId,
		"expiration_date": expirationTimestamp, // 存储 Unix 时间戳
	})
	if res.Err() != nil {
		log.Println("failed to set cache:", res.Err())
		return res.Err()
	}

	// 计算从当前时间到过期日期的时间差并加上一天
	timeRemaining := expirationDate.Sub(time.Now()) + 24*time.Hour

	// 设置过期时间
	expireRes := dao.Rdb.Expire(ctx, redisKeys, timeRemaining)
	if expireRes.Err() != nil {
		log.Println("failed to set expiration date:", expireRes.Err())
		return expireRes.Err()
	}

	// 检查是否成功设置过期时间
	if !expireRes.Val() {
		log.Println("failed to set expiration date, Redis returned false")
		return fmt.Errorf("failed to set expiration date")
	}

	return nil
}

//func SetActivationRecordCache2Redis(ctx context.Context, clientId, keyContent string, expirationDate *time.Time) error {
//	log.Println("client_id: ", clientId)
//	log.Println("key_content: ", keyContent)
//	var redisKeys string = "user_property:active_activation_record_cache:" + clientId
//
//	// 设置 Redis 键值对
//	res := dao.Rdb.HSet(ctx, redisKeys, map[string]any{
//		"key_content":     keyContent,
//		"client_id":       clientId,
//		"expiration_date": expirationDate,
//	})
//	if res.Err() != nil {
//		log.Println("failed to set cache:", res.Err())
//		return res.Err()
//	}
//
//	// 计算从当前时间到过期日期的时间差并加上一天
//	timeRemaining := expirationDate.Sub(time.Now()) + 24*time.Hour
//
//	// 设置过期时间
//	expireRes := dao.Rdb.Expire(ctx, redisKeys, timeRemaining)
//	if expireRes.Err() != nil {
//		log.Println("failed to set expiration date:", expireRes.Err())
//		return expireRes.Err()
//	}
//
//	// 检查是否成功设置过期时间
//	if !expireRes.Val() {
//		log.Println("failed to set expiration date, Redis returned false")
//		return fmt.Errorf("failed to set expiration date")
//	}
//
//	return nil
//}

// CheckActivationFromCache 检查缓存中的激活信息
// CheckActivationFromCache 检查缓存中的激活记录并返回状态码和过期时间
func CheckActivationFromCache(ctx context.Context, clientId, keyContent string) (code int32, expirationDate *time.Time) {
	// 生成 Redis 键
	log.Println(clientId, keyContent)
	var redisKeys string = "user_property:active_activation_record_cache:" + clientId

	// 从 Redis 获取客户端对应的哈希数据
	result, err := dao.Rdb.HGetAll(ctx, redisKeys).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// 如果 Redis 返回 Nil，表示没有找到该键
			log.Println("Cache not found for clientId:", clientId)
			return http.StatusNotFound, nil // 404: not found
		}
		// 如果其他错误发生，返回 500
		log.Println("Error fetching data from Redis:", err)
		return http.StatusInternalServerError, nil // 500: internal server error
	}

	// 如果 Redis 中没有找到对应的数据
	if len(result) == 0 {
		log.Println("No activation record found in cache for clientId:", clientId)
		return http.StatusNotFound, nil // 404: not found
	}

	// 获取缓存中的 "key_content" 和 "expirationDate"
	storedKeyContent, keyExists := result["key_content"]
	storedExpirationDate, expExists := result["expiration_date"]
	if !keyExists || !expExists {
		// 如果缓存中没有 "key_content" 或 "expirationDate" 字段，返回 500
		log.Println("key_content or expirationDate not found in cache for clientId:", clientId)
		return http.StatusInternalServerError, nil // 500: internal server error
	}

	// 将 "expirationDate" 从 Unix 时间戳解析为 time.Time 类型
	expirationUnix, err := strconv.ParseInt(storedExpirationDate, 10, 64)
	if err != nil {
		log.Println("Error parsing expirationDate timestamp:", err)
		return http.StatusInternalServerError, nil // 500: internal server error
	}

	expirationTime := time.Unix(expirationUnix, 0) // 转换为 time.Time

	// 检查缓存中的 "key_content" 是否与传入的 keyContent 匹配
	if storedKeyContent == keyContent {
		// 如果 "key_content" 匹配，返回 200 和过期时间
		log.Println("Activation successful for clientId:", clientId)
		return http.StatusOK, &expirationTime // 200: success
	}

	// 如果 "key_content" 不匹配，返回 409 和过期时间
	log.Println("Key content mismatch for clientId:", clientId)
	return http.StatusConflict, &expirationTime // 409: conflict
}
