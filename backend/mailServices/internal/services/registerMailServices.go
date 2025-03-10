package services

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	pb "mailServices/api/proto"
	"mailServices/internal/dao"
	"mailServices/internal/smtp"
	"net/http"
	"time"
)

type RegisterMail struct {
	Name string
	Code string
	URL  string
}

// RenderEmailTemplate 渲染注册验证的邮件
func (r *RegisterMail) RenderEmailTemplate(templatePath string) (string, error) {
	log.Println("renderEmailTemplate", r)
	// 解析模板文件
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template file: %v", err)
	}

	// 渲染模板，替换占位符
	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, r); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return rendered.String(), nil
}

// SendRegisterVerifyCode2Email 注册的验证码
func (s *MailServices) SendRegisterVerifyCode2Email(context context.Context, request *pb.SendRegisterVerifyCode2EmailRequest) (*pb.SendRegisterVerifyCode2EmailResponse, error) {
	userMail := request.Email
	generatedCode := RandCode()
	generatedTime := time.Now()
	log.Println("调用时间戳: ", generatedTime.Format("2006-01-02 15:04:05"))

	// 存入 Redis
	log.Println("存入redis")
	err := dao.Rdb.HSet(context, "verify_codes", userMail, generatedCode).Err()
	if err != nil {
		log.Println("存入 Redis 失败：", err)
		return &pb.SendRegisterVerifyCode2EmailResponse{
			Code: http.StatusInternalServerError,
			Msg:  "发送邮件失败" + err.Error(),
		}, nil
	}
	// 设置过期时间
	err = dao.Rdb.Expire(context, "verify_codes", time.Minute*5).Err()
	if err != nil {
		log.Println("设置过期时间失败：", err)
	}
	// 发送邮件
	//log.Println("发送邮件", vEmail)
	var regMail RegisterMail = RegisterMail{
		Name: request.Email,
		Code: generatedCode,
		URL:  "http://example.com",
	}

	// 渲染模板，替换占位符
	renderedEmail, err := regMail.RenderEmailTemplate(s.FilePrefix + "/template/default/VerifyCode.html")
	if err != nil {
		log.Println("渲染模版数据失败", err.Error())
		return &pb.SendRegisterVerifyCode2EmailResponse{
			Code: http.StatusInternalServerError,
			Msg:  "渲染模版数据失败" + err.Error(),
		}, nil
	}

	// 发送邮件
	if err := smtp.SendEmail(userMail, "VerifyCodeMail", renderedEmail); err != nil {
		log.Println("发送邮件失败", err.Error())
		return &pb.SendRegisterVerifyCode2EmailResponse{
			Code: http.StatusInternalServerError,
			Msg:  "发送邮件失败" + err.Error(),
		}, nil
	}

	//err = HandSendEmailCode(request.Email, vEmail.VerifyCode, "../sendmail/template/default/verifyMail.html", "verifyMail")
	//if err != nil {
	//	log.Println(err)
	//	return &pb.SendRegisterVerifyCode2EmailResponse{
	//		Code: http.StatusInternalServerError,
	//		Msg:  "发送邮件失败" + err.Error(),
	//	}, nil
	//}
	// pass
	return &pb.SendRegisterVerifyCode2EmailResponse{
		Code: http.StatusOK,
		Msg:  "邮件已发送",
	}, nil
}
