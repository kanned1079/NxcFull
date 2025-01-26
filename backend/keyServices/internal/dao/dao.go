package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"keyServices/internal/config/remote"
	"log"
)

var (
	Db *gorm.DB
	//Tx  *gorm.Tx
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
		log.Println("初始化数据库成功")
	}
	//if err := Db.AutoMigrate(&model.Orders{}); err != nil {
	//	panic(err)
	//}
	//if err := Db.AutoMigrate(&model.ActiveOrders{}); err != nil {
	//	panic(err)
	//}
	//if err := Db.AutoMigrate(&model.Coupon{}); err != nil {
	//	panic(err)
	//}
	//if err := Db.AutoMigrate(&model.CouponUsage{}); err != nil {
	//	panic(err)
	//}
	//if err := Db.AutoMigrate(&model.Plan{}); err != nil {
	//	panic(err)
	//}
	//if err := Db.AutoMigrate(&model.User{}); err != nil {
	//	panic(err)
	//}
	//if err := Db.AutoMigrate(&model.Keys{}); err != nil {
	//	panic(err)
	//}
	//if err := Db.AutoMigrate(&model.ActivateRecord{}); err != nil {
	//	panic(err)
	//}

}
