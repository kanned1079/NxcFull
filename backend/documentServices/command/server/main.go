package main

import (
	"NxcFull/backend/documentServices/internal/config"
	"NxcFull/backend/documentServices/internal/dao"
	"NxcFull/backend/documentServices/internal/etcd"
	"NxcFull/backend/documentServices/internal/handler"
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

}

func main() {
	dao.InitMysqlServer() // 初始化主数据库
	go etcd.RegisterService2Etcd("api.ikanned.com:22379", 60, "documentServices", "localhost:50002")
	handler.RunGRPCServer()

}
