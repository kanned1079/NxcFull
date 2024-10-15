package main

import (
	"log"
	"service-configurator-backend/internal/etcd"
	"service-configurator-backend/internal/routers"
)

func init() {
	if err := etcd.InitEtcdCli(); err != nil {
		log.Println(err)
		panic("Etcd服务器连接失败")
	} else {
		log.Println("Etcd服务器连接成功")
	}

	//etcd.RegisterConfig2Etcd("mysql1", "aaaaaa")
}

func main() {
	routers.StartRoute()

	defer func() {
		if err := etcd.EtcdCli.Close(); err != nil {
			log.Println(err)
		}
	}()
}
