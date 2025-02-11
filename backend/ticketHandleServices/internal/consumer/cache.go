package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"ticketHandleServices/internal/dao"
	"ticketHandleServices/internal/model"
	"time"
)

//func SaveMSgInRedis(ticketId int64, msg string) error {
//	// 将消息转换为 JSON 字符串
//	messageData, err := json.Marshal(chatMsg)
//	if err != nil {
//		tx.Rollback()
//		return &pb.CreateNewTicketResponse{
//			Code:    http.StatusInternalServerError,
//			Msg:     "消息序列化失败: " + err.Error(),
//			Created: false,
//		}, nil
//	}
//
//	// Redis List 键，使用工单 ID 作为 List 名称
//	cacheKey := "ticket:messages:" + strconv.FormatInt(newTicket.Id, 10)
//
//	// 初始化 Redis List 并插入第一条消息
//	redisContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	err = dao.Rdb.RPush(redisContext, cacheKey, messageData).Err()
//	if err != nil {
//		tx.Rollback()
//		return &pb.CreateNewTicketResponse{
//			Code:    http.StatusInternalServerError,
//			Msg:     "Redis 插入数据失败: " + err.Error(),
//			Created: false,
//		}, nil
//	}
//}

// 缓存聊天记录到 Redis
func cacheChatMessageToRedis(ctx context.Context, chatMsg model.Chat) error {
	chatMsgJson, err := json.Marshal(chatMsg)
	if err != nil {
		return fmt.Errorf("转换聊天消息为 JSON 失败: %v", err)
	}

	chatKey := fmt.Sprintf("ticket:%d:chats", chatMsg.TicketId)
	if err := dao.Rdb.RPush(ctx, chatKey, chatMsgJson).Err(); err != nil {
		return fmt.Errorf("缓存聊天消息到 Redis 失败: %v", err)
	}

	// 设置 Redis 过期时间（24 小时）
	if err := dao.Rdb.Expire(ctx, chatKey, 24*time.Hour).Err(); err != nil {
		return fmt.Errorf("设置 Redis 过期时间失败: %v", err)
	}

	return nil
}
