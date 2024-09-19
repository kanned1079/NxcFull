package routers

import (
	"NxcFull/nxc-backend/Document"
	"NxcFull/nxc-backend/auth"
	"NxcFull/nxc-backend/subscribePlan"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartAdminReq() {
	r := gin.Default()

	// 中间件 过OPTIONS预检
	r.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
			return
		}
		context.Next()
	})

	publicRoutes := r.Group("/api")

	// 这里是获取网站的标题 颜色 说明 配置
	publicRoutes.GET("/admin/getStartTheme", getStartTheme)

	publicRoutes.GET("/app/preference/")

	adminPublic := publicRoutes.Group("/admin/v1")
	{
		// 管理员的公共路由
		adminPublic.POST("/login", handleUserLogin)
	}

	userPublic := publicRoutes.Group("/user/v1")
	{
		// 一般用户登录和注册相关路由
		userPublic.POST("/login", handleUserLogin)
		userPublic.POST("/register/vc/get", SendVerifyEmail)
		userPublic.POST("/register/vc/verify", HandleVerifyEmailCode)
		userPublic.POST("/register/register", HandleUserRegister)
	}

	protectedRoutes := r.Group("/api")
	protectedRoutes.Use(auth.AuthMiddleware())
	protectedRoutes.Use(auth.RoleMiddleware())

	adminAuthorized := protectedRoutes.Group("/admin/v1", auth.RoleMiddleware())
	{
		// 管理员特定的路由
		adminAuthorized.GET("/os/status/get", handleGetServerInfo) // 获取当前系统的负载
		// adminAuthorized.POST("admin/save-setting", handleUpdateSystemSettings)
		adminAuthorized.POST("/settings/set", handleUpdateSingleOptions)     // 设置单个配置项目
		adminAuthorized.GET("/settings/get", handleGetSystemSetting)         // 获取所有的系统设置
		adminAuthorized.GET("/users/get", handleGetUserList)                 // 获取所有用户的列表
		adminAuthorized.GET("/notices/get", HandleGetAllNotices)             // 获取所有通知列表
		adminAuthorized.POST("/notice/add", HandleAddNotice)                 // 添加一条新的通知
		adminAuthorized.POST("/mail/test", HandleSendTestMail)               // 发送测试邮件
		adminAuthorized.DELETE("/notice/delete", HandleDeleteNotice)         // 删除一条通知
		adminAuthorized.GET("/plans/get", subscribePlan.HandleGetAllPlans)   // 获取所有的订阅
		adminAuthorized.POST("/document/add", Document.HandleAddNewDocument) // 添加一条说明文档

	}

	userAuthorized := protectedRoutes.Group("/user/v1", auth.RoleMiddleware())
	{
		// 用户特定的路由
		userAuthorized.GET("/profile")
		//...其他用户路由
		// 开始首页
		userAuthorized.GET("/notices/get", HandleGetAllNotices)
		userAuthorized.GET("/document/get", Document.HandleGetAllDocument)
		userAuthorized.GET("/plan/get", subscribePlan.HandleGetAllPlans)
	}

	r.Run("localhost:8080")
}
