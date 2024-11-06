package routers

import (
	"NxcFull/nxc-backend/dao"
	"NxcFull/nxc-backend/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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
	var newUserInfo user.User = user.User{
		Email:        postForm.Email,
		InviteUserID: postForm.InviteUserId,
		//LicenseExpiration: nil,
	}
	var newUserAuth user.Auth = user.Auth{
		Email:    postForm.Email,
		Password: postForm.Password,
	}
	result1 := dao.Db.Model(&user.User{}).Create(&newUserInfo)
	if result1.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "插入数据错误",
			"error": result1.Error.Error(),
		})
	}
	result2 := dao.Db.Model(&user.Auth{}).Create(&newUserAuth)
	if result2.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "插入验证数据错误",
			"error": result2.Error.Error(),
		})
	}
	if result1.Error == nil && result2.Error == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "新建用户成功",
		})
	}
}
