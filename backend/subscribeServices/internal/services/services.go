package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	pb "subscribeServices/api/proto"
	"subscribeServices/internal/dao"
	"subscribeServices/internal/model"
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
	//var myPlanResult []*pb.PlanInfo
	var myPlanResult []struct {
		//string plan_name = 1;         // 订阅计划名称
		//string expiration_date = 2;   // 到期日期
		PlanName       string `json:"plan_name"`
		ExpirationDate string `json:"expiration_date"`
	}

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
		//myPlanResult = append(myPlanResult, &pb.PlanInfo{
		//	PlanName:       plan.Name,
		//	ExpirationDate: order.ExpirationDate.Format("2006-01-02 15:04:05"),
		//})
		myPlanResult = append(myPlanResult, struct {
			PlanName       string `json:"plan_name"`
			ExpirationDate string `json:"expiration_date"`
		}{
			PlanName:       plan.Name,
			ExpirationDate: order.ExpirationDate.Format("2006-01-02 15:04:05"),
		})
	}
	if myPlanResultJson, err := json.Marshal(myPlanResult); err != nil {
		return &pb.GetActivePlanListByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "转换格式失败",
		}, nil
	} else {
		return &pb.GetActivePlanListByUserIdResponse{
			Code:    http.StatusOK,
			MyPlans: myPlanResultJson,
			Msg:     "获取成功",
		}, nil
	}

}

// CommitNewOrder 提交新订单
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
			return &pb.CommitNewOrderResponse{
				Code: http.StatusInternalServerError,
				Msg:  "添加优惠券使用记录出错",
			}, nil
		}
	}

	if result := dao.Db.Model(&model.Orders{}).Create(&order); result.Error != nil {
		return &pb.CommitNewOrderResponse{
			Code: http.StatusInternalServerError,
			Msg:  "创建新订单出错",
		}, nil
	}
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
		return &pb.GetOrdersResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查找用户订单错误",
		}, nil
	}
	//if len(orders) == 0 {
	//	return &pb.GetOrdersResponse{
	//		Code: http.StatusNotFound,
	//		Msg:  "用户还没有订单",
	//	}, nil
	//}
	if ordersJson, err := json.Marshal(orders); err != nil {
		return &pb.GetOrdersResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查找用户订单错误",
		}, nil
	} else {
		return &pb.GetOrdersResponse{
			Code:      http.StatusOK,
			OrderList: ordersJson,
			Msg:       "查询成功",
		}, nil
	}
	//return &pb.GetOrdersResponse{
	//	Code:      http.StatusInternalServerError,
	//	OrderList: ConvertOrder2pbOrder(orders),
	//	Msg:       "查找用户订单错误",
	//}, nil
}

func (s *SubscribeServices) GetAllPlanKeyName(context context.Context, request *pb.GetAllPlanKeyNameRequest) (*pb.GetAllPlanKeyNameResponse, error) {
	//type planInfo struct {
	//	Id     int64  `json:"id"`
	//	Name   string `json:"name"`
	//	IsSale bool   `json:"is_sale"`
	//}
	var planArr []struct {
		Id     int64  `json:"id"`
		Name   string `json:"name"`
		IsSale bool   `json:"is_sale"`
	}
	if result := dao.Db.Model(&model.Plan{}).Select("id, name, is_sale").Order("sort ASC").Find(&planArr); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllPlanKeyNameResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查找计划订阅键值错误",
		}, nil
	}
	log.Println(planArr)

	if planArrJson, err := json.Marshal(planArr); err != nil {
		return &pb.GetAllPlanKeyNameResponse{
			Code: http.StatusInternalServerError,
			Msg:  "转换格式失败" + err.Error(),
		}, nil
	} else {
		return &pb.GetAllPlanKeyNameResponse{
			Code:  http.StatusOK,
			Msg:   "查询成功",
			Plans: planArrJson,
		}, nil
	}
}

