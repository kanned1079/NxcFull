package dao

import (
	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
)

func InitRedisServer() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "api.ikanned.com:26379",
		//Username: "kanna",
		//Password: "SHY963852741.Year",
		DB: 0,
	})
}
