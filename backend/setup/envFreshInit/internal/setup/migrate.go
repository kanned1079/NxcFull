package setup

import (
	"envFreshInit/internal/dao"
	"envFreshInit/internal/model"
)

func StartAutoMigrate() error {
	var err error
	err = dao.Db.AutoMigrate(
		&model.User{},
		&model.Auth{},
		&model.TwoFA{},
		&model.Avatar{},
		&model.Chat{},
		&model.Ticket{},

		&model.Plan{},

		&model.Coupon{},
		&model.CouponUsage{},

		&model.Keys{},
		&model.Group{},
		//&model.Orders{},
		//&model.ActiveOrders{},
		//&model.ActivateRecord{},

		&model.Document{},
		&model.ApiLog{},
		&model.PublicNotices{},
		&model.PaymentSettings{},
		&model.SiteSetting{},
	)
	err = dao.Db.AutoMigrate(&model.Orders{})
	err = dao.Db.AutoMigrate(&model.ActiveOrders{})
	err = dao.Db.AutoMigrate(&model.ActivateRecord{})
	return err
}
