package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword 接收明文密码，返回哈希后的密码
func HashPassword(password string) (string, error) {
	saltRounds := 10 // bcrypt 的计算强度
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), saltRounds)
	if err != nil {
		return "", err
	}
	// 返回哈希后的密码
	return string(hashedPassword), nil
}
