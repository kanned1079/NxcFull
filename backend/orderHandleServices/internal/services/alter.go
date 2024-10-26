package services

import (
	sysContext "context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
	"net/http"
	pb "orderHandleServices/api/proto"
	"orderHandleServices/internal/dao"
	"orderHandleServices/internal/model"
	"time"
)

//func (s *OrderHandleServices) CancelOrder(context *context.Context, request *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
//	var err error
//	log.Println("取消订单信息: ", request.UserId, request.OrderId)
//	var pendingOrder model.Orders
//	redisKey := fmt.Sprintf("pending_order:%d:%s", request.UserId, request.OrderId)
//	// 从redis中取出数据 反序列化为pendingOrder
//
//	err = dao.Db.Transaction(func(tx *gorm.DB) error {
//		// 这里可以插入到order表中
//	})
//	if err != nil {
//		return &pb.CancelOrderResponse{
//			Code:      http.StatusInternalServerError,
//			Cancelled: false,
//			Msg:       "取消订单可能没有成功",
//		}, nil
//	}
//
//	return &pb.CancelOrderResponse{
//		Code:      http.StatusOK,
//		Cancelled: true,
//		Msg:       "订单已取消",
//	}, nil
//}

func (s *OrderHandleServices) CancelOrder(context context.Context, request *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	var err error
	log.Println("取消订单信息: ", request.UserId, request.OrderId)

	var pendingOrder model.Orders
	redisKey := fmt.Sprintf("pending_order:%d:%s", request.UserId, request.OrderId)

	// 从 Redis 中取出数据并反序列化
	redisCtx, cancel := sysContext.WithTimeout(sysContext.Background(), 5*time.Second)
	defer cancel()
	val, err := dao.Rdb.Get(redisCtx, redisKey).Result()
	if errors.Is(err, redis.Nil) {
		log.Println("订单不存在于 Redis 中")
		return &pb.CancelOrderResponse{
			Code:      http.StatusNotFound,
			Cancelled: false,
			Msg:       "订单未找到",
		}, nil
	} else if err != nil {
		log.Println("从 Redis 获取订单失败:", err)
		return &pb.CancelOrderResponse{
			Code:      http.StatusInternalServerError,
			Cancelled: false,
			Msg:       "获取订单信息失败" + err.Error(),
		}, nil
	}

	// 反序列化 Redis 中的订单数据
	err = json.Unmarshal([]byte(val), &pendingOrder)
	if err != nil {
		log.Println("订单反序列化失败:", err)
		return &pb.CancelOrderResponse{
			Code:      http.StatusInternalServerError,
			Cancelled: false,
			Msg:       "订单反序列化失败",
		}, nil
	}

	// 启动数据库事务
	err = dao.Db.Transaction(func(tx *gorm.DB) error {
		// 将 pendingOrder 插入订单表
		pendingOrder.IsFinished = true       // 标记订单为已取消
		pendingOrder.FailureReason = "订单已取消" // 取消原因

		if err := tx.Create(&pendingOrder).Error; err != nil {
			log.Println("订单插入数据库失败:", err)
			return err
		}

		// 删除 Redis 中的订单
		if err := dao.Rdb.Del(redisCtx, redisKey).Err(); err != nil {
			log.Println("从 Redis 删除订单失败:", err)
			return err
		}

		return nil
	})

	// 事务操作失败处理
	if err != nil {
		log.Println("取消订单失败:", err)
		return &pb.CancelOrderResponse{
			Code:      http.StatusInternalServerError,
			Cancelled: false,
			Msg:       "取消订单失败",
		}, nil
	}

	// 事务成功，返回成功响应
	return &pb.CancelOrderResponse{
		Code:      http.StatusOK,
		Cancelled: true,
		Msg:       "订单已取消",
	}, nil
}
