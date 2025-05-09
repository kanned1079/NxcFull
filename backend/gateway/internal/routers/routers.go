package routers

import (
	"fmt"
	"gateway/internal/config"
	"gateway/internal/handler"
	"gateway/internal/middleware"
	"gateway/internal/utils"
	"log"
	"strconv"
)

func (inst *GatewayApp) StartApiGateways() {
	//router := gin.Default()

	inst.Router.Use(middleware.ProtocolAllowance())      // 設置後端允許的請求方式和过OPTIONS预检
	inst.Router.Use(middleware.AccessLoggerMiddleware()) // 統計API訪問紀錄 提交到日誌為服務
	inst.Router.Use(middleware.APICountMiddleware())     // 對API訪問進行計數

	publicRoutes := inst.Router.Group("/api")

	// 这里是获取网站的标题 颜色 说明 配置
	publicRoutes.GET("/app/preference/")
	publicRoutes.GET("/app/v1/env/root", handler.HandleGetAppRuntimeEnv)
	publicRoutes.GET("/app/v1/env/welcome", handler.HandleGetWelcomeConfig)
	publicRoutes.GET("/app/v1/env/register", handler.HandleGetRegisterEnv)
	publicRoutes.POST("/user/v1/activation/bind", handler.HandleBindKeyToThirdExternalApp)
	publicRoutes.POST("/user/v1/activation/status/check", handler.HandleCheckKeyValidation)

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

	protectedRoutes := inst.Router.Group("/api")
	protectedRoutes.Use(middleware.AuthMiddleware()) // 验证Token中间价
	protectedRoutes.Use(middleware.RoleMiddleware()) // 鉴权中间件
	adminAuthorized := protectedRoutes.Group("/admin/v1", middleware.RoleMiddleware())
	{
		//// 管理员特定的路由
		adminAuthorized.GET("/server/latency/test", utils.ConnectionTestForAdmin)
		adminAuthorized.GET("/server/status/fetch", handler.HandleGetSystemStatus)
		adminAuthorized.DELETE("/server/log/delete", handler.HandleClearPreviousLog)
		adminAuthorized.GET("/os/status/get") // 获取当前系统的负载
		adminAuthorized.GET("/server/status")
		adminAuthorized.GET("/infrastructure/status", handler.GetSystemInfrastructureInfo)
		//// adminAuthorized.POST("admin/save-settings", handleUpdateSystemSettings)

		adminAuthorized.GET("/app/overview", handler.HandleGetAppOverview)
		adminAuthorized.GET("/app/users/layout", handler.HandleGetUserLayout)
		adminAuthorized.GET("/app/common/config", handler.HandleFetchServerSystemPartConfig) // 获取所有用户的列表

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

		adminAuthorized.GET("/document", handler.HandleGetAllDocumentsAdmin)  // 获取文档列表
		adminAuthorized.PUT("/document", handler.HandleUpdateDocument)        // 更新指定id的文档
		adminAuthorized.PATCH("/document", handler.HandleUpdateDocumentShow)  // 更新是否展示指定id的文档给用户
		adminAuthorized.DELETE("/document", handler.HandleDeleteDocumentById) // 按照id删除文档
		adminAuthorized.POST("/document", handler.HandleAddNewDocument)       // 添加一条说明文档

		adminAuthorized.POST("/groups", handler.HandleAddNewGroup)     // 添加权限组
		adminAuthorized.GET("/groups", handler.HandleGetAllGroups)     // 获取所有权限组列表
		adminAuthorized.GET("groups/kv", handler.HandleGetAllGroupsKv) // 快速查询 权限组的键值
		adminAuthorized.PUT("/groups", handler.HandleUpdateGroup)      // 更新权限组名称
		adminAuthorized.DELETE("/groups", handler.HandleDeleteGroup)   // 删除指定id的权限组

		adminAuthorized.POST("/coupon", handler.HandleAddNewCoupon)       // 新建优惠券
		adminAuthorized.GET("/coupon", handler.HandleGetAllCoupons)       // 获取优惠券列表
		adminAuthorized.PUT("/coupon", handler.HandleUpdateCoupons)       // 修改優惠券
		adminAuthorized.PUT("/coupon/status", handler.HandleActiveCoupon) // 按照id开启或关闭优惠券
		adminAuthorized.DELETE("/coupon", handler.HandleDeleteCoupon)     // 删除指定id的优惠券

		adminAuthorized.GET("/ticket", handler.HandleGetAllTickets) // 获取所有的工单列表 分组为活跃和历史
		// 工单的信息发送和接收更新由wsRoutes路由组进行传输 握手后升级为ws传输协议

		adminAuthorized.GET("users", handler.HandleGetAllUsers)                      // 获取所有用户列表 有搜索邮箱选项
		adminAuthorized.POST("users", handler.HandleAddUserManuallyFromAdmin)        // 管理员手动添加一个新用户
		adminAuthorized.PUT("users", handler.HandleUpdateUserInfoByIdFromAdmin)      // 修改指定id的用户的信息
		adminAuthorized.PATCH("users", handler.HandleBlock2UnblockUserByIdFromAdmin) // 封禁和解封用户

		adminAuthorized.GET("order", handler.HandleGetAllOrderByAdmin)       // 管理员获获取所有用户的订单 包括redis缓存的
		adminAuthorized.DELETE("order", handler.HandleManualCancelUserOrder) // 管理员手动取消用户的订单
		adminAuthorized.PUT("order", handler.HandleManualPassUserOrder)      // 管理员手动通过用户的订单

		adminAuthorized.GET("payment", handler.HandleGetAllPaymentMethodKv)                       // 获取所有存在的支付方式基础信息 名称 是否启用 优惠金额
		adminAuthorized.GET("payment/details", handler.HandleGetPaymentMethodDetailsBySystemName) // 获取详细信息 但是隐藏证书私钥等信息
		adminAuthorized.POST("payment", handler.HandleEditOrSavePaymentMethodBySystemName)        // 新建或编辑支付方式的配置
		adminAuthorized.PATCH("payment", handler.HandleSwitchPaymentMethodEnableBySystemName)     // 切换是否启用支付方式

		adminAuthorized.GET("activation", handler.GetAllActivateLogByAdmin)
		adminAuthorized.GET("keys", handler.HandleGetAllKeysByAdmin)
		adminAuthorized.DELETE("keys", handler.HandleAlterKeyIsValidByAdmin)

	}

	userAuthorized := protectedRoutes.Group("/user/v1", middleware.RoleMiddleware())
	{
		//// 用户特定的路由
		userAuthorized.GET("/profile")
		userAuthorized.GET("/user", handler.HandleUpdateUserInfo)
		userAuthorized.DELETE("/user/delete", handler.HandleDeleteUserAccount)
		userAuthorized.PATCH("/user/name", handler.HandleAlterUsername)
		userAuthorized.GET("/user/avatar", handler.HandleGetUserAvatar)
		userAuthorized.POST("/user/upload/avatar", handler.HandleUserUploadAvatar)

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

		// 充值
		//userAuthorized.POST("/balance/recharge", handler.HandleUserTopUp)

		userAuthorized.GET("/keys", handler.HandleGetAllMyKeys)
		userAuthorized.GET("/key/details", handler.HandleGetKeyDetailsById)

		userAuthorized.GET("ticket/check", handler.HandleCheckIsUserHaveOpeningTickets)
		userAuthorized.GET("/ticket", handler.HandleGetAllMyTickets)
		userAuthorized.POST("/ticket", handler.HandleCreateNewTicket) // 新建工单
		userAuthorized.DELETE("/ticket", handler.HandleCloseTicket)   // 关闭工单
		//userAuthorized.PUT("/ticket/chat", handler.HandleSendChatContent) // 发送消息
		//userAuthorized.GET("/ticket/chat", handler.HandleGetChatContent)  // 获取聊天消息内容

		userAuthorized.GET("/invite/banner", handler.HandleGetUserInviteBanner)
		userAuthorized.GET("/invite/code", handler.HandleGetUserInviteCodeByUserId)
		userAuthorized.POST("/invite/code", handler.HandleCreateUserInviteCodeByUserId)
		userAuthorized.GET("/invite/users", handler.HandleGetUserInvitedUserListByUserId)

		userAuthorized.GET("/payment/methods", handler.HandleGetAllPaymentMethodKv) // 获取所有存在的支付方式基础信息 名称 是否启用 优惠金额
		userAuthorized.POST("/payment/top-up", handler.HandleUserCommitNewTopUpOrder)
		userAuthorized.GET("/payment/top-up/check", handler.HandleQueryTopUpOrderStatus)

		userAuthorized.GET("/activation", handler.HandleGetAllMyActivationLogs)
		userAuthorized.DELETE("/activation", handler.HandleDisableBindKey)
		userAuthorized.PATCH("/activation/remark", handler.HandleAlterRemarkByUser)

		userAuthorized.GET("/app", handler.HandleGetAllAppDownloadLink)

	}

	// WebSocket 路由组
	wsRoutes := inst.Router.Group("/api/ws")
	//wsRoutes.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware())
	// 为 WebSocket 配置单独的用户路由
	userWsRoutes := wsRoutes.Group("/user/v1")
	userWsRoutes.GET("/chat", handler.HandleChatWs)

	if config.MyServerConfig.SSL {
		//config.MyServerConfig.Crt
		//config.MyServerConfig.Key
		// 启动启用 SSL 的服务器
		if err := inst.Router.RunTLS(
			fmt.Sprintf("%s:%s",
				config.MyServerConfig.ListenAddr,
				strconv.Itoa(int(config.MyServerConfig.ListenPort))),
			"/app/config/cert/"+config.MyServerConfig.Crt,
			"/app/config/cert/"+config.MyServerConfig.Key,
		); err != nil {
			log.Println("run server failure with ssl: ", err)
			return
		}

	} else {
		if err := inst.Router.Run(fmt.Sprintf("%s:%s", config.MyServerConfig.ListenAddr, strconv.Itoa(int(config.MyServerConfig.ListenPort)))); err != nil {
			log.Println("run server failure: ", err)
			return
		}
	}
}
