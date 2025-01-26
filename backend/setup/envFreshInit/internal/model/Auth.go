package model

import (
	"gorm.io/gorm"
	"time"
)

type Auth struct {
	Id       int64  `json:"id" gorm:"primary_key;"`   // 用户Id
	Email    string `json:"email" gorm:"not null"`    // 用户名
	Password string `json:"password" gorm:"not null"` // 用户密码

	CreatedAt time.Time      `json:"createdAt"` // 账户创建日期
	UpdatedAt time.Time      `json:"updatedAt"` // 更新日期
	DeletedAt gorm.DeletedAt `json:"deletedAt"` // 账户删除日期
}

func (Auth) TableName() string {
	return "x_auth"
}
