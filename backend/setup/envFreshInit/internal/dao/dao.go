package dao

import (
	"envFreshInit/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
	//Tx  *gorm.Tx
	err error
)

func InitMysqlServer() error {
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			//remote.MyDbConfig.Username,
			config.LocalCfg.Db.Username,
			config.LocalCfg.Db.Password,
			"tcp",
			config.LocalCfg.Db.Host,
			config.LocalCfg.Db.Port,
			config.LocalCfg.Db.Database),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	return err
}
