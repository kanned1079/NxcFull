package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	pb "subscribeServices/api/proto"
	"subscribeServices/internal/dao"
	"subscribeServices/internal/model"
	"time"
)

func (s *SubscribeServices) GetActivePlanListByUserId(ctx context.Context, request *pb.GetActivePlanListByUserIdRequest) (*pb.GetActivePlanListByUserIdResponse, error) {
	redisKey := fmt.Sprintf("user_property:active_order_summary:%d", request.UserId)

	cachedData, err := dao.Rdb.Get(ctx, redisKey).Result()
	if err == nil {
		log.Println("User's active orders summary in redis, skip query db.")
		var cachedResult []map[string]interface{}
		if err := json.Unmarshal([]byte(cachedData), &cachedResult); err == nil {
			return &pb.GetActivePlanListByUserIdResponse{
				Code:    http.StatusOK,
				Msg:     "获取成功 (from cache)",
				MyPlans: []byte(cachedData),
			}, nil
		}
	} else if !errors.Is(err, redis.Nil) {
		return &pb.GetActivePlanListByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to fetch data from Redis: " + err.Error(),
		}, nil
	}

	var activeOrders []model.ActiveOrders
	if err := dao.Db.Where("user_id = ? AND is_active = ?", request.UserId, true).Order(fmt.Sprintf("created_at DESC")).Find(&activeOrders).Error; err != nil {
		return &pb.GetActivePlanListByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to fetch active orders: " + err.Error(),
		}, nil
	}

	if len(activeOrders) == 0 {
		return &pb.GetActivePlanListByUserIdResponse{
			Code: http.StatusNotFound,
			Msg:  "No active plans found for the user",
		}, nil
	}

	var result []map[string]interface{}
	for _, order := range activeOrders {
		var plan model.Plan
		if err := dao.Db.Where("id = ?", order.PlanId).First(&plan).Error; err != nil {
			return &pb.GetActivePlanListByUserIdResponse{
				Code: http.StatusNotFound,
				Msg:  "Failed to fetch plan details: " + err.Error(),
			}, nil
		}

		result = append(result, map[string]interface{}{
			"id":              order.ID,
			"plan_name":       plan.Name,
			"expiration_date": order.ExpirationDate,
		})
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return &pb.GetActivePlanListByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to marshal result: " + err.Error(),
		}, nil
	}

	if err := dao.Rdb.Set(ctx, redisKey, resultJson, time.Hour).Err(); err != nil {
		return &pb.GetActivePlanListByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to cache result in Redis: " + err.Error(),
		}, nil
	}

	return &pb.GetActivePlanListByUserIdResponse{
		Code:    http.StatusOK,
		Msg:     "获取成功",
		MyPlans: resultJson,
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

// v1
//func (s *SubscribeServices) GetAllPlans(ctx context.Context, request *pb.GetAllPlansRequest) (*pb.GetAllPlansResponse, error) {
//	var plans []model.Plan
//	var redisKey string
//
//	if request.IsUser {
//		redisKey = "plans:user:all"
//	} else {
//		redisKey = "plans:admin:all"
//	}
//
//	//request.Page
//	//request.Size
//	// 如果管理员请求了该函数获取所有计划列表 需要按照page和size进行分页
//	// 并且计算最后的PageCount
//	// 如果是用户请求那么PageCount直接为0即可
//
//	// 尝试从 Redis 中获取数据
//	redisData, err := dao.Rdb.Get(ctx, redisKey).Result()
//	if err == nil && redisData != "" {
//		// 如果 Redis 中有数据，直接返回
//		log.Println("Plans existed in redis.")
//		return &pb.GetAllPlansResponse{
//			Code:  http.StatusOK,
//			Msg:   "Get plans successfully.",
//			Plans: []byte(redisData),
//		}, nil
//	}
//	log.Println("Plans not in redis, query in mysql.")
//
//	// Redis 中无数据，查询 MySQL
//	db := dao.Db.Model(&model.Plan{}).Order("sort ASC")
//	if request.IsUser {
//		db = db.Where("is_sale = ?", true)
//	}
//
//	if result := db.Find(&plans); result.Error != nil {
//		log.Println("query failure:", result.Error)
//		return &pb.GetAllPlansResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "Get plans failure.",
//		}, nil
//	}
//
//	// 将查询结果序列化为 JSON
//	plansJson, err := json.Marshal(plans)
//	if err != nil {
//		log.Println("JSON Marshal failure:", err)
//		return &pb.GetAllPlansResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "JSON Marshal failure",
//		}, nil
//	}
//
//	// 将查询结果缓存到 Redis
//	if err := dao.Rdb.Set(ctx, redisKey, plansJson, time.Hour).Err(); err != nil {
//		log.Println("Store in redis failure:", err)
//	}
//
//	// 返回响应
//	return &pb.GetAllPlansResponse{
//		Code:      http.StatusOK,
//		Msg:       "Get plans successfully.",
//		Plans:     plansJson,
//		PageCount: 10,
//	}, nil
//}

//func (s *SubscribeServices) GetAllPlans(ctx context.Context, request *pb.GetAllPlansRequest) (*pb.GetAllPlansResponse, error) {
//	var plans []model.Plan
//	var redisKey string
//
//	if request.IsUser {
//		redisKey = "plans:user:all"
//	} else {
//		redisKey = "plans:admin:all"
//	}
//
//	// 尝试从 Redis 中获取数据
//	redisData, err := dao.Rdb.Get(ctx, redisKey).Result()
//	if err == nil && redisData != "" {
//		// 如果 Redis 中有数据，直接返回
//		log.Println("Plans existed in redis.")
//		return &pb.GetAllPlansResponse{
//			Code:  http.StatusOK,
//			Msg:   "Get plans successfully.",
//			Plans: []byte(redisData),
//		}, nil
//	}
//	log.Println("Plans not in redis, query in mysql.")
//
//	// Redis 中无数据，查询 MySQL
//	db := dao.Db.Model(&model.Plan{}).Order("sort ASC")
//	if request.IsUser {
//		db = db.Where("is_sale = ?", true)
//	}
//
//	// 如果是管理员请求，执行分页
//	var pageCount int64 = 0
//	if !request.IsUser {
//		var totalPlans int64
//		// 查询总数以计算 PageCount
//		if err := db.Count(&totalPlans).Error; err != nil {
//			log.Println("count failure:", err)
//			return &pb.GetAllPlansResponse{
//				Code: http.StatusInternalServerError,
//				Msg:  "Count plans failure.",
//			}, nil
//		}
//
//		// 计算 PageCount
//		pageCount = int64((totalPlans + int64(request.Size) - 1) / int64(request.Size))
//
//		// 应用分页限制
//		offset := (request.Page - 1) * request.Size
//		db = db.Limit(int(request.Size)).Offset(int(offset))
//	}
//
//	// 查询分页后的结果
//	if result := db.Find(&plans); result.Error != nil {
//		log.Println("query failure:", result.Error)
//		return &pb.GetAllPlansResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "Get plans failure.",
//		}, nil
//	}
//
//	// 将查询结果序列化为 JSON
//	plansJson, err := json.Marshal(plans)
//	if err != nil {
//		log.Println("JSON Marshal failure:", err)
//		return &pb.GetAllPlansResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "JSON Marshal failure",
//		}, nil
//	}
//
//	// 将查询结果缓存到 Redis
//	if err := dao.Rdb.Set(ctx, redisKey, plansJson, time.Hour).Err(); err != nil {
//		log.Println("Store in redis failure:", err)
//	}
//
//	// 返回响应
//	return &pb.GetAllPlansResponse{
//		Code:      http.StatusOK,
//		Msg:       "Get plans successfully.",
//		Plans:     plansJson,
//		PageCount: pageCount,
//	}, nil
//}

//func (s *SubscribeServices) GetAllPlans(ctx context.Context, request *pb.GetAllPlansRequest) (*pb.GetAllPlansResponse, error) {
//	var plans []model.Plan
//	var redisKey string
//
//	if request.IsUser {
//		redisKey = "plans:user:all"
//	} else {
//		// 如果管理员请求了分页数据，将分页参数添加到 Redis 键
//		if request.Page > 0 && request.Size > 0 {
//			redisKey = fmt.Sprintf("plans:admin:page:%d:size:%d", request.Page, request.Size)
//		} else {
//			redisKey = "plans:admin:all"
//		}
//	}
//
//	// 尝试从 Redis 中获取数据
//	redisData, err := dao.Rdb.Get(ctx, redisKey).Result()
//	if err == nil && redisData != "" {
//		// 如果 Redis 中有数据，直接返回
//		log.Println("Plans existed in redis.")
//		return &pb.GetAllPlansResponse{
//			Code:  http.StatusOK,
//			Msg:   "Get plans successfully.",
//			Plans: []byte(redisData),
//		}, nil
//	}
//	log.Println("Plans not in redis, query in mysql.")
//
//	// Redis 中无数据，查询 MySQL
//	db := dao.Db.Model(&model.Plan{}).Order("sort ASC")
//	if request.IsUser {
//		db = db.Where("is_sale = ?", true)
//	}
//
//	// 如果是管理员请求，执行分页
//	var pageCount int64 = 0
//	if !request.IsUser {
//		var totalPlans int64
//		// 查询总数以计算 PageCount
//		if err := db.Count(&totalPlans).Error; err != nil {
//			log.Println("count failure:", err)
//			return &pb.GetAllPlansResponse{
//				Code: http.StatusInternalServerError,
//				Msg:  "Count plans failure.",
//			}, nil
//		}
//
//		// 计算 PageCount
//		if request.Size > 0 {
//			pageCount = (totalPlans + int64(request.Size) - 1) / int64(request.Size)
//		}
//
//		// 应用分页限制（仅在请求中指定了分页参数时）
//		if request.Page > 0 && request.Size > 0 {
//			offset := (request.Page - 1) * request.Size
//			db = db.Limit(int(request.Size)).Offset(int(offset))
//		}
//	}
//
//	// 查询分页后的结果
//	if result := db.Find(&plans); result.Error != nil {
//		log.Println("query failure:", result.Error)
//		return &pb.GetAllPlansResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "Get plans failure.",
//		}, nil
//	}
//
//	// 将查询结果序列化为 JSON
//	plansJson, err := json.Marshal(plans)
//	if err != nil {
//		log.Println("JSON Marshal failure:", err)
//		return &pb.GetAllPlansResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "JSON Marshal failure",
//		}, nil
//	}
//
//	// 将查询结果缓存到 Redis
//	if err := dao.Rdb.Set(ctx, redisKey, plansJson, time.Hour).Err(); err != nil {
//		log.Println("Store in redis failure:", err)
//	}
//
//	// 返回响应
//	return &pb.GetAllPlansResponse{
//		Code:      http.StatusOK,
//		Msg:       "Get plans successfully.",
//		Plans:     plansJson,
//		PageCount: pageCount,
//	}, nil
//}

func (s *SubscribeServices) GetAllPlans(ctx context.Context, request *pb.GetAllPlansRequest) (*pb.GetAllPlansResponse, error) {
	var plans []model.Plan
	var redisKey string

	if request.IsUser {
		redisKey = "plans:user:all"
	} else {
		// 如果管理员请求了分页数据，将分页参数添加到 Redis 键
		if request.Page > 0 && request.Size > 0 {
			redisKey = fmt.Sprintf("plans:admin:page:%d:size:%d", request.Page, request.Size)
		} else {
			redisKey = "plans:admin:all"
		}
	}

	// 尝试从 Redis 中获取数据
	redisData, err := dao.Rdb.Get(ctx, redisKey).Result()
	if err == nil && redisData != "" {
		log.Println("Plans existed in redis.")

		// 如果是管理员请求，需要计算 PageCount
		var pageCount int64 = 0
		if !request.IsUser && request.Page > 0 && request.Size > 0 {
			var totalPlans int64
			if err := dao.Db.Model(&model.Plan{}).Count(&totalPlans).Error; err != nil {
				log.Println("count failure:", err)
				return &pb.GetAllPlansResponse{
					Code: http.StatusInternalServerError,
					Msg:  "Count plans failure.",
				}, nil
			}
			pageCount = (totalPlans + int64(request.Size) - 1) / int64(request.Size)
		}

		// 直接返回 Redis 中的缓存数据
		return &pb.GetAllPlansResponse{
			Code:      http.StatusOK,
			Msg:       "Get plans successfully.",
			Plans:     []byte(redisData),
			PageCount: pageCount,
		}, nil
	}
	log.Println("Plans not in redis, query in mysql.")

	// Redis 中无数据，查询 MySQL
	db := dao.Db.Model(&model.Plan{}).Order("sort ASC")
	if request.IsUser {
		db = db.Where("is_sale = ?", true)
	}

	// 如果是管理员请求，执行分页
	var pageCount int64 = 0
	if !request.IsUser {
		var totalPlans int64
		// 查询总数以计算 PageCount
		if err := db.Count(&totalPlans).Error; err != nil {
			log.Println("count failure:", err)
			return &pb.GetAllPlansResponse{
				Code: http.StatusInternalServerError,
				Msg:  "Count plans failure.",
			}, nil
		}

		// 计算 PageCount
		if request.Size > 0 {
			pageCount = (totalPlans + int64(request.Size) - 1) / int64(request.Size)
		}

		// 应用分页限制（仅在请求中指定了分页参数时）
		if request.Page > 0 && request.Size > 0 {
			offset := (request.Page - 1) * request.Size
			db = db.Limit(int(request.Size)).Offset(int(offset))
		}
	}

	// 查询分页后的结果
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
		Code:      http.StatusOK,
		Msg:       "Get plans successfully.",
		Plans:     plansJson,
		PageCount: pageCount,
	}, nil
}
