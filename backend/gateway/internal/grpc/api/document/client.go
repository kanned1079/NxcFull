package document

import (
	"NxcFull/backend/gateway/internal/etcd"
	pb "NxcFull/backend/gateway/internal/grpc/api/document/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewDocumentClient() (client pb.DocumentServiceClient) {
	log.Println("NewDocumentClient")
	clientConn, err := grpc.NewClient(etcd.GetServiceAddress("documentServices"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接到user grpc服务器失败", err.Error())
		return
	}
	log.Println("连接到documentServices成功")
	return pb.NewDocumentServiceClient(clientConn)
}
