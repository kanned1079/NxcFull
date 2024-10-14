package handler

import (
	"NxcFull/nxc-backend/dao"
	sysContext "context"
	"encoding/json"
	pb "gateway/internal/grpc/api/sendmail/proto"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// HandleSendTestMail 管理员测试邮件配置 同步操作
func HandleSendTestMail(context *gin.Context) {
	postData := &struct {
		Email string `json:"email"`
	}{}
	if err := context.ShouldBindJSON(postData); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}
	resp, err := grpcClient.MailServicesClient.SendTestEmail(sysContext.Background(), &pb.SendTestEmailRequest{
		Email: postData.Email,
	})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//log.Println(string(resp.Info))
	var smtpConfig map[string]any
	if err := json.Unmarshal(resp.Info, &smtpConfig); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	log.Println(smtpConfig)
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
		"info": smtpConfig,
	})
}

// HandleSendVerifyEmail 用于生成验证码并发送邮件到用户邮箱
func HandleSendVerifyEmail(context *gin.Context) {
	postForm := &struct {
		Email string `json:"email"`
	}{}
	if err := context.ShouldBind(&postForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"msg":   "解析邮件地址失败",
			"error": err.Error(),
		})
	}
	resp, err := grpcClient.MailServicesClient.SendVerifyEmail()
	//var vEmail VerifyEmail = VerifyEmail{
	//	Email:      postForm.Email,
	//	CreatedAt:  time.Now(), // 创建日期
	//	VerifyCode: RandCode(), // 生成的验证码
	//}
	//// 存入 Redis
	//log.Println("存入redis", vEmail)
	//err := dao.Rdb.HSet(context, "verify_codes", vEmail.Email, vEmail.VerifyCode).Err()
	//if err != nil {
	//	log.Println("存入 Redis 失败：", err)
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"code":  http.StatusInternalServerError,
	//		"msg":   "发送邮件失败",
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//// 设置过期时间
	//err = dao.Rdb.Expire(context, "verify_codes", time.Minute*5).Err()
	//if err != nil {
	//	log.Println("设置过期时间失败：", err)
	//}
	//// 发送邮件
	//log.Println("发送邮件", vEmail)
	//err = HandSendEmailCode(postForm.Email, vEmail.VerifyCode, "../sendmail/template/default/verifyMail.html", "verifyMail")
	//if err != nil {
	//	log.Println(err)
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"code":  http.StatusInternalServerError,
	//		"msg":   "邮件发送失败",
	//		"error": err.Error(),
	//	})
	//}
	//// pass
	//
	//context.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "邮件已发送",
	//})
}

// HandleVerifyEmailCode 用于从redis中取出与用户邮件地址对应的验证码来验证
func HandleVerifyEmailCode(context *gin.Context) {
	postForm := &struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}{}
	if err := context.ShouldBind(&postForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"msg":   "解析验证信息失败",
			"error": err.Error(),
		})
		return
	}
	log.Println("验证邮箱", postForm)
	storedCode, err := dao.Rdb.HGet(context, "verify_codes", postForm.Email).Result()
	if err == redis.Nil {
		context.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "未找到对应验证码",
		})
		return
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "验证失败",
			"error": err.Error(),
		})
		return
	}
	log.Println("redis的验证码", storedCode)
	log.Println("用户输入的", postForm.Code)
	if storedCode != postForm.Code {
		log.Println("验证码不一致")
		context.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "验证码错误",
		})
		return
	}
	log.Println("验证码一致")
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "验证码正确",
	})
}

// RandCode 生成验证码
func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
