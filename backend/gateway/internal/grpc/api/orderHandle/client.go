package orderHandle

import (
	"gateway/internal/etcd"
	pb "gateway/internal/grpc/api/orderHandle/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

func NewOrderHandleClient() (client pb.OrderHandleServiceClient) {
	log.Println("NewOrderHandleClient")
	keepAliveParams := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             2 * time.Second,
		PermitWithoutStream: true,
	}
	clientConn, err := grpc.NewClient(
		etcd.GetServiceAddress("orderHandleServices"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepAliveParams),
	)
	if err != nil {
		log.Println("连接到orderHandle grpc服务器失败", err.Error())
		return
	}
	log.Println("连接到orderHandleServices成功")
	return pb.NewOrderHandleServiceClient(clientConn)
}
