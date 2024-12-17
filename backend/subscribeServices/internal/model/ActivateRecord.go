package model

import (
	"gorm.io/gorm"
	"time"
)

type ActivateRecord struct {
	Id            int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId        int64          `json:"user_id"`
	Email         string         `json:"email"`
	OrderId       string         `json:"order_id"`
	KeyId         int64          `json:"key_id"`
	RequestAt     *time.Time     `json:"request_at"`
	ClientVersion string         `json:"client_version"`
	OsType        string         `json:"os_type"`
	Remark        string         `json:"remark" gorm:"type:TEXT"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	// 外鍵關聯
	User         User         `gorm:"foreignKey:UserId;references:Id"`
	KeyInfo      Keys         `gorm:"foreignKey:KeyId;references:Id"`
	ActiveOrders ActiveOrders `gorm:"foreignKey:OrderId;references:OrderId"`
	/*
	  string email = 1; // 用户邮箱
	  string password = 2;  // 密码
	  string key = 3; // 密钥key
	  int64 request_at = 4; // 时间戳
	  string client_version = 5;  // 客户端版本
	  string os_type = 6; // 操作系统类型
	*/
}

func (ActivateRecord) TableName() string {
	return "x_activate_record"
}
