package handler

import (
	pb "NxcFull/backend/gateway/internal/grpc/api/coupon/proto"
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
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
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
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	log.Println(resp)
	//// 验证通过，返回优惠信息
	context.JSON(http.StatusOK, gin.H{
		"code":        resp.Code,
		"verified":    true,
		"msg":         "优惠券有效",
		"percent_off": resp.PercentOff,
		"name":        resp.Name,
		"id":          resp.Id,
	})
}

func HandleGetAllCoupons(context *gin.Context) {
	resp, err := grpcClient.CouponServiceClient.GetAllCoupons(sysContext.Background(), &pb.GetAllCouponsRequest{})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":        resp.Code,
		"coupon_list": resp.CouponList,
	})

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
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Message,
	})
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
	resp, err := grpcClient.CouponServiceClient.DeleteCoupon(sysContext.Background(), &pb.DeleteCouponRequest{
		CouponId: int64(deleteCoupon.CouponId),
	})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  "删除成功",
	})
}
