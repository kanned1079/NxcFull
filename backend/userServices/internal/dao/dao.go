package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"userServices/internal/config/remote"
)

const (
	dbUser  = "nxc"          // 使用哪个用户来连接数据库
	dbPass  = "Password1!"   // 用户的密码
	dbProto = "tcp"          // 连接的协议
	dbHost  = "ikanned.com"  // 数据库主机的IP或域名
	dbPort  = "24407"        // 数据库连接端口
	dbName  = "nxc_networks" // 操作哪一个数据库
)

var (
	Db  *gorm.DB
	err error
)

func InitMysqlServer() {
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			remote.MyDbConfig.Username,
			remote.MyDbConfig.Password,
			remote.MyDbConfig.Protocol,
			remote.MyDbConfig.Host,
			remote.MyDbConfig.Port,
			remote.MyDbConfig.Database),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		log.Println("初始化数据库失败 err: ", err)
		panic(err)
	} else {
		//if err := Db.AutoMigrate(model.User{}); err != nil {
		//	//log.Println("AutoMigrate Failure", err)
		//	panic("AutoMigrate Failure" + err.Error())
		//}
		//if err := Db.AutoMigrate(model.Auth{}); err != nil {
		//	//log.Println("AutoMigrate Failure", err)
		//	panic("AutoMigrate Failure" + err.Error())
		//}
		//if err := Db.AutoMigrate(model.TwoFA{}); err != nil {
		//	//log.Println("AutoMigrate Failure", err)
		//	panic("AutoMigrate Failure" + err.Error())
		//}
		log.Println("初始化数据库成功")
	}
}

//func init() {
//	log.Println("初始化Mysql")
//	log.Println("初始化Redis")
//	InitRedisServer() // 初始化Redis
//}
