package services

import (
	pb "NxcFull/backend/subscribeServices/api/proto"
	"NxcFull/backend/subscribeServices/internal/dao"
	"NxcFull/backend/subscribeServices/internal/model"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
)

type SubscribeServices struct {
	pb.UnimplementedSubscriptionServiceServer
}

func NewOrderServices() *SubscribeServices {
	return &SubscribeServices{}
}

func (s *SubscribeServices) GetActivePlanListByUserId(ctx context.Context, request *pb.GetActivePlanListByUserIdRequest) (*pb.GetActivePlanListByUserIdResponse, error) {
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

func (s *SubscribeServices) CommitNewOrder(ctx context.Context, request *pb.CommitNewOrderRequest) (*pb.CommitNewOrderResponse, error) {
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

	var couponInfo model.Coupon
	if request.CouponId != 0 {
		if result := dao.Db.Model(&model.Coupon{}).Where("id = ?", request.CouponId).First(&couponInfo); result.Error != nil {
			log.Println(result.Error.Error())
		}
	}

	var plan model.Plan
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
		if result := dao.Db.Model(&model.CouponUsage{}).Create(&couponUsed); result.Error != nil {
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
	//	"order_id":        subscription.Id,
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

func (s *SubscribeServices) GetOrders(ctx context.Context, request *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
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

func (s *SubscribeServices) GetAllPlanKeyName(context context.Context, request *pb.GetAllPlanKeyNameRequest) (*pb.GetAllPlanKeyNameResponse, error) {
	type planInfo struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}
	var planArr []*pb.PlanKV
	if result := dao.Db.Model(&model.Plan{}).Select("id, name").Order("sort ASC").Find(&planArr); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllPlanKeyNameResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查找计划订阅键值错误",
		}, nil
	}
	log.Println(planArr)
	//context.JSON(http.StatusOK, gin.H{
	//	"code":  http.StatusOK,
	//	"plans": planArr,
	//})
	return &pb.GetAllPlanKeyNameResponse{
		Code:  http.StatusOK,
		Msg:   "查询成功",
		Plans: planArr,
	}, nil
}

func (s *SubscribeServices) GetAllPlans(ctx context.Context, request *pb.GetAllPlansRequest) (*pb.GetAllPlansResponse, error) {
	var plans []model.Plan
	//if result := dao.Db.Model(&Plan{}).Find(&plans); result.Error != nil {
	if result := dao.Db.Model(&model.Plan{}).Order("sort ASC").Find(&plans); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllPlansResponse{
			Code: http.StatusInternalServerError,
			Msg:  "获取订阅列表失败",
		}, nil
	}
	return &pb.GetAllPlansResponse{
		Code:  http.StatusOK,
		Msg:   "获取订阅列表成功",
		Plans: ConvertPlan2pbPlan(plans),
	}, nil
}

func (s *SubscribeServices) AddNewPlan(ctx context.Context, request *pb.AddNewPlanRequest) (*pb.AddNewPlanResponse, error) {
	var newPlan = model.Plan{
		Name:          request.Name,
		Describe:      request.Describe,
		IsSale:        request.IsSale,
		IsRenew:       request.IsRenew,
		GroupId:       request.GroupId,
		CapacityLimit: request.CapacityLimit,
		Residue:       request.CapacityLimit,
		MonthPrice:    request.MonthPrice,
		QuarterPrice:  request.QuarterPrice,
		HalfYearPrice: request.HalfYearPrice,
		YearPrice:     request.YearPrice,
		// 以下Sort用于在前端进行排序 优先级从低到高，添加新订阅时，它的优先级为上一个的优先级+1，如果表内为空不存在数据，则第一个添加的sort为1000
		// Sort:
		Sort: request.Sort,
	}
	if result := dao.Db.Create(&newPlan); result.Error != nil {
		log.Println(result.Error)
		//context.JSON(http.StatusInternalServerError, gin.H{
		//	"code":  http.StatusInternalServerError,
		//	"msg":   "插入数据库失败",
		//	"error": result.Error.Error(),
		//})
		return &pb.AddNewPlanResponse{
			Code: http.StatusInternalServerError,
			Msg:  "插入数据库失败",
		}, nil
	}
	//context.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "插入数据成功",
	//})
	return &pb.AddNewPlanResponse{
		Code: http.StatusOK,
		Msg:  "插入数据成功",
	}, nil
}

// pass start
//func HandleGetAllPlans(context *gin.Context) {
//	var plans []model.Plan
//	//if result := dao.Db.Model(&Plan{}).Find(&plans); result.Error != nil {
//	if result := dao.Db.Model(&model.Plan{}).Order("sort ASC").Find(&plans); result.Error != nil {
//		log.Println(result.Error)
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code": http.StatusInternalServerError,
//			"msg":  "获取订阅列表失败",
//		})
//	}
//	log.Println(plans)
//	context.JSON(http.StatusOK, gin.H{
//		"code":  http.StatusOK,
//		"plans": plans,
//	})
//}
//
//func HandleAddNewPlan(context *gin.Context) {
//	postData := &struct {
//		Name          string  `json:"name"`
//		Describe      string  `json:"describe"`
//		IsSale        bool    `json:"is_sale"`
//		IsRenew       bool    `json:"is_renew"`
//		GroupId       int64   `json:"group_id"`
//		CapacityLimit int64   `json:"capacity_limit"`
//		MonthPrice    float64 `json:"month_price"`
//		QuarterPrice  float64 `json:"quarter_price"`
//		HalfYearPrice float64 `json:"half_year_price"`
//		YearPrice     float64 `json:"year_price"`
//		Sort          int64   `json:"sort"`
//	}{}
//	if err := context.ShouldBind(&postData); err != nil {
//		context.JSON(http.StatusBadRequest, gin.H{
//			"code":  http.StatusBadRequest,
//			"msg":   "绑定数据失败",
//			"error": err.Error(),
//		})
//	}
//	log.Println(postData)
//	var newPlan = model.Plan{
//		Name:          postData.Name,
//		Describe:      postData.Describe,
//		IsSale:        postData.IsSale,
//		IsRenew:       postData.IsRenew,
//		GroupId:       postData.GroupId,
//		CapacityLimit: postData.CapacityLimit,
//		Residue:       postData.CapacityLimit,
//		MonthPrice:    postData.MonthPrice,
//		QuarterPrice:  postData.QuarterPrice,
//		HalfYearPrice: postData.HalfYearPrice,
//		YearPrice:     postData.YearPrice,
//		// 以下Sort用于在前端进行排序 优先级从低到高，添加新订阅时，它的优先级为上一个的优先级+1，如果表内为空不存在数据，则第一个添加的sort为1000
//		// Sort:
//		Sort: postData.Sort,
//	}
//	if result := dao.Db.Create(&newPlan); result.Error != nil {
//		log.Println(result.Error)
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "插入数据库失败",
//			"error": result.Error.Error(),
//		})
//	}
//	context.JSON(http.StatusOK, gin.H{
//		"code": http.StatusOK,
//		"msg":  "插入数据成功",
//	})
//}

//func GetAllPlanKeyName(context *gin.Context) {
//	type planInfo struct {
//		Id   int64  `json:"id"`
//		Name string `json:"name"`
//	}
//	var planArr = []planInfo{}
//	if result := dao.Db.Model(&Plan{}).Select("id, name").Order("sort ASC").Find(&planArr); result.Error != nil {
//		log.Println(result.Error)
//	}
//	log.Println(planArr)
//	context.JSON(http.StatusOK, gin.H{
//		"code":  http.StatusOK,
//		"plans": planArr,
//	})
//}
