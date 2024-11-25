package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// SiteSetting 这是数据库中的定义
type SiteSetting struct {
	Id        uint            `gorm:"primaryKey" json:"id"`
	Category  string          `gorm:"size:50;not null;index" json:"category"`
	Key       string          `gorm:"size:100;not null;index" json:"key"`
	Value     json.RawMessage `gorm:"type:json;not null" json:"value"`
	CreatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt  `json:"deleted_at"`
}

func (SiteSetting) TableName() string {
	return "x_site_settings"
}

func CreateSetting(db *gorm.DB) {
	value := json.RawMessage("my-app")
	setting := SiteSetting{
		Category: "test",
		Key:      "app_name",
		Value:    value,
	}
	if err := db.Create(&setting).Error; err != nil {
		fmt.Println("Error creating settings:", err)
	} else {
		fmt.Println("Setting created successfully")
	}
}
