package routers

import (
	"fmt"
	"gateway/internal/config"
	"gateway/internal/handler"
	"gateway/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func StartApiGateways() {
	router := gin.Default()
	// 中间件 过OPTIONS预检
	router.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
			return
		}
		context.Next()
	})

	publicRoutes := router.Group("/api")
	//publicRoutes := router.Group("/ws")

	// 这里是获取网站的标题 颜色 说明 配置
	//publicRoutes.GET("/admin/getStartTheme", getStartTheme)

	publicRoutes.GET("/app/preference/")

	adminPublic := publicRoutes.Group("/admin/v1")
	{
		//	// 管理员的公共路由
		adminPublic.POST("/login", handler.HandleUserLogin)
	}

	userPublic := publicRoutes.Group("/user/v1")
	{

		// 一般用户登录和注册相关路由
		userPublic.POST("/login", handler.HandleUserLogin) // rpc实现

		userPublic.POST("/mail/register/get", handler.HandleSendRegisterVerifyEmail) // 发送的是注册的验证码
		userPublic.POST("mail/retrieve/get", handler.HandleSendRetrieveVerifyEmail)  // 忘记密码的验证码
		userPublic.POST("/mail/register/verify", handler.HandleVerifyEmailCode)      // 这个可以是通用的函数
		userPublic.POST("/register/register", handler.HandleUserRegister)            // rpc实现
	}

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.AuthMiddleware()) // 验证Token中间价
	protectedRoutes.Use(middleware.RoleMiddleware()) // 鉴权中间件
	//protectedRoutesWs := router.Group("/ws")
	//protectedRoutesWs.Use(middleware.AuthMiddleware())
	//protectedRoutesWs.Use(middleware.RoleMiddleware())

	adminAuthorized := protectedRoutes.Group("/admin/v1", middleware.RoleMiddleware())
	{
		//// 管理员特定的路由
		adminAuthorized.GET("/os/status/get") // 获取当前系统的负载
		//// adminAuthorized.POST("admin/save-settings", handleUpdateSystemSettings)
		//
		adminAuthorized.PUT("/setting", handler.HandleUpdateSingleOptions) // 修改单个配置项目
		adminAuthorized.GET("/setting", handler.HandleGetSystemSetting)    // 获取所有的系统设置
		//
		//adminAuthorized.GET("/users", handleGetUserList) // 获取所有用户的列表
		//
		adminAuthorized.GET("/notice", handler.HandleGetAllNotices)             // 获取所有通知列表
		adminAuthorized.POST("/notice", handler.HandleAddNotice)                // 添加一条新的通知
		adminAuthorized.PUT("notice", handler.HandleUpdateNotice)               // 修改通知信息
		adminAuthorized.PUT("/notice/status", handler.HandleUpdateNoticeStatus) // 更新通知的状态
		adminAuthorized.DELETE("/notice", handler.HandleDeleteNotice)           // 删除一条通知

		adminAuthorized.POST("/mail/test", handler.HandleSendTestMail) // 发送测试邮件

		adminAuthorized.GET("/plan", handler.HandleGetAllPlans)          // 获取所有的订阅
		adminAuthorized.POST("/plan", handler.HandleAddNewPlan)          // 添加新的订阅
		adminAuthorized.PUT("/plan", handler.HandleUpdatePlan)           // 修改订阅的细节
		adminAuthorized.DELETE("/plan", handler.HandleDeletePlan)        // 删除指定的订阅
		adminAuthorized.GET("/plan/kv", handler.HandleGetAllPlanKeyName) // 获取订阅计划的键值对
		adminAuthorized.PUT("plan/sale", handler.HandleUpdatePlanSale)   // 更新是否启用销售
		adminAuthorized.PUT("plan/renew", handler.HandleUpdatePlanRenew) // 更新是否允许续费

		adminAuthorized.GET("/document", handler.HandleGetAllDocumentsAdmin)
		adminAuthorized.PUT("/document", handler.HandleUpdateDocument)
		adminAuthorized.PATCH("/document", handler.HandleUpdateDocumentShow)
		adminAuthorized.DELETE("/document", handler.HandleDeleteDocumentById)
		adminAuthorized.POST("/document", handler.HandleAddNewDocument) // 添加一条说明文档
		//
		adminAuthorized.POST("/groups", handler.HandleAddNewGroup) // 添加权限组
		adminAuthorized.GET("/groups", handler.HandleGetAllGroups) // 获取所有权限组列表
		adminAuthorized.GET("groups/kv", handler.HandleGetAllGroupsKv)
		adminAuthorized.PUT("/groups", handler.HandleUpdateGroup)
		adminAuthorized.DELETE("/groups", handler.HandleDeleteGroup)
		//
		adminAuthorized.POST("/coupon", handler.HandleAddNewCoupon)       // 新建优惠券	POST
		adminAuthorized.GET("/coupon", handler.HandleGetAllCoupons)       // 获取优惠券列表 GET
		adminAuthorized.PUT("/coupon/status", handler.HandleActiveCoupon) // PUT
		adminAuthorized.DELETE("/coupon", handler.HandleDeleteCoupon)

		adminAuthorized.GET("/ticket", handler.HandleGetAllTickets)

	}

	userAuthorized := protectedRoutes.Group("/user/v1", middleware.RoleMiddleware())
	{
		//// 用户特定的路由
		userAuthorized.GET("/profile")
		userAuthorized.GET("/user", handler.HandleUpdateUserInfo)

		userAuthorized.GET("/notice", handler.HandleGetAllNotices)        // 获取所有的通知列表
		userAuthorized.GET("/document", handler.HandleGetAllDocuments)    // 按照分类获取所有的文档列表	// rpc实现
		userAuthorized.GET("/plan", handler.HandleGetAllPlans)            // 获取所有的有效订阅
		userAuthorized.POST("/coupon/verify", handler.HandleVerifyCoupon) // 验证优惠券可用性	// POST
		userAuthorized.POST("/auth/passcode/verify", handler.HandleCheckPreviousPassword)
		userAuthorized.POST("/auth/passcode/update", handler.HandleApplyNewPassword)

		userAuthorized.GET("/auth/2fa/status", handler.HandleGet2FAStatus)            // 检查是否启用了2FA
		userAuthorized.POST("/auth/2fa/setup", handler.HandleSetup2FA)                // 新建2fa请求
		userAuthorized.POST("/auth/2fa/setup/test", handler.HandleTest2FA)            // 测试2FA是否可用
		userAuthorized.DELETE("/auth/2fa/setup/cancel", handler.HandleCancelSetup2FA) // 取消2fa在redis中的暂存
		userAuthorized.DELETE("auth/2fa/disable", handler.HandleDisable2FA)

		userAuthorized.GET("/plan/summary/fetch", handler.GetActivePlanListByUserId)

		// 这里将生成订单号后推送至消息队列 交由第二个进程进行处理
		userAuthorized.GET("/orders", handler.HandleGetAllMyOrders)
		userAuthorized.GET("/order/status", handler.HandleCheckOrderStatus)
		userAuthorized.POST("/order", handler.HandleCommitNewOrder) // 下单新的订阅
		userAuthorized.DELETE("/order", handler.HandleCancelOrder)  // 取消订阅
		userAuthorized.PUT("/order", handler.HandlePlaceOrder)      // 确认并下单

		userAuthorized.GET("/keys", handler.HandleGetAllMyKeys)

		userAuthorized.GET("/ticket", handler.HandleGetAllMyTickets)
		userAuthorized.POST("/ticket", handler.HandleCreateNewTicket) // 新建工单
		userAuthorized.DELETE("/ticket", handler.HandleCloseTicket)   // 关闭工单
		//userAuthorized.PUT("/ticket/chat", handler.HandleSendChatContent) // 发送消息
		//userAuthorized.GET("/ticket/chat", handler.HandleGetChatContent)  // 获取聊天消息内容

	}

	// WebSocket 路由组
	wsRoutes := router.Group("/ws")
	//wsRoutes.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware())
	// 为 WebSocket 配置单独的用户路由
	userWsRoutes := wsRoutes.Group("/user/v1")
	userWsRoutes.GET("/chat", handler.HandleChatWs)

	//userAuthorizedWs := protectedRoutes.Group("/user/v1", middleware.RoleMiddleware())
	//{
	//	userAuthorizedWs.GET("/chat", handler.HandleChatWs)
	//}

	if err := router.Run(fmt.Sprintf("%s:%s", config.MyServerConfig.ListenAddr, strconv.Itoa(int(config.MyServerConfig.ListenPort)))); err != nil {
		log.Println("启动服务器失败", err)
	}
}
