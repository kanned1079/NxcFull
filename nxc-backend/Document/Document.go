package Document

import (
	"gorm.io/gorm"
	"time"
)

type Document struct {
	Id       int64          `json:"id" gorm:"primary_key"`
	Language string         `json:"language"`
	Category string         `json:"category"`
	Title    string         `json:"title"`
	Body     string         `json:"body"`
	Sort     int            `json:"sort"`
	Show     bool           `json:"show"`
	CreateAt time.Time      `json:"create_at"`
	UpdateAt time.Time      `json:"update_at"`
	DeleteAt gorm.DeletedAt `json:"delete_at"`
}

func (Document) TableName() string {
	return "x_document"
}
