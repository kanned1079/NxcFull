package utils

import (
	"fmt"
	"time"
)

func GenerateOrderId(userId int64) string {
	now := time.Now()

	// 获取当天的 0 点时间
	beginningOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 拼接订单号，格式为：<年月日><用户ID><当天0点起的时间戳>
	orderId := fmt.Sprintf(
		"%s%d%d",
		now.Format("20060102"), // 当前年月日
		userId,                 // 用户ID
		now.UnixMilli()-beginningOfDay.UnixMilli(), // 当天0点起的毫秒时间戳
	)

	return orderId
}

func GenerateTopUpOrderId(userId int64) string {
	now := time.Now()

	// 获取当天的 0 点时间
	beginningOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 拼接订单号，格式为：<年月日><用户ID><当天0点起的时间戳>
	return "t" + fmt.Sprintf(
		"%s%d%d",
		now.Format("20060102"), // 当前年月日
		userId,                 // 用户ID
		now.UnixMilli()-beginningOfDay.UnixMilli(), // 当天0点起的毫秒时间戳
	)
}
