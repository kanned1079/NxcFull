package main

import (
	"NxcFull/nxc-backend/Document"
	"NxcFull/nxc-backend/coupon"
	"NxcFull/nxc-backend/dao"
	"NxcFull/nxc-backend/finance"
	"NxcFull/nxc-backend/keys"
	"NxcFull/nxc-backend/orders"
	"NxcFull/nxc-backend/privilege"
	"NxcFull/nxc-backend/publicNotice"
	"NxcFull/nxc-backend/routers"
	"NxcFull/nxc-backend/settings"
	"NxcFull/nxc-backend/subscribePlan"
	"NxcFull/nxc-backend/user"
)

type adminConfig struct {
	Name     string `yaml:"name"`
	Email    string `yaml:"email"`
	IsAdmin  bool   `yaml:"is_admin"`
	Password string `yaml:"password"`
}

func main() {
	// 自动迁移
	if err := dao.Db.AutoMigrate(&settings.SiteSetting{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(&user.User{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(&user.Auth{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(&finance.Subscribes{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(&publicNotice.PublicNotices{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(subscribePlan.Plan{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(Document.Document{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(coupon.Coupon{}); err != nil {
		panic("failed to migrate database")
	}
	if err := dao.Db.AutoMigrate(coupon.CouponUsage{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(privilege.Group{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(orders.Orders{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(orders.ActiveOrders{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(keys.Keys{}); err != nil {
		panic("failed to migrate database")
	}

	routers.StartReq() // 此函数中定义了gin实现路由
}
