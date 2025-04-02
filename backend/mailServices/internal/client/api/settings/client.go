package settings

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	pb "mailServices/internal/client/api/settings/proto"
	"mailServices/internal/etcd"
	"time"
)

func NewSettingClient() (client pb.SettingsServiceClient) {
	log.Println("NewSettingClient")
	keepAliveParams := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             2 * time.Second,
		PermitWithoutStream: true,
	}
	clientConn, err := grpc.NewClient(
		etcd.GetServiceAddress("settingServices"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepAliveParams),
	)
	if err != nil {
		log.Println("连接到user grpc服务器失败", err.Error())
		return
	}
	log.Println("settingServices")
	return pb.NewSettingsServiceClient(clientConn)
}