//func (s *SubscribeServices) GetAllPlans(ctx context.Context, request *pb.GetAllPlansRequest) (*pb.GetAllPlansResponse, error) {
//	var plans []model.Plan
//	//if result := dao.Db.Model(&Plan{}).Find(&plans); result.Error != nil {
//	log.Println("request.User: ", request.IsUser)
//	if request.IsUser == true {
//		// 如果是用户的请求
//		//log.Println("用户的请求")
//		if result := dao.Db.Model(&model.Plan{}).Where("is_sale = ?", true).Order("sort ASC").Find(&plans); result.Error != nil {
//			log.Println(result.Error)
//			return &pb.GetAllPlansResponse{
//				Code: http.StatusInternalServerError,
//				Msg:  "Get plans failure.",
//			}, nil
//		}
//	} else {
//		// 如果是管理员的请求
//		//log.Println("Request from admin")
//		if result := dao.Db.Model(&model.Plan{}).Order("sort ASC").Find(&plans); result.Error != nil {
//			log.Println(result.Error)
//			return &pb.GetAllPlansResponse{
//				Code: http.StatusInternalServerError,
//				Msg:  "Get plans failure.",
//			}, nil
//		}
//	}
//	if plansJson, err := json.Marshal(plans); err != nil {
//		return &pb.GetAllPlansResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "Marshal data failure.",
//		}, nil
//	} else {
//		return &pb.GetAllPlansResponse{
//			Code:  http.StatusOK,
//			Msg:   "Get plans ok",
//			Plans: plansJson,
//		}, nil
//	}
//
//}

func (s *SubscribeServices) GetAllPlans(ctx context.Context, request *pb.GetAllPlansRequest) (*pb.GetAllPlansResponse, error) {
	var plans []model.Plan
	var redisKey string

	if request.IsUser {
		redisKey = "plans:user:all"
	} else {
		redisKey = "plans:admin:all"
	}

	// 尝试从 Redis 中获取数据
	redisData, err := dao.Rdb.Get(ctx, redisKey).Result()
	if err == nil && redisData != "" {
		// 如果 Redis 中有数据，直接返回
		log.Println("Plans existed in redis.")
		return &pb.GetAllPlansResponse{
			Code:  http.StatusOK,
			Msg:   "Get plans successfully.",
			Plans: []byte(redisData),
		}, nil
	}
	log.Println("Plans not in redis, query in mysql.")

	// Redis 中无数据，查询 MySQL
	db := dao.Db.Model(&model.Plan{}).Order("sort ASC")
	if request.IsUser {
		db = db.Where("is_sale = ?", true)
	}

	if result := db.Find(&plans); result.Error != nil {
		log.Println("query failure:", result.Error)
		return &pb.GetAllPlansResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Get plans failure.",
		}, nil
	}

	// 将查询结果序列化为 JSON
	plansJson, err := json.Marshal(plans)
	if err != nil {
		log.Println("JSON Marshal failure:", err)
		return &pb.GetAllPlansResponse{
			Code: http.StatusInternalServerError,
			Msg:  "JSON Marshal failure",
		}, nil
	}

	// 将查询结果缓存到 Redis
	if err := dao.Rdb.Set(ctx, redisKey, plansJson, time.Hour).Err(); err != nil {
		log.Println("Store in redis failure:", err)
	}

	// 返回响应
	return &pb.GetAllPlansResponse{
		Code:  http.StatusOK,
		Msg:   "Get plans successfully.",
		Plans: plansJson,
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
		return &pb.AddNewPlanResponse{
			Code: http.StatusInternalServerError,
			Msg:  "插入数据库失败",
		}, nil
	}
	// 在这里清空redis的plans
	//var redisKey string
	//redisKey = "plans:user:all"
	//redisKey = "plans:admin:all"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := ClearPlansRedisCache(ctx); err != nil {
		log.Println(err)
		return &pb.AddNewPlanResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}, nil
	}

	// 在这里清空redis的plans
	// 清空用户计划缓存

	return &pb.AddNewPlanResponse{
		Code: http.StatusOK,
		Msg:  "插入数据成功",
	}, nil
}

