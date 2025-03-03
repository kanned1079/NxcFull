package utils

import (
	"crypto/rand"
	"fmt"
	"keyServices/internal/dao"
	"keyServices/internal/model"
	"log"
	"math/big"
	"strings"
)

const clientIdLength = 64
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateClientId 生成一个唯一的 clientId
func GenerateClientId() (string, error) {
	// 用于生成 clientId 的字符数组
	var clientIdBuilder strings.Builder
	for i := 0; i < clientIdLength; i++ {
		// 生成一个随机数
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %v", err)
		}
		// 使用该随机数来选择字符
		clientIdBuilder.WriteByte(charset[index.Int64()])
	}

	clientId := clientIdBuilder.String()

	// 检查数据库中是否有重复的 clientId
	var count int64
	if err := dao.Db.Model(&model.Keys{}).Where("`client_id` = ?", clientId).Count(&count).Error; err != nil {
		log.Printf("数据库查询错误: %v\n", err)
		return "", fmt.Errorf("database error: %v", err)
	}

	// 如果 clientId 已经存在，重新生成
	if count > 0 {
		return GenerateClientId() // 递归调用直到生成一个唯一的 clientId
	}

	return clientId, nil
}
