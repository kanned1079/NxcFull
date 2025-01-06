package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"mailServices/internal/dao"
	"mailServices/internal/model"
	"mailServices/internal/smtp"
	"math/rand"
	"sync"
	"time"
)

const CodeLength = 6

//// 渲染 HTML 模板，替换占位符
//func (v *VerifyEmailData) RenderEmailTemplate(templatePath string) (string, error) {
//	log.Println("renderEmailTemplate", v)
//	// 解析模板文件
//	tmpl, err := template.ParseFiles(templatePath)
//	if err != nil {
//		return "", fmt.Errorf("failed to parse template file: %v", err)
//	}
//
//	// 渲染模板，替换占位符
//	var rendered bytes.Buffer
//	if err := tmpl.Execute(&rendered, v); err != nil {
//		return "", fmt.Errorf("failed to execute template: %v", err)
//	}
//
//	return rendered.String(), nil
//}

type SendMailConfigCache struct {
	Config     *sync.Map
	LastUpdate *time.Time
}

// GetConfigFromDB 从数据库获取所有相关的邮件配置项
func GetConfigFromDB() (model.SMTPConfig, error) {
	log.Println("从数据库中读取邮件配置")
	smtpConfig := model.SMTPConfig{}

	// 从数据库逐个读取配置
	// 邮件主机
	host, err := smtp.ReadSettingFromDB("sendmail", "email_host")
	if err != nil {
		return model.SMTPConfig{}, err
	}
	if err := json.Unmarshal(host, &smtpConfig.Host); err != nil {
		log.Println(err)
	}

	// 邮件端口
	port, err := smtp.ReadSettingFromDB("sendmail", "email_port")
	if err != nil {
		return model.SMTPConfig{}, err
	}
	if err := json.Unmarshal(port, &smtpConfig.Port); err != nil {
		log.Println(err)
	}

	// 邮件用户名
	username, err := smtp.ReadSettingFromDB("sendmail", "email_username")
	if err != nil {
		return model.SMTPConfig{}, err
	}
	if err := json.Unmarshal(username, &smtpConfig.Username); err != nil {
		log.Println(err)
	}

	// 邮件密码
	password, err := smtp.ReadSettingFromDB("sendmail", "email_password")
	if err != nil {
		return model.SMTPConfig{}, err
	}
	if err := json.Unmarshal(password, &smtpConfig.Password); err != nil {
		log.Println(err)
	}

	// 发件人地址
	from, err := smtp.ReadSettingFromDB("sendmail", "email_from_address")
	if err != nil {
		return model.SMTPConfig{}, err
	}
	if err := json.Unmarshal(from, &smtpConfig.From); err != nil {
		log.Println(err)
	}

	// 邮件模版
	mailTemplate, err := smtp.ReadSettingFromDB("sendmail", "email_template")
	if err != nil {
		return model.SMTPConfig{}, err
	}
	if err := json.Unmarshal(mailTemplate, &smtpConfig.MailTemplate); err != nil {
		log.Println(err)
	}

	// 加密方式 (SSL / TLS)
	encryption, err := smtp.ReadSettingFromDB("sendmail", "email_encryption")
	if err != nil {
		return model.SMTPConfig{}, err
	}
	var encryptionMethod string
	if err := json.Unmarshal(encryption, &encryptionMethod); err != nil {
		log.Println(err)
	}

	if encryptionMethod == "SSL" {
		smtpConfig.UseSSL = true
	} else if encryptionMethod == "TLS" {
		smtpConfig.UseTLS = true
	}

	return smtpConfig, nil
}

// GetSMTPConfig 从 Redis 缓存中获取邮件配置，如果 Redis 为空则从数据库获取并缓存
func GetSMTPConfig() (model.SMTPConfig, error) {
	log.Println("从Redis中直接读取邮件配置")
	ctx := context.Background()

	// 尝试从 Redis 获取缓存的 SMTP 配置
	result, err := dao.Rdb.Get(ctx, "smtpConfig").Result()
	if errors.Is(err, redis.Nil) {
		// 如果 Redis 中没有，读取数据库并缓存
		smtpConfig, err := GetConfigFromDB()
		if err != nil {
			return model.SMTPConfig{}, err
		}

		// 将配置转换为 JSON 并缓存到 Redis
		smtpConfigJSON, err := json.Marshal(smtpConfig)
		if err != nil {
			return model.SMTPConfig{}, err
		}

		err = dao.Rdb.Set(ctx, "smtpConfig", smtpConfigJSON, time.Hour).Err() // 设置过期时间为 1 小时
		if err != nil {
			return model.SMTPConfig{}, err
		}

		return smtpConfig, nil
	} else if err != nil {
		return model.SMTPConfig{}, err
	}

	// 从 Redis 缓存中解析 SMTP 配置
	var smtpConfig model.SMTPConfig
	err = json.Unmarshal([]byte(result), &smtpConfig)
	if err != nil {
		return model.SMTPConfig{}, err
	}

	return smtpConfig, nil
}

// RandCode 生成验证码
func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
