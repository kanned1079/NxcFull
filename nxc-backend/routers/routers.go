package routers

import (
	"NxcFull/nxc-backend/Document"
	"NxcFull/nxc-backend/auth"
	"NxcFull/nxc-backend/coupon"
	"NxcFull/nxc-backend/keys"
	"NxcFull/nxc-backend/orders"
	"NxcFull/nxc-backend/privilege"
	"NxcFull/nxc-backend/subscribePlan"
	"NxcFull/nxc-backend/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func StartReq() {
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
		// adminAuthorized.POST("admin/save-settings", handleUpdateSystemSettings)

		adminAuthorized.PUT("/settings", handleUpdateSingleOptions) // 修改单个配置项目
		adminAuthorized.GET("/settings", handleGetSystemSetting)    // 获取所有的系统设置

		adminAuthorized.GET("/users", handleGetUserList) // 获取所有用户的列表

		adminAuthorized.GET("/notices/get", HandleGetAllNotices)     // 获取所有通知列表
		adminAuthorized.POST("/notice/add", HandleAddNotice)         // 添加一条新的通知
		adminAuthorized.DELETE("/notice/delete", HandleDeleteNotice) // 删除一条通知

		adminAuthorized.POST("/mail/test", HandleSendTestMail) // 发送测试邮件

		adminAuthorized.GET("/plans/get", subscribePlan.HandleGetAllPlans)   // 获取所有的订阅
		adminAuthorized.POST("/plans/add", subscribePlan.HandleAddNewPlan)   // 添加新的订阅
		adminAuthorized.POST("/document/add", Document.HandleAddNewDocument) // 添加一条说明文档
		adminAuthorized.POST("/coupon/add", coupon.HandleAddNewCoupon)       // 新建优惠券

		adminAuthorized.POST("/groups", privilege.HandleAddNewGroup) // 添加权限组
		adminAuthorized.GET("/groups", privilege.HandleGetAllGroups) // 获取所有权限组列表
		adminAuthorized.PUT("/groups", privilege.HandleUpdateGroup)
		adminAuthorized.DELETE("/groups", privilege.HandleDeleteGroup)

		adminAuthorized.GET("/coupon/fetch", coupon.HandleGetAllCoupons)
		adminAuthorized.POST("/coupon/status/update", coupon.HandleActiveCoupon)
		adminAuthorized.GET("/plans/kv/fetch", subscribePlan.HandleGetAllPlanKeyName)

	}

	userAuthorized := protectedRoutes.Group("/user/v1", auth.RoleMiddleware())
	{
		// 用户特定的路由
		userAuthorized.GET("/profile")
		//...其他用户路由
		// 开始首页
		userAuthorized.GET("/notices/get", HandleGetAllNotices)            // 获取所有的通知列表
		userAuthorized.GET("/document/get", Document.HandleGetAllDocument) // 按照分类获取所有的文档列表
		userAuthorized.GET("/plan/get", subscribePlan.HandleGetAllPlans)   // 获取所有的有效订阅
		userAuthorized.POST("/coupon/verify", coupon.HandleVerifyCoupon)   // 验证优惠券可用性
		userAuthorized.POST("/auth/passcode/verify", user.HandleCheckOldPassword)
		userAuthorized.POST("/auth/passcode/update", user.HandleApplyNewPassword)
		userAuthorized.GET("/plan/info/fetch", orders.GetActivePlanListByUserId)
		userAuthorized.GET("/keys", keys.HandleGetAllUserKeys)

		userAuthorized.PUT("/order", orders.HandleCommitNewOrder) // 下单新的订阅
		// 这里将生成订单号后推送至消息队列 交由第二个进程进行处理

		userAuthorized.GET("/orders", orders.HandleGetOrders)
	}

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}

}
