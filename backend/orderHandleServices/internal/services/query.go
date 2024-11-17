package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"net/http"
	pb "orderHandleServices/api/proto"
	"orderHandleServices/internal/consumer"
	"orderHandleServices/internal/dao"
	"orderHandleServices/internal/model"
)

//func (s *OrderHandleServices) GetOrderStatus(context context.Context, request *pb.GetOrderStatusRequest) (*pb.GetOrderStatusResponse, error) {
//	var order model.Orders
//	// redis中的键值路径定义为 "pending_orders" / <用户ID> / <订单号>
//	log.Printf("[订单号 %s] [用户Id %d]\n", request.OrderId, request.UserId)
//	orderKey := fmt.Sprintf("%s:%d:%s", consumer.PendingOrderKeyPrefix, request.UserId, request.OrderId)
//	val, err := dao.Rdb.Get(context, orderKey).Result()
//	if errors.Is(err, redis.Nil) {
//		// 如果 Redis 中没有此订单，说明可能已处理，直接返回
//		log.Println("redis中订单信息不存在可能已经被处理")
//		return &pb.GetOrderStatusResponse{
//			Code:       http.StatusNotFound,
//			IsFinished: true,  // 订单已经完成
//			IsSuccess:  false, // 但是失败了
//			Msg:        "没有该订单请重试下单",
//		}, nil
//	} else if err != nil {
//		// 处理 Redis 错误
//		log.Println("Failed to fetch from Redis:", err)
//		return &pb.GetOrderStatusResponse{
//			Code:       http.StatusInternalServerError,
//			IsFinished: true,  // 订单已经完成
//			IsSuccess:  false, // 但是失败了
//			Msg:        "未知错误" + err.Error(),
//		}, nil
//	}
//	if err := json.Unmarshal([]byte(val), &order); err != nil {
//		return &pb.GetOrderStatusResponse{
//			Code:       http.StatusInternalServerError,
//			IsFinished: true,  // 订单已经完成
//			IsSuccess:  false, // 但是失败了
//			Msg:        "未知错误" + err.Error(),
//		}, nil
//	}
//	return &pb.GetOrderStatusResponse{
//		Code:       http.StatusOK,    // 查询成功
//		Msg:        "Query order ok", // 查询结果信息
//		IsFinished: order.IsFinished, // 订单是否结束
//		IsSuccess:  order.IsSuccess,  // 订单是否成功
//		OrderInfo:  []byte(val),      // 订单信息
//	}, nil
//}

//func (s *OrderHandleServices) GetOrderStatus(ctx context.Context, request *pb.GetOrderStatusRequest) (*pb.GetOrderStatusResponse, error) {
//	// 日志记录查询信息
//	log.Printf("[订单查询] 用户ID: %d, 订单号: %s\n", request.UserId, request.OrderId)
//
//	// 构造 Redis 键
//	orderKey := fmt.Sprintf("%s:%d:%s", consumer.PendingOrderKeyPrefix, request.UserId, request.OrderId)
//
//	// 从 Redis 中获取订单状态
//	val, err := dao.Rdb.Get(ctx, orderKey).Result()
//	if errors.Is(err, redis.Nil) {
//		// 如果 Redis 中没有该订单，说明可能已经被处理或订单无效
//		log.Println("[订单查询] Redis 中未找到订单，可能已处理或不存在")
//		return &pb.GetOrderStatusResponse{
//			Code:       http.StatusNotFound,
//			IsFinished: true,
//			IsSuccess:  false,
//			Msg:        "订单已处理或不存在，请重试下单",
//		}, nil
//	} else if err != nil {
//		// 处理 Redis 获取失败的情况
//		log.Printf("[订单查询] Redis 查询失败: %v\n", err)
//		return &pb.GetOrderStatusResponse{
//			Code:       http.StatusInternalServerError,
//			IsFinished: true,
//			IsSuccess:  false,
//			Msg:        "订单查询失败: " + err.Error(),
//		}, nil
//	}
//
//	// 尝试解析订单数据
//	var order model.Orders
//	if err := json.Unmarshal([]byte(val), &order); err != nil {
//		log.Printf("[订单查询] 订单数据解析失败: %v\n", err)
//		return &pb.GetOrderStatusResponse{
//			Code:       http.StatusInternalServerError,
//			IsFinished: true,
//			IsSuccess:  false,
//			Msg:        "订单数据解析失败: " + err.Error(),
//		}, nil
//	}
//
//	// 返回订单状态
//	return &pb.GetOrderStatusResponse{
//		Code:       http.StatusOK,    // 成功状态码
//		Msg:        "订单查询成功",         // 查询成功提示
//		IsFinished: order.IsFinished, // 是否已完成
//		IsSuccess:  order.IsSuccess,  // 是否成功
//		OrderInfo:  []byte(val),      // 原始订单数据
//	}, nil
//}

