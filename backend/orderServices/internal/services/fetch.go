package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	pb "orderServices/api/proto"
	"orderServices/internal/dao"
	"orderServices/internal/model"
)

// GetAllMyOrders 返回用户所有的订单 包括所有提交的 支付成功或失败的
// 请求参数 用户id 返回一个序列化的对象数组
func (s *OrderServices) GetAllMyOrders(context context.Context, request *pb.GetAllMyOrdersRequest) (*pb.GetAllMyOrdersResponse, error) {
	var orders []model.Orders
	if result := dao.Db.Model(&model.Orders{}).Where("user_id = ?", request.UserId).Find(&orders); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllMyOrdersResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询订单信息错误" + result.Error.Error(),
		}, nil
	}
	if tempJson, err := json.Marshal(orders); err != nil {
		log.Println(err.Error())
		return &pb.GetAllMyOrdersResponse{
			Code: http.StatusInternalServerError,
			Msg:  "序列化数据失败" + err.Error(),
		}, nil
	} else {
		return &pb.GetAllMyOrdersResponse{
			Code:      http.StatusOK,
			Msg:       "查询成功",
			OrderList: tempJson,
		}, nil
	}
}

// GetActivePlanListByUserId 获取用户的有效订单 其中包含其密钥id
// 请求参数 用户id 返回一个序列化的对象数组 和GetAllMyOrders函数相比只返回有效的订单
func (s *OrderServices) GetActivePlanListByUserId(context context.Context, request *pb.GetActivePlanListByUserIdRequest) (*pb.GetActivePlanListByUserIdResponse, error) {
	var activeOrders []model.ActiveOrders
	// Query the ActiveOrders table for the user_id and where IsActive is true
	if err := dao.Db.Where("user_id = ? AND is_active = ?", request.UserId, true).Find(&activeOrders).Error; err != nil {
		//context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch active orders"})
		//return
		return &pb.GetActivePlanListByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to fetch active orders" + err.Error(),
		}, nil
	}

	if len(activeOrders) == 0 { // 如果用户的没有过订单
		//context.JSON(http.StatusNotFound, gin.H{"error": "No active plans found for the user"})
		//return
		return &pb.GetActivePlanListByUserIdResponse{
			Code: http.StatusNotFound, // 404
			Msg:  "No active plans found for the user",
		}, nil
	}

	// Create a result array to hold plan_name and expiration_date
	var result []map[string]interface{}

	// Iterate through the active orders and fetch plan details
	for _, order := range activeOrders {
		var plan model.Plan
		// Fetch the plan details using the PlanId
		if err := dao.Db.Where("id = ?", order.PlanId).First(&plan).Error; err != nil {
			//context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch plan details"})
			//return
			return &pb.GetActivePlanListByUserIdResponse{
				Code: http.StatusNotFound, // 404
				Msg:  "Failed to fetch plan details" + err.Error(),
			}, nil
		}

		// Add the plan_name and expiration_date to the result
		result = append(result, map[string]interface{}{
			"plan_name":       plan.Name,
			"expiration_date": order.ExpirationDate,
		})
	}

	if tempJson, err := json.Marshal(result); err != nil {
		return &pb.GetActivePlanListByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to marshal result" + err.Error(),
		}, nil
	} else {
		return &pb.GetActivePlanListByUserIdResponse{
			Code:    http.StatusOK,
			Msg:     "获取成功",
			MyPlans: tempJson,
		}, nil
	}
}
