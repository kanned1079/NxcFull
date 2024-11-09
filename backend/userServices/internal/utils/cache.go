package utils

import (
	"context"
	"fmt"
	"log"
	"time"
	"userServices/internal/dao"
)

func ClearUserCacheByEmail(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userKey := fmt.Sprintf("user:%s", email)
	authKey := fmt.Sprintf("auth:%s", email)

	err := dao.Rdb.Del(ctx, userKey, authKey).Err()
	if err != nil {
		return err
	}
	log.Println("Clear user cache successfully: ", email)

	return nil
}
