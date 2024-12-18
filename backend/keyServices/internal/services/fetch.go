package services

import (
	"context"
	"encoding/json"
	"fmt"
	pb "keyServices/api/proto"
	"keyServices/internal/dao"
	"keyServices/internal/model"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetAllMyKeysByUserIdAsc 获取用户的密钥列表
// v3 优化数据库查询速度 并添加redis缓存
//func (s *KeyServices) GetAllMyKeysByUserIdAsc(ctx context.Context, request *pb.GetAllMyKeysByUserIdAscRequest) (*pb.GetAllMyKeysByUserIdAscResponse, error) {
//	log.Println(request.UserId, request.Page, request.Size)
//	size := request.Size
//	if size <= 0 {
//		size = 10 // 默认每页大小为10
//	}
//	page := request.Page
//	if page < 1 {
//		page = 1 // 默认页码为1
//	}
//	offset := (page - 1) * size
//
//	// Redis 缓存的键值
//	cacheKey := fmt.Sprintf("user_property:user_keys:%d:page:%d:size:%d", request.UserId, page, size)
//
//	// 尝试从 Redis 缓存中获取数据
//	cacheData, err := dao.Rdb.HGetAll(ctx, cacheKey).Result()
//	if err != nil {
//		log.Println("Redis 获取缓存失败:", err)
//	}
//	if len(cacheData) > 0 {
//		// 如果缓存存在，直接解析并返回
//		log.Println("Keys existed in redis.")
//		var responseDataList []struct {
//			OrderId             string `json:"order_id"`
//			KeyId               int64  `json:"key_id"`
//			Key                 string `json:"key"`
//			PlanId              int64  `json:"plan_id"`
//			PlanName            string `json:"plan_name"`
//			IsActive            bool   `json:"is_active"`
//			IsValid             bool   `json:"is_valid"`
//			IsUsed              bool   `json:"is_used"`
//			ClientId            string `json:"client_id"`
//			CreatedAt           string `json:"created_at"`
//			ExpirationDate      string `json:"expiration_date"`
//			ExpirationDateStamp int64  `json:"expiration_date_stamp"`
//		}
//
//		// 从缓存中提取 JSON 并解析
//		err := json.Unmarshal([]byte(cacheData["data"]), &responseDataList)
//		if err != nil {
//			log.Println("缓存解析错误:", err)
//		} else {
//			totalPages, _ := strconv.ParseInt(cacheData["total_pages"], 10, 64)
//			return &pb.GetAllMyKeysByUserIdAscResponse{
//				Code:      http.StatusOK,
//				Msg:       "获取成功（缓存）",
//				Keys:      []byte(cacheData["data"]), // fixed
//				PageCount: totalPages,
//			}, nil
//		}
//	}
//	log.Println("Keys not in redis, query mysql.")
//
//	// 缓存未命中，查询数据库
//	var responseDataList []struct {
//		OrderId             string `json:"order_id"`
//		KeyId               int64  `json:"key_id"`
//		Key                 string `json:"key"`
//		PlanId              int64  `json:"plan_id"`
//		PlanName            string `json:"plan_name"`
//		IsActive            bool   `json:"is_active"`
//		IsValid             bool   `json:"is_valid"`
//		IsUsed              bool   `json:"is_used"`
//		ClientId            string `json:"client_id"`
//		CreatedAt           string `json:"created_at"`
//		ExpirationDate      string `json:"expiration_date"`
//		ExpirationDateStamp int64  `json:"expiration_date_stamp"`
//	}
//
//	activeOrdersTable := model.ActiveOrders{}.TableName()
//	keysTable := model.Keys{}.TableName()
//	plansTable := model.Plan{}.TableName()
//
//	if result := dao.Db.Table(activeOrdersTable).
//		Select(fmt.Sprintf("%s.order_id, %s.key_id, %s.plan_id, %s.is_active, %s.is_valid, %s.is_used, %s.expiration_date, %s.client_id, %s.key, %s.created_at, %s.name AS plan_name",
//			activeOrdersTable, activeOrdersTable, activeOrdersTable, activeOrdersTable, activeOrdersTable, activeOrdersTable, activeOrdersTable,
//			keysTable, keysTable, keysTable, plansTable)).
//		Joins(fmt.Sprintf("JOIN `%s` ON %s.key_id = `%s`.id", keysTable, activeOrdersTable, keysTable)).
//		Joins(fmt.Sprintf("JOIN `%s` ON %s.plan_id = `%s`.id", plansTable, activeOrdersTable, plansTable)).
//		Where(fmt.Sprintf("%s.user_id = ?", activeOrdersTable), request.UserId).
//		Order(fmt.Sprintf("%s.created_at DESC", activeOrdersTable)).
//		Limit(int(size)).
//		Offset(int(offset)).
//		Scan(&responseDataList); result.Error != nil {
//		log.Println(result.Error)
//		return &pb.GetAllMyKeysByUserIdAscResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "查询错误",
//		}, nil
//	}
//
//	// 计算总页数
//	var totalRecords int64
//	dao.Db.Model(&model.ActiveOrders{}).Where("user_id = ?", request.UserId).Count(&totalRecords)
//	totalPages := (totalRecords + int64(size) - 1) / int64(size)
//
//	// 序列化数据并存入 Redis 缓存
//	if myKeyListJson, err := json.Marshal(responseDataList); err != nil {
//		log.Println("序列化错误:", err)
//		return &pb.GetAllMyKeysByUserIdAscResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "序列化错误: " + err.Error(),
//		}, nil
//	} else {
//		// 使用 HASH 结构缓存查询结果
//		cacheData := map[string]interface{}{
//			"data":        myKeyListJson,
//			"total_pages": totalPages,
//			"created_at":  time.Now(),
//		}
//		err = dao.Rdb.HMSet(ctx, cacheKey, cacheData).Err()
//		if err != nil {
//			log.Println("Redis 缓存设置失败:", err)
//		}
//
//		// 设置缓存过期时间
//		dao.Rdb.Expire(ctx, cacheKey, time.Hour) // 设定 1 小时过期时间
//
//		// 返回响应
//		return &pb.GetAllMyKeysByUserIdAscResponse{
//			Code:      http.StatusOK,
//			Msg:       "获取成功",
//			Keys:      myKeyListJson,
//			PageCount: totalPages,
//		}, nil
//	}
//}

func (s *KeyServices) GetAllMyKeysByUserIdAsc(ctx context.Context, request *pb.GetAllMyKeysByUserIdAscRequest) (*pb.GetAllMyKeysByUserIdAscResponse, error) {
	log.Println(request.UserId, request.Page, request.Size)
	size := request.Size
	if size <= 0 {
		size = 10 // 默认每页大小为10
	}
	page := request.Page
	if page < 1 {
		page = 1 // 默认页码为1
	}
	offset := (page - 1) * size

	// Redis 缓存的键值
	cacheKey := fmt.Sprintf("user_property:user_keys:%d:page:%d:size:%d", request.UserId, page, size)

	// 尝试从 Redis 缓存中获取数据
	cacheData, err := dao.Rdb.HGetAll(ctx, cacheKey).Result()
	if err != nil {
		log.Println("Redis 获取缓存失败:", err)
	}
	if len(cacheData) > 0 {
		// 如果缓存存在，直接解析并返回
		log.Println("Keys existed in redis.")
		var responseDataList []struct {
			OrderId             string `json:"order_id"`
			KeyId               int64  `json:"key_id"`
			Key                 string `json:"key"`
			PlanId              int64  `json:"plan_id"`
			PlanName            string `json:"plan_name"`
			IsActive            bool   `json:"is_active"`
			IsValid             bool   `json:"is_valid"`
			IsUsed              bool   `json:"is_used"`
			ClientId            string `json:"client_id"`
			CreatedAt           string `json:"created_at"`
			ExpirationDate      string `json:"expiration_date"`
			ExpirationDateStamp int64  `json:"expiration_date_stamp"`
		}

		// 从缓存中提取 JSON 并解析
		err := json.Unmarshal([]byte(cacheData["data"]), &responseDataList)
		if err != nil {
			log.Println("缓存解析错误:", err)
		} else {
			// 计算ExpirationDateStamp
			for i := range responseDataList {
				if expirationDate, err := time.Parse(time.RFC3339, responseDataList[i].ExpirationDate); err == nil {
					responseDataList[i].ExpirationDateStamp = expirationDate.UnixMilli()
				} else {
					log.Println("解析ExpirationDate错误:", err)
				}
			}

			totalPages, _ := strconv.ParseInt(cacheData["total_pages"], 10, 64)
			return &pb.GetAllMyKeysByUserIdAscResponse{
				Code:      http.StatusOK,
				Msg:       "获取成功（缓存）",
				Keys:      []byte(cacheData["data"]), // fixed
				PageCount: totalPages,
			}, nil
		}
	}
	log.Println("Keys not in redis, query mysql.")

	// 缓存未命中，查询数据库
	var responseDataList []struct {
		OrderId             string `json:"order_id"`
		KeyId               int64  `json:"key_id"`
		Key                 string `json:"key"`
		PlanId              int64  `json:"plan_id"`
		PlanName            string `json:"plan_name"`
		IsActive            bool   `json:"is_active"`
		IsValid             bool   `json:"is_valid"`
		IsUsed              bool   `json:"is_used"`
		ClientId            string `json:"client_id"`
		CreatedAt           string `json:"created_at"`
		ExpirationDate      string `json:"expiration_date"`
		ExpirationDateStamp int64  `json:"expiration_date_stamp"`
	}

	activeOrdersTable := model.ActiveOrders{}.TableName()
	keysTable := model.Keys{}.TableName()
	plansTable := model.Plan{}.TableName()

	if result := dao.Db.Table(activeOrdersTable).
		Select(fmt.Sprintf("%s.order_id, %s.key_id, %s.plan_id, %s.is_active, %s.is_valid, %s.is_used, %s.expiration_date, %s.client_id, %s.key, %s.created_at, %s.name AS plan_name",
			activeOrdersTable, activeOrdersTable, activeOrdersTable, activeOrdersTable, activeOrdersTable, activeOrdersTable, activeOrdersTable,
			keysTable, keysTable, keysTable, plansTable)).
		Joins(fmt.Sprintf("JOIN `%s` ON %s.key_id = `%s`.id", keysTable, activeOrdersTable, keysTable)).
		Joins(fmt.Sprintf("JOIN `%s` ON %s.plan_id = `%s`.id", plansTable, activeOrdersTable, plansTable)).
		Where(fmt.Sprintf("%s.user_id = ?", activeOrdersTable), request.UserId).
		Order(fmt.Sprintf("%s.created_at DESC", activeOrdersTable)).
		Limit(int(size)).
		Offset(int(offset)).
		Scan(&responseDataList); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllMyKeysByUserIdAscResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询错误",
		}, nil
	}

	// 计算总页数
	var totalRecords int64
	dao.Db.Model(&model.ActiveOrders{}).Where("user_id = ?", request.UserId).Count(&totalRecords)
	totalPages := (totalRecords + int64(size) - 1) / int64(size)

	// 计算ExpirationDateStamp
	for i := range responseDataList {
		if expirationDate, err := time.Parse(time.RFC3339, responseDataList[i].ExpirationDate); err == nil {
			responseDataList[i].ExpirationDateStamp = expirationDate.UnixMilli()
		} else {
			log.Println("解析ExpirationDate错误:", err)
		}
	}

	// 序列化数据并存入 Redis 缓存
	if myKeyListJson, err := json.Marshal(responseDataList); err != nil {
		log.Println("序列化错误:", err)
		return &pb.GetAllMyKeysByUserIdAscResponse{
			Code: http.StatusInternalServerError,
			Msg:  "序列化错误: " + err.Error(),
		}, nil
	} else {
		// 使用 HASH 结构缓存查询结果
		cacheData := map[string]interface{}{
			"data":        myKeyListJson,
			"total_pages": totalPages,
			"created_at":  time.Now(),
		}
		err = dao.Rdb.HMSet(ctx, cacheKey, cacheData).Err()
		if err != nil {
			log.Println("Redis 缓存设置失败:", err)
		}

		// 设置缓存过期时间
		dao.Rdb.Expire(ctx, cacheKey, time.Hour) // 设定 1 小时过期时间

		// 返回响应
		return &pb.GetAllMyKeysByUserIdAscResponse{
			Code:      http.StatusOK,
			Msg:       "获取成功",
			Keys:      myKeyListJson,
			PageCount: totalPages,
		}, nil
	}
}
