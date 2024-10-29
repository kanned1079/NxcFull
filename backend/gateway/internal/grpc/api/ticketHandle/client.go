package ticketHandle

import (
	"gateway/internal/etcd"
	pb "gateway/internal/grpc/api/ticketHandle/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

func NewTicketHandleClient() (client pb.TicketHandleServiceClient) {
	log.Println("NewTicketHandleClient", etcd.GetServiceAddress("ticketHandleServices"))
	keepAliveParams := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             2 * time.Second,
		PermitWithoutStream: true,
	}
	clientConn, err := grpc.NewClient(
		etcd.GetServiceAddress("ticketHandleServices"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepAliveParams),
	)
	if err != nil {
		log.Println("连接到user grpc服务器失败", err.Error())
		return
	}
	log.Println("连接到SubscriptionService成功")
	return pb.NewTicketHandleServiceClient(clientConn)
}
