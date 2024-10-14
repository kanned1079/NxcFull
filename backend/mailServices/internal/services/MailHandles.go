package services

import (
	"log"
	"mailServices/internal/smtp"
)

//var smtpConfig SMTPConfig
//
//// SMTP 配置信息
//type SMTPConfig struct {
//	Host         string `json:"host"`     // SMTP 服务器地址
//	Port         int    `json:"port"`     // SMTP 服务器端口（25、465、587）
//	Username     string `json:"username"` // SMTP 用户名
//	Password     string `json:"password"` // SMTP 密码
//	From         string `json:"from"`     // 发件人地址
//	UseSSL       bool   `json:"use_ssl"`  // 是否使用 SSL 加密
//	UseTLS       bool   `json:"use_tls"`  // 是否使用 TLS 加密
//	MailTemplate string `json:"mail_template"`
//}

//func InitMailConfig() {
//	var kv map[string]string
//}

// gin
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
			renderedEmail, err = verifyEmail.RenderEmailTemplate(templatePath)
			if err != nil {
				return err
			}
		}
	case "notifyEmail":
		{
			log.Println("通知邮件")
		}

	}
	// 发送邮件
	if err := smtp.SendEmail(to, "邮箱验证码", renderedEmail); err != nil {
		return err
	}

	return nil

}
