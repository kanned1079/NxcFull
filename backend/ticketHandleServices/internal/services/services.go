package services

import (
	"context"
	"encoding/json"
	"fmt"
	pb "ticketHandleServices/api/proto"
	"ticketHandleServices/internal/dao"
	"ticketHandleServices/internal/model"
	"time"
)

type TicketHandleServices struct {
	pb.UnimplementedTicketHandleServiceServer
}

func NewTicketHandleServices() *TicketHandleServices {
	return &TicketHandleServices{}
}

// 将聊天记录缓存到 Redis
func (s *TicketHandleServices) cacheChatHistoryToRedis(ctx context.Context, chatCacheKey string, chatHistory []model.Chat) error {
	// 使用 pipeline 批量执行写入
	//cacheCtx, cacheCancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cacheCancel()
	pipe := dao.Rdb.Pipeline()
	for _, msg := range chatHistory {
		chatJson, _ := json.Marshal(msg)
		pipe.RPush(ctx, chatCacheKey, chatJson) // 将聊天记录添加到 Redis
	}
	// 设置 Redis 过期时间，防止长期占用
	pipe.Expire(ctx, chatCacheKey, 24*time.Hour)
	if _, err := pipe.Exec(ctx); err != nil {
		return fmt.Errorf("缓存聊天记录到 Redis 失败: %v", err)
	}
	return nil
}