func (s *OrderHandleServices) GetOrderStatus(ctx context.Context, request *pb.GetOrderStatusRequest) (*pb.GetOrderStatusResponse, error) {
	// 日志记录查询信息
	log.Printf("[订单查询] 用户ID: %d, 订单号: %s\n", request.UserId, request.OrderId)

	// 构造 Redis 键
	orderKey := fmt.Sprintf("%s:%d:%s", consumer.PendingOrderKeyPrefix, request.UserId, request.OrderId)

	// 优先从 Redis 中获取订单状态
	val, err := dao.Rdb.Get(ctx, orderKey).Result()
	if errors.Is(err, redis.Nil) {
		// 如果 Redis 中没有该订单，尝试从数据库查询
		log.Println("[订单查询] Redis 中未找到订单，尝试从数据库查询")
		var order model.Orders
		if err := dao.Db.Model(&model.Orders{}).Where("user_id = ? AND order_id = ?", request.UserId, request.OrderId).First(&order).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据库中也找不到订单
				log.Printf("[订单查询] 数据库中未找到订单: 用户ID=%d, 订单号=%s\n", request.UserId, request.OrderId)
				return &pb.GetOrderStatusResponse{
					Code:       http.StatusNotFound,
					IsFinished: true,
					IsSuccess:  false,
					Msg:        "订单已处理或不存在，请重试下单",
				}, nil
			}
			// 数据库查询发生错误
			log.Printf("[订单查询] 数据库查询失败: %v\n", err)
			return &pb.GetOrderStatusResponse{
				Code:       http.StatusInternalServerError,
				IsFinished: true,
				IsSuccess:  false,
				Msg:        "订单查询失败: " + err.Error(),
			}, nil
		}

		//// 返回数据库中的订单状态
		//return &pb.GetOrderStatusResponse{
		//	Code:       http.StatusOK,
		//	Msg:        "订单查询成功",
		//	IsFinished: order.IsFinished,
		//	IsSuccess:  order.IsSuccess,
		//	OrderInfo:  nil, // 如果需要，可以序列化数据库订单信息
		//}, nil
		// 返回数据库中的订单状态
		orderData, err := json.Marshal(order)
		if err != nil {
			log.Printf("[订单查询] 数据库订单序列化失败: %v\n", err)
			return &pb.GetOrderStatusResponse{
				Code:       http.StatusInternalServerError,
				IsFinished: true,
				IsSuccess:  false,
				Msg:        "订单数据序列化失败: " + err.Error(),
			}, nil
		}

		return &pb.GetOrderStatusResponse{
			Code:       http.StatusOK,
			Msg:        "订单查询成功",
			IsFinished: order.IsFinished,
			IsSuccess:  order.IsSuccess,
			OrderInfo:  orderData, // 返回序列化的订单信息
		}, nil

	} else if err != nil {
		// Redis 查询发生错误
		log.Printf("[订单查询] Redis 查询失败: %v\n", err)
		return &pb.GetOrderStatusResponse{
			Code:       http.StatusInternalServerError,
			IsFinished: true,
			IsSuccess:  false,
			Msg:        "订单查询失败: " + err.Error(),
		}, nil
	}

	// 如果 Redis 中查到订单，尝试反序列化
	var order model.Orders
	if err := json.Unmarshal([]byte(val), &order); err != nil {
		log.Printf("[订单查询] 订单数据解析失败: %v\n", err)
		return &pb.GetOrderStatusResponse{
			Code:       http.StatusInternalServerError,
			IsFinished: true,
			IsSuccess:  false,
			Msg:        "订单数据解析失败: " + err.Error(),
		}, nil
	}

	// 返回 Redis 中的订单状态
	return &pb.GetOrderStatusResponse{
		Code:       http.StatusOK,
		Msg:        "订单查询成功",
		IsFinished: order.IsFinished,
		IsSuccess:  order.IsSuccess,
		OrderInfo:  []byte(val),
	}, nil
}
