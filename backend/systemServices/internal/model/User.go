package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id           int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 用户id
	InviteUserID string         `json:"invite_user_id"`                       // 邀请人id
	InviteCode   string         `json:"invite_code"`                          // 邀请码
	GroupId      int64          `json:"group_id"`                             // 用戶組id
	Name         string         `json:"name" `                                // 用户名
	Email        string         `json:"email" gorm:"unique;not null"`         // 登录邮箱
	IsAdmin      bool           `json:"is_admin"`                             // 是否是管理员
	IsStaff      bool           `json:"is_staff"`                             // 是否是员工
	Balance      float32        `json:"balance"`                              // 余额
	LastLogin    *time.Time     `json:"last_login"`                           // 上一次登录时间
	LastLoginIp  string         `json:"last_login_ip"`                        // 上一次登录Ip
	IsTwoFA      bool           `json:"is_two_fa"`                            // 启用两步验证
	Status       bool           `json:"status"`                               // 帐户状态 true正常 false封禁
	CreatedAt    time.Time      `json:"created_at"`                           // 账户创建日期
	UpdatedAt    time.Time      `json:"updated_at"`                           // 更新日期
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`                           // 账户删除日期
	//Auth         Auth           `gorm:"foreignKey:Email;references:Email"`    // Auth 表的外键关联
	//TwoFA        TwoFA          `gorm:"foreignKey:Email;references:Email"`    // TwoFA 表的外键关联
}

func (User) TableName() string {
	return "x_users"
}
