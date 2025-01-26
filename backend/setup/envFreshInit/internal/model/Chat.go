package model

import (
	"gorm.io/gorm"
	"time"
)

type Chat struct {
	gorm.Model
	Id         int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId     int64          `json:"user_id"`
	TicketId   int64          `json:"ticket_id"`
	SenderType string         `json:"sender_type"`
	Content    string         `json:"content" gorm:"type:TEXT"`
	SentAt     *time.Time     `json:"sent_at"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	User       User           `gorm:"foreignKey:UserId;references:Id"`
	Ticket     Ticket         `gorm:"foreignKey:TicketId;references:Id"`
}

func (Chat) TableName() string {
	return "x_chat"
}
