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

	//cacheKey := "ticket:messages:" + strconv.FormatInt(chatMsg.TicketId, 10)
	//messageJSON, err := json.Marshal(chatMsg)
	//if err != nil {
	//	return err
	//}
	//// 将消息推送到 Redis List 尾部
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()
	//if err := dao.Rdb.RPush(ctx, cacheKey, messageJSON).Err(); err != nil {
	//	log.Println("Redis 消息插入失败:", err)
	//	return err
	//}

	if result := dao.Db.Model(&model.Chat{}).Create(&chatMsg); result.Error != nil {
		log.Println("消息插入数据库失败")
		return result.Error
	}
	go func() {
		if messageData.Role == "admin" {
			if result := dao.Db.Model(&model.Ticket{}).Where("id = ?", messageData.TicketId).Update("last_reply", sentAt); result.Error != nil {
				//return result.Error
				log.Println("Update last response failure")
			}
		}
	}()

	return nil

}

func ProcessErrorOrder(badOrder *model.Orders) {
	if result := dao.Db.Model(&model.Orders{}).Create(badOrder); result.Error != nil {
		log.Println(result.Error.Error())
	}
}
