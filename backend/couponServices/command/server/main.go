package main

import (
	"NxcFull/backend/couponServices/internal/config"
	"NxcFull/backend/couponServices/internal/dao"
	"NxcFull/backend/couponServices/internal/etcd"
	"NxcFull/backend/couponServices/internal/handler"
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
	go etcd.RegisterService2Etcd("api.ikanned.com:22379", 60, "couponServices", "localhost:50006")
	handler.RunGRPCServer()

}
