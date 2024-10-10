package grpc

import (
	"NxcFull/backend/gateway/internal/grpc/api/coupon"
	couponPb "NxcFull/backend/gateway/internal/grpc/api/coupon/proto"
	"NxcFull/backend/gateway/internal/grpc/api/document"
	documentPb "NxcFull/backend/gateway/internal/grpc/api/document/proto"
	"NxcFull/backend/gateway/internal/grpc/api/group"
	groupPb "NxcFull/backend/gateway/internal/grpc/api/group/proto"
	"NxcFull/backend/gateway/internal/grpc/api/notice"
	noticePb "NxcFull/backend/gateway/internal/grpc/api/notice/proto"
	"NxcFull/backend/gateway/internal/grpc/api/setting"
	settingPb "NxcFull/backend/gateway/internal/grpc/api/setting/proto"
	"NxcFull/backend/gateway/internal/grpc/api/subscription"
	subscribePb "NxcFull/backend/gateway/internal/grpc/api/subscription/proto"
	"NxcFull/backend/gateway/internal/grpc/api/user"
	userPb "NxcFull/backend/gateway/internal/grpc/api/user/proto"
	//"google.golang.org/grpc"
)

type Clients struct {
	UserServiceClient         userPb.UserServiceClient
	DocumentServiceClient     documentPb.DocumentServiceClient
	NoticeServiceClient       noticePb.NoticeServiceClient
	GroupServiceClient        groupPb.GroupServiceClient
	SubscriptionServiceClient subscribePb.SubscriptionServiceClient
	CouponServiceClient       couponPb.CouponServiceClient
	SettingServiceClient      settingPb.SettingsServiceClient
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
		SettingServiceClient:      setting.NewSettingClient(),
	}

}
