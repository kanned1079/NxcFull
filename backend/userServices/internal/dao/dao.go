package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"userServices/internal/config"
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
		DSN: fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.EnvConfig.MySql.Username,
			config.EnvConfig.MySql.Password,
			config.EnvConfig.MySql.Protocol,
			config.EnvConfig.MySql.Host,
			config.EnvConfig.MySql.Port,
			config.EnvConfig.MySql.Database),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Println("初始化数据库失败 err: ", err)
		panic(err)
	} else {
		log.Println("初始化数据库成功")
	}
}

//func init() {
//	log.Println("初始化Mysql")
//	log.Println("初始化Redis")
//	InitRedisServer() // 初始化Redis
//}
