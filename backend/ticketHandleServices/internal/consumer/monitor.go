package consumer

import (
	"context"
	"encoding/json"
	"log"
	"ticketHandleServices/internal/dao"
	"ticketHandleServices/internal/model"
	"time"
)

// MonitorAndWriteToDatabase 监控 Redis List 并批量写入数据库
func MonitorAndWriteToDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("监控已停止:", ctx.Err())
			return
		default:
			// 查找 Redis 中所有的工单 List 键
			keys, err := dao.Rdb.Keys(ctx, "ticket:messages:*").Result()
			if err != nil {
				log.Println("Redis 键查找失败:", err)
				continue
			}

			for _, key := range keys {
				// 获取所有消息并从 Redis 中移除
				messages, err := dao.Rdb.LRange(ctx, key, 0, -1).Result()
				if err != nil {
					log.Println("读取 Redis List 失败:", err)
					continue
				}

				if len(messages) == 0 {
					continue
				}

				// 解析并将消息批量写入数据库
				var chatMessages []model.Chat
				for _, msg := range messages {
					var chatMsg model.Chat
					if err := json.Unmarshal([]byte(msg), &chatMsg); err != nil {
						log.Println("消息解析失败:", err)
						continue
					}
					chatMessages = append(chatMessages, chatMsg)
				}

				// 批量插入数据库
				if err := dao.Db.WithContext(ctx).Model(&model.Chat{}).Create(&chatMessages).Error; err != nil {
					log.Println("批量插入数据库失败:", err)
					continue
				}

				// 从 Redis List 中删除已经写入的消息
				if err := dao.Rdb.Del(ctx, key).Err(); err != nil {
					log.Println("Redis List 删除失败:", err)
				}
			}
			time.Sleep(5 * time.Second) // 设置为每隔 5 秒检查一次
		}
	}
}
