package model

import (
	"gorm.io/gorm"
	"time"
)

type Document struct {
	Id        int64          `json:"id" gorm:"primary_key"`
	Language  string         `json:"language"`
	Category  string         `json:"category"`
	Title     string         `json:"title"`
	Body      string         `json:"body" gorm:"type:TEXT"`
	Sort      int64          `json:"sort"`
	Show      bool           `json:"show"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Document) TableName() string {
	return "x_document"
}
