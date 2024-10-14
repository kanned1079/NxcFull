package services

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	pb "mailServices/api/proto"
	"mailServices/internal/dao"
	"net/http"
)

func (s *MailServices) VerifyEmailCode(context context.Context, request *pb.VerifyEmailCodeRequest) (*pb.VerifyEmailCodeResponse, error) {
	log.Println("验证邮箱", request.Email)
	storedCode, err := dao.Rdb.HGet(context, "verify_codes", request.Email).Result()
	if errors.Is(err, redis.Nil) {
		return &pb.VerifyEmailCodeResponse{
			Code:   http.StatusOK,
			Msg:    "验证码可能过期 请重新获取",
			Passed: false,
		}, nil
	} else if err != nil {
		return &pb.VerifyEmailCodeResponse{
			Code:   http.StatusOK,
			Msg:    "未知错误 请重试",
			Passed: false,
		}, nil
	}
	log.Println("redis的验证码", storedCode)
	log.Println("用户输入的", request.Code)
	if storedCode != request.Code {
		log.Println("验证码不一致")
		//context.JSON(http.StatusUnauthorized, gin.H{
		//	"code": http.StatusUnauthorized,
		//	"msg":  "验证码错误",
		//})
		return &pb.VerifyEmailCodeResponse{
			Code:   http.StatusOK,
			Msg:    "验证码错误",
			Passed: false,
		}, nil
	}
	log.Println("验证码一致")
	//context.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "验证码正确",
	//})
	return &pb.VerifyEmailCodeResponse{
		Code:   http.StatusOK,
		Msg:    "验证码正确验证通过",
		Passed: true,
	}, nil
}

// HandleVerifyEmailCode 用于从redis中取出与用户邮件地址对应的验证码来验证
//func HandleVerifyEmailCode(context *gin.Context) {
//	postForm := &struct {
//		Email string `json:"email"`
//		Code  string `json:"code"`
//	}{}
//	if err := context.ShouldBind(&postForm); err != nil {
//		context.JSON(http.StatusBadRequest, gin.H{
//			"code":  http.StatusBadRequest,
//			"msg":   "解析验证信息失败",
//			"error": err.Error(),
//		})
//		return
//	}
//	log.Println("验证邮箱", postForm)
//	storedCode, err := dao.Rdb.HGet(context, "verify_codes", postForm.Email).Result()
//	if errors.Is(err, redis.Nil) {
//		context.JSON(http.StatusNotFound, gin.H{
//			"code": http.StatusNotFound,
//			"msg":  "未找到对应验证码",
//		})
//		return
//	} else if err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "验证失败",
//			"error": err.Error(),
//		})
//		return
//	}
//	log.Println("redis的验证码", storedCode)
//	log.Println("用户输入的", postForm.Code)
//	if storedCode != postForm.Code {
//		log.Println("验证码不一致")
//		context.JSON(http.StatusUnauthorized, gin.H{
//			"code": http.StatusUnauthorized,
//			"msg":  "验证码错误",
//		})
//		return
//	}
//	log.Println("验证码一致")
//	context.JSON(http.StatusOK, gin.H{
//		"code": http.StatusOK,
//		"msg":  "验证码正确",
//	})
//}

// RandCode 生成验证码
//func RandCode() string {
//	s := "1234567890"
//	code := ""
//	rand.Seed(time.Now().UnixNano())
//	for i := 0; i < CodeLength; i++ {
//		code += string(s[rand.Intn(len(s))])
//	}
//	return code
//}
//
//// VerifyEmailData 验证码邮件的替换数据
//type VerifyEmailData struct {
//	Code string
//	URL  string
//	Name string
//}
//
//// RenderEmailTemplate 渲染 HTML 模板，替换占位符 验证码邮件
//func (v *VerifyEmailData) RenderEmailTemplate(templatePath string) (string, error) {
//	log.Println("renderEmailTemplate", v)
//	// 解析模板文件
//	tmpl, err := template.ParseFiles(templatePath)
//	if err != nil {
//		return "", fmt.Errorf("failed to parse template file: %v", err)
//	}
//
//	// 渲染模板，替换占位符
//	var rendered bytes.Buffer
//	if err := tmpl.Execute(&rendered, v); err != nil {
//		return "", fmt.Errorf("failed to execute template: %v", err)
//	}
//
//	return rendered.String(), nil
//}
