package model

import (
	"gorm.io/gorm"
	"time"
)

type Orders struct {
	gorm.Model
	Id             int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 订单数据Id
	OrderId        string         `json:"order_id" gorm:"unique"`               // 订单号码
	UserId         int64          `json:"user_id"`                              // 用户ID
	GroupId        int64          `json:"group_id"`                             // 群组id
	Email          string         `json:"email"`                                // 用户邮箱地址
	PlanId         int64          `json:"plan_id"`                              // 选择的订阅计划ID
	PlanName       string         `json:"plan_name"`                            // 订阅计划名
	PaymentMethod  string         `json:"payment_method"`                       // 付款方式
	Period         string         `json:"period"`                               // 订阅周期id
	CouponId       int64          `json:"coupon_id"`                            // 如使用优惠券的id
	CouponName     string         `json:"coupon_name"`                          // 如果使用了优惠券这里为优惠券的名称
	Status         int8           `json:"status" gorm:"COMMENT:'1新购2续费3升级'"`    // 1新购2续费3升级
	IsFinished     bool           `json:"is_finished"`                          // 该订单是否完成
	IsSuccess      bool           `json:"is_success"`                           // 是否支付成功
	FailureReason  string         `json:"failure_reason"`                       // 如果支付失败，记录失败原因
	DiscountAmount float64        `json:"discount_amount"`                      // 抵折金额
	Amount         float64        `json:"amount"`                               // 订单金额
	Price          float64        `json:"price"`                                // 这是订单原始的价格
	PaidAt         *time.Time     `json:"paid_at"`                              // 支付时间
	CreatedAt      time.Time      `json:"created_at"`                           // 下单时间
	UpdatedAt      time.Time      `json:"updated_at"`                           // 更新时间
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`                           // 软删除时间
	User           User           `gorm:"foreignKey:UserId;references:Id"`      // 多对一关联，User 表中的用户
	//ActiveOrder    ActiveOrders   `gorm:"foreignKey:OrderId;references:OrderId"` // 外键关联到 Orders 表中的 OrderId
}

func (Orders) TableName() string {
	return "x_orders"
}
