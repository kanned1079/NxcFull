package main

import (
	rpcClient "keyServices/internal/client"
	"keyServices/internal/config/local"
	"keyServices/internal/config/remote"
	"keyServices/internal/dao"
	"keyServices/internal/etcd"
	"keyServices/internal/handler"
	"log"
	"sync"
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
	if err := remote.MyDbConfig.Get(); err != nil {
		panic(err)
	}
	if err := remote.MyRedisConfig.Get(); err != nil {
		panic(err)
	}
	if err := remote.MyMqConfig.Get(); err != nil {
		panic(err)
	}

	rpcClient.GrpcClient = rpcClient.NewClients()

	dao.InitMysqlServer() // 初始化主数据库
	dao.InitRedisServer() // 初始化Redis服务器
	//dao.InitMq()          // 初始化RabbitMQ消息队列
	dao.InitMq()
}

func main() {

	// 使用 WaitGroup 来等待所有 goroutine
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		go etcd.RegisterService2Etcd(86400)
	}()

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	//go etcd.RegisterService2Etcd(86400)
	//	//go consumer.StartConsumeOrders()
	//}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		//go etcd.RegisterService2Etcd(86400)
		//go consumer.StartConsumeOrders()
		handler.RunGRPCServer()
	}()

	wg.Wait()

	//go etcd.RegisterService2Etcd(86400)
	//go consumer.StartConsumeOrders()
	//handler.RunGRPCServer()

}
