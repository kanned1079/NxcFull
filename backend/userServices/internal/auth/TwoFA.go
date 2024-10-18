package auth

import (
	"github.com/pquerna/otp/totp"
)

// Verify2FACode 验证用户输入的 TOTP 验证码
func Verify2FACode(userSecret, userOTP string) bool {
	//log.Println(time.Now().Format("2006-01-02 15:04:05"), userSecret, userOTP)
	//log.Println(totp.Validate(userOTP, userSecret))
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
