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
	Sort      int            `json:"sort"`
	Show      bool           `json:"show"`
	CreatedAt time.Time      `json:"create_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at"`
}

func (Document) TableName() string {
	return "x_document"
}
