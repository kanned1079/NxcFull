package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	pb "orderServices/api/proto"
	"orderServices/internal/puiblisher"
	"orderServices/internal/utils"
	"time"
)

// CommitNewOrder 用戶提交新的訂單 目前是同步处理 后续需要使用mq实现异步
func (s *OrderServices) CommitNewOrder(context context.Context, request *pb.CommitNewOrderRequest) (*pb.CommitNewOrderResponse, error) {
	log.Printf("UserId %v\n PlanId %v\n Period %v\n CouponId %v\n", request.UserId, request.PlanId, request.Period, request.CouponId)
	orderIdGenerated := utils.GenerateOrderId(request.UserId)
	log.Println("newOrderId: ", orderIdGenerated)

	// 推送到消息队列的数据
	postOrderData := &struct {
		OrderId  string `json:"order_id"`
		UserId   int64  `json:"user_id"`
		PlanId   int64  `json:"plan_id"`
		Period   string `json:"period"`
		CouponId int64  `json:"coupon_id"`
	}{
		OrderId:  orderIdGenerated,
		UserId:   request.UserId,
		PlanId:   request.PlanId,
		Period:   request.Period,
		CouponId: request.CouponId,
	}
	if postOrderDataJSON, err := json.Marshal(postOrderData); err != nil {
		return &pb.CommitNewOrderResponse{
			Code: http.StatusInternalServerError,
			Msg:  "内部错误" + err.Error(),
		}, nil
	} else {
		if err = puiblisher.PublishOrderMessage(postOrderDataJSON); err != nil {
			log.Println("puiblisher.PublishOrderMessage error:", err)
			return &pb.CommitNewOrderResponse{
				Code: http.StatusInternalServerError,
				Msg:  "内部错误" + err.Error(),
			}, nil
		}
		time.Sleep(time.Millisecond * 500) // 延迟500ms
	}

	// 这里将直接返回订单号

	//var userMail string
	//if result := dao.Db.Model(&model.User{}).Where("id = ?", request.UserId).Select("email").Limit(1).Find(&userMail); result.Error != nil {
	//	log.Println("获取邮箱地址失败")
	//}
	//log.Println(userMail)
	//
	//var couponInfo model.Coupon
	//if request.CouponId != 0 {
	//	if result := dao.Db.Model(&model.Coupon{}).Where("id = ?", request.CouponId).First(&couponInfo); result.Error != nil {
	//		log.Println(result.Error.Error())
	//	}
	//}
	//
	//var plan model.Plan
	//if result := dao.Db.Where("id = ?", request.PlanId).First(&plan); result.Error != nil {
	//	log.Println(result.Error.Error())
	//}
	//
	//if plan.Residue <= 0 {
	//	log.Println("订阅剩余不足")
	//	return &pb.CommitNewOrderResponse{
	//		Code: http.StatusInternalServerError,
	//		Msg:  "订阅剩余不足",
	//	}, nil
	//}
	//var originalPrice float64
	//switch request.Period {
	//case "month":
	//	originalPrice = plan.MonthPrice
	//case "quarter":
	//	originalPrice = plan.QuarterPrice
	//case "half-year":
	//	originalPrice = plan.HalfYearPrice
	//case "year":
	//	originalPrice = plan.YearPrice
	//}
	//log.Println("原始价格", originalPrice)
	//
	//var discountPercent float64 = couponInfo.PercentOff / 100
	//
	//var disCountAmount float64 = originalPrice * discountPercent
	//
	//log.Println(discountPercent, disCountAmount, originalPrice-disCountAmount)
	//
	//order := model.Orders{
	//	//Id:             int64(orderIdGenerated),        // 订单id
	//	// 注意这里的修改后 密钥的处理查询OrderId而不是Id
	//	OrderId:        orderIdGenerated,
	//	UserId:         request.UserId,                 // 用户id
	//	PlanId:         request.PlanId,                 // 订阅计划id
	//	Period:         request.Period,                 // 付款周期
	//	Email:          userMail,                       // 用户邮箱
	//	DiscountAmount: disCountAmount,                 // 抵折金额
	//	Amount:         originalPrice - disCountAmount, // 最终应该支付金额
	//	IsFinished:     false,
	//	Status:         0, // 新购买
	//	CouponId:       couponInfo.Id,
	//}
	//
	//log.Println(order)
	//
	//couponUsed := model.CouponUsage{
	//	CouponId: couponInfo.Id,
	//	UserId:   request.UserId,
	//	UsedAt:   time.Now(),
	//}
	//
	//if request.CouponId != 0 {
	//	if result := dao.Db.Model(&model.CouponUsage{}).Create(&couponUsed); result.Error != nil {
	//		return &pb.CommitNewOrderResponse{
	//			Code: http.StatusInternalServerError,
	//			Msg:  "添加优惠券使用记录出错" + result.Error.Error(),
	//		}, nil
	//	}
	//}
	//
	//if result := dao.Db.Model(&model.Orders{}).Create(&order); result.Error != nil {
	//	return &pb.CommitNewOrderResponse{
	//		Code: http.StatusInternalServerError,
	//		Msg:  "创建新订单出错" + result.Error.Error(),
	//	}, nil
	//}
	//
	return &pb.CommitNewOrderResponse{
		Code:    http.StatusOK,
		OrderId: orderIdGenerated,
		Msg:     "ok",
		//PlanName:       plan.Name,
		//OriginalPrice:  float32(originalPrice),
		//DiscountAmount: float32(disCountAmount),
		//PayPrice:       float32(originalPrice - discountPercent),
		//Period:         request.Period,
		//CreatedAt:      time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
