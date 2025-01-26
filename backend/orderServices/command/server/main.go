package main

import (
	"log"
	"orderServices/internal/config/local"
	"orderServices/internal/config/remote"
	"orderServices/internal/dao"
	"orderServices/internal/etcd"
	"orderServices/internal/handler"
	"orderServices/internal/payment"
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

	dao.InitMysqlServer() // 初始化主数据库
	dao.InitRedisServer() // 初始化Redis服务器
	dao.InitMq()

	// 初始化支付方式
	payment.InitPaymentConf()

	payment.PaymentInstanceRoot.Initial() // 初始化所有被启用的支付方式

	//if err = dao.Db.Model(model.User{}).AutoMigrate(); err != nil {
	//	panic(err)
	//}
	//if err = dao.Db.Model(model.Coupon{}).AutoMigrate(); err != nil {
	//	panic(err)
	//}
	//if err = dao.Db.Model(model.CouponUsage{}).AutoMigrate(); err != nil {
	//	panic(err)
	//}
	//if err = dao.Db.Model(model.Plan{}).AutoMigrate(); err != nil {
	//	panic(err)
	//}
	//if err = dao.Db.Model(model.ActiveOrders{}).AutoMigrate(); err != nil {
	//	panic(err)
	//}
	//if err = dao.Db.Model(model.Orders{}).AutoMigrate(); err != nil {
	//	panic(err)
	//}

}

func main() {
	go etcd.RegisterService2Etcd(86400)
	handler.RunGRPCServer()

}
