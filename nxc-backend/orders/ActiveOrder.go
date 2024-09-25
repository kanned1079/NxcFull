package orders

import (
	"gorm.io/gorm"
	"time"
)

type ActiveOrders struct {
	ID             int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 订阅记录 ID
	UserId         int64          `json:"user_id"`                              // 用户 ID
	GroupId        int64          `json:"group_id"`                             // 對應權限組
	Email          string         `json:"email"`                                // 邮箱地址
	PlanId         int64          `json:"plan_id"`                              // 购买的订阅计划 ID
	KeyId          int64          `json:"key_id"`                               // 对应表x_keys的密钥Id
	IsActive       bool           `json:"is_active"`                            // 订阅是否有效
	IsUsed         bool           `json:"is_used"`                              // 订阅密钥是否被使用过
	CreatedAt      time.Time      `json:"created_at"`                           // 创建日期
	ExpirationDate *time.Time     `json:"expiration_date"`                      // 到期日期
	UpdatedAt      time.Time      `json:"updated_at"`                           // 更新时间
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`                           // 删除时间
}

func (ActiveOrders) TableName() string {
	return "x_active_orders"
}
