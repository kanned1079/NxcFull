package main

import (
	"fmt"
	"log"
	rpcClient "mailServices/internal/client"
	"mailServices/internal/config/local"
	"mailServices/internal/config/remote"
	"mailServices/internal/dao"
	"mailServices/internal/etcd"
	"mailServices/internal/handler"
	"os"
)

var err error

func listCurrentDirectory() {
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// 打印当前工作目录
	fmt.Println("Current directory:", dir)

	// 打开当前目录
	files, err := os.ReadDir(dir + "/config")
	if err != nil {
		log.Fatal(err)
	}

	// 遍历并打印目录中的每一个文件/文件夹
	fmt.Println("Contents of the current directory:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func init() {
	log.Println("读取配置文件")

	listCurrentDirectory()

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

	rpcClient.GrpcClient = rpcClient.NewClients()

	dao.InitMysqlServer() // 初始化主数据库
	dao.InitRedisServer() // 初始化Redis

}

func main() {
	go etcd.RegisterService2Etcd(86400)
	handler.RunGRPCServer()

}
