package handler

import (
	pb "NxcFull/backend/gateway/internal/grpc/api/subscription/proto"

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
	context.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"my_plans": resp.MyPlans,
		"msg":      resp.Msg,
	})

	//var activeOrders []ActiveOrders
	//// Query the ActiveOrders table for the user_id and where IsActive is true
	//if err := dao.Db.Where("user_id = ? AND is_active = ?", userID, true).Find(&activeOrders).Error; err != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch active orders"})
	//	return
	//}
	//
	//if len(activeOrders) == 0 {
	//	context.JSON(http.StatusNotFound, gin.H{"error": "No active plans found for the user"})
	//	return
	//}
	//
	//// Create a result array to hold plan_name and expiration_date
	//var result []map[string]interface{}
	//
	//// Iterate through the active orders and fetch plan details
	//for _, subscription := range activeOrders {
	//	var plan subscribePlan.Plan
	//	// Fetch the plan details using the PlanId
	//	if err := dao.Db.Where("id = ?", subscription.PlanId).First(&plan).Error; err != nil {
	//		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch plan details"})
	//		return
	//	}
	//
	//	// Add the plan_name and expiration_date to the result
	//	result = append(result, map[string]interface{}{
	//		"plan_name":       plan.Name,
	//		"expiration_date": subscription.ExpirationDate,
	//	})
	//}
	//
	//// Return the result as a JSON response
	//context.JSON(http.StatusOK, gin.H{
	//	"code":     http.StatusOK,
	//	"my_plans": result,
	//	"msg":      "获取数据成功",
	//})
}

func HandleGetAllPlanKeyName(context *gin.Context) {
	//type planInfo struct {
	//	Id   int64  `json:"id"`
	//	Name string `json:"name"`
	//}
	//var planArr = []planInfo{}
	//if result := dao.Db.Model(&Plan{}).Select("id, name").Order("sort ASC").Find(&planArr); result.Error != nil {
	//	log.Println(result.Error)
	//}
	//log.Println(planArr)
	//context.JSON(http.StatusOK, gin.H{
	//	"code":  http.StatusOK,
	//	"plans": planArr,
	//})
	resp, err := grpcClient.SubscriptionServiceClient.GetAllPlanKeyName(sysContext.Background(), &pb.GetAllPlanKeyNameRequest{})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code":  resp.Code,
		"plans": resp.Plans,
	})
}

