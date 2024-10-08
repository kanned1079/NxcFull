package services

import (
	pb "NxcFull/backend/orderServices/api/proto"
	"NxcFull/backend/orderServices/internal/dao"
	"NxcFull/backend/orderServices/internal/model"
	"NxcFull/nxc-backend/coupon"
	"NxcFull/nxc-backend/subscribePlan"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
)

type OrderServices struct {
	pb.UnimplementedSubscriptionServiceServer
}

func NewOrderServices() *OrderServices {
	return &OrderServices{}
}

func (s *OrderServices) GetActivePlanListByUserId(ctx context.Context, request *pb.GetActivePlanListByUserIdRequest) (*pb.GetActivePlanListByUserIdResponse, error) {
	var activeOrders []model.ActiveOrders
	// Query the ActiveOrders table for the user_id and where IsActive is true
	if err := dao.Db.Where("user_id = ? AND is_active = ?", request.UserId, true).Find(&activeOrders).Error; err != nil {
		//context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch active orders"})
		//return
		return &pb.GetActivePlanListByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "获取有效的订单失败",
		}, nil
	}

	if len(activeOrders) == 0 {
		return &pb.GetActivePlanListByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "获取有效的订单失败",
		}, nil
	}

	// Create a result array to hold plan_name and expiration_date
	var result []map[string]interface{}
	var myPlanResult []*pb.PlanInfo

	// Iterate through the active orders and fetch plan details
	for _, order := range activeOrders {
		var plan model.Plan
		// Fetch the plan details using the PlanId
		if err := dao.Db.Where("id = ?", order.PlanId).First(&plan).Error; err != nil {
			//context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch plan details"})
			//return
			return &pb.GetActivePlanListByUserIdResponse{
				Code: http.StatusInternalServerError,
				Msg:  "获取订阅的细节失败",
			}, nil
		}

		// Add the plan_name and expiration_date to the result
		result = append(result, map[string]interface{}{
			"plan_name":       plan.Name,
			"expiration_date": order.ExpirationDate,
		})
		myPlanResult = append(myPlanResult, &pb.PlanInfo{
			PlanName:       plan.Name,
			ExpirationDate: order.ExpirationDate.Format("2006-01-02 15:04:05"),
		})
	}

	// Return the result as a JSON response
	//context.JSON(http.StatusOK, gin.H{
	//	"code":     http.StatusOK,
	//	"my_plans": result,
	//	"msg":      "获取数据成功",
	//})
	return &pb.GetActivePlanListByUserIdResponse{
		Code:    http.StatusOK,
		MyPlans: myPlanResult,
		Msg:     "获取成功",
	}, nil
}

func (s *OrderServices) CommitNewOrder(ctx context.Context, request *pb.CommitNewOrderRequest) (*pb.CommitNewOrderResponse, error) {
	var err error
	log.Printf("UserId %v\n PlanId %v\n Period %v\n CouponId %v\n", request.UserId, request.PlanId, request.Period, request.CouponId)
	beginningOfDay := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	var orderIdGenerated int
	if orderIdGenerated, err = strconv.Atoi(time.Now().Format("20060102") + strconv.FormatInt(time.Now().UnixMilli()-beginningOfDay.UnixMilli(), 10)); err != nil {
		log.Println(err)
	}
	log.Println("newOrderId: ", orderIdGenerated)
	var userMail string
	if result := dao.Db.Model(&model.User{}).Where("id = ?", request.UserId).Select("email").Limit(1).Find(&userMail); result.Error != nil {
		log.Println("获取邮箱地址失败")
	}
	log.Println(userMail)

	var couponInfo coupon.Coupon
	if request.CouponId != 0 {
		if result := dao.Db.Model(&coupon.Coupon{}).Where("id = ?", request.CouponId).First(&couponInfo); result.Error != nil {
			log.Println(result.Error.Error())
		}
	}

	var plan subscribePlan.Plan
	if result := dao.Db.Where("id = ?", request.PlanId).First(&plan); result.Error != nil {
		log.Println(result.Error.Error())
	}

	if plan.Residue <= 0 {
		log.Println("订阅剩余不足")
		//context.JSON(http.StatusOK, gin.H{
		//	"code": http.StatusInternalServerError,
		//	"msg":  "该计划余量不足",
		//})
		return &pb.CommitNewOrderResponse{
			Code: http.StatusInternalServerError,
			Msg:  "订阅剩余不足",
		}, nil
	}
	var originalPrice float64
	switch request.Period {
	case "month":
		originalPrice = plan.MonthPrice
	case "quarter":
		originalPrice = plan.QuarterPrice
	case "half-year":
		originalPrice = plan.HalfYearPrice
	case "year":
		originalPrice = plan.YearPrice
	}
	log.Println("原始价格", originalPrice)

	var discountPercent float64 = couponInfo.PercentOff / 100

	var disCountAmount float64 = originalPrice * discountPercent

	log.Println(discountPercent, disCountAmount, originalPrice-disCountAmount)

	order := model.Orders{
		Id:             int64(orderIdGenerated),        // 订单id
		UserId:         request.UserId,                 // 用户id
		PlanId:         request.PlanId,                 // 订阅计划id
		Period:         request.Period,                 // 付款周期
		Email:          userMail,                       // 用户邮箱
		DiscountAmount: disCountAmount,                 // 抵折金额
		Amount:         originalPrice - disCountAmount, // 最终应该支付金额
		IsSuccess:      false,
		Status:         0,
		CouponId:       couponInfo.Id,
	}

	log.Println(order)

	couponUsed := model.CouponUsage{
		CouponId: couponInfo.Id,
		UserId:   request.UserId,
		UsedAt:   time.Now(),
	}

	if request.CouponId != 0 {
		if result := dao.Db.Model(&coupon.CouponUsage{}).Create(&couponUsed); result.Error != nil {
			//context.JSON(http.StatusOK, gin.H{
			//	"code":  http.StatusInternalServerError,
			//	"error": result.Error.Error(),
			//	"msg":   "添加优惠券使用记录出错",
			//})
			//return
			return &pb.CommitNewOrderResponse{
				Code: http.StatusInternalServerError,
				Msg:  "添加优惠券使用记录出错",
			}, nil
		}
	}

	if result := dao.Db.Model(&model.Orders{}).Create(&order); result.Error != nil {
		//context.JSON(http.StatusOK, gin.H{
		//	"code":  http.StatusInternalServerError,
		//	"error": result.Error.Error(),
		//	"msg":   "创建新订单出错",
		//})
		return &pb.CommitNewOrderResponse{
			Code: http.StatusInternalServerError,
			Msg:  "创建新订单出错",
		}, nil
	}

	//context.JSON(http.StatusOK, gin.H{
	//	"code":            http.StatusOK,
	//	"order_id":        subscribe.Id,
	//	"plan_name":       plan.Name,
	//	"original_price":  originalPrice,
	//	"discount_amount": disCountAmount,
	//	"pay_price":       originalPrice - discountPercent,
	//	"period":          postData.Period,
	//	"created_at":      time.Now().Format("2006-01-02 15:04:05"),
	//})
	return &pb.CommitNewOrderResponse{
		Code:           http.StatusInternalServerError,
		OrderId:        order.Id,
		PlanName:       plan.Name,
		OriginalPrice:  float32(originalPrice),
		DiscountAmount: float32(disCountAmount),
		PayPrice:       float32(originalPrice - discountPercent),
		Period:         request.Period,
		CreatedAt:      time.Now().Format("2006-01-02 15:04:05"),
		Msg:            "订单创建成功",
	}, nil
}

