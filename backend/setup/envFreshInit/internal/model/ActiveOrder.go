package model

import (
	"gorm.io/gorm"
	"time"
)

type ActiveOrders struct {
	Id             int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 订阅记录 ID
	OrderId        string         `json:"order_id" gorm:"unique;index"`         // 订单号，唯一索引
	UserId         int64          `json:"user_id" gorm:"index"`                 // 用户 ID，普通索引
	GroupId        int64          `json:"group_id" gorm:"index"`                // 权限组，普通索引
	Email          string         `json:"email"`                                // 邮箱地址
	PlanId         int64          `json:"plan_id" gorm:"index"`                 // 订阅计划 ID，普通索引
	KeyId          int64          `json:"key_id" gorm:"unique;index"`           // 密钥 ID，唯一索引
	IsActive       bool           `json:"is_active"`                            // 订阅是否有效
	IsUsed         bool           `json:"is_used"`                              // 订阅是否被使用
	IsValid        bool           `json:"is_valid"`                             // 该密钥是否有效
	CreatedAt      time.Time      `json:"created_at"`                           // 创建日期
	ExpirationDate *time.Time     `json:"expiration_date"`                      // 到期日期
	UpdatedAt      time.Time      `json:"updated_at"`                           // 更新时间
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`                           // 删除时间
	User           User           `gorm:"foreignKey:UserId;references:Id"`      // 用户外键关联
	KeyInfo        Keys           `gorm:"foreignKey:KeyId;references:Id"`       // 密钥外键关联
}

func (ActiveOrders) TableName() string {
	return "x_active_orders"
}
