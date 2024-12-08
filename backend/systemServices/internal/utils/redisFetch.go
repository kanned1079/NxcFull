package utils

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"systemServices/internal/dao"
	"time"
)

// GetLast7DaysAccessApiCounts 从redis获取最近7天的api接口请求次数记录
func GetLast7DaysAccessApiCounts() ([]int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 获取最近 7 天的日期
	now := time.Now()
	keys := make([]string, 7)
	for i := 0; i < 7; i++ {
		date := now.AddDate(0, 0, -i).Format("20060102")
		keys[i] = fmt.Sprintf("access:api:%s", date)
	}

	// 从 Redis 批量获取数据
	values, err := dao.Rdb.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch access counts: %v", err)
	}

	// 将结果解析为 []int64
	counts := make([]int64, 7)
	for i, value := range values {
		if value != nil {
			// 将 value 转换为字符串再解析为 int64
			strValue, ok := value.(string)
			if !ok {
				log.Printf("Error converting value for key %s: not a string", keys[i])
				counts[i] = 0
				continue
			}

			count, convErr := strconv.ParseInt(strValue, 10, 64)
			if convErr != nil {
				log.Printf("Error parsing value for key %s: %v", keys[i], convErr)
				counts[i] = 0
				continue
			}

			counts[i] = count
		} else {
			// 如果 Redis 键不存在，设置为 0
			counts[i] = 0
			if err := dao.Rdb.Set(ctx, keys[i], 0, 7*24*time.Hour).Err(); err != nil {
				log.Printf("Error setting default value for key %s: %v\n", keys[i], err)
			}
		}
	}

	log.Printf("Fetched access counts: %v", counts)

	return counts, nil
}
