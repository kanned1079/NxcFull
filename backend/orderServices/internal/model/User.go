package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id           int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 用户id
	InviteUserID string         `json:"invite_user_id"`                       // 邀请人id
	GroupId      int64          `json:"group_id"`                             // 用戶組id
	Name         string         `json:"name" `                                // 用户名
	Email        string         `json:"email" gorm:"unique"`                  // 登录邮箱
	IsAdmin      bool           `json:"isAdmin"`                              // 是否是管理员
	IsStaff      bool           `json:"is_staff"`                             // 是否是员工
	Balance      float32        `json:"balance"`                              // 余额
	LastLogin    *time.Time     `json:"last_login"`                           // 上一次登录时间
	LastLoginIp  string         `json:"last_login_ip"`                        // 上一次登录Ip
	CreatedAt    time.Time      `json:"createdAt"`                            // 账户创建日期
	UpdatedAt    time.Time      `json:"updatedAt"`                            // 更新日期
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`                            // 账户删除日期
}

func (User) TableName() string {
	return "x_users"
}
