package grpc

import (
	"NxcFull/backend/gateway/internal/grpc/api/proto/user"
	"google.golang.org/grpc"
)

type Clients struct {
	//userServiceClient pb.UserServiceClient
	UserServiceClient    *grpc.ClientConn
	CouponServiceClient  *grpc.ClientConn
	NoticesServiceClient *grpc.ClientConn
	OrderServiceClient   *grpc.ClientConn
}

func NewClients() Clients {
	// 初始化gRPC客户端
	return Clients{
		UserServiceClient: user.NewUserClient(),
	}

}
