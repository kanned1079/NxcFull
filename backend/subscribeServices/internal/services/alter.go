package services

import (
	"context"
	"log"
	"net/http"
	"strconv"
	pb "subscribeServices/api/proto"
	"subscribeServices/internal/dao"
	"subscribeServices/internal/model"
	"time"
)

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

func (s *SubscribeServices) AddNewPlan(ctx context.Context, request *pb.AddNewPlanRequest) (*pb.AddNewPlanResponse, error) {
	var newPlan = model.Plan{
		Name:          request.Name,
		Describe:      request.Describe,
		IsSale:        request.IsSale,
		IsRenew:       request.IsRenew,
		GroupId:       request.GroupId,
		CapacityLimit: request.CapacityLimit,
		Residue:       request.Residue,
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
	// 创建一个包含所有更新字段的 map
	updates := map[string]interface{}{
		"name":            request.Name,
		"describe":        request.Describe,
		"is_sale":         request.IsSale,
		"is_renew":        request.IsRenew,
		"group_id":        request.GroupId,
		"capacity_limit":  request.CapacityLimit,
		"residue":         request.Residue,
		"month_price":     request.MonthPrice,
		"quarter_price":   request.QuarterPrice,
		"half_year_price": request.HalfYearPrice,
		"year_price":      request.YearPrice,
		"sort":            request.Sort,
	}

	// 设置请求超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// 清除 Redis 缓存
	if err := ClearPlansRedisCache(ctx); err != nil {
		log.Println(err)
		return &pb.UpdatePlanResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}, nil
	}

	// 使用 map 进行更新
	if result := dao.Db.Model(&model.Plan{}).Where("`id` = ?", request.Id).Updates(updates); result.Error != nil {
		log.Println(result.Error)
		return &pb.UpdatePlanResponse{
			Code: http.StatusInternalServerError,
			Msg:  "更新失败" + result.Error.Error(),
		}, nil
	}

	// 返回成功响应
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
