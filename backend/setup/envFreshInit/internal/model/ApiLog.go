package model

import (
	"gorm.io/gorm"
	"time"
)

type ApiLog struct {
	Id           int64          `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Level        string         `json:"level"`         // 日誌等級
	StatusCode   int32          `json:"status_code"`   // 請求狀態碼
	RequestAt    time.Time      `json:"request_at"`    // 請求時間
	ResponseTime float32        `json:"response_time"` // 響應時間
	ClientIp     string         `json:"client_ip"`     // 客戶端IP
	Method       string         `json:"method"`        // 請求方法
	Path         string         `json:"path"`          // 請求路徑
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func (ApiLog) TableName() string {
	return "x_api_log"
}
