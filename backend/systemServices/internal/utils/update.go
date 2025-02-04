package utils

import (
	"context"
	"systemServices/internal/dao"
	"time"
)

// NotifyPaymentConfigUpdate 管理员更新支付配置时，向 Redis 发布通知
func NotifyPaymentConfigUpdate() error {
	notifyCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return dao.Rdb.Publish(notifyCtx, "PAYMENT_CONFIG_UPDATED", "reload").Err()
}
