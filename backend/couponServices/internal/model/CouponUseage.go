package model

import "time"

type CouponUsage struct {
	Id        int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 使用记录ID
	UserId    int64     `json:"user_id"`                              // 用户ID
	CouponId  int64     `json:"coupon_id"`                            // 优惠券ID
	UsedAt    time.Time `json:"used_at"`                              // 使用时间
	CreatedAt time.Time `json:"created_at"`                           // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                           // 更新时间
	Coupon    Coupon    `gorm:"foreignKey:CouponId;references:Id"`    // 一个Coupon可以有多个CouponUseage
	//User      User      `gorm:"foreignKey:UserId;references:UserId"`
}

func (CouponUsage) TableName() string {
	return "x_coupon_usage"
}
