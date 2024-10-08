package model

import (
	"gorm.io/gorm"
	"time"
)

type Orders struct {
	Id             int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 订单ID
	UserId         int64          `json:"user_id"`                              // 用户ID
	Email          string         `json:"email"`                                // 用户邮箱地址
	PlanId         int64          `json:"plan_id"`                              // 选择的订阅计划ID
	PaymentMethod  string         `json:"payment_method"`                       // 付款方式
	Period         string         `json:"period"`                               // 订阅周期id
	CouponId       int64          `json:"coupon_id"`                            // 如使用优惠券的id
	Status         int64          `json:"status" gorm:"COMMENT:'1新购2续费3升级'"`    // 1新购2续费3升级
	IsSuccess      bool           `json:"is_success"`                           // 是否支付成功
	FailureReason  string         `json:"failure_reason"`                       // 如果支付失败，记录失败原因
	DiscountAmount float64        `json:"discount_amount"`                      // 抵折金额
	Amount         float64        `json:"amount"`                               // 订单金额
	PaidAt         *time.Time     `json:"paid_at"`                              // 支付时间
	CreatedAt      time.Time      `json:"created_at"`                           // 下单时间
	UpdatedAt      time.Time      `json:"updated_at"`                           // 更新时间
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`                           // 软删除时间
}

func (Orders) TableName() string {
	return "x_orders"
}
