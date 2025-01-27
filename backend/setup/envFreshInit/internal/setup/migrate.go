package setup

import (
	"envFreshInit/internal/dao"
	"envFreshInit/internal/model"
	"envFreshInit/internal/utils"
	"fmt"
	"log"
)

func migrateTable(table interface{}, tableName string) error {
	err := dao.Db.AutoMigrate(table)
	if err != nil {
		utils.ShowFailure() // 添加失败提示
		log.Printf("Error during %s table migration: %v", tableName, err)
		return err
	}
	utils.ShowOK() // 添加成功提示
	fmt.Printf("%s table migrated successfully.\n", tableName)
	return nil
}

func StartAutoMigrate() error {
	// 定义所有需要迁移的表及其对应名称
	tables := []struct {
		model interface{}
		name  string
	}{
		{&model.User{}, "User"},
		{&model.Auth{}, "Auth"},
		{&model.TwoFA{}, "TwoFA"},
		{&model.Avatar{}, "Avatar"},
		{&model.Chat{}, "Chat"},
		{&model.Ticket{}, "Ticket"},
		{&model.Plan{}, "Plan"},
		{&model.Coupon{}, "Coupon"},
		{&model.CouponUsage{}, "CouponUsage"},
		{&model.Keys{}, "Keys"},
		{&model.Group{}, "Group"},
		{&model.Document{}, "Document"},
		{&model.ApiLog{}, "ApiLog"},
		{&model.PublicNotices{}, "PublicNotices"},
		{&model.PaymentSettings{}, "PaymentSettings"},
		{&model.SiteSetting{}, "SiteSetting"},
		{&model.Orders{}, "Orders"},
		{&model.ActiveOrders{}, "ActiveOrders"},
		{&model.ActivateRecord{}, "ActivateRecord"},
	}

	// 循环迁移所有表
	for _, table := range tables {
		if err := migrateTable(table.model, table.name); err != nil {
			return err
		}
	}

	return nil
}
