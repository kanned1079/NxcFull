package main

import (
	"NxcFull/backend/groupServices/internal/config"
	"NxcFull/backend/groupServices/internal/dao"
	"NxcFull/backend/groupServices/internal/etcd"
	"NxcFull/backend/groupServices/internal/handler"
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
	go etcd.RegisterService2Etcd("api.ikanned.com:22379", 60, "groupServices", "localhost:50004")
	handler.RunGRPCServer()
}
