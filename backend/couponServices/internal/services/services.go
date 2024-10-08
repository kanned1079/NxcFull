package services

import (
	pb "NxcFull/backend/couponServices/api/proto"
	"NxcFull/backend/couponServices/internal/dao"
	"NxcFull/backend/couponServices/internal/model"
	"context"
	"log"
	"net/http"
	"time"
)

type CouponService struct {
	pb.UnimplementedCouponServiceServer
}

func NewCouponService() *CouponService {
	return &CouponService{}
}

func (s *CouponService) AddNewCoupon(ctx context.Context, request *pb.AddNewCouponRequest) (*pb.AddNewCouponResponse, error) {
	startTime := time.Unix(0, request.StartTime*int64(time.Millisecond))
	endTime := time.Unix(0, request.EndTime*int64(time.Millisecond))
	var newCoupon model.Coupon = model.Coupon{
		Name:         request.Name,
		Code:         request.Code,
		PercentOff:   float64(request.PercentOff),
		Capacity:     request.Capacity,
		Residue:      request.Capacity,
		PerUserLimit: request.PerUserLimit,
		PlanLimit:    request.PlanLimit,
		StartTime:    &startTime,
		EndTime:      &endTime,
	}
	//log.Println(postData)
	if err := dao.Db.Model(&model.Coupon{}).Create(&newCoupon).Error; err != nil {
		//context.JSON(http.StatusOK, gin.H{
		//	"code":  http.StatusInternalServerError,
		//	"msg":   "插入数据失败",
		//	"error": err.Error(),
		//})
		return &pb.AddNewCouponResponse{
			Code: http.StatusInternalServerError,
			Msg:  "插入数据失败",
		}, nil
	}
	//context.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "插入数据成功",
	//})
	return &pb.AddNewCouponResponse{
		Code: http.StatusOK,
		Msg:  "插入数据失败",
	}, nil
}

func (s *CouponService) VerifyCoupon(ctx context.Context, request *pb.VerifyCouponRequest) (*pb.VerifyCouponResponse, error) {
	//// 接收前端传递的数据
	//postData := &struct {
	//	CouponCode string `json:"coupon_code"` // 优惠券代码
	//	PlanId     int64  `json:"plan_id"`     // 订阅计划ID
	//	UserId     int64  `json:"user_id"`     // 用户ID
	//}{}
	//
	//// 数据绑定错误处理
	//if err := context.ShouldBind(&postData); err != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code":  http.StatusInternalServerError,
	//		"msg":   "请求数据无效",
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//
	//log.Println("优惠券认证信息", postData)

	// 查找优惠券
	var coupon model.Coupon
	if err := dao.Db.Model(&model.Coupon{}).First(&coupon, "code = ?", request.CouponCode).Error; err != nil {
		//context.JSON(http.StatusOK, gin.H{
		//	"code":  http.StatusInternalServerError,
		//	"msg":   "优惠券无效或不存在",
		//	"error": err.Error(),
		//})
		return &pb.VerifyCouponResponse{
			Code: http.StatusInternalServerError,
			Msg:  "优惠券无效或不存在",
		}, nil
	}

	if !coupon.Enabled {
		//context.JSON(http.StatusOK, gin.H{
		//	"code": http.StatusInternalServerError,
		//	"meg":  "优惠券未启用",
		//})
		return &pb.VerifyCouponResponse{
			Code: http.StatusInternalServerError,
			Msg:  "优惠券未启用",
		}, nil
	}

	// 检查优惠券是否有指定的订阅计划限制
	if coupon.PlanLimit != 0 && coupon.PlanLimit != request.PlanId {
		//context.JSON(http.StatusOK, gin.H{
		//	"code": http.StatusInternalServerError,
		//	"msg":  "优惠券不适用于此订阅计划",
		//})
		//return
		return &pb.VerifyCouponResponse{
			Code: http.StatusInternalServerError,
			Msg:  "优惠券不适用于此订阅计划",
		}, nil
	}

	// 检查优惠券是否已过期或未开始
	now := time.Now()
	if (coupon.StartTime != nil && now.Before(*coupon.StartTime)) || (coupon.EndTime != nil && now.After(*coupon.EndTime)) {
		//context.JSON(http.StatusOK, gin.H{
		//	"code": http.StatusInternalServerError,
		//	"msg":  "优惠券不在有效期内",
		//})
		//return
		return &pb.VerifyCouponResponse{
			Code: http.StatusInternalServerError,
			Msg:  "优惠券不在有效期内",
		}, nil
	}

	// 检查优惠券的剩余数量
	if coupon.Residue <= 0 {
		//context.JSON(http.StatusOK, gin.H{
		//	"code": http.StatusInternalServerError,
		//	"msg":  "优惠券已使用完",
		//})
		//return
		return &pb.VerifyCouponResponse{
			Code: http.StatusInternalServerError,
			Msg:  "优惠券已使用完",
		}, nil
	}

	// 查询当前用户使用此优惠券的次数
	var usedCount int64
	dao.Db.Model(&model.CouponUsage{}).Where("user_id = ? AND coupon_id = ?", request.UserId, coupon.Id).Count(&usedCount)
	if coupon.PerUserLimit > 0 && usedCount >= coupon.PerUserLimit {
		//context.JSON(http.StatusOK, gin.H{
		//	"code": http.StatusBadRequest,
		//	"msg":  "您已达到优惠券的使用次数限制",
		//})
		//return
		return &pb.VerifyCouponResponse{
			Code: http.StatusInternalServerError,
			Msg:  "您已达到优惠券的使用次数限制",
		}, nil
	}

	// 验证通过，返回优惠信息
	//context.JSON(http.StatusOK, gin.H{
	//	"code":        http.StatusOK,
	//	"verified":    true,
	//	"msg":         "优惠券有效",
	//	"percent_off": coupon.PercentOff,
	//	"name":        coupon.Name,
	//	"id":          coupon.Id,
	//})
	return &pb.VerifyCouponResponse{
		Code:       http.StatusOK,
		Msg:        "优惠券有效",
		Verified:   true,
		PercentOff: float32(coupon.PercentOff),
		Name:       coupon.Name,
		Id:         coupon.Id,
	}, nil
}

