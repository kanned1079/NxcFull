package routers

import (
	"NxcFull/nxc-backend/auth"
	"NxcFull/nxc-backend/dao"
	"NxcFull/nxc-backend/system"
	"NxcFull/nxc-backend/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LoginMsg struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handleUserLogin(context *gin.Context) {
	var req LoginMsg
	if err := context.ShouldBind(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//log.Println(req)
	//log.Println("输入 ", req.Email, req.Password)
	if user.IsUserExist(req.Email) == user.User_Exist {
		//log.Println("验证密码")
		// 验证用户密码
		if user.AuthUserInfo(req.Email, req.Password) == user.Auth_Pass {
			var thisUser user.User
			if result := dao.Db.Model(&user.User{}).Where("email = ?", req.Email).First(&thisUser); result.Error != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
					"code":     http.StatusInternalServerError,
					"isAuthed": false,
					"msg":      result.Error.Error(),
				})
			}
			token, err := auth.GenerateToken(req.Email)
			log.Println("用户Token: ", token)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
					"code":     http.StatusInternalServerError,
					"isAuthed": false,
					"msg":      err.Error(),
				})
			}
			context.JSON(http.StatusOK, gin.H{
				"code":      http.StatusOK,
				"isAuthed":  true,
				"msg":       "验证通过",
				"token":     token,
				"user_data": thisUser,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"code":     http.StatusOK,
				"isAuthed": false,
				"msg":      "incorrect_password",
			})
		}
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":     http.StatusOK,
			"isAuthed": false,
			"msg":      "user_not_exist",
		})
	}
}

func handleGetServerInfo(context *gin.Context) {
	var systemOverLook system.OsInfo
	systemOverLook.GetOsInfo()
	//log.Println("系统详情", systemOverLook)

	context.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"osInfo": systemOverLook,
	})
}

func handleGetUserList(context *gin.Context) {
	var users []user.User
	if result := dao.Db.Model(&user.User{}).Find(&users); result.Error != nil {
		log.Println("读取用户列表错误")
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "读取用户列表失败",
			"err":  result.Error.Error(),
			//"users": users,
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  "ok",
		//"err":  result.Error.Error(),
		"users": users,
	})

}

func getSettingSit(context *gin.Context) {

}

func responseLoginMsg(context *gin.Context, status any, authed bool, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code":    status,
		"authed":  authed,
		"message": msg,
	})
}
