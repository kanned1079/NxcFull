package services

import (
	pb "NxcFull/backend/orderServices/api/proto"
	"NxcFull/backend/orderServices/internal/model"
	"time"
)

func ConvertOrder2pbOrder(getModel []model.Orders) []*pb.OrderInfo {
	var pbOrders []*pb.OrderInfo

	for _, order := range getModel {
		pbOrder := &pb.OrderInfo{
			Id:             order.Id,
			UserId:         order.UserId,
			Email:          order.Email,
			PlanId:         order.PlanId,
			PaymentMethod:  order.PaymentMethod,
			Period:         order.Period,
			CouponId:       order.CouponId,
			Status:         order.Status,
			IsSuccess:      order.IsSuccess,
			FailureReason:  order.FailureReason,
			DiscountAmount: float32(order.DiscountAmount),
			Amount:         float32(order.Amount),
			PaidAt:         formatTime(order.PaidAt), // 将 *time.Time 转换为 string
			CreatedAt:      order.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:      order.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		pbOrders = append(pbOrders, pbOrder)
	}

	return pbOrders
}

// formatTime 用于处理 *time.Time 类型的时间并转换为 string
func formatTime(t *time.Time) string {
	if t != nil {
		return t.Format("2006-01-02 15:04:05")
	}
	return ""
}
