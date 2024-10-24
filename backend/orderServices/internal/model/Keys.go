package model

import (
	"gorm.io/gorm"
	"time"
)

type Keys struct {
	Id        int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Key       string         `json:"key" gorm:"unique"`
	ClientId  string         `json:"client_id" gorm:"unique"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Keys) TableName() string {
	return "x_keys"
}
