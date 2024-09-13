package routers

import (
	"NxcFull/nxc-backend/dao"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"net/http"
	"time"
)

var smtpConfig SMTPConfig

// SMTP 配置信息
type SMTPConfig struct {
	Host         string `json:"host"`     // SMTP 服务器地址
	Port         int    `json:"port"`     // SMTP 服务器端口（25、465、587）
	Username     string `json:"username"` // SMTP 用户名
	Password     string `json:"password"` // SMTP 密码
	From         string `json:"from"`     // 发件人地址
	UseSSL       bool   `json:"use_ssl"`  // 是否使用 SSL 加密
	UseTLS       bool   `json:"use_tls"`  // 是否使用 TLS 加密
	MailTemplate string `json:"mail_template"`
}

// 定义模板数据结构
type VerifyEmailData struct {
	Code string
	URL  string
	Name string
}

type RemindExpireMail struct {
	Email string `json:"email"`
}

// 渲染 HTML 模板，替换占位符
func (v *VerifyEmailData) renderEmailTemplate(templatePath string) (string, error) {
	log.Println("renderEmailTemplate", v)
	// 解析模板文件
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template file: %v", err)
	}

	// 渲染模板，替换占位符
	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, v); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return rendered.String(), nil
}

// 从数据库获取所有相关的邮件配置项
func getConfigFromDB() (SMTPConfig, error) {
	log.Println("从数据库中读取邮件配置")
	smtpConfig := SMTPConfig{}

	// 从数据库逐个读取配置
	// 邮件主机
	host, err := readSettingFromDB("sendmail", "email_host")
	if err != nil {
		return SMTPConfig{}, err
	}
	json.Unmarshal(host, &smtpConfig.Host)

	// 邮件端口
	port, err := readSettingFromDB("sendmail", "email_port")
	if err != nil {
		return SMTPConfig{}, err
	}
	json.Unmarshal(port, &smtpConfig.Port)

	// 邮件用户名
	username, err := readSettingFromDB("sendmail", "email_username")
	if err != nil {
		return SMTPConfig{}, err
	}
	json.Unmarshal(username, &smtpConfig.Username)

	// 邮件密码
	password, err := readSettingFromDB("sendmail", "email_password")
	if err != nil {
		return SMTPConfig{}, err
	}
	json.Unmarshal(password, &smtpConfig.Password)

	// 发件人地址
	from, err := readSettingFromDB("sendmail", "email_from_address")
	if err != nil {
		return SMTPConfig{}, err
	}
	json.Unmarshal(from, &smtpConfig.From)

	// 邮件模版
	mailTemplate, err := readSettingFromDB("sendmail", "email_template")
	if err != nil {
		return SMTPConfig{}, err
	}
	json.Unmarshal(mailTemplate, &smtpConfig.MailTemplate)

	// 加密方式 (SSL / TLS)
	encryption, err := readSettingFromDB("sendmail", "email_encryption")
	if err != nil {
		return SMTPConfig{}, err
	}
	var encryptionMethod string
	json.Unmarshal(encryption, &encryptionMethod)

	if encryptionMethod == "SSL" {
		smtpConfig.UseSSL = true
	} else if encryptionMethod == "TLS" {
		smtpConfig.UseTLS = true
	}

	return smtpConfig, nil
}

// 从 Redis 缓存中获取邮件配置，如果 Redis 为空则从数据库获取并缓存
func getSMTPConfig() (SMTPConfig, error) {
	log.Println("从Redis中直接读取邮件配置")
	ctx := context.Background()

	// 尝试从 Redis 获取缓存的 SMTP 配置
	result, err := dao.Rdb.Get(ctx, "smtpConfig").Result()
	if err == redis.Nil {
		// 如果 Redis 中没有，读取数据库并缓存
		smtpConfig, err := getConfigFromDB()
		if err != nil {
			return SMTPConfig{}, err
		}

		// 将配置转换为 JSON 并缓存到 Redis
		smtpConfigJSON, err := json.Marshal(smtpConfig)
		if err != nil {
			return SMTPConfig{}, err
		}

		err = dao.Rdb.Set(ctx, "smtpConfig", smtpConfigJSON, time.Hour).Err() // 设置过期时间为 1 小时
		if err != nil {
			return SMTPConfig{}, err
		}

		return smtpConfig, nil
	} else if err != nil {
		return SMTPConfig{}, err
	}

	// 从 Redis 缓存中解析 SMTP 配置
	var smtpConfig SMTPConfig
	err = json.Unmarshal([]byte(result), &smtpConfig)
	if err != nil {
		return SMTPConfig{}, err
	}

	return smtpConfig, nil
}

//func InitMailConfig() {
//	var kv map[string]string
//}

func HandSendEmailCode(to string, code string, templatePath string, templateType string) (err error) {
	emailData := &struct {
		Code string `json:"code"`
		URL  string `json:"url"`
		Name string `json:"name"`
	}{
		Code: code,
		URL:  "https://example.com",
		Name: "Example App",
	}
	//smtpConfig, err = getSMTPConfig()
	//if err != nil {
	//	return err
	//}
	var renderedEmail string
	// 渲染模板，替换占位符
	switch templateType {
	case "verifyMail":
		{
			log.Println("验证码邮件")
			verifyEmail := VerifyEmailData{
				Code: emailData.Code,
				URL:  emailData.URL,
				Name: emailData.Name,
			}
			renderedEmail, err = verifyEmail.renderEmailTemplate(templatePath)
			if err != nil {
				return err
			}
		}
	case "notifyEmail":
		{
			log.Println("通知邮件")
		}

	}
	//renderedEmail, err := renderEmailTemplate(templatePath, emailData)
	//if err != nil {
	//	return err
	//}

	// 发送邮件
	if err := sendEmail(to, "邮箱验证码", renderedEmail); err != nil {
		return err
	}

	return nil

}

func HandleSendTestMail(context *gin.Context) {
	postForm := &struct {
		Email string `json:"email"`
	}{}
	if err := context.ShouldBindJSON(postForm); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "获取表单失败",
			"error": err.Error(),
		})
	}
	log.Println("发送测试邮件")
	// 目标收件人
	to := postForm.Email
	subject := "测试邮件"

	// 模板文件路径
	templatePath := "../sendmail/template/default/verifyMail.html"

	// 定义要替换的数据
	emailData := VerifyEmailData{
		Code: "000111",
		URL:  "https://example.com",
		Name: "Example App",
	}

	smtpConfig, err := getSMTPConfig()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "获取邮件配置失败",
			"error": err.Error(),
		})
	}

	// 渲染模板，替换占位符
	renderedEmail, err := emailData.renderEmailTemplate(templatePath)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "Error rendering email template",
			"error": err.Error(),
		})
		log.Fatalf("Error rendering email template: %v", err)
	}

	// 发送邮件
	if err := sendEmail(to, subject, renderedEmail); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "Error sending email",
			"error": err.Error(),
		})
		log.Fatalf("Error sending email: %v", err)
	}

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "发送测试邮件成功",
		"info": smtpConfig,
	})
}

// 发送邮件
func sendEmail(to string, subject string, body string) error {
	// 获取 SMTP 配置
	config, err := getSMTPConfig()
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
