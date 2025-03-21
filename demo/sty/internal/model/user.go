package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int64          `json:"id"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func TableName() string {
	return ""
}
