package main

import (
	"log"
	"systemServices/internal/config"
	"systemServices/internal/dao"
	"systemServices/internal/etcd"
	"systemServices/internal/handler"
)

var err error

func init() {
	//config.SaveConfigToYaml(&config.EnvConfig, "./config/config.yaml")\
	log.Println("读取配置文件")
	if config.EnvConfig, err = config.LoadConfigFromYaml("./config/config.yaml"); err != nil {
		log.Println("读取配置文件出错", err)
		panic(err)
	}
	dao.InitMysqlServer()
}

func main() {
	go etcd.RegisterService2Etcd("api.ikanned.com:22379", 3600, "settingServices", "localhost:50007")
	handler.RunGRPCServer()
}
