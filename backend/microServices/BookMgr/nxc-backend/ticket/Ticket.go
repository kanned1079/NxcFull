package ticket

import "time"

type Ticket struct {
	Id          int64      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId      int64      `json:"user_id"`     // 用户ID
	Title       string     `json:"title"`       // 工单主题
	Urgency     uint8      `json:"urgency"`     // 紧急程度 1低 2中 3高
	Description string     `json:"description"` // 问题描述
	Status      string     `json:"status"`      // 工单状态（Open, Closed）
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	ClosedAt    *time.Time `json:"closed_at"` // 工单关闭时间
}

type Message struct {
	Id        int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	TicketId  int64     `json:"ticket_id"`  // 关联工单ID
	SenderId  uint8     `json:"sender_id"`  // 发送者ID 0代表用户 1代表管理员
	Content   string    `json:"content"`    // 消息内容
	CreatedAt time.Time `json:"created_at"` // 发送时间
}
