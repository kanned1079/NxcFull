package consumer

import (
	"encoding/json"
)

//const (
//	PendingOrderKeyPrefix = "pending_order"
//)

func ProcessTopUpOrder(jsonData []byte) error {

	var postOrderData = &struct {
		OrderId  string `json:"order_id"`
		UserId   int64  `json:"user_id"`
		PlanId   int64  `json:"plan_id"`
		Period   string `json:"period"`
		CouponId int64  `json:"coupon_id"`
	}{}
	err := json.Unmarshal(jsonData, &postOrderData)
	if err != nil {
		return err
	}

	return nil

}
