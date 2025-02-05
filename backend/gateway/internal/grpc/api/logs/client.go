package logs

import (
	"gateway/internal/etcd"
	pb "gateway/internal/grpc/api/logs/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

func NewLogsClient() (client pb.LogServiceClient) {
	log.Println("NewLogsClient")
	keepAliveParams := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             2 * time.Second,
		PermitWithoutStream: true,
	}
	clientConn, err := grpc.NewClient(
		etcd.GetServiceAddress("logServices"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepAliveParams),
	)
	if err != nil {
		log.Println("连接到log grpc服务器失败", err.Error())
		return
	}
	log.Println("连接到logServices成功")
	return pb.NewLogServiceClient(clientConn)
}
