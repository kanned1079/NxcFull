package dao

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"ticketHandleServices/internal/config/remote"
)

var (
	Rdb *redis.Client
)

// InitRedisServer 初始化Redis数据库
func InitRedisServer() {
	addr := fmt.Sprintf("%s:%d", remote.MyRedisConfig.Host, remote.MyRedisConfig.Port)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       remote.MyRedisConfig.Database,
		Username: remote.MyRedisConfig.Username,
		Password: remote.MyDbConfig.Password,
	})

	// 测试与 Redis 服务器的连接
	pong, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		//return fmt.Errorf()
		panic("初始化Redis连接错误" + err.Error())
	}

	log.Println("初始化Redis连接成功", pong)

}
