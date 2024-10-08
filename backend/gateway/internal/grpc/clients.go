package grpc

import (
	"NxcFull/backend/gateway/internal/grpc/api/coupon"
	"NxcFull/backend/gateway/internal/grpc/api/document"
	"NxcFull/backend/gateway/internal/grpc/api/group"
	"NxcFull/backend/gateway/internal/grpc/api/notice"
	pb "NxcFull/backend/gateway/internal/grpc/api/proto"
	"NxcFull/backend/gateway/internal/grpc/api/subscribe"
	"NxcFull/backend/gateway/internal/grpc/api/user"
	//"google.golang.org/grpc"
)

type Clients struct {
	//userServiceClient pb.UserServiceClient
	UserServiceClient pb.UserServiceClient

	//CouponServiceClient  *grpc.ClientConn
	//NoticesServiceClient *grpc.ClientConn
	//OrderServiceClient   *grpc.ClientConn
	DocumentServiceClient     pb.DocumentServiceClient
	NoticeServiceClient       pb.NoticeServiceClient
	GroupServiceClient        pb.GroupServiceClient
	SubscriptionServiceClient pb.SubscriptionServiceClient
	CouponServiceClient       pb.CouponServiceClient
}

func NewClients() Clients {
	// 初始化gRPC客户端
	return Clients{
		UserServiceClient:         user.NewUserClient(),
		DocumentServiceClient:     document.NewDocumentClient(),
		NoticeServiceClient:       notice.NewNoticeClient(),
		GroupServiceClient:        group.NewGroupClient(),
		SubscriptionServiceClient: subscribe.NewSubscriptionServiceClient(),
		CouponServiceClient:       coupon.NewCouponServiceClient(),
	}

}
