package handler

import (
	pb "gateway/internal/grpc/api/subscription/proto"
	"strconv"

	sysContext "context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetActivePlanListByUserId 获取用户的有效订阅
func GetActivePlanListByUserId(context *gin.Context) {
	// Parse the user_id from the query parameters
	userID := context.Query("user_id")
	if userID == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	resp, err := grpcClient.SubscriptionServiceClient.GetActivePlanListByUserId(sysContext.Background(), &pb.GetActivePlanListByUserIdRequest{
		UserId: userID,
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
		"code":     resp.Code,
		"my_plans": resp.MyPlans,
		"msg":      resp.Msg,
	})
}

func HandleGetAllPlanKeyName(context *gin.Context) {
	resp, err := grpcClient.SubscriptionServiceClient.GetAllPlanKeyName(sysContext.Background(), &pb.GetAllPlanKeyNameRequest{})
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
		"code":  resp.Code,
		"plans": resp.Plans,
	})
}

func HandleGetAllPlans(context *gin.Context) {
	var isUserQ bool
	var err error
	log.Println("is_user: ", context.Query("is_user"))
	if isUserQ, err = strconv.ParseBool(context.Query("is_user")); err != nil && !isUserQ {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "转换格式错误" + err.Error(),
		})
		return
	}
	resp, err := grpcClient.SubscriptionServiceClient.GetAllPlans(sysContext.Background(), &pb.GetAllPlansRequest{
		//User: isUser, // true用户 false管理员
		IsUser: isUserQ,
	})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":  resp.Code,
		"plans": resp.Plans,
		"msg":   resp.Msg,
	})
}

func HandleAddNewPlan(context *gin.Context) {
	postData := &struct {
		Name          string  `json:"name"`
		Describe      string  `json:"describe"`
		IsSale        bool    `json:"is_sale"`
		IsRenew       bool    `json:"is_renew"`
		GroupId       int64   `json:"group_id"`
		CapacityLimit int64   `json:"capacity_limit"`
		MonthPrice    float64 `json:"month_price"`
		QuarterPrice  float64 `json:"quarter_price"`
		HalfYearPrice float64 `json:"half_year_price"`
		YearPrice     float64 `json:"year_price"`
		Sort          int64   `json:"sort"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"msg":   "绑定数据失败",
			"error": err.Error(),
		})
	}
	log.Println(postData)
	resp, err := grpcClient.SubscriptionServiceClient.AddNewPlan(sysContext.Background(), &pb.AddNewPlanRequest{
		Name:          postData.Name,
		Describe:      postData.Describe,
		IsSale:        postData.IsSale,
		IsRenew:       postData.IsRenew,
		GroupId:       postData.GroupId,
		CapacityLimit: postData.CapacityLimit,
		MonthPrice:    postData.MonthPrice,
		QuarterPrice:  postData.QuarterPrice,
		HalfYearPrice: postData.HalfYearPrice,
		YearPrice:     postData.YearPrice,
		Sort:          postData.Sort,
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

func HandleUpdatePlan(context *gin.Context) {
	postData := &struct {
		Id            int64   `json:"id"`
		Name          string  `json:"name"`
		Describe      string  `json:"describe"`
		IsSale        bool    `json:"is_sale"`
		IsRenew       bool    `json:"is_renew"`
		GroupId       int64   `json:"group_id"`
		CapacityLimit int64   `json:"capacity_limit"`
		MonthPrice    float64 `json:"month_price"`
		QuarterPrice  float64 `json:"quarter_price"`
		HalfYearPrice float64 `json:"half_year_price"`
		YearPrice     float64 `json:"year_price"`
		Sort          int64   `json:"sort"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"msg":   "绑定数据失败",
			"error": err.Error(),
		})
	}
	log.Println(postData)
	resp, err := grpcClient.SubscriptionServiceClient.UpdatePlan(sysContext.Background(), &pb.UpdatePlanRequest{
		Id:            postData.Id,
		Name:          postData.Name,
		Describe:      postData.Describe,
		IsSale:        postData.IsSale,
		IsRenew:       postData.IsRenew,
		GroupId:       postData.GroupId,
		CapacityLimit: postData.CapacityLimit,
		MonthPrice:    postData.MonthPrice,
		QuarterPrice:  postData.QuarterPrice,
		HalfYearPrice: postData.HalfYearPrice,
		YearPrice:     postData.YearPrice,
		Sort:          postData.Sort,
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

func HandleDeletePlan(context *gin.Context) {
	var planId int
	var err error
	if planId, err = strconv.Atoi(context.Query("plan_id")); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	log.Println(planId)
	resp, err := grpcClient.SubscriptionServiceClient.DeletePlan(sysContext.Background(), &pb.DeletePlanRequest{
		PlanId: int64(planId),
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

func HandleUpdatePlanSale(context *gin.Context) {
	postData := &struct {
		Id     int64 `json:"id"`
		IsSale bool  `json:"is_sale"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "绑定数据失败",
		})
		return
	}
	log.Println("更新Sale:", postData)
	// UpdatePlanIsSale 调用gRPC服务器来更新销售状态
	resp, err := grpcClient.SubscriptionServiceClient.UpdatePlanIsSale(sysContext.Background(), &pb.UpdatePlanIsSaleRequest{
		Id:     postData.Id,
		IsSale: postData.IsSale,
	})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败" + err.Error(),
		})
		return
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

func HandleUpdatePlanRenew(context *gin.Context) {
	postData := &struct {
		Id      int64 `json:"id"`
		IsRenew bool  `json:"is_renew"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "绑定数据失败",
		})
		return
	}
	log.Println("更新Renew:", postData)

	// UpdatePlanIsSale 调用gRPC服务器来更新是否启用续费
	resp, err := grpcClient.SubscriptionServiceClient.UpdatePlanRenew(sysContext.Background(), &pb.UpdatePlanRenewRequest{
		Id:      postData.Id,
		IsRenew: postData.IsRenew,
	})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败" + err.Error(),
		})
		return
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
