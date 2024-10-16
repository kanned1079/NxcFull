package main

import (
	"log"
	"userServices/internal/config/local"
	"userServices/internal/config/remote"
	"userServices/internal/dao"
	"userServices/internal/etcd"
	"userServices/internal/handler"
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

}

func main() {
	//serverAddr := fmt.Sprintf("%s:%d", config.MyLocalConfig.RpcConfig.ListenAddr, config.MyLocalConfig.RpcConfig.ListenPort)
	//log.Println(local.MyLocalConfig)
	//log.Println(remote.MyRedisConfig)
	//log.Println(remote.MyDbConfig)
	go etcd.RegisterService2Etcd(3600)
	handler.RunGRPCServer()

}
