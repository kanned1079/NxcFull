package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mailServices/internal/config/remote"
)

var (
	Db  *gorm.DB
	err error
)

// InitMysqlServer 初始化Mysql服务器
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
