package model

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	Id        int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Group) TableName() string {
	return "x_group"
}
