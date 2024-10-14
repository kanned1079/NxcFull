package main

import (
	"log"
	"mailServices/internal/config"
	"mailServices/internal/dao"
	"mailServices/internal/etcd"
	"mailServices/internal/handler"
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
	dao.InitRedisServer()
}

func main() {
	go etcd.RegisterService2Etcd("api.ikanned.com:22379", 60, "mailServices", "localhost:50008")
	handler.RunGRPCServer()
}
