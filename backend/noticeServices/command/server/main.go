package main

import (
	"NxcFull/backend/noticeServices/internal/config"
	"NxcFull/backend/noticeServices/internal/dao"
	"NxcFull/backend/noticeServices/internal/etcd"
	"NxcFull/backend/noticeServices/internal/handler"
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
	go etcd.RegisterService2Etcd("api.ikanned.com:22379", 3600, "noticeServices", "localhost:50003")
	handler.RunGRPCServer()
}
