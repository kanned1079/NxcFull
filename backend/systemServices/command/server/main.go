package main

import (
	"log"
	"systemServices/internal/config/local"
	"systemServices/internal/config/remote"
	"systemServices/internal/dao"
	"systemServices/internal/etcd"
	"systemServices/internal/handler"
	"systemServices/internal/model"
	"systemServices/internal/services"
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
	dao.InitRedisServer()

	if err = dao.Db.AutoMigrate(&model.SiteSetting{}); err != nil {
		panic(err)
	}

	// 刷新缓存
	if err = services.MakeSettingsCache(); err != nil {
		log.Panicln(err.Error())
	}
	// 自动迁移

}

func main() {
	go etcd.RegisterService2Etcd(86400)
	handler.RunGRPCServer()
}
