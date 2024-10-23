package handler

import (
	sysContext "context"
	"encoding/json"
	orderPb "gateway/internal/grpc/api/order/proto"
	orderHandlePb "gateway/internal/grpc/api/orderHandle/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HandleGetAllMyOrders(context *gin.Context) {
	var userId int
	var err error
	if userId, err = strconv.Atoi(context.Query("user_id")); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "解析数据失败",
		})
	}
	resp, err := grpcClient.OrderServicesClient.GetAllMyOrders(sysContext.Background(), &orderPb.GetAllMyOrdersRequest{
		UserId: int64(userId),
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "rpc服务调用出错" + err.Error(),
		})
		return
	}
	if resp == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "rpc调用错误无返回值",
		})
		return
	}
	var orders []map[string]any
	if err := json.Unmarshal(resp.OrderList, &orders); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "转换数据格式失败" + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"order_list": orders,
	})
}

func HandleCommitNewOrder(context *gin.Context) {
	//var err error
	postData := &struct {
		UserId   int64  `json:"user_id"`   // 用户Id
		PlanId   int64  `json:"plan_id"`   // 订阅计划Id
		Period   string `json:"period"`    // 付款周期 month, quater, half-year, year
		CouponId int64  `json:"coupon_id"` // 如果没使用优惠券这就是0否则为优惠券的id
	}{}
	if err := context.ShouldBindJSON(postData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "解析表单错误",
		})
		return
	}
	log.Printf("UserId %v\n PlanId %v\n Period %v\n CouponId %v\n", postData.UserId, postData.PlanId, postData.Period, postData.CouponId)
	resp, err := grpcClient.OrderServicesClient.CommitNewOrder(sysContext.Background(), &orderPb.CommitNewOrderRequest{
		UserId:   postData.UserId,
		PlanId:   postData.PlanId,
		Period:   postData.Period,
		CouponId: postData.CouponId,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "rpc服务调用出错" + err.Error(),
		})
		return
	}
	if resp == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "rpc调用错误无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":            resp.Code,
		"order_id":        resp.OrderId,
		"plan_name":       resp.PlanName,
		"original_price":  resp.OriginalPrice,
		"discount_amount": resp.DiscountAmount,
		"pay_price":       resp.PayPrice,
		"period":          resp.Period,
		"created_at":      resp.CreatedAt,
	})

}

// HandleCheckOrderStatus 轮训检查用户的订单状态
func HandleCheckOrderStatus(context *gin.Context) {
	orderId := context.Query("order_id")
	userIdStr := context.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if orderId == "" || userIdStr == "" || err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "提供的信息不正确",
		})
		return
	}
	resp, err := grpcClient.OrderHandleServicesClient.GetOrderStatus(sysContext.Background(), &orderHandlePb.GetOrderStatusRequest{
		OrderId: orderId,
		UserId:  userId,
	})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "rpc服务调用出错" + err.Error(),
		})
		return
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusNotFound,
			"msg":  "rpc调用失败无返回值",
		})
		return
	}
	var orderInfoKvData map[string]any
	if err := json.Unmarshal(resp.OrderInfo, &orderInfoKvData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "格式化失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":           resp.Code,
		"msg":            resp.Msg,
		"is_finished":    resp.IsFinished,
		"is_success":     resp.IsSuccess,
		"failure_reason": resp.FailureReason,
		"order_info":     orderInfoKvData,
	})

}
