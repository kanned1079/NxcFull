package model

import (
	"encoding/json"
	"log"
)

// ProcessOrder 处理订单的业务逻辑
//func ProcessOrder(order model.Orders) error {
//	// 将订单写入数据库
//	log.Println("数据库操作...")
//	//result := Db.Create(&order)
//	//if result.Error != nil {
//	//	return fmt.Errorf("订单写入数据库失败: %v", result.Error)
//	//}
//
//	// 其他业务逻辑，比如优惠券更新等
//	return nil
//}

func ProcessOrder(jsonData []byte) error {
	var postOrderData = &struct {
		UserId   int64  `json:"user_id"`
		PlanId   int64  `json:"plan_id"`
		Period   string `json:"period"`
		CouponId int64  `json:"coupon_id"`
	}{}
	err := json.Unmarshal(jsonData, &postOrderData)
	if err != nil {
		return err
	}
	log.Println(postOrderData)
	return nil
}
