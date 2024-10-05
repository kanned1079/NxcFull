package routers

import (
	"NxcFull/backend/gateway/internal/handler"
	"NxcFull/nxc-backend/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func StartApiGateways() {
	router := gin.Default()
	// 中间件 过OPTIONS预检
	router.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
			return
		}
		context.Next()
	})

	publicRoutes := router.Group("/api")

	// 这里是获取网站的标题 颜色 说明 配置
	//publicRoutes.GET("/admin/getStartTheme", getStartTheme)

	publicRoutes.GET("/app/preference/")

	//adminPublic := publicRoutes.Group("/admin/v1")
	//{
	//	// 管理员的公共路由
	//	adminPublic.POST("/login", handleUserLogin)
	//}

	userPublic := publicRoutes.Group("/user/v1")
	{

		// 一般用户登录和注册相关路由
		userPublic.POST("/login", handler.HandleUserLogin) // rpc实现

		//userPublic.POST("/register/vc/get", SendVerifyEmail)
		//userPublic.POST("/register/vc/verify", HandleVerifyEmailCode)
		userPublic.POST("/register/register", handler.HandleUserRegister) // rpc实现
	}

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(auth.AuthMiddleware()) // 验证Token中间价
	protectedRoutes.Use(auth.RoleMiddleware()) // 鉴权中间件

	adminAuthorized := protectedRoutes.Group("/admin/v1", auth.RoleMiddleware())
	{
		//// 管理员特定的路由
		adminAuthorized.GET("/os/status/get") // 获取当前系统的负载
		//// adminAuthorized.POST("admin/save-setting", handleUpdateSystemSettings)
		//
		//adminAuthorized.PUT("/settings", handleUpdateSingleOptions) // 修改单个配置项目
		//adminAuthorized.GET("/settings", handleGetSystemSetting)    // 获取所有的系统设置
		//
		//adminAuthorized.GET("/users", handleGetUserList) // 获取所有用户的列表
		//
		//adminAuthorized.GET("/notices/get", HandleGetAllNotices)             // 获取所有通知列表
		//adminAuthorized.POST("/notice/add", HandleAddNotice)                 // 添加一条新的通知
		//adminAuthorized.POST("/mail/test", HandleSendTestMail)               // 发送测试邮件
		//adminAuthorized.DELETE("/notice/delete", HandleDeleteNotice)         // 删除一条通知
		//adminAuthorized.GET("/plans/get", subscribePlan.HandleGetAllPlans)   // 获取所有的订阅
		//adminAuthorized.POST("/plans/add", subscribePlan.HandleAddNewPlan)   // 添加新的订阅
		adminAuthorized.POST("/document/add", handler.HandleAddNewDocument) // 添加一条说明文档
		//adminAuthorized.POST("/coupon/add", coupon.HandleAddNewCoupon)       // 新建优惠券
		//
		//adminAuthorized.POST("/groups", privilege.HandleAddNewGroup) // 添加权限组
		//adminAuthorized.GET("/groups", privilege.HandleGetAllGroups) // 获取所有权限组列表
		//adminAuthorized.PUT("/groups", privilege.HandleUpdateGroup)
		//adminAuthorized.DELETE("/groups", privilege.HandleDeleteGroup)
		//
		//adminAuthorized.GET("/coupon/fetch", coupon.HandleGetAllCoupons)
		//adminAuthorized.POST("/coupon/status/update", coupon.HandleActiveCoupon)
		//adminAuthorized.GET("/plans/kv/fetch", subscribePlan.HandleGetAllPlanKeyName)

	}

	userAuthorized := protectedRoutes.Group("/user/v1", auth.RoleMiddleware())
	{
		//// 用户特定的路由
		userAuthorized.GET("/profile")
		////...其他用户路由
		//// 开始首页
		//userAuthorized.GET("/notices/get", HandleGetAllNotices)            // 获取所有的通知列表
		userAuthorized.GET("/document", handler.HandleGetAllDocuments) // 按照分类获取所有的文档列表	// rpc实现
		//userAuthorized.GET("/plan/get", subscribePlan.HandleGetAllPlans)   // 获取所有的有效订阅
		//userAuthorized.POST("/coupon/verify", coupon.HandleVerifyCoupon)   // 验证优惠券可用性
		userAuthorized.POST("/auth/passcode/verify", handler.HandleCheckPreviousPassword)
		userAuthorized.POST("/auth/passcode/update", handler.HandleApplyNewPassword)
		//userAuthorized.GET("/plan/info/fetch", orders.GetActivePlanListByUserId)
		//userAuthorized.GET("/keys", keys.HandleGetAllUserKeys)
		//
		//userAuthorized.PUT("/order", orders.HandleCommitNewOrder) // 下单新的订阅
		//// 这里将生成订单号后推送至消息队列 交由第二个进程进行处理
		//
		//userAuthorized.GET("/orders", orders.HandleGetOrders)
	}

	if err := router.Run("localhost:8081"); err != nil {
		log.Println("启动服务器失败", err)
	}
}
