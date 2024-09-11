package routers

import (
	"NxcFull/nxc-backend/dao"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const CodeLength = 6

type VerifyEmail struct {
	Email      string    `json:"email"`       // 用户邮箱
	VerifyCode string    `json:"verify_code"` // 生成的验证码
	CreatedAt  time.Time `json:"created_at"`  // 验证码创建时间戳
}

// SendVerifyEmail 用于生成验证码并发送邮件到用户邮箱
func SendVerifyEmail(context *gin.Context) {
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
	var vEmail VerifyEmail = VerifyEmail{
		Email:      postForm.Email,
		CreatedAt:  time.Now(), // 创建日期
		VerifyCode: RandCode(), // 生成的验证码
	}
	// 存入 Redis
	log.Println("存入redis", vEmail)
	err := dao.Rdb.HSet(context, "verify_codes", vEmail.Email, vEmail.VerifyCode).Err()
	if err != nil {
		log.Println("存入 Redis 失败：", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "发送邮件失败",
			"error": err.Error(),
		})
		return
	}
	// 设置过期时间
	err = dao.Rdb.Expire(context, "verify_codes", time.Minute*5).Err()
	if err != nil {
		log.Println("设置过期时间失败：", err)
	}
	// 发送邮件
	log.Println("发送邮件", vEmail)
	// pass

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "邮件已发送",
	})
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

func HandleUserRegister(context *gin.Context) {
	postForm := &struct {
		Email        string `json:"email"`
		Password     string `json:"password"`
		InviteUserId string `json:"invite_user_id"`
	}{}
	if err := context.ShouldBind(&postForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"msg":   "数据绑定错误",
			"error": err.Error(),
		})
	}
	log.Println("最终注册", postForm)
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
