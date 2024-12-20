package model

import (
	"gorm.io/gorm"
	"time"
)

type TwoFA struct {
	Id        int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Email     string         `json:"email" gorm:"not null"`
	TwoFAKey  string         `json:"two_fa_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (TwoFA) TableName() string {
	return "x_two_fa"
}
