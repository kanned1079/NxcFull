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
	// redis中的键值路径定义为 "pending_orders" / <用户ID> / <订单号>
	log.Printf("[订单号 %s] [用户Id %d]\n", request.OrderId, request.UserId)
	orderKey := fmt.Sprintf("%s:%d:%s", consumer.PendingOrderKeyPrefix, request.UserId, request.OrderId)
	val, err := dao.Rdb.Get(context, orderKey).Result()
	if errors.Is(err, redis.Nil) {
		// 如果 Redis 中没有此订单，说明可能已处理，直接返回
		return &pb.GetOrderStatusResponse{
			Code:       http.StatusNotFound,
			IsFinished: true,  // 订单已经完成
			IsSuccess:  false, // 但是失败了
			Msg:        "没有该订单请重试下单",
		}, nil
	} else if err != nil {
		// 处理 Redis 错误
		log.Println("Failed to fetch from Redis:", err)
		return &pb.GetOrderStatusResponse{
			Code:       http.StatusInternalServerError,
			IsFinished: true,  // 订单已经完成
			IsSuccess:  false, // 但是失败了
			Msg:        "未知错误" + err.Error(),
		}, nil
	}
	if err := json.Unmarshal([]byte(val), &order); err != nil {
		return &pb.GetOrderStatusResponse{
			Code:       http.StatusInternalServerError,
			IsFinished: true,  // 订单已经完成
			IsSuccess:  false, // 但是失败了
			Msg:        "未知错误" + err.Error(),
		}, nil
	}
	return &pb.GetOrderStatusResponse{
		Code:       http.StatusOK,    // 查询成功
		Msg:        "Query order ok", // 查询结果信息
		IsFinished: order.IsFinished, // 订单是否结束
		IsSuccess:  order.IsSuccess,  // 订单是否成功
		OrderInfo:  []byte(val),      // 订单信息
	}, nil
}
