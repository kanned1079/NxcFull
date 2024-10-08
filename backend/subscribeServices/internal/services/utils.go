package services

import (
	pb "NxcFull/backend/subscribeServices/api/proto"
	"NxcFull/backend/subscribeServices/internal/model"
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

func ConvertPlan2pbPlan(getModel []model.Plan) []*pb.Plan {
	var pbPlans []*pb.Plan

	for _, plan := range getModel {
		pbPlan := pb.Plan{
			Id:            plan.Id,
			GroupId:       plan.GroupId,
			Name:          plan.Name,
			IsSale:        plan.IsSale,
			IsRenew:       plan.IsRenew,
			CapacityLimit: plan.CapacityLimit,
			Residue:       plan.Residue,
			Describe:      plan.Describe,
			MonthPrice:    plan.MonthPrice,
			QuarterPrice:  plan.QuarterPrice,
			HalfYearPrice: plan.HalfYearPrice,
			YearPrice:     plan.YearPrice,
			Sort:          plan.Sort,
			CreatedAt:     plan.CreatedAt.Format("2006-01-02 15:04:05"),      // 将时间转换为 string
			UpdatedAt:     plan.UpdatedAt.Format("2006-01-02 15:04:05"),      // 将时间转换为 string
			DeletedAt:     plan.DeletedAt.Time.Format("2006-01-02 15:04:05"), // DeletedAt 是 gorm.DeletedAt 类型，需检查是否有效
		}

		// 如果 DeletedAt 时间无效，不返回任何值
		if plan.DeletedAt.Valid {
			pbPlan.DeletedAt = plan.DeletedAt.Time.Format("2006-01-02 15:04:05")
		} else {
			pbPlan.DeletedAt = ""
		}

		pbPlans = append(pbPlans, &pbPlan)
	}

	return pbPlans
}