func (s *OrderServices) GetOrders(ctx context.Context, request *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	var orders []model.Orders
	if result := dao.Db.Model(&model.Orders{}).Where("user_id = ?", request.UserId).Find(&orders); result.Error != nil {
		log.Println(result.Error)
		//context.JSON(http.StatusOK, gin.H{
		//	"code":  http.StatusNotFound,
		//	"error": result.Error.Error(),
		//	"msg":   "查找用户订单错误",
		//})
		//return
		return &pb.GetOrdersResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查找用户订单错误",
		}, nil
	}
	//context.JSON(http.StatusOK, gin.H{
	//	"code":       http.StatusOK,
	//	"order_list": orders,
	//})
	return &pb.GetOrdersResponse{
		Code:      http.StatusInternalServerError,
		OrderList: ConvertOrder2pbOrder(orders),
		Msg:       "查找用户订单错误",
	}, nil
}

//// GetActivePlanListByUserId 获取用户的有效订阅
//func GetActivePlanListByUserId(context *gin.Context) {
//	// Parse the user_id from the query parameters
//	userID := context.Query("user_id")
//	if userID == "" {
//		context.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
//		return
//	}
//
//	var activeOrders []ActiveOrders
//	// Query the ActiveOrders table for the user_id and where IsActive is true
//	if err := dao.Db.Where("user_id = ? AND is_active = ?", userID, true).Find(&activeOrders).Error; err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch active orders"})
//		return
//	}
//
//	if len(activeOrders) == 0 {
//		context.JSON(http.StatusNotFound, gin.H{"error": "No active plans found for the user"})
//		return
//	}
//
//	// Create a result array to hold plan_name and expiration_date
//	var result []map[string]interface{}
//
//	// Iterate through the active orders and fetch plan details
//	for _, subscribe := range activeOrders {
//		var plan model.Plan
//		// Fetch the plan details using the PlanId
//		if err := dao.Db.Where("id = ?", subscribe.PlanId).First(&plan).Error; err != nil {
//			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch plan details"})
//			return
//		}
//
//		// Add the plan_name and expiration_date to the result
//		result = append(result, map[string]interface{}{
//			"plan_name":       plan.Name,
//			"expiration_date": subscribe.ExpirationDate,
//		})
//	}
//
//	// Return the result as a JSON response
//	context.JSON(http.StatusOK, gin.H{
//		"code":     http.StatusOK,
//		"my_plans": result,
//		"msg":      "获取数据成功",
//	})
//}

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
//	subscribe := Orders{
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
//	log.Println(subscribe)
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
//	if result := dao.Db.Model(&Orders{}).Create(&subscribe); result.Error != nil {
//		context.JSON(http.StatusOK, gin.H{
//			"code":  http.StatusInternalServerError,
//			"error": result.Error.Error(),
//			"msg":   "创建新订单出错",
//		})
//	}
//
//	context.JSON(http.StatusOK, gin.H{
//		"code":            http.StatusOK,
//		"order_id":        subscribe.Id,
//		"plan_name":       plan.Name,
//		"original_price":  originalPrice,
//		"discount_amount": disCountAmount,
//		"pay_price":       originalPrice - discountPercent,
//		"period":          postData.Period,
//		"created_at":      time.Now().Format("2006-01-02 15:04:05"),
//	})
//}

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
