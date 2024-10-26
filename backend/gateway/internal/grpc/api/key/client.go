package key

import (
	"gateway/internal/etcd"
	pb "gateway/internal/grpc/api/key/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

func NewKeyClient() (client pb.KeyServiceClient) {
	log.Println("NewKeyClient")
	keepAliveParams := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             2 * time.Second,
		PermitWithoutStream: true,
	}
	clientConn, err := grpc.NewClient(
		etcd.GetServiceAddress("keyServices"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepAliveParams),
	)
	if err != nil {
		log.Println("连接到key grpc服务器失败", err.Error())
		return
	}
	log.Println("连接到noticeServices成功")
	return pb.NewKeyServiceClient(clientConn)
}
