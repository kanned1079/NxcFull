package services

//func HandSendEmailCode(to string, code string, templatePath string, templateType string) (err error) {
//	emailData := &struct {
//		Code string `json:"code"`
//		URL  string `json:"url"`
//		Name string `json:"name"`
//	}{
//		Code: code,
//		URL:  "https://example.com",
//		Name: "Example App",
//	}
//	var renderedEmail string
//	// 渲染模板，替换占位符
//	switch templateType {
//	case "verifyMail":
//		{
//			log.Println("验证码邮件")
//			verifyEmail := VerifyEmailData{
//				Code: emailData.Code,
//				URL:  emailData.URL,
//				Name: emailData.Name,
//			}
//			renderedEmail, err = verifyEmail.RenderEmailTemplate(templatePath)
//			if err != nil {
//				return err
//			}
//		}
//	case "notifyEmail":
//		{
//			log.Println("通知邮件")
//		}
//
//	}
//	// 发送邮件
//	if err := smtp.SendEmail(to, "邮箱验证码", renderedEmail); err != nil {
//		return err
//	}
//
//	return nil
//
//}
//
//
