package user

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	pb "keyServices/internal/client/api/user/proto"
	"keyServices/internal/etcd"
	"log"
	"time"
)

func NewUserClient() (client pb.UserServiceClient) {
	log.Println("NewUserClient")
	keepAliveParams := keepalive.ClientParameters{
		Time:                10 * time.Second, // 10秒内无操作则发送KeepAlive ping
		Timeout:             2 * time.Second,  // 超过2秒未响应则认为连接不可用
		PermitWithoutStream: true,             // 即使没有活跃的RPC流也发送KeepAlive ping
	}
	clientConn, err := grpc.NewClient(
		etcd.GetServiceAddress("userServices"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepAliveParams))
	if err != nil {
		log.Println("连接到user grpc服务器失败", err.Error())
		return
	}
	log.Println("连接到userServices成功")
	return pb.NewUserServiceClient(clientConn)
}
