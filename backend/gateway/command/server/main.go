package main

import (
	grpc2 "NxcFull/backend/gateway/internal/grpc"
	pb "NxcFull/backend/gateway/internal/grpc/api/proto"
	"NxcFull/backend/gateway/internal/routers"
)

var GrpcClients grpc2.Clients

// var grpcClient pb.UserServiceClient
var err error
var userServiceClient *pb.UserServiceClient

func init() {
	//var grpcConn *grpc.ClientConn
	//if grpcConn, err = grpc.Dial("localhost:50001", grpc.WithInsecure()); err != nil {
	//	log.Println("grpc连接失败", err)
	//}
	//defer func() {
	//	if err := grpcConn.Close(); err != nil {
	//		log.Println("grpc关闭连接失败", err)
	//	}
	//}()
	//grpcClient = pb.NewUserServiceClient(grpcConn)
	GrpcClients = grpc2.NewClients()
	//userServiceClient = pb.NewUserServiceClient(GrpcClients.UserServiceClient)
}

func main() {
	routers.StartApiGateways()
}
