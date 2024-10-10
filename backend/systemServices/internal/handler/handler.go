package handler

import (
	"NxcFull/backend/systemServices/api/proto"
	"NxcFull/backend/systemServices/internal/services"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

// 日志拦截器
func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("gRPC method called: %s, with request: %v", info.FullMethod, req)

	// 调用实际的 gRPC 方法
	resp, err := handler(ctx, req)

	// 返回响应和错误
	return resp, err
}

func RunGRPCServer() {
	port := 50007
	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		log.Println("设置监听端口失败", err)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			log.Println(err)
		}
	}()
	// 创建grpc服务器实例
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor), // 	启用日志拦截器
		grpc.Creds(insecure.NewCredentials()))
	// 注册反射服务 方便grpc调试工具使用
	reflection.Register(grpcServer)
	// 注册UserService到grpcServer服务器
	settingServices := services.NewSettingServices()
	proto.RegisterSettingsServiceServer(grpcServer, settingServices)
	log.Printf("gRPC 服务器正在 %v 端口监听...", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Println("启动gRPC服务器失败", err)
	}
}
