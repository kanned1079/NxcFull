package routers

import (
	"NxcFull/nxc-backend/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartAdminReq() {
	r := gin.Default()

	// 中间件
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

	r.POST("/api/admin/login", handleUserLogin)
	//r.GET("/api/admin/getSysInfo", handleGetServerInfo)

	r.GET("/api/admin/getStartTheme", getStartTheme)

	authorized := r.Group("/api", auth.AuthMiddleware())
	//	后续所有的都在这里执行
	{
		authorized.GET("/admin/getSysInfo", handleGetServerInfo)
		authorized.POST("admin/save-setting", handleUpdateSystemSettings)
		authorized.POST("/admin/setSingleSetting", handleUpdateSingleOptions)
		authorized.GET("/admin/get-setting", handleGetSystemSetting)
		authorized.GET("/admin/getUserList", handleGetUserList)
		authorized.GET("/admin/get-all-notices", HandleGetAllNotices)
		authorized.POST("/admin/add-notice", HandleAddNotice)
		authorized.DELETE("/admin/delete-notice", HandleDeleteNotice)
	}

	r.Run("localhost:8080")
}
