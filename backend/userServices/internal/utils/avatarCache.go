package utils

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"userServices/internal/dao"
)

// GetAvatarCache 读取用户头像数据从 Redis
func GetAvatarCache(ctx context.Context, userId int64) ([]byte, string, error) {
	redisKeyFileName := fmt.Sprintf("user_avatar:%d:file_name", userId)
	redisKeyFileData := fmt.Sprintf("user_avatar:%d:file_data", userId)

	// 读取文件名
	fileName, err := dao.Rdb.Get(ctx, redisKeyFileName).Result()
	if errors.Is(err, redis.Nil) {
		// 如果没有文件名缓存
		return nil, "", nil
	}
	if err != nil {
		return nil, "", err
	}

	// 读取文件内容
	fileData, err := dao.Rdb.Get(ctx, redisKeyFileData).Result()
	if errors.Is(err, redis.Nil) {
		// 如果没有文件数据缓存
		return nil, "", nil
	}
	if err != nil {
		return nil, "", err
	}

	return []byte(fileData), fileName, nil
}

// SetAvatarCache 将用户头像数据写入 Redis
func SetAvatarCache(ctx context.Context, userId int64, fileName string, fileData []byte) error {
	redisKeyFileName := fmt.Sprintf("user_avatar:%d:file_name", userId)
	redisKeyFileData := fmt.Sprintf("user_avatar:%d:file_data", userId)

	// 将文件名和文件内容存储到 Redis
	if err := dao.Rdb.Set(ctx, redisKeyFileName, fileName, 24*time.Hour).Err(); err != nil {
		return err
	}

	// 将文件内容存储为 base64 编码字符串，防止二进制数据的问题
	encodedFileData := base64.StdEncoding.EncodeToString(fileData)
	if err := dao.Rdb.Set(ctx, redisKeyFileData, encodedFileData, 24*time.Hour).Err(); err != nil {
		return err
	}

	return nil
}
