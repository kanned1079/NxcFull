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

//func (s *KeyServices) GetAllMyKeysByUserIdAsc(context context.Context, request *pb.GetAllMyKeysByUserIdAscRequest) (*pb.GetAllMyKeysByUserIdAscResponse, error) {
//	//request.Size
//	//request.Page
//	// 查询用户的所有有效订单
//	var activeOrders []model.ActiveOrders
//	if result := dao.Db.Where("user_id = ?", request.UserId).Find(&activeOrders); result.Error != nil {
//		log.Println(result.Error)
//		//context.JSON(http.StatusInternalServerError, gin.H{
//		//	"code":  http.StatusInternalServerError,
//		//	"msg":   "获取用户的有效订单错误",
//		//	"error": result.Error.Error(),
//		//})
//		return &pb.GetAllMyKeysByUserIdAscResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "查询错误",
//		}, nil
//	}
//	if len(activeOrders) == 0 {
//		return &pb.GetAllMyKeysByUserIdAscResponse{
//			Code: http.StatusNotFound,
//			Msg:  "还没有有效的订单",
//		}, nil
//	}
//	// 创建返回的 key 列表
//	var responseDataList []struct {
//		OrderId        int64  `json:"order_id"`
//		KeyId          int64  `json:"key_id"`
//		Key            string `json:"key"`
//		PlanId         int64  `json:"plan_id"`
//		PlanName       string `json:"plan_name"`
//		IsActive       bool   `json:"is_active"`
//		IsUsed         bool   `json:"is_used"`
//		ClientId       string `json:"client_id"`
//		CreatedAt      string `json:"created_at"`
//		ExpirationDate string `json:"expiration_date"`
//	}
//
//	for _, order := range activeOrders {
//		// 根据 key_id 查询密钥信息
//		var key model.Keys
//		if result := dao.Db.Where("id = ?", order.KeyId).First(&key); result.Error != nil {
//			log.Println(result.Error)
//			continue
//		}
//
//		// 根据 plan_id 查询计划名称
//		var plan model.Plan
//		if result := dao.Db.Where("id = ?", order.PlanId).First(&plan); result.Error != nil {
//			log.Println(result.Error)
//			continue
//		}
//
//		// 组装订单和密钥的信息
//		responseData := struct {
//			OrderId        int64  `json:"order_id"`
//			KeyId          int64  `json:"key_id"`
//			Key            string `json:"key"`
//			PlanId         int64  `json:"plan_id"`
//			PlanName       string `json:"plan_name"`
//			IsActive       bool   `json:"is_active"`
//			IsUsed         bool   `json:"is_used"`
//			ClientId       string `json:"client_id"`
//			CreatedAt      string `json:"created_at"`
//			ExpirationDate string `json:"expiration_date"`
//		}{
//			OrderId:  order.ID,
//			KeyId:    order.KeyId,
//			Key:      key.Key,
//			PlanId:   order.PlanId,
//			PlanName: plan.Name,
//			IsActive: order.IsActive,
//			IsUsed:   order.IsUsed,
//			ClientId: key.ClientId,
//			//CreatedAt:      key.CreatedAt,
//			CreatedAt:      key.CreatedAt.Format("2006-01-02 15:04:05"),
//			ExpirationDate: order.ExpirationDate.Format("2006-01-02 15:04:05"),
//		}
//
//		// 将每个订单信息添加到 responseDataList
//		responseDataList = append(responseDataList, responseData)
//	}
//
//	if myKeyListJson, err := json.Marshal(responseDataList); err != nil {
//		log.Println(err)
//		return &pb.GetAllMyKeysByUserIdAscResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "序列化错误" + err.Error(),
//		}, nil
//	} else {
//		return &pb.GetAllMyKeysByUserIdAscResponse{
//			Code: http.StatusOK,
//			Msg:  "获取成功",
//			Keys: myKeyListJson,
//			PageCount: 计算这里的页数
//		}, nil
//	}
//
//}

