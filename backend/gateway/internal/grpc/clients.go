package grpc

import (
	"gateway/internal/grpc/api/coupon"
	couponPb "gateway/internal/grpc/api/coupon/proto"
	"gateway/internal/grpc/api/document"
	documentPb "gateway/internal/grpc/api/document/proto"
	"gateway/internal/grpc/api/group"
	groupPb "gateway/internal/grpc/api/group/proto"
	"gateway/internal/grpc/api/key"
	keyPb "gateway/internal/grpc/api/key/proto"
	"gateway/internal/grpc/api/logs"
	logPb "gateway/internal/grpc/api/logs/proto"
	"gateway/internal/grpc/api/notice"
	noticePb "gateway/internal/grpc/api/notice/proto"
	"gateway/internal/grpc/api/order"
	orderPb "gateway/internal/grpc/api/order/proto"
	"gateway/internal/grpc/api/orderHandle"
	orderHandlePb "gateway/internal/grpc/api/orderHandle/proto"
	"gateway/internal/grpc/api/sendmail"
	mailPb "gateway/internal/grpc/api/sendmail/proto"
	"gateway/internal/grpc/api/settings"
	settingPb "gateway/internal/grpc/api/settings/proto"
	"gateway/internal/grpc/api/subscription"
	subscribePb "gateway/internal/grpc/api/subscription/proto"
	"gateway/internal/grpc/api/ticketHandle"
	ticketHandlePb "gateway/internal/grpc/api/ticketHandle/proto"
	"gateway/internal/grpc/api/user"
	userPb "gateway/internal/grpc/api/user/proto"
)

type Clients struct {
	UserServiceClient         userPb.UserServiceClient
	DocumentServiceClient     documentPb.DocumentServiceClient
	NoticeServiceClient       noticePb.NoticeServiceClient
	GroupServiceClient        groupPb.GroupServiceClient
	SubscriptionServiceClient subscribePb.SubscriptionServiceClient
	CouponServiceClient       couponPb.CouponServiceClient
	SettingServiceClient      settingPb.SettingsServiceClient
	MailServicesClient        mailPb.MailServiceClient
	OrderServicesClient       orderPb.OrderServiceClient
	OrderHandleServicesClient orderHandlePb.OrderHandleServiceClient
	KeyServicesClient         keyPb.KeyServiceClient
	TicketHandleClient        ticketHandlePb.TicketHandleServiceClient
	LogServiceClient          logPb.LogServiceClient
}

func NewClients() Clients {
	// 初始化gRPC客户端
	return Clients{
		UserServiceClient:         user.NewUserClient(),
		DocumentServiceClient:     document.NewDocumentClient(),
		NoticeServiceClient:       notice.NewNoticeClient(),
		GroupServiceClient:        group.NewGroupClient(),
		SubscriptionServiceClient: subscription.NewSubscriptionClient(),
		CouponServiceClient:       coupon.NewCouponServiceClient(),
		SettingServiceClient:      settings.NewSettingClient(),
		MailServicesClient:        sendmail.NewMailClient(),
		OrderServicesClient:       order.NewOrderClient(),
		OrderHandleServicesClient: orderHandle.NewOrderHandleClient(),
		KeyServicesClient:         key.NewKeyClient(),
		TicketHandleClient:        ticketHandle.NewTicketHandleClient(),
		LogServiceClient:          logs.NewLogsClient(),
	}
}
