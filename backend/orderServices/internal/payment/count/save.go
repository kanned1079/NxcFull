package count

import (
	"context"
	"fmt"
	"log"
	"math"
	"orderServices/internal/dao"
	"time"
)

// SaveIncomeCount2Redis 如果用户支付成功则存储或更新redis中的当天收入的总计
// 注意这存储的是金额*100后的数值 使用的时候需要取出再/100
func SaveIncomeCount2Redis(incomeAmount float32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 将浮点数转化为整数
	incomeCents := int64(math.Round(float64(incomeAmount * 100)))

	key := fmt.Sprintf("access:income:%s", time.Now().Format("20060102"))

	// 使用整数存储
	if err := dao.Rdb.IncrBy(ctx, key, incomeCents).Err(); err != nil {
		log.Printf("Error incrementing income for key %s: %v\n", key, err)
		return
	}

	// 设置过期时间
	if err := dao.Rdb.Expire(ctx, key, 30*24*time.Hour).Err(); err != nil {
		log.Printf("Error setting expiration for key %s: %v\n", key, err)
	}
}