// HandleCommitNewOrder 用户提交一个新的订单
//func HandleCommitNewOrder(context *gin.Context) {
//	var err error
//	postData := &struct {
//		UserId   int64  `json:"user_id"`   // 用户Id
//		PlanId   int64  `json:"plan_id"`   // 订阅计划Id
//		Period   string `json:"period"`    // 付款周期 month, quater, half-year, year
//		CouponId int64  `json:"coupon_id"` // 如果没使用优惠券这就是0否则为优惠券的id
//	}{}
//	if err := context.ShouldBindJSON(postData); err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"error": err.Error(),
//			"msg":   "解析表单错误",
//		})
//	}
//	log.Printf("UserId %v\n PlanId %v\n Period %v\n CouponId %v\n", postData.UserId, postData.PlanId, postData.Period, postData.CouponId)
//	beginningOfDay := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
//	var orderIdGenerated int
//	if orderIdGenerated, err = strconv.Atoi(time.Now().Format("20060102") + strconv.FormatInt(time.Now().UnixMilli()-beginningOfDay.UnixMilli(), 10)); err != nil {
//		log.Println(err)
//	}
//	log.Println("newOrderId: ", orderIdGenerated)
//	var userMail string
//	if result := dao.Db.Model(&user.User{}).Where("id = ?", postData.UserId).Select("email").Limit(1).Find(&userMail); result.Error != nil {
//		log.Println("获取邮箱地址失败")
//	}
//	log.Println(userMail)
//
//	var couponInfo coupon.Coupon
//	if postData.CouponId != 0 {
//		if result := dao.Db.Model(&coupon.Coupon{}).Where("id = ?", postData.CouponId).First(&couponInfo); result.Error != nil {
//			log.Println(result.Error.Error())
//		}
//	}
//
//	var plan subscribePlan.Plan
//	if result := dao.Db.Where("id = ?", postData.PlanId).First(&plan); result.Error != nil {
//		log.Println(result.Error.Error())
//	}
//
//	if plan.Residue <= 0 {
//		log.Println("订阅剩余不足")
//		context.JSON(http.StatusOK, gin.H{
//			"code": http.StatusInternalServerError,
//			"msg":  "该计划余量不足",
//		})
//	}
//	var originalPrice float64
//	switch postData.Period {
//	case "month":
//		originalPrice = plan.MonthPrice
//	case "quarter":
//		originalPrice = plan.QuarterPrice
//	case "half-year":
//		originalPrice = plan.HalfYearPrice
//	case "year":
//		originalPrice = plan.YearPrice
//	}
//	log.Println("原始价格", originalPrice)
//
//	var discountPercent float64 = couponInfo.PercentOff / 100
//
//	var disCountAmount float64 = originalPrice * discountPercent
//
//	log.Println(discountPercent, disCountAmount, originalPrice-disCountAmount)
//
//	subscription := Orders{
//		Id:             int64(orderIdGenerated),        // 订单id
//		UserId:         postData.UserId,                // 用户id
//		PlanId:         postData.PlanId,                // 订阅计划id
//		Period:         postData.Period,                // 付款周期
//		Email:          userMail,                       // 用户邮箱
//		DiscountAmount: disCountAmount,                 // 抵折金额
//		Amount:         originalPrice - disCountAmount, // 最终应该支付金额
//		IsSuccess:      false,
//		Status:         0,
//		CouponId:       couponInfo.Id,
//	}
//
//	log.Println(subscription)
//
//	couponUsed := coupon.CouponUsage{
//		CouponId: couponInfo.Id,
//		UserId:   postData.UserId,
//		UsedAt:   time.Now(),
//	}
//
//	if postData.CouponId != 0 {
//		if result := dao.Db.Model(&coupon.CouponUsage{}).Create(&couponUsed); result.Error != nil {
//			context.JSON(http.StatusOK, gin.H{
//				"code":  http.StatusInternalServerError,
//				"error": result.Error.Error(),
//				"msg":   "添加优惠券使用记录出错",
//			})
//			return
//		}
//	}
//
//	if result := dao.Db.Model(&Orders{}).Create(&subscription); result.Error != nil {
//		context.JSON(http.StatusOK, gin.H{
//			"code":  http.StatusInternalServerError,
//			"error": result.Error.Error(),
//			"msg":   "创建新订单出错",
//		})
//	}
//
//	context.JSON(http.StatusOK, gin.H{
//		"code":            http.StatusOK,
//		"order_id":        subscription.Id,
//		"plan_name":       plan.Name,
//		"original_price":  originalPrice,
//		"discount_amount": disCountAmount,
//		"pay_price":       originalPrice - discountPercent,
//		"period":          postData.Period,
//		"created_at":      time.Now().Format("2006-01-02 15:04:05"),
//	})
//}
//
//func HandleGetOrders(context *gin.Context) {
//	var userId int
//	var err error
//	if userId, err = strconv.Atoi(context.Query("user_id")); err != nil {
//		log.Println(err)
//	}
//	log.Println("user_id", userId)
//	var orders []Orders
//	if result := dao.Db.Model(&Orders{}).Where("user_id = ?", userId).Find(&orders); result.Error != nil {
//		log.Println(result.Error)
//		context.JSON(http.StatusOK, gin.H{
//			"code":  http.StatusNotFound,
//			"error": result.Error.Error(),
//			"msg":   "查找用户订单错误",
//		})
//		return
//	}
//	context.JSON(http.StatusOK, gin.H{
//		"code":       http.StatusOK,
//		"order_list": orders,
//	})
//}

func HandleGetAllPlans(context *gin.Context) {
	//var plans []Plan
	////if result := dao.Db.Model(&Plan{}).Find(&plans); result.Error != nil {
	//if result := dao.Db.Model(&Plan{}).Order("sort ASC").Find(&plans); result.Error != nil {
	//	log.Println(result.Error)
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "获取订阅列表失败",
	//	})
	//}
	//log.Println(plans)
	//context.JSON(http.StatusOK, gin.H{
	//	"code":  http.StatusOK,
	//	"plans": plans,
	//})
	resp, err := grpcClient.SubscriptionServiceClient.GetAllPlans(sysContext.Background(), &pb.GetAllPlansRequest{})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
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
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
	//var newPlan = Plan{
	//	Name:          postData.Name,
	//	Describe:      postData.Describe,
	//	IsSale:        postData.IsSale,
	//	IsRenew:       postData.IsRenew,
	//	GroupId:       postData.GroupId,
	//	CapacityLimit: postData.CapacityLimit,
	//	Residue:       postData.CapacityLimit,
	//	MonthPrice:    postData.MonthPrice,
	//	QuarterPrice:  postData.QuarterPrice,
	//	HalfYearPrice: postData.HalfYearPrice,
	//	YearPrice:     postData.YearPrice,
	//	// 以下Sort用于在前端进行排序 优先级从低到高，添加新订阅时，它的优先级为上一个的优先级+1，如果表内为空不存在数据，则第一个添加的sort为1000
	//	// Sort:
	//	Sort: postData.Sort,
	//}
	//if result := dao.Db.Create(&newPlan); result.Error != nil {
	//	log.Println(result.Error)
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"code":  http.StatusInternalServerError,
	//		"msg":   "插入数据库失败",
	//		"error": result.Error.Error(),
	//	})
	//}
	//context.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "插入数据成功",
	//})
}
