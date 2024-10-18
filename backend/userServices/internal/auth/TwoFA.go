package auth

import (
	"context"
	"errors"
	"github.com/pquerna/otp/totp"
	"log"
	"time"
	"userServices/internal/dao"
	"userServices/internal/model"
)

// Verify2FACode 验证用户输入的 TOTP 验证码
func Verify2FACode(userSecret, userOTP string) bool {
	return totp.Validate(userOTP, userSecret) // 验证用户输入的验证码
}

// Generate2FASecret 生成 TOTP 秘钥和二维码 URL
func Generate2FASecret(appName, username string) (secret string, url string, err error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      appName,  // 显示在验证界面的网站或应用名
		AccountName: username, // 显示在验证界面的用户名
	})
	if err != nil {
		return "", "", err
	}
	return key.Secret(), key.URL(), nil // 返回秘钥和二维码 URL
}

func GetSecretAndVerify2FACode(email, code string) (bool, error) {
	log.Println(email, code)
	if code == "" {
		return false, errors.New("请提供2FA代码")
	}
	// 从 Redis 中获取 2FA 信息
	key := "2fa:" + email
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := dao.Rdb.HMGet(ctx, key, "email", "secret").Result()
	if err != nil {
		return false, err
	}

	// 如果 Redis 中有数据
	if len(result) > 0 && result[1] != nil {
		secret := result[1].(string)
		// 验证 2FA code
		if Verify2FACode(secret, code) {
			// 重新设置有效期为一天
			ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			err = dao.Rdb.Expire(ctx, key, 24*time.Hour).Err()
			if err != nil {
				return false, errors.New("设置过期时间错误" + err.Error())
			}
			return true, nil
		}
		return false, nil
	}

	// 如果 Redis 中没有数据，去数据库中查询
	var secret string
	if err := dao.Db.Model(&model.TwoFA{}).
		Where("email = ?", email).
		Select("two_fa_key").Scan(&secret).Error; err != nil {
		return false, errors.New("在数据库中查询密钥失败" + err.Error())
	}

	// 验证 2FA code
	if !Verify2FACode(secret, code) {
		return false, nil
	}

	// 将密钥存入 Redis，并设置有效期为一天
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = dao.Rdb.HMSet(ctx, key, "email", email, "secret", secret).Err()
	if err != nil {
		return false, errors.New("设置Redis键值失败" + err.Error())
	}
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = dao.Rdb.Expire(ctx, key, 24*time.Hour).Err()
	if err != nil {
		return false, errors.New("设置过期时间错误" + err.Error())
	}

	return true, nil
}
