package subscription

import (
	"NxcFull/backend/gateway/internal/etcd"
	pb "NxcFull/backend/gateway/internal/grpc/api/subscription/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewSubscriptionClient() (client pb.SubscriptionServiceClient) {
	log.Println("NewSubscriptionServiceClient", etcd.GetServiceAddress("subscribeServices"))
	clientConn, err := grpc.NewClient(etcd.GetServiceAddress("subscribeServices"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接到user grpc服务器失败", err.Error())
		return
	}
	log.Println("连接到SubscriptionService成功")
	return pb.NewSubscriptionServiceClient(clientConn)
}
