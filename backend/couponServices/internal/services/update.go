package services

import (
	"context"
	pb "couponServices/api/proto"
	"couponServices/internal/dao"
	"couponServices/internal/model"
	"net/http"
	"time"
)

func (s *CouponService) UpdateCoupon(ctx context.Context, request *pb.UpdateCouponRequest) (*pb.UpdateCouponResponse, error) {
	if request.StartTime > request.EndTime { // 确保启用时间小于结束时间
		return &pb.UpdateCouponResponse{
			Code: http.StatusConflict,
			Msg:  "start time must be less than end time",
		}, nil
	}

	startTime := time.Unix(request.StartTime/1000, (request.StartTime%1000)*int64(time.Millisecond)) // 处理毫秒部分
	endTime := time.Unix(request.EndTime/1000, (request.EndTime%1000)*int64(time.Millisecond))       // 处理毫秒部分

	if result := dao.Db.Model(&model.Coupon{}).Where("`id` = ?", request.Id).Updates(map[string]interface{}{
		"code":           request.Code,
		"name":           request.Name,
		"percent_off":    float64(request.PercentOff),
		"start_time":     &startTime,
		"end_time":       &endTime,
		"per_user_limit": request.PerUserLimit,
		"capacity":       request.Capacity,
		"residue":        request.Residue,
		"plan_limit":     request.PlanLimit,
	}); result.RowsAffected == 0 {
		return &pb.UpdateCouponResponse{
			Code: http.StatusInternalServerError,
			Msg:  result.Error.Error(),
		}, nil
	}

	return &pb.UpdateCouponResponse{
		Code: http.StatusOK,
		Msg:  "success",
	}, nil
}
