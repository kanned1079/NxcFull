package coupon

import (
	"NxcFull/backend/gateway/internal/etcd"
	pb "NxcFull/backend/gateway/internal/grpc/api/coupon/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewCouponServiceClient() (client pb.CouponServiceClient) {
	log.Println("NewCouponServiceClient")
	clientConn, err := grpc.NewClient(etcd.GetServiceAddress("couponServices"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接到user grpc服务器失败", err.Error())
		return
	}
	log.Println("couponServices")
	return pb.NewCouponServiceClient(clientConn)
}
