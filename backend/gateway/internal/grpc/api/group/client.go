package group

import (
	"NxcFull/backend/gateway/internal/etcd"
	pb "NxcFull/backend/gateway/internal/grpc/api/group/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewGroupClient() (client pb.GroupServiceClient) {
	log.Println("NewNoticeClient")
	clientConn, err := grpc.NewClient(etcd.GetServiceAddress("groupServices"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接到user grpc服务器失败", err.Error())
		return
	}
	log.Println("连接到noticeServices成功")
	return pb.NewGroupServiceClient(clientConn)
}
