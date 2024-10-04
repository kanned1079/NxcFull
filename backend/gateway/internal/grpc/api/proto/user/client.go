package user

import (
	pb "NxcFull/backend/gateway/internal/grpc/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewUserClient() (client pb.UserServiceClient) {
	clientConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接到user grpc服务器失败", err.Error())
		panic(err)
	}

	return pb.NewUserServiceClient(clientConn)
}
