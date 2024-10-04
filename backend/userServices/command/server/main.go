package main

import (
	"NxcFull/backend/userServices/internal/dao"
	"NxcFull/backend/userServices/internal/handler"
)

func main() {
	dao.InitMysqlServer() // 初始化主数据库
	handler.RunGRPCServer()
	// 创建grpc服务器
	//grpcServer := grpc.NewServer()
	//// 注册gRPC服务
	//userServices := service.NewUserServices()
}
