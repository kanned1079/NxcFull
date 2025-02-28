package main

import (
	"gateway/internal/config"
	"gateway/internal/config/remote"
	"gateway/internal/dao"
	"gateway/internal/etcd"
	"gateway/internal/handler"
	"gateway/internal/middleware"
	"gateway/internal/routers"
	"github.com/gin-gonic/gin"
	"log"
)

//var GrpcClients grpc2.Clients

// var grpcClient pb.UserServiceClient
var err error

//var userServiceClient *pb.UserServiceClient

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
	etcd.InitEtcdClient() // 初始化etcd服务器
	//GrpcClients = grpc2.NewClients() // 初始化所有的rpc客户端连接
	//userServiceClient = pb.NewUserServiceClient(GrpcClients.UserServiceClient)
	if err = remote.MyRedisConfig.Get(); err != nil {
		panic(err)
	}
	dao.InitRedisServer()

	if err = config.GetServerEnv(); err != nil {
		log.Println("获取运行环境参数错误")
		panic(err)
	}

}

func main() {

	// 使用 WaitGroup 来等待所有 goroutine
	//var wg sync.WaitGroup
	//
	////wg.Add(1)
	////go func() {
	////	defer wg.Done()
	////	go middleware.StartApiAccessCountCron()
	////}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	go routers.StartApiGateways()
	//}()
	//
	//wg.Wait()
	go middleware.StartApiAccessCountCron()

	go handler.StartLogFlushLog(60)

	MyApp1 := routers.NewGatewayApp(1, gin.TestMode)

	//routers.StartApiGateways()

	MyApp1.StartApiGateways()

}
