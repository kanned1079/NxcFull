package consumer

import (
	"encoding/json"
	"log"
	"ticketHandleServices/internal/dao"
	"ticketHandleServices/internal/model"
	"time"
)

const (
	PendingOrderKeyPrefix = "pending_order"
)

func ProcessTicket(jsonData []byte) error {
	// 消息格式
	var messageData = &struct {
		UserId   int64     `json:"user_id"`   // 用户或管理员的id
		TicketId int64     `json:"ticket_id"` // 对应哪一个工单的id
		Role     string    `json:"role"`      // 管理员admin或user
		Content  string    `json:"content"`   // 消息内容
		SentAt   time.Time `json:"sent_at"`   // 发送时间
	}{}
	err := json.Unmarshal(jsonData, &messageData)
	if err != nil {
		return err
	}
	log.Println("聊天消息", messageData)

	sentAt := time.Now()
	var chatMsg = model.Chat{
		UserId:     messageData.UserId,
		TicketId:   messageData.TicketId,
		Content:    messageData.Content,
		SenderType: messageData.Role,
		SentAt:     &sentAt,
	}
	if result := dao.Db.Model(&model.Chat{}).Create(&chatMsg); result.Error != nil {
		log.Println("消息插入数据库失败")
		return result.Error
	}
	if messageData.Role == "admin" {
		if result := dao.Db.Model(&model.Ticket{}).Where("id = ?", messageData.TicketId).Update("last_reply", sentAt); result.Error != nil {
			return result.Error
		}
	}
	return nil

}

func ProcessErrorOrder(badOrder *model.Orders) {
	if result := dao.Db.Model(&model.Orders{}).Create(badOrder); result.Error != nil {
		log.Println(result.Error.Error())
	}
}