func (s *SubscribeServices) UpdatePlan(ctx context.Context, request *pb.UpdatePlanRequest) (*pb.UpdatePlanResponse, error) {
	var plan = model.Plan{
		Id:            request.Id,
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
		Sort: request.Sort,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := ClearPlansRedisCache(ctx); err != nil {
		log.Println(err)
		return &pb.UpdatePlanResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}, nil
	}

	if result := dao.Db.Model(&model.Plan{}).Where("id = ?", plan.Id).Updates(&plan); result.Error != nil {
		log.Println(result.Error)
		return &pb.UpdatePlanResponse{
			Code: http.StatusInternalServerError,
			Msg:  "更新失败" + result.Error.Error(),
		}, nil
	}
	return &pb.UpdatePlanResponse{
		Code: http.StatusOK,
		Msg:  "更新成功",
	}, nil
}

func (s *SubscribeServices) DeletePlan(ctx context.Context, request *pb.DeletePlanRequest) (*pb.DeletePlanResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := ClearPlansRedisCache(ctx); err != nil {
		log.Println(err)
		return &pb.DeletePlanResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}, nil
	}
	if result := dao.Db.Model(&model.Plan{}).Where("id = ?", request.PlanId).Delete(&model.Plan{Id: request.PlanId}); result.Error != nil {
		return &pb.DeletePlanResponse{
			Code: http.StatusInternalServerError,
			Msg:  "删除订阅错误" + result.Error.Error(),
		}, nil
	}
	return &pb.DeletePlanResponse{
		Code: http.StatusOK,
		Msg:  "删除成功",
	}, nil
}

func (s *SubscribeServices) UpdatePlanIsSale(ctx context.Context, request *pb.UpdatePlanIsSaleRequest) (*pb.UpdatePlanIsSaleResponse, error) {
	log.Println("UpdatePlanIsSale", request.Id, request.IsSale)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := ClearPlansRedisCache(ctx); err != nil {
		log.Println(err)
		return &pb.UpdatePlanIsSaleResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}, nil
	}

	var plan = model.Plan{}
	if result := dao.Db.Model(&model.Plan{}).Where("id = ?", request.Id).First(&plan); result.Error != nil {
		return &pb.UpdatePlanIsSaleResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询失败" + result.Error.Error(),
		}, nil
	}
	plan.IsSale = request.IsSale
	if result := dao.Db.Model(&model.Plan{}).Where("id = ?", request.Id).Update("is_sale", request.IsSale); result.Error != nil {
		return &pb.UpdatePlanIsSaleResponse{
			Code: http.StatusInternalServerError,
			Msg:  "保存失败" + result.Error.Error(),
		}, nil
	}
	return &pb.UpdatePlanIsSaleResponse{
		Code: http.StatusOK,
		Msg:  "保存成功",
	}, nil
}

func (s *SubscribeServices) UpdatePlanRenew(ctx context.Context, request *pb.UpdatePlanRenewRequest) (*pb.UpdatePlanRenewResponse, error) {
	log.Println("UpdatePlanRenew", request.Id, request.IsRenew)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := ClearPlansRedisCache(ctx); err != nil {
		log.Println(err)
		return &pb.UpdatePlanRenewResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}, nil
	}

	var plan = model.Plan{}
	if result := dao.Db.Model(&model.Plan{}).Where("id = ?", request.Id).Update("is_renew", request.IsRenew); result.Error != nil {
		return &pb.UpdatePlanRenewResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询失败" + result.Error.Error(),
		}, nil
	}
	plan.IsRenew = request.IsRenew
	if result := dao.Db.Model(&model.Plan{}).Where("id = ?", request.Id).Updates(&plan); result.Error != nil {
		return &pb.UpdatePlanRenewResponse{
			Code: http.StatusInternalServerError,
			Msg:  "保存失败" + result.Error.Error(),
		}, nil
	}
	return &pb.UpdatePlanRenewResponse{
		Code: http.StatusOK,
		Msg:  "保存成功",
	}, nil
}
