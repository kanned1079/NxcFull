package model

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

//
//// VerifyEmailData 定义模板数据结构
//type VerifyEmailData struct {
//	Code string
//	URL  string
//	Name string
//}
//
//// RemindExpireMail 提示过期邮件
//type RemindExpireMail struct {
//	Email string `json:"email"`
//}
