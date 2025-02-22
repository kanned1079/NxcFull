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
	if err = failOnRpcError(err, context.Request.RequestURI); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//var orders []map[string]any
	//if err := json.Unmarshal(resp.OrderList, &orders); err != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "转换数据格式失败" + err.Error(),
	//	})
	//	return
	//}
	context.JSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"order_list": json.RawMessage(resp.OrderList),
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
	if err = failOnRpcError(err, context.Request.RequestURI); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
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
	if err = failOnRpcError(err, context.Request.RequestURI); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//var orderInfoKvData map[string]interface{}
	//if resp.Code == 200 {
	//	if err := json.Unmarshal(resp.OrderInfo, &orderInfoKvData); err != nil {
	//		context.JSON(http.StatusOK, gin.H{
	//			"code": http.StatusInternalServerError,
	//			"msg":  "格式化失败",
	//		})
	//		return
	//	}
	//}
	context.JSON(http.StatusOK, gin.H{
		"code":           resp.Code,
		"msg":            resp.Msg,
		"is_finished":    resp.IsFinished,
		"is_success":     resp.IsSuccess,
		"failure_reason": resp.FailureReason,
		"order_info":     json.RawMessage(resp.OrderInfo),
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
	//var orderMap []map[string]interface{}
	//if json.Unmarshal(resp.Orders, &orderMap) != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  err.Error(),
	//	})
	//	return
	//}
	context.JSON(http.StatusOK, gin.H{
		"code":       resp.Code,
		"orders":     json.RawMessage(resp.Orders),
		"msg":        resp.Msg,
		"page_count": resp.PageCount,
	})
}

// HandleManualCancelUserOrder DELETE请求
func HandleManualCancelUserOrder(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Query("user_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":      http.StatusBadRequest,
			"cancelled": false,
			"msg":       "Bad request",
		})
		return
	}
	orderId := context.Query("order_id")
	log.Println(userId, orderId)
	resp, err := grpcClient.OrderHandleServicesClient.ManualCancelOrderPayment(sysContext.Background(), &orderHandlePb.ManualCancelOrderPaymentRequest{
		OrderId: orderId,
		UserId:  userId,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":      http.StatusInternalServerError,
			"cancelled": false,
			"msg":       err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":      resp.Code,
		"cancelled": resp.Cancelled,
		"msg":       resp.Msg,
	})
}

// HandleManualPassUserOrder PUT请求
func HandleManualPassUserOrder(context *gin.Context) {
	postData := &struct {
		UserId  int64  `json:"user_id"`
		OrderId string `json:"order_id"`
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":      http.StatusBadRequest,
			"cancelled": false,
			"msg":       "Bad request",
		})
		return
	}
	resp, err := grpcClient.OrderHandleServicesClient.ManualPassOrderPayment(sysContext.Background(), &orderHandlePb.ManualPassOrderPaymentRequest{
		OrderId: postData.OrderId,
		UserId:  postData.UserId,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":      http.StatusInternalServerError,
			"cancelled": false,
			"msg":       err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":   resp.Code,
		"passed": resp.Passed,
		"msg":    resp.Msg,
	})
}

//func HandleUserTopUp(context *gin.Context) {
//	var err error
//	postData := &struct {
//		UserId        int64   `json:"user_id"`
//		PaymentMethod string  `json:"payment_method"`
//		TopUpAmount   float32 `json:"top_up_amount"`
//		Discount      float32 `json:"discount"`
//		Confirmed     bool    `json:"confirmed"`
//		CommittedAt
//	}{}
//	if err = context.ShouldBind(&postData); err != nil {
//		context.JSON(http.StatusOK, gin.H{
//			"code":    http.StatusBadRequest,
//			"created": false,
//			"msg":     "Bad request" + err.Error(),
//		})
//	}
//	resp, err := grpcClient.OrderServicesClient.CommitNewTopUpOrder(sysContext.Background(), &orderPb.CommitNewTopUpOrderRequest{
//		UserId:        postData.UserId,
//		PaymentMethod: postData.PaymentMethod,
//		Amount:        postData.TopUpAmount,
//		Discount:      postData.Discount,
//	})
//	if err := failOnRpcError(err, resp); err != nil {
//		context.JSON(http.StatusOK, gin.H{
//			"code":    http.StatusInternalServerError,
//			"created": false,
//			"msg":     err.Error(),
//		})
//		return
//	}
//	context.JSON(http.StatusOK, gin.H{
//		"code":     resp.Code,
//		"created":  resp.Created,
//		"qr_code":  resp.QrCode,
//		"order_id": resp.OrderId,
//		"msg":      resp.Msg,
//	})
//}
