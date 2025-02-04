package utils

import (
	"context"
	"systemServices/internal/dao"
)

// NotifyPaymentConfigUpdate 管理员更新支付配置时，向 Redis 发布通知
func NotifyPaymentConfigUpdate(ctx context.Context) error {
	return dao.Rdb.Publish(ctx, "PAYMENT_CONFIG_UPDATED", "reload").Err()
}
