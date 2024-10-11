package main

import (
	"NxcFull/backend/userServices/internal/config"
	"NxcFull/backend/userServices/internal/dao"
	"NxcFull/backend/userServices/internal/etcd"
	"NxcFull/backend/userServices/internal/handler"
	"log"
)

var err error

func init() {
	//config.SaveConfigToYaml(&config.EnvConfig, "./config/config.yaml")\
	log.Println("读取配置文件")
	if config.EnvConfig, err = config.LoadConfigFromYaml("./config/config.yaml"); err != nil {
		log.Println("读取配置文件出错", err)
		return
	}
	dao.InitMysqlServer() // 初始化主数据库

}

func main() {
	go etcd.RegisterService2Etcd("api.ikanned.com:22379", 3600, "userServices", "localhost:50001")
	handler.RunGRPCServer()

}
