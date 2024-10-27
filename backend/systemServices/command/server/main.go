package main

import (
	"log"
	"systemServices/internal/config/local"
	"systemServices/internal/config/remote"
	"systemServices/internal/dao"
	"systemServices/internal/etcd"
	"systemServices/internal/handler"
	"systemServices/internal/model"
)

var err error

func init() {
	log.Println("读取配置文件")

	// 读取本地文件获取rpc地址和etcd连接信息
	if err = local.MyLocalConfig.LoadConfigFromXml("./config/config.xml"); err != nil {
		log.Println(err)
		panic("读取配置文件xml错误" + err.Error())
	}
	//log.Println(2)

	// 使用本地配置文件来初始化Etcd服务器来获取配置
	if err = etcd.InitEtcdClient(); err != nil {
		panic(err)
	}
	if err := remote.MyDbConfig.Get(); err != nil {
		panic(err)
	}

	if err := remote.MyRedisConfig.Get(); err != nil {
		panic(err)
	}

	dao.InitMysqlServer() // 初始化主数据库

	if err = dao.Db.Model(model.SiteSetting{}).AutoMigrate(); err != nil {
		panic(err)
	}

}

func main() {
	//serverAddr := fmt.Sprintf("%s:%d", config.MyLocalConfig.RpcConfig.ListenAddr, config.MyLocalConfig.RpcConfig.ListenPort)
	//log.Println(local.MyLocalConfig)
	//log.Println(remote.MyRedisConfig)
	//log.Println(remote.MyDbConfig)
	go etcd.RegisterService2Etcd(86400)
	handler.RunGRPCServer()

}
