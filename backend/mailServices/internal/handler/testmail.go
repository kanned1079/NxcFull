package handler

//import (
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//)
//
//func SendTestMail(context *gin.Context) {
//	postForm := &struct {
//		Email string `json:"email"`
//	}{}
//	if err := context.ShouldBindJSON(postForm); err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "获取表单失败",
//			"error": err.Error(),
//		})
//	}
//	log.Println("发送测试邮件")
//	// 目标收件人
//	to := postForm.Email
//	subject := "测试邮件"
//
//	// 模板文件路径
//
//	// pass
//	templatePath := "./template/"
//	// pass
//
//	// 定义要替换的数据
//	emailData := VerifyEmailData{
//		Code: "000111",
//		URL:  "https://example.com",
//		Name: "Nxc Networks",
//	}
//
//	smtpConfig, err := getSMTPConfig()
//	if err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "获取邮件配置失败",
//			"error": err.Error(),
//		})
//	}
//
//	// 渲染模板，替换占位符
//	renderedEmail, err := emailData.renderEmailTemplate(templatePath)
//	if err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "Error rendering email template",
//			"error": err.Error(),
//		})
//		log.Fatalf("Error rendering email template: %v", err)
//	}
//
//	// 发送邮件
//	if err := sendEmail(to, subject, renderedEmail); err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "Error sending email",
//			"error": err.Error(),
//		})
//		log.Fatalf("Error sending email: %v", err)
//	}
//
//	context.JSON(http.StatusOK, gin.H{
//		"code": http.StatusOK,
//		"msg":  "发送测试邮件成功",
//		"info": smtpConfig,
//	})
//}