func (s *CouponService) GetAllCoupons(ctx context.Context, request *pb.GetAllCouponsRequest) (*pb.GetAllCouponsResponse, error) {
	var coupons []model.Coupon
	if err := dao.Db.Model(&model.Coupon{}).Find(&coupons).Error; err != nil {
		//context.JSON(http.StatusOK, gin.H{
		//	"code": http.StatusInternalServerError,
		//	"msg":  "获取优惠券列表失败",
		//})
		return &pb.GetAllCouponsResponse{
			Code: http.StatusInternalServerError,
			Msg:  "获取优惠券列表失败",
		}, nil
	}
	//context.JSON(http.StatusOK, gin.H{
	//	"code":        http.StatusOK,
	//	"coupon_list": coupons,
	//})
	return &pb.GetAllCouponsResponse{
		Code:       http.StatusOK,
		CouponList: Convert2rpcCoupons(coupons),
		Msg:        "查询成功",
	}, nil

	//return &pb.GetAllCouponsResponse{}, nil
}

func (s *CouponService) DeleteCoupon(ctx context.Context, request *pb.DeleteCouponRequest) (*pb.DeleteCouponResponse, error) {
	if result := dao.Db.Model(&model.Coupon{}).Delete(&model.Coupon{Id: int64(request.CouponId)}); result.Error != nil {
		log.Println(result.Error.Error())
		return &pb.DeleteCouponResponse{
			Code: http.StatusInternalServerError,
			Msg:  "删除错误",
		}, nil

	}

	return &pb.DeleteCouponResponse{
		Code: http.StatusOK,
		Msg:  "删除完成",
	}, nil

}

func (s *CouponService) ActivateCoupon(ctx context.Context, request *pb.ActivateCouponRequest) (*pb.ActivateCouponResponse, error) {
	var coupon model.Coupon
	if err := dao.Db.First(&coupon, request.Id).Error; err != nil {
		//context.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
		//return
		return &pb.ActivateCouponResponse{
			Code:    http.StatusInternalServerError,
			Message: "找不到记录",
		}, nil
	}

	// Update the Enabled field
	coupon.Enabled = request.Status
	log.Println(coupon)

	// Save the updated coupon
	if err := dao.Db.Model(&model.Coupon{}).Where("id = ?", coupon.Id).Save(&coupon).Error; err != nil {
		return &pb.ActivateCouponResponse{
			Code:    http.StatusInternalServerError,
			Message: "更新失败",
		}, nil
	}
	return &pb.ActivateCouponResponse{
		Code:    http.StatusOK,
		Message: "更新状态成功",
	}, nil

}
