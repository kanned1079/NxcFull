package user

import (
	"NxcFull/backend/gateway/internal/etcd"
	pb "NxcFull/backend/gateway/internal/grpc/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewUserClient() (client pb.UserServiceClient) {
	log.Println("NewUserClient")
	clientConn, err := grpc.NewClient(etcd.GetServiceAddress("userServices"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接到user grpc服务器失败", err.Error())
		return
	}
	log.Println("连接到userServices成功")
	return pb.NewUserServiceClient(clientConn)
}