// GetAllMyKeysByUserIdAsc 查询用户所有的密钥
// v1 没有优化数据库查询速度的版本
//func (s *KeyServices) GetAllMyKeysByUserIdAsc(context context.Context, request *pb.GetAllMyKeysByUserIdAscRequest) (*pb.GetAllMyKeysByUserIdAscResponse, error) {
//	log.Println(request.UserId, request.Page, request.Size)
//	size := request.Size
//	if size <= 0 {
//		size = 10 // 默认每页大小为10
//	}
//	page := request.Page
//	if page < 1 {
//		page = 1 // 默认页码为1
//	}
//
//	// 查询用户的所有有效订单
//	var activeOrders []model.ActiveOrders
//	if result := dao.Db.Where("user_id = ?", request.UserId).Order("created_at DESC").Find(&activeOrders); result.Error != nil {
//		log.Println(result.Error)
//		return &pb.GetAllMyKeysByUserIdAscResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "查询错误",
//		}, nil
//	}
//
//	// 计算总页数
//	totalRecords := int64(len(activeOrders))
//	totalPages := (totalRecords + size - 1) / size
//
//	// 获取当前页的数据
//	offset := (page - 1) * size
//	end := offset + size
//	if end > totalRecords {
//		end = totalRecords
//	}
//	paginatedOrders := activeOrders[offset:end]
//
//	// 创建返回的 key 列表
//	var responseDataList []struct {
//		OrderId        string `json:"order_id"`
//		KeyId          int64  `json:"key_id"`
//		Key            string `json:"key"`
//		PlanId         int64  `json:"plan_id"`
//		PlanName       string `json:"plan_name"`
//		IsActive       bool   `json:"is_active"`
//		IsValid        bool   `json:"is_valid"`
//		IsUsed         bool   `json:"is_used"`
//		ClientId       string `json:"client_id"`
//		CreatedAt      string `json:"created_at"`
//		ExpirationDate string `json:"expiration_date"`
//	}
//
//	for _, order := range paginatedOrders {
//		// 这里获取密钥列表
//		var key model.Keys
//		if result := dao.Db.Where("id = ?", order.KeyId).First(&key); result.Error != nil {
//			log.Println(result.Error)
//			continue
//		}
//		// 这里获取密钥对应的订阅计划的信息
//		var plan model.Plan
//		if result := dao.Db.Where("id = ?", order.PlanId).First(&plan); result.Error != nil {
//			log.Println(result.Error)
//			continue
//		}
//
//		responseData := struct {
//			OrderId        string `json:"order_id"`
//			KeyId          int64  `json:"key_id"`
//			Key            string `json:"key"`
//			PlanId         int64  `json:"plan_id"`
//			PlanName       string `json:"plan_name"`
//			IsActive       bool   `json:"is_active"`
//			IsValid        bool   `json:"is_valid"`
//			IsUsed         bool   `json:"is_used"`
//			ClientId       string `json:"client_id"`
//			CreatedAt      string `json:"created_at"`
//			ExpirationDate string `json:"expiration_date"`
//		}{
//			OrderId:        order.OrderId,
//			KeyId:          order.KeyId,
//			Key:            key.Key,
//			PlanId:         order.PlanId,
//			PlanName:       plan.Name,
//			IsActive:       order.IsActive,
//			IsUsed:         order.IsUsed,
//			IsValid:        order.IsValid,
//			ClientId:       key.ClientId,
//			CreatedAt:      key.CreatedAt.Format("2006-01-02 15:04:05"),
//			ExpirationDate: order.ExpirationDate.Format("2006-01-02 15:04:05"),
//		}
//
//		responseDataList = append(responseDataList, responseData)
//	}
//	if myKeyListJson, err := json.Marshal(responseDataList); err != nil {
//		log.Println(err)
//		return &pb.GetAllMyKeysByUserIdAscResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "序列化错误" + err.Error(),
//		}, nil
//	} else {
//		return &pb.GetAllMyKeysByUserIdAscResponse{
//			Code:      http.StatusOK,
//			Msg:       "获取成功",
//			Keys:      myKeyListJson,
//			PageCount: totalPages,
//		}, nil
//	}
//}

