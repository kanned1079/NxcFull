package notice

import (
	"NxcFull/backend/gateway/internal/etcd"
	pb "NxcFull/backend/gateway/internal/grpc/api/notice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewNoticeClient() (client pb.NoticeServiceClient) {
	log.Println("NewNoticeClient")
	clientConn, err := grpc.NewClient(etcd.GetServiceAddress("noticeServices"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接到user grpc服务器失败", err.Error())
		return
	}
	log.Println("连接到noticeServices成功")
	return pb.NewNoticeServiceClient(clientConn)
}
