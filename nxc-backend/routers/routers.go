package routers

import (
	"NxcFull/nxc-backend/auth"
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
		adminPublic.POST("/api/admin/v1/login", handleUserLogin)
	}

	userPublic := publicRoutes.Group("/user/v1")
	{
		// 一般用户登录和注册相关路由
		userPublic.POST("/login", handleUserLogin)
		userPublic.POST("/register/vc/get", SendVerifyEmail)
		userPublic.POST("/register/vc/verify", HandleVerifyEmailCode)
		userPublic.POST("/register/register", HandleUserRegister)
	}

	//// 一般用户
	//r.POST("/api/v1/user/login")
	//// 注册时邮件验证
	//r.POST("/api/user/v1/register/vc/get", SendVerifyEmail)
	//r.POST("/api/user/v1/register/vc/verify", HandleVerifyEmailCode)
	//r.POST("/api/user/v1/register/register", HandleUserRegister)

	// 使用token验证和角色选择中间件
	//r.Use(auth.AuthMiddleware())
	//r.Use(auth.RoleMiddleware())

	protectedRoutes := r.Group("/api")
	protectedRoutes.Use(auth.AuthMiddleware())
	protectedRoutes.Use(auth.RoleMiddleware())

	//authorized := r.Group("/api", auth.AuthMiddleware())
	////	管理员认证后所有的都在这里执行
	//{
	//	authorized.GET("/admin/v1/os/status/get", handleGetServerInfo)
	//	// authorized.POST("admin/save-setting", handleUpdateSystemSettings)
	//	authorized.POST("/admin/v1/settings/set", handleUpdateSingleOptions)
	//	authorized.GET("/admin/v1/settings/get", handleGetSystemSetting)
	//	authorized.GET("/admin/v1/users/get", handleGetUserList)
	//	authorized.GET("/admin/v1/notices/get", HandleGetAllNotices)
	//	authorized.POST("/admin/v1/notice/add", HandleAddNotice)
	//	authorized.DELETE("/admin/v1/notice/delete", HandleDeleteNotice)
	//}

	// new
	//authorized := r.Group("/api", auth.AuthMiddleware())

	adminAuthorized := protectedRoutes.Group("/admin/v1", auth.RoleMiddleware())
	{
		// 管理员特定的路由
		adminAuthorized.GET("/os/status/get", handleGetServerInfo)
		// adminAuthorized.POST("admin/save-setting", handleUpdateSystemSettings)
		adminAuthorized.POST("/settings/set", handleUpdateSingleOptions)
		adminAuthorized.GET("/settings/get", handleGetSystemSetting)
		adminAuthorized.GET("/users/get", handleGetUserList)
		adminAuthorized.GET("/notices/get", HandleGetAllNotices)
		adminAuthorized.POST("/notice/add", HandleAddNotice)
		adminAuthorized.DELETE("/notice/delete", HandleDeleteNotice)
	}

	userAuthorized := protectedRoutes.Group("/user/v1", auth.RoleMiddleware())
	{
		// 用户特定的路由
		userAuthorized.GET("/profile")
		//...其他用户路由
	}

	r.Run("localhost:8080")
}
