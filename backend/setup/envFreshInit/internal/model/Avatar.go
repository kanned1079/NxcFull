package model

import (
	"gorm.io/gorm"
	"time"
)

type Avatar struct {
	Id        int64          `json:"id"`
	UserId    int64          `json:"user_id"`
	FileName  string         `json:"file_name"`
	FileData  []byte         `json:"file_data"`
	User      User           `json:"user" gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Avatar) TableName() string {
	return "x_avatars"
}
