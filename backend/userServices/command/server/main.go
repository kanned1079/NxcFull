package main

import (
	"log"
	"userServices/internal/config"
	"userServices/internal/dao"
	"userServices/internal/etcd"
	"userServices/internal/handler"
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
