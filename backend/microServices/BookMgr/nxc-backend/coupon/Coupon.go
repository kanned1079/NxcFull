package coupon

import (
	"gorm.io/gorm"
	"time"
)

type Coupon struct {
	Id           int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 优惠券Id
	Name         string         `json:"name"`                                 // 优惠券名称
	Code         string         `json:"code"`                                 // 优惠码
	Enabled      bool           `json:"enabled"`                              // 启用优惠券
	PercentOff   float64        `json:"percent_off"`                          // 优惠百分比
	StartTime    *time.Time     `json:"start_time"`                           // 可用起始时间
	EndTime      *time.Time     `json:"end_time"`                             // 过期时间
	PerUserLimit int64          `json:"per_user_limit"`                       // 限制每个用户使用次数 0为不限制
	Capacity     int64          `json:"capacity"`                             // 总容量限制
	Residue      int64          `json:"residue"`                              // 剩余数量
	PlanLimit    int64          `json:"plan_limit"`                           // 指定订阅可用 0为不使用
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func (Coupon) TableName() string {
	return "x_coupons"
}
