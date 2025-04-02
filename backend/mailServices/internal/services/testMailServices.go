package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	pb "mailServices/api/proto"
	"mailServices/internal/client"
	settingsPb "mailServices/internal/client/api/settings/proto"
	"mailServices/internal/smtp"
	"net/http"
)

// 这里是发送的管理员测试邮件

// TestEmailData 测试邮件
type TestEmailData struct {
	Name string
	URL  string
}

// RenderEmailTemplate 渲染 HTML 模板，替换占位符 验证码邮件
func (t *TestEmailData) RenderEmailTemplate(templatePath string) (string, error) {
	log.Println("renderEmailTemplate", t)
	// 解析模板文件
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template file: %v", err)
	}

	// 渲染模板，替换占位符
	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, t); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return rendered.String(), nil
}

func (s *MailServices) SendTestEmail(ctx context.Context, request *pb.SendTestEmailRequest) (*pb.SendTestEmailResponse, error) {
	log.Println("发送测试邮件")
	// 目标收件人
	to := request.Email
	subject := "TestMail"

	// 模板文件路径
	templatePath := s.FilePrefix + "/template/default/TestMail.html"

	var appName, appUrl string
	sendMailTemplateResponse, err := client.GrpcClient.SystemServicesClient.GetSendMailTemplateFillContent(ctx, &settingsPb.GetSendMailTemplateFillContentRequest{})
	if err != nil || sendMailTemplateResponse.Code != http.StatusOK {
		log.Println("fetch app detail err, use default to process.")
		appName = "App"
		appUrl = "https://example.com"
	} else {
		appName = sendMailTemplateResponse.AppName
		appUrl = sendMailTemplateResponse.AppUrl
	}

	// 定义要替换的数据
	emailData := TestEmailData{
		URL:  appUrl,
		Name: appName,
	}

	smtpConfig, err := smtp.GetSMTPConfig()
	if err != nil {
		log.Println("获取发件配置失败", err.Error())
		return &pb.SendTestEmailResponse{
			Code: http.StatusInternalServerError,
			Msg:  "获取发件配置失败" + err.Error(),
		}, nil
	}

	// 渲染模板，替换占位符
	renderedEmail, err := emailData.RenderEmailTemplate(templatePath)
	if err != nil {
		log.Println("渲染模版数据失败", err.Error())
		return &pb.SendTestEmailResponse{
			Code: http.StatusInternalServerError,
			Msg:  "渲染模版数据失败" + err.Error(),
		}, nil
	}

	// 发送邮件
	if err := smtp.SendEmail(to, subject, renderedEmail); err != nil {
		log.Println("发送邮件失败", err.Error())
		return &pb.SendTestEmailResponse{
			Code: http.StatusInternalServerError,
			Msg:  "发送邮件失败" + err.Error(),
		}, nil
	}
	smtpConfigJson, err := json.Marshal(smtpConfig)
	if err != nil {
		return &pb.SendTestEmailResponse{
			Code: http.StatusInternalServerError,
			Msg:  "邮件配置序列化失败" + err.Error(),
		}, nil
	}
	//log.Println(string(smtpConfigJson))
	return &pb.SendTestEmailResponse{
		Code: http.StatusOK,
		Msg:  "发送测试邮件成功",
		Info: smtpConfigJson,
	}, nil
}
