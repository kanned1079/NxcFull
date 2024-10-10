package setting

import (
	"NxcFull/backend/gateway/internal/etcd"
	pb "NxcFull/backend/gateway/internal/grpc/api/setting/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewSettingClient() (client pb.SettingsServiceClient) {
	log.Println("NewSettingClient")
	clientConn, err := grpc.NewClient(etcd.GetServiceAddress("settingServices"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接到user grpc服务器失败", err.Error())
		return
	}
	log.Println("settingServices")
	return pb.NewSettingsServiceClient(clientConn)
}