// GetAllMyKeysByUserIdAsc 获取用户的密钥列表
// v2 加速查询速度
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
//
//	// 计算分页的偏移量
//	offset := (page - 1) * size
//
//	// 使用 JOIN 查询从三个表中一次性获取数据
//	var responseDataList []struct {
//		OrderId        string `json:"order_id"`
//		KeyId          int64  `json:"key_id"`
//		Key            string `json:"key"`
//		PlanId         int64  `json:"plan_id"`
//		PlanName       string `json:"plan_name"`
//		IsActive       bool   `json:"is_active"`
//		IsValid        bool   `json:"is_valid"`
//		IsUsed         bool   `json:"is_used"`
//		ClientId       string `json:"client_id"`
//		CreatedAt      string `json:"created_at"`
//		ExpirationDate string `json:"expiration_date"`
//	}
//
//	// 执行多表 JOIN 查询，按 `created_at` 倒序并应用分页
//	// 获取表名
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
//	// 计算总记录数以获取总页数
//	var totalRecords int64
//	dao.Db.Model(&model.ActiveOrders{}).Where("user_id = ?", request.UserId).Count(&totalRecords)
//	totalPages := (totalRecords + int64(size) - 1) / int64(size)
//
//	// 序列化并返回响应
//	if myKeyListJson, err := json.Marshal(responseDataList); err != nil {
//		log.Println(err)
//		return &pb.GetAllMyKeysByUserIdAscResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "序列化错误: " + err.Error(),
//		}, nil
//	} else {
//		return &pb.GetAllMyKeysByUserIdAscResponse{
//			Code:      http.StatusOK,
//			Msg:       "获取成功",
//			Keys:      myKeyListJson,
//			PageCount: totalPages,
//		}, nil
//	}
//}

// GetAllMyKeysByUserIdAsc 获取用户的密钥列表
// v3 优化数据库查询速度 并添加redis缓存
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
	cacheKey := fmt.Sprintf("user_keys:%d:page:%d:size:%d", request.UserId, page, size)

	// 尝试从 Redis 缓存中获取数据
	cacheData, err := dao.Rdb.HGetAll(ctx, cacheKey).Result()
	if err != nil {
		log.Println("Redis 获取缓存失败:", err)
	}
	if len(cacheData) > 0 {
		// 如果缓存存在，直接解析并返回
		log.Println("Keys existed in redis.")
		var responseDataList []struct {
			OrderId        string `json:"order_id"`
			KeyId          int64  `json:"key_id"`
			Key            string `json:"key"`
			PlanId         int64  `json:"plan_id"`
			PlanName       string `json:"plan_name"`
			IsActive       bool   `json:"is_active"`
			IsValid        bool   `json:"is_valid"`
			IsUsed         bool   `json:"is_used"`
			ClientId       string `json:"client_id"`
			CreatedAt      string `json:"created_at"`
			ExpirationDate string `json:"expiration_date"`
		}

		// 从缓存中提取 JSON 并解析
		err := json.Unmarshal([]byte(cacheData["data"]), &responseDataList)
		if err != nil {
			log.Println("缓存解析错误:", err)
		} else {
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
		OrderId        string `json:"order_id"`
		KeyId          int64  `json:"key_id"`
		Key            string `json:"key"`
		PlanId         int64  `json:"plan_id"`
		PlanName       string `json:"plan_name"`
		IsActive       bool   `json:"is_active"`
		IsValid        bool   `json:"is_valid"`
		IsUsed         bool   `json:"is_used"`
		ClientId       string `json:"client_id"`
		CreatedAt      string `json:"created_at"`
		ExpirationDate string `json:"expiration_date"`
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
