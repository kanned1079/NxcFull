package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	pb "orderHandleServices/api/proto"
	"orderHandleServices/internal/consumer"
	"orderHandleServices/internal/dao"
	"orderHandleServices/internal/model"
)

func (s *OrderHandleServices) GetOrderStatus(context context.Context, request *pb.GetOrderStatusRequest) (*pb.GetOrderStatusResponse, error) {
	var order model.Orders
	orderKey := fmt.Sprintf("%s:%s", consumer.PendingOrderKeyPrefix, request.OrderId)
	val, err := dao.Rdb.Get(context, orderKey).Result()
	if errors.Is(err, redis.Nil) {
		// 如果 Redis 中没有此订单，说明可能已处理，直接返回
		return &pb.GetOrderStatusResponse{
			Code: http.StatusNotFound,
			Msg:  "没有该订单请重试下单",
		}, nil
	} else if err != nil {
		// 处理 Redis 错误
		log.Println("Failed to fetch from Redis:", err)
		return &pb.GetOrderStatusResponse{
			Code: http.StatusInternalServerError,
			Msg:  "未知错误" + err.Error(),
		}, nil
	}
	if err := json.Unmarshal([]byte(val), &order); err != nil {
		return &pb.GetOrderStatusResponse{
			Code: http.StatusInternalServerError,
			Msg:  "未知错误" + err.Error(),
		}, nil
	}
	return &pb.GetOrderStatusResponse{
		Code:       http.StatusOK,
		Msg:        "Query order ok",
		IsFinished: order.IsFinished,
	}, nil
}
