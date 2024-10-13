package model

// VerifyEmailData 定义模板数据结构
type VerifyEmailData struct {
	Code string
	URL  string
	Name string
}

// RemindExpireMail 提示过期邮件
type RemindExpireMail struct {
	Email string `json:"email"`
}
