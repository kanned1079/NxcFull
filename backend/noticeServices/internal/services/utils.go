package services

import (
	"context"
	"fmt"
	"log"
	pb "noticeServices/api/proto"
	"noticeServices/internal/dao"
	"noticeServices/internal/model"
	"time"
)

// 转换函数: 将 []model.PublicNotices 转换为 []*pb.PublicNotice
func convertToProtoNotices(notices []model.PublicNotices) []*pb.PublicNotice {
	protoNotices := make([]*pb.PublicNotice, len(notices))
	for i, notice := range notices {
		protoNotices[i] = &pb.PublicNotice{
			Id:        notice.Id,
			Title:     notice.Title,
			Content:   notice.Content,
			Show:      notice.Show,
			ImgUrl:    notice.ImgUrl,
			Tags:      notice.Tags,
			CreatedAt: notice.CreatedAt.Format(time.RFC3339), // 时间转为字符串
			UpdatedAt: notice.UpdatedAt.Format(time.RFC3339),
			DeletedAt: func() string {
				if !notice.DeletedAt.Valid {
					return ""
				}
				return notice.DeletedAt.Time.Format(time.RFC3339)
			}(),
		}
	}
	return protoNotices
}

func ClearNoticeRedisCache(ctx context.Context) {
	// 使用 Redis 客户端的 Del 方法来删除指定的键
	redisKey := "notices:user"
	result, err := dao.Rdb.Del(ctx, redisKey).Result()
	if err != nil {
		log.Println("Error clearing cache:", err)
	}
	// 如果 result 是 1，表示成功删除了该键
	if result == 1 {
		fmt.Println("Cache cleared successfully.")
	} else {
		fmt.Println("Cache key does not exist.")
	}
}
