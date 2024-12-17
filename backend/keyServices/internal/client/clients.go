package client

import (
	"keyServices/internal/client/api/user"
	userPb "keyServices/internal/client/api/user/proto"
)

type Clients struct {
	UserServiceClient userPb.UserServiceClient
}

func NewClients() Clients {
	// 初始化gRPC客户端
	return Clients{
		UserServiceClient: user.NewUserClient(),
	}

}
