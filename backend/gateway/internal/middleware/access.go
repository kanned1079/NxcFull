package middleware

import (
	logPb "gateway/internal/grpc/api/logs/proto"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

var (
	LogBuffer []*logPb.ApiLogEntry
	LogMutex  sync.Mutex
)

// AccessLoggerMiddleware 收集请求日志的中间件
func AccessLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 收集日志信息
		duration := time.Since(startTime).Seconds() * 1000 // 响应时间 (毫秒)
		logEntry := &logPb.ApiLogEntry{                    // 使用指针创建实例
			Level:        "INFO",
			StatusCode:   int32(c.Writer.Status()),
			RequestAt:    startTime.UnixNano(),
			ResponseTime: float32(duration),
			ClientIp:     c.ClientIP(),
			Method:       c.Request.Method,
			Path:         c.Request.URL.Path,
		}

		// 将日志添加到缓冲区
		LogMutex.Lock()
		LogBuffer = append(LogBuffer, logEntry)
		LogMutex.Unlock()
	}
}
