package services

import (
	"context"
	"errors"
	"log"
	"subscribeServices/internal/dao"
)

//func ConvertOrder2pbOrder(getModel []model.Orders) []*pb.OrderInfo {
//	var pbOrders []*pb.OrderInfo
//	for _, order := range getModel {
//		pbOrder := &pb.OrderInfo{
//			Id:             order.Id,
//			UserId:         order.UserId,
//			Email:          order.Email,
//			PlanId:         order.PlanId,
//			PaymentMethod:  order.PaymentMethod,
//			Period:         order.Period,
//			CouponId:       order.CouponId,
//			Status:         order.Status,
//			IsSuccess:      order.IsSuccess,
//			FailureReason:  order.FailureReason,
//			DiscountAmount: float32(order.DiscountAmount),
//			Amount:         float32(order.Amount),
//			PaidAt:         formatTime(order.PaidAt), // 将 *time.Time 转换为 string
//			CreatedAt:      order.CreatedAt.Format("2006-01-02 15:04:05"),
//			UpdatedAt:      order.UpdatedAt.Format("2006-01-02 15:04:05"),
//		}
//
//		pbOrders = append(pbOrders, pbOrder)
//	}
//
//	return pbOrders
//}
//
//// formatTime 用于处理 *time.Time 类型的时间并转换为 string
//func formatTime(t *time.Time) string {
//	if t != nil {
//		return t.Format("2006-01-02 15:04:05")
//	}
//	return ""
//}
//
//func ConvertPlan2pbPlan(getModel []model.Plan) []*pb.Plan {
//	var pbPlans []*pb.Plan
//
//	for _, plan := range getModel {
//		pbPlan := pb.Plan{
//			Id:            plan.Id,
//			GroupId:       plan.GroupId,
//			Name:          plan.Name,
//			IsSale:        plan.IsSale,
//			IsRenew:       plan.IsRenew,
//			CapacityLimit: plan.CapacityLimit,
//			Residue:       plan.Residue,
//			Describe:      plan.Describe,
//			MonthPrice:    plan.MonthPrice,
//			QuarterPrice:  plan.QuarterPrice,
//			HalfYearPrice: plan.HalfYearPrice,
//			YearPrice:     plan.YearPrice,
//			Sort:          plan.Sort,
//			CreatedAt:     plan.CreatedAt.Format("2006-01-02 15:04:05"),      // 将时间转换为 string
//			UpdatedAt:     plan.UpdatedAt.Format("2006-01-02 15:04:05"),      // 将时间转换为 string
//			DeletedAt:     plan.DeletedAt.Time.Format("2006-01-02 15:04:05"), // DeletedAt 是 gorm.DeletedAt 类型，需检查是否有效
//		}
//
//		// 如果 DeletedAt 时间无效，不返回任何值
//		if plan.DeletedAt.Valid {
//			pbPlan.DeletedAt = plan.DeletedAt.Time.Format("2006-01-02 15:04:05")
//		} else {
//			pbPlan.DeletedAt = ""
//		}
//
//		pbPlans = append(pbPlans, &pbPlan)
//	}
//
//	return pbPlans
//}

// ClearPlansRedisCache 用于清除缓存
//func ClearPlansRedisCache(ctx context.Context) (err error) {
//	// 清空用户计划缓存
//	if err = dao.Rdb.Del(ctx, "plans:user:all").Err(); err != nil {
//		//log.Println("Failed to delete user plans from Redis:", err)
//		return errors.New("Failed to delete user plans from Redis")
//	}
//	// 清空管理员计划缓存
//	if err = dao.Rdb.Del(ctx, "plans:admin:all").Err(); err != nil {
//		//log.Println("Failed to delete admin plans from Redis:", err)
//		return errors.New("Failed to delete admin plans from Redis")
//	}
//	return
//}

func ClearPlansRedisCache(ctx context.Context) error {
	var cursor uint64
	//var err error

	for {
		// 使用 SCAN 查找带有 "plans:" 前缀的所有键
		keys, nextCursor, scanErr := dao.Rdb.Scan(ctx, cursor, "plans:*", 100).Result()
		if scanErr != nil {
			return errors.New("Failed to scan Redis for plan keys")
		}

		// 删除找到的键
		if len(keys) > 0 {
			if delErr := dao.Rdb.Del(ctx, keys...).Err(); delErr != nil {
				return errors.New("Failed to delete plan keys from Redis")
			}
		}

		// 更新游标，如果游标为 0 表示扫描完毕
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	log.Println("清除plans缓存成功")
	return nil
}
