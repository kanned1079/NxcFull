package main

import (
	"NxcFull/backend/mailServices/internal/config"
	"NxcFull/backend/mailServices/internal/etcd"
	"NxcFull/backend/mailServices/internal/handler"
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
	go etcd.RegisterService2Etcd("api.ikanned.com:22379", 60, "mailServices", "localhost:50002")
	handler.RunGRPCServer()
}
