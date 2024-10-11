package model

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"userServices/internal/dao"
)

// IsUserExist 指定邮箱的用户是否存在
func IsUserExist(email string) (code int) {
	user := Auth{
		Email: email,
	}
	result := dao.Db.Model(&Auth{}).Where("email = ?", user.Email).First(&user)
	if result.RowsAffected > 0 {
		code = User_Exist
	} else {
		code = User_Not_Exist
	}
	return
}

// AuthUserInfo 校验用户凭据
func AuthUserInfo(email string, password string) (code int) {
	var userAuth Auth
	result := dao.Db.Model(&Auth{}).Where("email = ?", email).First(&userAuth)
	if result.Error != nil {
		log.Println("未知错误")
	}
	log.Println("查询到的用户凭证", userAuth)
	if password == userAuth.Password {
		code = Auth_Pass
	} else {
		code = Auth_Failure
	}
	return
}

// HandleCheckOldPassword 验证旧密码
func HandleCheckOldPassword(context *gin.Context) {
	// 绑定 JSON 请求体中的数据
	postData := &struct {
		Email       string `json:"email"`
		OldPassword string `json:"old_password"` // 已经哈希过的密码
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "绑定数据失败",
			"error": err.Error(),
		})
		return
	}

	log.Println(postData)

	// 查找用户
	var auth Auth
	if err := dao.Db.Model(&Auth{}).Where("email = ?", postData.Email).First(&auth).Error; err != nil {
		// 如果用户不存在
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "用户不存在",
		})
		return
	}

	// 直接比较哈希过的旧密码
	if postData.OldPassword != auth.Password {
		// 如果哈希后的密码不匹配
		log.Println("密码不匹配")
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "旧密码不正确",
		})
		return
	}

	// 如果密码验证通过
	context.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"verified": true,
		"msg":      "旧密码验证通过",
	})
}

// HandleApplyNewPassword 更新用户密码
func HandleApplyNewPassword(context *gin.Context) {
	// 绑定 JSON 请求体中的数据
	postData := &struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password"` // 新密码已经哈希过
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "绑定数据失败",
			"error": err.Error(),
		})
		return
	}

	log.Println(postData)

	// 查找用户
	var auth Auth
	if err := dao.Db.Model(&Auth{}).Where("email = ?", postData.Email).First(&auth).Error; err != nil {
		// 如果用户不存在
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "用户不存在",
		})
		return
	}

	// 更新密码，直接存储哈希过的新密码
	auth.Password = postData.NewPassword
	if err := dao.Db.Save(&auth).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "更新密码失败",
			"error": err.Error(),
		})
		return
	}

	log.Println("修改密码成功")
	// 密码更新成功
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"updated": true,
		"msg":     "密码更新成功",
	})
}
