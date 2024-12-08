package middleware

import (
	"context"
	"fmt"
	"gateway/internal/dao"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

var (
	totalCount int64
	mu         sync.Mutex
)

// APICountMiddleware is a Gin middleware to count API accesses
func APICountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		totalCount++
		mu.Unlock()
		c.Next()
	}
}

// StartApiAccessCountCron initializes and starts the cron job to refresh counts into Redis
func StartApiAccessCountCron() {
	log.Println("Starting API access count cron job")
	cronJob := cron.New()

	_, err := cronJob.AddFunc("@every 10m", func() {
		mu.Lock()
		count := totalCount
		totalCount = 0
		mu.Unlock()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		key := fmt.Sprintf("access:api:%s", time.Now().Format("20060102"))
		if err := dao.Rdb.IncrBy(ctx, key, count).Err(); err != nil {
			log.Printf("Error incrementing count for key %s: %v\n", key, err)
		}

		if err := dao.Rdb.Expire(ctx, key, 30*24*time.Hour).Err(); err != nil {
			log.Printf("Error setting expiration for key %s: %v\n", key, err)
		}
	})

	if err != nil {
		log.Printf("Error scheduling cron job: %v\n", err)
		return
	}

	cronJob.Start()
	select {} // Block forever
}
