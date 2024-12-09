package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"math"
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
			if err := dao.Rdb.Set(ctx, keys[i], 0, 30*24*time.Hour).Err(); err != nil {
				log.Printf("Error setting default value for key %s: %v\n", keys[i], err)
			}
		}
	}

	//log.Printf("Fetched access counts: %v", counts)

	return counts, nil
}

// GetLast7DaysAndMonthIncome 返回本周每天收入的数组 昨日收入 当月收入 是否有错误
func GetLast7DaysAndMonthIncome() (data []float32, yesterdayIncome float32, monthIncome float32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 获取最近 7 天的日期，包括当天
	now := time.Now()
	keys := make([]string, 7)
	for i := 0; i < 7; i++ {
		date := now.AddDate(0, 0, -i).Format("20060102")
		keys[i] = fmt.Sprintf("access:income:%s", date)
	}

	// 从 Redis 批量获取最近 7 天的数据
	values, err := dao.Rdb.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to fetch income data: %v", err)
	}

	// 将结果解析为 []float32
	incomes := make([]float32, 7)
	for i, value := range values {
		if value != nil {
			strValue, ok := value.(string)
			if !ok {
				log.Printf("Error converting value for key %s: not a string", keys[i])
				incomes[i] = 0
				continue
			}

			intValue, convErr := strconv.ParseInt(strValue, 10, 64)
			if convErr != nil {
				log.Printf("Error parsing value for key %s: %v", keys[i], convErr)
				incomes[i] = 0
				continue
			}

			// 恢复原始金额（除以 100 得到浮点数）
			incomes[i] = float32(intValue) / 100
		} else {
			// 如果 Redis 键不存在，收入设置为 0
			incomes[i] = 0
			if err := dao.Rdb.Set(ctx, keys[i], 0, 30*24*time.Hour).Err(); err != nil {
				log.Printf("Error setting default value for key %s: %v\n", keys[i], err)
			}
		}
	}

	log.Printf("Fetched income data: %v", incomes)

	// 获取昨日收入
	yesterdayIncome = incomes[1] // 昨日的收入是从 7 天中倒数第二天的数据

	// 计算本月收入（本月从 1 号到今天）
	monthIncome = float32(0)
	for i := 0; i < 7; i++ {
		date := now.AddDate(0, 0, -i)
		if date.Month() == now.Month() {
			monthIncome += incomes[i]
		}
	}

	// 补充计算 Redis 缺失的天数
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	for d := startOfMonth; d.Before(now.AddDate(0, 0, -7)); d = d.AddDate(0, 0, 1) {
		key := fmt.Sprintf("access:income:%s", d.Format("20060102"))
		val, err := dao.Rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			// 如果 Redis 中没有这一天的数据，默认收入为 0
			continue
		} else if err != nil {
			log.Printf("Error fetching value for key %s: %v\n", key, err)
			continue
		}

		intValue, convErr := strconv.ParseInt(val, 10, 64)
		if convErr != nil {
			log.Printf("Error parsing value for key %s: %v\n", key, convErr)
			continue
		}

		monthIncome += float32(intValue) / 100
	}

	// 格式化为两位小数
	monthIncome = float32(math.Round(float64(monthIncome)*100) / 100)
	yesterdayIncome = float32(math.Round(float64(yesterdayIncome)*100) / 100)

	return incomes, yesterdayIncome, monthIncome, nil
}
