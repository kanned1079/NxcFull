package smtp

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	//"mailServices/internal/model"
)

//var smtpConfig model.SMTPConfig

// SendEmail 发送邮件序列化字符串后的数据
func SendEmail(to string, subject string, body string) error {
	// 获取 SMTP 配置
	config, err := GetSMTPConfig()
	if err != nil {
		return fmt.Errorf("failed to get SMTP config: %v", err)
	}

	// 创建邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", config.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body) // 设置为 HTML 格式

	// 创建 SMTP 拨号器
	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)

	// 如果选择使用 SSL
	if config.UseSSL {
		d.SSL = true
	}

	// 如果选择使用 TLS
	if config.UseTLS {
		d.SSL = false
		d.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	log.Println("Email sent successfully!")
	return nil
}
