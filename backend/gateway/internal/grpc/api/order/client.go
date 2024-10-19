package order

import (
	"gateway/internal/etcd"
	pb "gateway/internal/grpc/api/order/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

func NewOrderClient() (client pb.OrderServiceClient) {
	log.Println("NewOrderClient")
	keepAliveParams := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             2 * time.Second,
		PermitWithoutStream: true,
	}
	clientConn, err := grpc.NewClient(
		etcd.GetServiceAddress("orderServices"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepAliveParams),
	)
	if err != nil {
		log.Println("连接到order grpc服务器失败", err.Error())
		return
	}
	log.Println("连接到orderServices成功")
	return pb.NewOrderServiceClient(clientConn)
}
