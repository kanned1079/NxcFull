package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"ticketHandleServices/internal/dao"
	"ticketHandleServices/internal/model"
	"time"
)

const (
	PendingOrderKeyPrefix = "pending_order"
)

//func ProcessTicket(jsonData []byte) error {
//	// 消息格式
//	var messageData = &struct {
//		UserId   int64     `json:"user_id"`   // 用户或管理员的id
//		TicketId int64     `json:"ticket_id"` // 对应哪一个工单的id
//		Role     string    `json:"role"`      // 管理员admin或user
//		Content  string    `json:"content"`   // 消息内容
//		SentAt   time.Time `json:"sent_at"`   // 发送时间
//	}{}
//	err := json.Unmarshal(jsonData, &messageData)
//	if err != nil {
//		return err
//	}
//	log.Println("聊天消息", messageData)
//
//	sentAt := time.Now()
//	var chatMsg = model.Chat{
//		UserId:     messageData.UserId,   // 用户id
//		TicketId:   messageData.TicketId, // 工单id
//		Content:    messageData.Content,  // 消息内容
//		SenderType: messageData.Role,     // 是谁发送的 user/admin
//		SentAt:     &sentAt,              // 发送时间
//	}
//
//	//dao.Rdb
//
//	if result := dao.Db.Model(&model.Chat{}).Create(&chatMsg); result.Error != nil {
//		log.Println("消息插入数据库失败")
//		return result.Error
//	}
//	go func() {
//		if messageData.Role == "admin" {
//			if result := dao.Db.Model(&model.Ticket{}).Where("id = ?", messageData.TicketId).Update("last_reply", sentAt); result.Error != nil {
//				//return result.Error
//				log.Println("Update last response failure")
//			}
//		}
//	}()
//
//	return nil
//}

//func ProcessTicket(jsonData []byte) error {
//	// 创建一个带有超时的 context
//	redisCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	// 消息格式
//	var messageData = &struct {
//		UserId   int64     `json:"user_id"`   // 用户或管理员的id
//		TicketId int64     `json:"ticket_id"` // 对应哪一个工单的id
//		Role     string    `json:"role"`      // 管理员admin或user
//		Content  string    `json:"content"`   // 消息内容
//		SentAt   time.Time `json:"sent_at"`   // 发送时间
//	}{}
//	err := json.Unmarshal(jsonData, &messageData)
//	if err != nil {
//		return err
//	}
//	log.Println("聊天消息", messageData)
//
//	sentAt := time.Now()
//	var chatMsg = model.Chat{
//		UserId:     messageData.UserId,   // 用户id
//		TicketId:   messageData.TicketId, // 工单id
//		Content:    messageData.Content,  // 消息内容
//		SenderType: messageData.Role,     // 是谁发送的 user/admin
//		SentAt:     &sentAt,              // 发送时间
//	}
//
//	// 插入聊天记录到数据库
//	if result := dao.Db.Model(&model.Chat{}).Create(&chatMsg); result.Error != nil {
//		log.Println("消息插入数据库失败")
//		return result.Error
//	}
//
//	// 将聊天消息对象转为 JSON 字符串
//	chatMsgJson, err := json.Marshal(chatMsg)
//	if err != nil {
//		log.Println("转换聊天消息为 JSON 失败:", err)
//		return err
//	}
//
//	// 缓存聊天记录到 Redis 列表（使用 LPush）
//	chatKey := fmt.Sprintf("ticket:%d:chats", messageData.TicketId)
//	//if err := dao.Rdb.LPush(redisCtx, chatKey, chatMsgJson).Err(); err != nil {
//	if err := dao.Rdb.RPush(redisCtx, chatKey, chatMsgJson).Err(); err != nil {
//		log.Println("缓存聊天消息到 Redis 失败:", err)
//	}
//
//	// 如果消息是管理员发送的，更新工单的最后回复时间到 Redis
//	if messageData.Role == "admin" {
//		// 将最后回复时间存储到 Redis，设置 24 小时过期
//		ticketKey := fmt.Sprintf("ticket:%d:last_reply", messageData.TicketId)
//		if err := dao.Rdb.Set(redisCtx, ticketKey, sentAt.Format(time.RFC3339), 24*time.Hour).Err(); err != nil {
//			log.Println("缓存工单最后回复时间到 Redis 失败:", err)
//		}
//		go func() {
//			// 更新数据库中的工单的最后回复时间
//			if result := dao.Db.Model(&model.Ticket{}).Where("id = ?", messageData.TicketId).Update("last_reply", sentAt); result.Error != nil {
//				log.Println("Update last response failure")
//			}
//		}()
//	}
//
//	return nil
//}

func ProcessTicket(jsonData []byte) error {
	// 创建一个带有超时的 context
	redisCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
		UserId:     messageData.UserId,   // 用户id
		TicketId:   messageData.TicketId, // 工单id
		Content:    messageData.Content,  // 消息内容
		SenderType: messageData.Role,     // 是谁发送的 user/admin
		SentAt:     &sentAt,              // 发送时间
	}

	// 插入聊天记录到数据库
	if result := dao.Db.Model(&model.Chat{}).Create(&chatMsg); result.Error != nil {
		log.Println("消息插入数据库失败")
		return result.Error
	}

	// 缓存聊天记录到 Redis
	if err := cacheChatMessageToRedis(redisCtx, chatMsg); err != nil {
		log.Println("缓存聊天消息到 Redis 失败:", err)
	}

	// 如果消息是管理员发送的，更新工单的最后回复时间到 Redis
	if messageData.Role == "admin" {
		// 更新 Redis 的最后回复时间
		ticketKey := fmt.Sprintf("ticket:%d:last_reply", messageData.TicketId)
		if err := dao.Rdb.Set(redisCtx, ticketKey, sentAt.Format(time.RFC3339), 24*time.Hour).Err(); err != nil {
			log.Println("缓存工单最后回复时间到 Redis 失败:", err)
		}

		// 异步更新数据库中的工单最后回复时间
		go func() {
			if result := dao.Db.Model(&model.Ticket{}).Where("id = ?", messageData.TicketId).Update("last_reply", sentAt); result.Error != nil {
				log.Println("更新工单最后回复时间失败:", result.Error)
			}
		}()
	}

	return nil
}
