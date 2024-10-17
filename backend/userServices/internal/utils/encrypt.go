package utils

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// 哈希密码函数
func HashPassword(password string) (string, error) {
	// 使用 bcrypt 生成哈希值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// 返回哈希后的密码
	return string(hashedPassword), nil
}

// ComparePasswordHash 验证密码函数 参数1为明文密码 参数2为hash密码
func ComparePasswordHash(password, hash string) bool {
	// 使用 bcrypt 比对密码与哈希值
	log.Println("password: ", password)
	log.Println("hash: ", hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// DecodeBase64 Base64解码函数
func DecodeBase64(encodedText string) (string, error) {
	// Base64 解码
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedText)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
