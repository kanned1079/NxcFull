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

func init() {

}

var activeOrdersTable = model.ActiveOrders{}.TableName()
var keysTable = model.Keys{}.TableName()
var plansTable = model.Plan{}.TableName()

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

	//activeOrdersTable := model.ActiveOrders{}.TableName()
	//keysTable := model.Keys{}.TableName()
	//plansTable := model.Plan{}.TableName()

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

//func (s *KeyServices) GetAllKeysByAdminDesc(ctx context.Context, request *pb.GetAllKeysByAdminDescRequest) (*pb.GetAllKeysByAdminDescResponse, error) {
//	// request.Page 分页查询参数
//	// request.Size 分页查询参数
//	// request.SearchKeyId 查询key的id
//	// request.SearchKeyContent 模糊查询key的内容
//	// 如果SearchKeyId大于0则按照这个查询 就不需要使用SearchKeyContent查询了
//	// 只有当SearchKeyId为小于等于0的时候才需要根据SearchKeyContent查询
//	var responseDataList []struct { // 这里是需要查询出的数据
//		OrderId             string `json:"order_id"`
//		KeyId               int64  `json:"key_id"`
//		Key                 string `json:"key"`
//		PlanId              int64  `json:"plan_id"`
//		PlanName            string `json:"plan_name"`
//		ClientId            string `json:"client_id"`
//		CreatedAt           string `json:"created_at"`
//		ExpirationDate      string `json:"expiration_date"`       // 这个应该是从数据库里取出来额
//		ExpirationDateStamp int64  `json:"expiration_date_stamp"` // 随后这个是根据上面的进行计算
//	}
//
//	return &pb.GetAllKeysByAdminDescResponse{
//		Code: http.StatusOK,
//		Msg:  "success",
//		//Keys:	// 这里需要将上面的列表序列化后传入
//		//PageCount:	// 这里根据上面的page和size计算总共的页数
//	}, nil
//}

func (s *KeyServices) GetAllKeysByAdminDesc(ctx context.Context, request *pb.GetAllKeysByAdminDescRequest) (*pb.GetAllKeysByAdminDescResponse, error) {
	// 定义返回的数据结构，将 ExpirationDate 定义为 time.Time 类型
	var responseDataList []struct {
		OrderId             string    `json:"order_id"`
		Email               string    `json:"email"`
		KeyId               int64     `json:"key_id"`
		Key                 string    `json:"key"`
		IsValid             bool      `json:"is_valid"`
		IsUsed              bool      `json:"is_used"`
		PlanId              int64     `json:"plan_id"`
		PlanName            string    `json:"plan_name"`
		ClientId            string    `json:"client_id"`
		CreatedAt           string    `json:"created_at"`
		ExpirationDate      time.Time `json:"expiration_date"` // 将 ExpirationDate 类型更改为 time.Time
		ExpirationDateStamp int64     `json:"expiration_date_stamp"`
	}

	// 构建查询的基本条件
	query := dao.Db.Table(activeOrdersTable).
		Select(fmt.Sprintf("%s.order_id, %s.email, %s.key_id, %s.plan_id, %s.client_id, %s.key, %s.is_valid, %s.is_used, %s.created_at, %s.expiration_date, %s.name AS plan_name",
			activeOrdersTable, activeOrdersTable, activeOrdersTable, activeOrdersTable, keysTable, keysTable, activeOrdersTable, activeOrdersTable, keysTable,
			activeOrdersTable, plansTable)).
		Joins(fmt.Sprintf("JOIN `%s` ON %s.key_id = `%s`.id", keysTable, activeOrdersTable, keysTable)).
		Joins(fmt.Sprintf("JOIN `%s` ON %s.plan_id = `%s`.id", plansTable, activeOrdersTable, plansTable)).
		Order(fmt.Sprintf("%s.created_at DESC", activeOrdersTable))

	// 根据 SearchKeyId 或 SearchKeyContent 设置查询条件
	if request.SearchKeyId > 0 {
		// 如果 SearchKeyId 大于 0，则根据它查询
		query = query.Where(fmt.Sprintf("%s.key_id = ?", activeOrdersTable), request.SearchKeyId)
	} else if len(request.SearchKeyContent) > 0 {
		// 如果 SearchKeyContent 不为空，则进行模糊查询
		query = query.Where(fmt.Sprintf("%s.key LIKE ?", activeOrdersTable), "%"+request.SearchKeyContent+"%")
	}

	// 分页查询
	query = query.Limit(int(request.Size)).Offset(int(request.Page-1) * int(request.Size))

	// 执行查询
	result := query.Scan(&responseDataList)
	if result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllKeysByAdminDescResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询错误" + result.Error.Error(),
		}, nil
	}

	// 将 ExpirationDate 转换为 ExpirationDateStamp
	for i := range responseDataList {
		// 直接获取 Unix 时间戳
		responseDataList[i].ExpirationDateStamp = responseDataList[i].ExpirationDate.UnixMilli()
	}

	// 计算总页数
	var totalCount int64
	countResult := dao.Db.Table(activeOrdersTable).
		Count(&totalCount)
	if countResult.Error != nil {
		log.Println(countResult.Error)
		return &pb.GetAllKeysByAdminDescResponse{
			Code: http.StatusInternalServerError,
			Msg:  "计算总页数失败" + countResult.Error.Error(),
		}, nil
	}
	pageCount := int(totalCount / request.Size)
	if totalCount%request.Size > 0 {
		pageCount++
	}

	// 转换为 JSON 格式
	jsonList, err := json.Marshal(responseDataList)
	if err != nil {
		return &pb.GetAllKeysByAdminDescResponse{
			Code:      http.StatusInternalServerError,
			Msg:       err.Error(),
			PageCount: 0,
		}, nil
	}

	// 返回响应
	return &pb.GetAllKeysByAdminDescResponse{
		Code:      http.StatusOK,
		Msg:       "success",
		Keys:      jsonList, // 返回包含 ExpirationDateStamp 的数据
		PageCount: int64(pageCount),
	}, nil
}
