package handler

import (
	pb "NxcFull/backend/gateway/internal/grpc/api/proto"
	sysContext "context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HandleAddNewCoupon(context *gin.Context) {
	postData := &struct {
		Name         string  `json:"name"`
		Code         string  `json:"code"`
		PercentOff   float64 `json:"percent_off"`
		Capacity     int64   `json:"capacity"`
		PerUserLimit int64   `json:"per_user_limit"`
		PlanLimit    int64   `json:"plan_limit"`
		StartTime    int64   `json:"start_time"`
		EndTime      int64   `json:"end_time"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "绑定数据失败",
			"error": err.Error()})
	}
	resp, err := grpcClient.CouponServiceClient.AddNewCoupon(sysContext.Background(), &pb.AddNewCouponRequest{
		Name:         postData.Name,
		Code:         postData.Code,
		PercentOff:   float32(postData.PercentOff),
		Capacity:     postData.Capacity,
		PerUserLimit: postData.PerUserLimit,
		PlanLimit:    postData.PlanLimit,
		StartTime:    postData.StartTime,
		EndTime:      postData.EndTime,
	})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
	//
	//startTime := time.Unix(0, postData.StartTime*int64(time.Millisecond))
	//endTime := time.Unix(0, postData.EndTime*int64(time.Millisecond))
	//var newCoupon Coupon = Coupon{
	//	Name:         postData.Name,
	//	Code:         postData.Code,
	//	PercentOff:   postData.PercentOff,
	//	Capacity:     postData.Capacity,
	//	Residue:      postData.Capacity,
	//	PerUserLimit: postData.PerUserLimit,
	//	PlanLimit:    postData.PlanLimit,
	//	StartTime:    &startTime,
	//	EndTime:      &endTime,
	//}
	//log.Println(postData)
	//if err := dao.Db.Model(&Coupon{}).Create(&newCoupon).Error; err != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code":  http.StatusInternalServerError,
	//		"msg":   "插入数据失败",
	//		"error": err.Error(),
	//	})
	//}
	//context.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "插入数据成功",
	//})
}

// HandleVerifyCoupon 验证优惠券
func HandleVerifyCoupon(context *gin.Context) {
	// 接收前端传递的数据
	postData := &struct {
		CouponCode string `json:"coupon_code"` // 优惠券代码
		PlanId     int64  `json:"plan_id"`     // 订阅计划ID
		UserId     int64  `json:"user_id"`     // 用户ID
	}{}

	// 数据绑定错误处理
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "请求数据无效",
			"error": err.Error(),
		})
		return
	}

	log.Println("优惠券认证信息", postData)

	resp, err := grpcClient.CouponServiceClient.VerifyCoupon(sysContext.Background(), &pb.VerifyCouponRequest{
		CouponCode: postData.CouponCode,
		PlanId:     postData.PlanId,
		UserId:     postData.UserId,
	})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	//// 验证通过，返回优惠信息
	context.JSON(http.StatusOK, gin.H{
		"code":        resp.Code,
		"verified":    true,
		"msg":         "优惠券有效",
		"percent_off": resp.PercentOff,
		"name":        resp.Name,
		"id":          resp.Id,
	})

	//// 查找优惠券
	//var coupon Coupon
	//if err := dao.Db.Model(&Coupon{}).First(&coupon, "code = ?", postData.CouponCode).Error; err != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code":  http.StatusInternalServerError,
	//		"msg":   "优惠券无效或不存在",
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//
	//if !coupon.Enabled {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"meg":  "优惠券未启用",
	//	})
	//}
	//
	//// 检查优惠券是否有指定的订阅计划限制
	//if coupon.PlanLimit != 0 && coupon.PlanLimit != postData.PlanId {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "优惠券不适用于此订阅计划",
	//	})
	//	return
	//}
	//
	//// 检查优惠券是否已过期或未开始
	//now := time.Now()
	//if (coupon.StartTime != nil && now.Before(*coupon.StartTime)) || (coupon.EndTime != nil && now.After(*coupon.EndTime)) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "优惠券不在有效期内",
	//	})
	//	return
	//}
	//
	//// 检查优惠券的剩余数量
	//if coupon.Residue <= 0 {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "优惠券已使用完",
	//	})
	//	return
	//}
	//
	//// 查询当前用户使用此优惠券的次数
	//var usedCount int64
	//dao.Db.Model(&CouponUsage{}).Where("user_id = ? AND coupon_id = ?", postData.UserId, coupon.Id).Count(&usedCount)
	//if coupon.PerUserLimit > 0 && usedCount >= coupon.PerUserLimit {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusBadRequest,
	//		"msg":  "您已达到优惠券的使用次数限制",
	//	})
	//	return
	//}
	//
	//// 验证通过，返回优惠信息
	//context.JSON(http.StatusOK, gin.H{
	//	"code":        http.StatusOK,
	//	"verified":    true,
	//	"msg":         "优惠券有效",
	//	"percent_off": coupon.PercentOff,
	//	"name":        coupon.Name,
	//	"id":          coupon.Id,
	//})
}

func HandleGetAllCoupons(context *gin.Context) {
	resp, err := grpcClient.CouponServiceClient.GetAllCoupons(sysContext.Background(), &pb.GetAllCouponsRequest{})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code":        resp.Code,
		"coupon_list": resp.CouponList,
	})
	//var coupons []Coupon
	//if err := dao.Db.Model(&Coupon{}).Find(&coupons).Error; err != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "获取优惠券列表失败",
	//	})
	//}
	//context.JSON(http.StatusOK, gin.H{
	//	"code":        http.StatusOK,
	//	"coupon_list": coupons,
	//})

}

// HandleActiveCoupon 修改优惠券的 Enabled字段
func HandleActiveCoupon(context *gin.Context) {
	postData := &struct {
		Id     int64 `json:"id"`
		Status bool  `json:"status"`
	}{}

	// Bind the incoming JSON to postData
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Invalid request body",
		})
		return
	}

	log.Println(postData.Id, postData.Status)

	resp, err := grpcClient.CouponServiceClient.ActivateCoupon(sysContext.Background(), &pb.ActivateCouponRequest{
		Id:     postData.Id,
		Status: postData.Status,
	})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Message,
	})

	//// Find the coupon by Id
	//var coupon Coupon
	//if err := dao.Db.First(&coupon, postData.Id).Error; err != nil {
	//	context.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
	//	return
	//}
	//
	//// Update the Enabled field
	//coupon.Enabled = postData.Status
	//log.Println(coupon)
	//
	//// Save the updated coupon
	//if err := dao.Db.Model(&Coupon{}).Where("id = ?", coupon.Id).Save(&coupon).Error; err != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update coupon status"})
	//	return
	//}
	//
	//// Return success response
	//context.JSON(http.StatusOK, gin.H{
	//	"code":    http.StatusOK,
	//	"message": "Coupon status updated successfully",
	//})
}

func HandleDeleteCoupon(context *gin.Context) {
	deleteCoupon := &struct {
		CouponId int `json:"coupon_id"`
	}{}
	log.Println(context.Query("coupon_id"))
	var err error
	if deleteCoupon.CouponId, err = strconv.Atoi(context.Query("coupon_id")); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
	}

	log.Println("消息id: ", deleteCoupon.CouponId)
	//if result := dao.Db.Model(&Coupon{}).Delete(&Coupon{Id: int64(deleteCoupon.CouponId)}); err != nil {
	//	log.Println(result.Error.Error())
	//}
	resp, err := grpcClient.CouponServiceClient.DeleteCoupon(sysContext.Background(), &pb.DeleteCouponRequest{
		CouponId: int64(deleteCoupon.CouponId),
	})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  "删除成功",
	})
}
