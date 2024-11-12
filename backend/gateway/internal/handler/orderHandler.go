package handler

import (
	sysContext "context"
	"encoding/json"
	"errors"
	orderPb "gateway/internal/grpc/api/order/proto"
	orderHandlePb "gateway/internal/grpc/api/orderHandle/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HandleGetAllMyOrders(context *gin.Context) {
	var userId int64
	var page int64
	var size int64
	var err error
	if userId, err = strconv.ParseInt(context.Query("user_id"), 10, 64); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "解析数据失败",
		})
	}
	if page, err = strconv.ParseInt(context.Query("page"), 10, 64); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "解析数据失败",
		})
	}
	if size, err = strconv.ParseInt(context.Query("size"), 10, 64); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "解析数据失败",
		})
	}
	resp, err := grpcClient.OrderServicesClient.GetAllMyOrders(sysContext.Background(), &orderPb.GetAllMyOrdersRequest{
		UserId: userId,
		Page:   page,
		Size:   size,
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
		"page_count": resp.PageCount,
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

func HandleCancelOrder(context *gin.Context) {
	if err, userId, orderId := getUserIdAndOrderIdFromContextViaGetMethod(context); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "提供的信息不正确",
		})
		return
	} else {
		resp, err := grpcClient.OrderHandleServicesClient.CancelOrder(sysContext.Background(), &orderHandlePb.CancelOrderRequest{
			OrderId: orderId,
			UserId:  userId,
		})
		if err := failOnRpcError(err, resp); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code":      resp.Code,
			"cancelled": resp.Cancelled,
			"msg":       resp.Msg,
		})
	}
}

func HandlePlaceOrder(context *gin.Context) {
	if err, userId, orderId := getUserIdAndOrderIdFromContextViaPostMethod(context); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "提供的信息不正确",
		})
		return
	} else {
		log.Println("用户确认下单请求", userId, orderId)
		resp, err := grpcClient.OrderHandleServicesClient.PlaceOrder(sysContext.Background(), &orderHandlePb.PlaceOrderRequest{
			OrderId: orderId,
			UserId:  userId,
		})
		if err := failOnRpcError(err, resp); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code":          resp.Code,
			"order_id":      resp.OrderId,
			"msg":           resp.Msg,
			"placed":        resp.Placed,
			"expired_date":  resp.ExpiredDate,
			"key_generated": resp.KeyGenerated,
			"user_balance":  resp.UserBalance,
		})
	}
}

func HandleGetAllOrderByAdmin(context *gin.Context) {
	err, page, size := getPage2Size(context)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	searchEmail := context.Query("search_email")
	sort := context.Query("sort")
	resp, err := grpcClient.OrderServicesClient.GetAllOrdersAdmin(sysContext.Background(), &orderPb.GetAllOrdersAdminRequest{
		Page:        page,
		Size:        size,
		SearchEmail: searchEmail,
		Sort:        sort,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	var orderMap []map[string]interface{}
	if json.Unmarshal(resp.Orders, &orderMap) != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":       resp.Code,
		"orders":     orderMap,
		"msg":        resp.Msg,
		"page_count": resp.PageCount,
	})

}

func getUserIdAndOrderIdFromContextViaGetMethod(context *gin.Context) (err error, userId int64, orderId string) {
	orderId = context.Query("order_id")
	userIdStr := context.Query("user_id")
	userId, err = strconv.ParseInt(userIdStr, 10, 64)
	if orderId == "" || userIdStr == "" || err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "提供的信息不正确",
		})
	}
	log.Println(userId, orderId, err)
	return
}

func getUserIdAndOrderIdFromContextViaPostMethod(context *gin.Context) (err error, userId int64, orderId string) {
	postData := &struct {
		UserId  int64  `json:"user_id"`
		OrderId string `json:"order_id"`
	}{}
	if err = context.BindJSON(postData); err != nil {
		err = errors.New("提供的信息不正确" + err.Error())
		return
	}
	userId = postData.UserId
	orderId = postData.OrderId
	return
}

//func HandleGetAllOrderByAdmin(context *gin.Context) {
//
//}
