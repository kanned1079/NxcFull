package main

import (
	"documentServices/internal/config/local"
	"documentServices/internal/config/remote"
	"documentServices/internal/dao"
	"documentServices/internal/etcd"
	"documentServices/internal/handler"
	"log"
)

var err error

func init() {
	log.Println("读取配置文件")

	// 读取本地文件获取rpc地址和etcd连接信息
	if err = local.MyLocalConfig.LoadConfigFromXml("./config/config.xml"); err != nil {
		log.Println(err)
		panic("读取配置文件xml错误" + err.Error())
	}

	// 使用本地配置文件来初始化Etcd服务器来获取配置
	if err = etcd.InitEtcdClient(); err != nil {
		panic(err)
	}
	if err = remote.MyDbConfig.Get(); err != nil {
		panic(err)
	}

	if err = remote.MyRedisConfig.Get(); err != nil {
		panic(err)
	}

	dao.InitMysqlServer() // 初始化主数据库
	dao.InitRedisServer() // 初始化redis

	//if err = dao.Db.Model(&model.Document{}).AutoMigrate(&model.Document{}); err != nil {
	//	panic(err)
	//}

}

func main() {
	go etcd.RegisterService2Etcd(86400)
	handler.RunGRPCServer()

}
