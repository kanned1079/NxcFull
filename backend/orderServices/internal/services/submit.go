package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	pb "orderServices/api/proto"
	"orderServices/internal/puiblisher"
	"orderServices/internal/utils"
	"time"
)

// CommitNewOrder 用戶提交新的訂單 目前是同步处理 后续需要使用mq实现异步
func (s *OrderServices) CommitNewOrder(context context.Context, request *pb.CommitNewOrderRequest) (*pb.CommitNewOrderResponse, error) {
	log.Printf("UserId %v\n PlanId %v\n Period %v\n CouponId %v\n", request.UserId, request.PlanId, request.Period, request.CouponId)
	orderIdGenerated := utils.GenerateOrderId(request.UserId)
	log.Println("newOrderId: ", orderIdGenerated)

	// 推送到消息队列的数据
	postOrderData := &struct {
		OrderId  string `json:"order_id"`
		UserId   int64  `json:"user_id"`
		PlanId   int64  `json:"plan_id"`
		Period   string `json:"period"`
		CouponId int64  `json:"coupon_id"`
	}{
		OrderId:  orderIdGenerated,
		UserId:   request.UserId,
		PlanId:   request.PlanId,
		Period:   request.Period,
		CouponId: request.CouponId,
	}
	if postOrderDataJSON, err := json.Marshal(postOrderData); err != nil {
		return &pb.CommitNewOrderResponse{
			Code: http.StatusInternalServerError,
			Msg:  "内部错误" + err.Error(),
		}, nil
	} else {
		if err = puiblisher.PublishOrderMessage(postOrderDataJSON); err != nil {
			log.Println("puiblisher.PublishOrderMessage error:", err)
			return &pb.CommitNewOrderResponse{
				Code: http.StatusInternalServerError,
				Msg:  "内部错误" + err.Error(),
			}, nil
		}
		time.Sleep(time.Millisecond * 500) // 延迟500ms
	}

	return &pb.CommitNewOrderResponse{
		Code:    http.StatusOK,
		OrderId: orderIdGenerated,
		Msg:     "ok",
	}, nil
}

func (s *OrderServices) CommitNewTopUpOrder(ctx context.Context, request *pb.CommitNewTopUpOrderRequest) (*pb.CommitNewTopUpOrderResponse, error) {
	postTopUpOrderData := &struct {
		UserId        int64   `json:"user_id"`
		TopUpAmount   float32 `json:"top_up_amount"`
		PaymentMethod string  `json:"payment_method"`
		OrderId       string  `json:"order_id"`
	}{
		UserId:        request.UserId,
		TopUpAmount:   request.TopUpAmount,
		PaymentMethod: request.PaymentMethod,
	}

	// 生成唯一订单号
	postTopUpOrderData.OrderId = utils.GenerateTopUpOrderId(request.UserId)

	if postTopUpOrderDataJson, err := json.Marshal(postTopUpOrderData); err != nil {
		return &pb.CommitNewTopUpOrderResponse{
			Code:    http.StatusInternalServerError,
			Msg:     "内部错误" + err.Error(),
			Created: false,
		}, nil
	} else {
		if err = puiblisher.PublishOrderMessage(postTopUpOrderDataJson); err != nil {
			log.Println("puiblisher.PublishOrderMessage error:", err)
			return &pb.CommitNewTopUpOrderResponse{
				Code: http.StatusInternalServerError,
				Msg:  "内部错误" + err.Error(),
			}, nil
		}
		time.Sleep(time.Millisecond * 500) // 延迟500ms
	}

	return &pb.CommitNewTopUpOrderResponse{}, nil
}
