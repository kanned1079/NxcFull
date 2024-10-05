package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type SiteSetting struct {
	ID       uint            `gorm:"primaryKey" json:"id"`
	Category string          `gorm:"size:50;not null;index" json:"category"`
	Key      string          `gorm:"size:100;not null;index" json:"key"`
	Value    json.RawMessage `gorm:"type:json;not null" json:"value"`
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
		fmt.Println("Error creating setting:", err)
	} else {
		fmt.Println("Setting created successfully")
	}
}
