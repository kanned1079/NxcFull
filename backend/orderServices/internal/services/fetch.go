package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	pb "orderServices/api/proto"
	"orderServices/internal/dao"
	"orderServices/internal/model"
	"strconv"
)

// v1
// GetAllMyOrders 返回用户所有的订单 包括所有提交的 支付成功或失败的
// 请求参数 用户id 返回一个序列化的对象数组
//func (s *OrderServices) GetAllMyOrders(context context.Context, request *pb.GetAllMyOrdersRequest) (*pb.GetAllMyOrdersResponse, error) {
//	var orders []model.Orders
//	if result := dao.Db.Model(&model.Orders{}).Where("user_id = ?", request.UserId).Find(&orders); result.Error != nil {
//		log.Println(result.Error)
//		return &pb.GetAllMyOrdersResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "查询订单信息错误" + result.Error.Error(),
//		}, nil
//	}
//	// 还需要增加redis中还未完成的订单 因为未超时的订单不存储在数据库中
//	// redis中的键路径为 "pending_order:" + <userId> + ":" + <订单号> 这里可能有好多个订单
//
//	if tempJson, err := json.Marshal(orders); err != nil {
//		log.Println(err.Error())
//		return &pb.GetAllMyOrdersResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "序列化数据失败" + err.Error(),
//		}, nil
//	} else {
//		return &pb.GetAllMyOrdersResponse{
//			Code:      http.StatusOK,
//			Msg:       "查询成功",
//			OrderList: tempJson,
//		}, nil
//	}
//}

func (s *OrderServices) GetAllMyOrders(context context.Context, request *pb.GetAllMyOrdersRequest) (*pb.GetAllMyOrdersResponse, error) {
	// 每页订单数量
	size := request.Size
	if size <= 0 {
		size = 10 // 默认每页显示10条
	}

	// 当前页码
	page := request.Page
	if page <= 0 {
		page = 1 // 默认第一页
	}

	// 查询数据库中的订单总数
	var totalOrders int64
	if result := dao.Db.Model(&model.Orders{}).Where("user_id = ?", request.UserId).Count(&totalOrders); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllMyOrdersResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询订单总数失败" + result.Error.Error(),
		}, nil
	}

	// 计算总页数
	pageCount := (totalOrders + int64(size) - 1) / int64(size)

	// 查询数据库中的订单（分页并按时间降序排序）
	var orders []model.Orders
	if result := dao.Db.Model(&model.Orders{}).
		Where("user_id = ?", request.UserId).
		Order("created_at DESC").
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&orders); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllMyOrdersResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询订单信息错误" + result.Error.Error(),
		}, nil
	}

	// 查询Redis中未完成的订单
	pendingOrders := make([]model.Orders, 0)
	if page == 1 {
		redisKeyPattern := "pending_order:" + strconv.FormatInt(request.UserId, 10) + ":*"
		iter := dao.Rdb.Scan(context, 0, redisKeyPattern, 0).Iterator()
		for iter.Next(context) {
			orderData, err := dao.Rdb.Get(context, iter.Val()).Result()
			if err != nil {
				log.Println("获取Redis订单失败:", err)
				continue
			}

			var order model.Orders
			if err := json.Unmarshal([]byte(orderData), &order); err != nil {
				log.Println("解析Redis订单失败:", err)
				continue
			}

			pendingOrders = append(pendingOrders, order)
		}
		if err := iter.Err(); err != nil {
			log.Println("Redis扫描出错:", err)
		}
	}

	// 合并 Redis 和数据库的订单，Redis 订单优先
	orders = append(pendingOrders, orders...)

	// 序列化订单列表
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
			PageCount: pageCount,
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
