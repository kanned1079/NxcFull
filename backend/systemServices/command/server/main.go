package main

import (
	"NxcFull/backend/systemServices/config"
	"NxcFull/backend/systemServices/internal/dao"
	"NxcFull/backend/systemServices/internal/etcd"
	"NxcFull/backend/systemServices/internal/handler"
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
	dao.InitMysqlServer()
}

func main() {
	go etcd.RegisterService2Etcd("api.ikanned.com:22379", 60, "subscribeServices", "localhost:50005")
	handler.RunGRPCServer()
}
