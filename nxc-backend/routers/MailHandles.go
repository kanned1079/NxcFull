package routers

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"log"
	"net/http"
	"os"
)

// SMTP 配置信息
type SMTPConfig struct {
	Host     string // SMTP 服务器地址
	Port     int    // SMTP 服务器端口（25、465、587）
	Username string // SMTP 用户名
	Password string // SMTP 密码
	From     string // 发件人地址
	UseSSL   bool   // 是否使用 SSL 加密
	UseTLS   bool   // 是否使用 TLS 加密
}

func HandleSendTestMail(context *gin.Context) {
	log.Println("发送测试邮件")

	smtpConfig := SMTPConfig{
		Host:     "smtp.exmail.qq.com",   // SMTP 服务器地址
		Port:     465,                    // SMTP 端口，465 通常是 SSL
		Username: "no-reply@ikanned.com", // SMTP 账号
		Password: "7gCepaWEcg2YuVYD",     // SMTP 密码
		From:     "no-reply@ikanned.com", // 发件人邮箱
		UseSSL:   true,                   // 是否使用 SSL
	}

	// 目标收件人
	to := "kanned1079@gmail.com"
	subject := "Test Email"
	htmlFilePath := "./email_template.html" // 静态 HTML 文件路径

	// 发送邮件
	if err := sendEmail(smtpConfig, to, subject, htmlFilePath); err != nil {
		fmt.Printf("Error sending email: %v\n", err)
	}

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func sendEmail(config SMTPConfig, to string, subject string, htmlFilePath string) error {
	// 读取静态 HTML 文件
	htmlContent, err := os.ReadFile(htmlFilePath) // 使用 os.ReadFile 代替 ioutil.ReadFile
	if err != nil {
		return fmt.Errorf("failed to read HTML file: %v", err)
	}

	// 创建邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", config.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", string(htmlContent))

	// 创建 SMTP 拨号器
	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)

	// 如果选择使用 SSL
	if config.UseSSL {
		d.SSL = true
	}

	// 如果选择使用 TLS
	if config.UseTLS {
		d.SSL = false // TLS 需要先通过非加密连接，然后再升级到加密连接
		d.TLSConfig = &tls.Config{
			InsecureSkipVerify: true, // 允许跳过证书验证（根据实际情况配置）
		}
	}

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully!")
	return nil
}
