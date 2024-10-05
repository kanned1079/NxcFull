package main

import (
	"NxcFull/backend/userServices/internal/dao"
	"NxcFull/backend/userServices/internal/etcd"
	"NxcFull/backend/userServices/internal/handler"
)

func main() {
	dao.InitMysqlServer() // 初始化主数据库
	go etcd.RegisterService2Etcd("api.ikanned.com:22379", 60, "userServices", "localhost:50001")
	handler.RunGRPCServer()

}
