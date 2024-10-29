package model

import (
	"gorm.io/gorm"
	"time"
)

type Ticket struct {
	gorm.Model
	Id        int64          `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	UserId    int64          `json:"user_id"`
	Subject   string         `json:"subject"`
	Urgency   int8           `json:"urgency"` // 1低 2中 3高
	Status    int32          `json:"status"`  // 1已创建 2已经关闭 3等待回复
	LastReply *time.Time     `json:"last_reply"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Ticket) TableName() string {
	return "x_ticket"
}
