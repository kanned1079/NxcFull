package services

import (
	"context"
	"encoding/json"
	pb "keyServices/api/proto"
	"keyServices/internal/dao"
	"keyServices/internal/model"
	"log"
	"net/http"
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

func (s *KeyServices) GetAllMyKeysByUserIdAsc(context context.Context, request *pb.GetAllMyKeysByUserIdAscRequest) (*pb.GetAllMyKeysByUserIdAscResponse, error) {
	log.Println(request.UserId, request.Page, request.Size)
	size := request.Size
	if size <= 0 {
		size = 10 // 默认每页大小为10
	}
	page := request.Page
	if page < 1 {
		page = 1 // 默认页码为1
	}

	// 查询用户的所有有效订单
	var activeOrders []model.ActiveOrders
	if result := dao.Db.Where("user_id = ?", request.UserId).Order("created_at DESC").Find(&activeOrders); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllMyKeysByUserIdAscResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询错误",
		}, nil
	}

	// 如果没有有效订单
	if len(activeOrders) == 0 {
		return &pb.GetAllMyKeysByUserIdAscResponse{
			Code: http.StatusNotFound,
			Msg:  "还没有有效的订单",
		}, nil
	}

	// 计算总页数
	totalRecords := int64(len(activeOrders))
	totalPages := (totalRecords + size - 1) / size

	// 获取当前页的数据
	offset := (page - 1) * size
	end := offset + size
	if end > totalRecords {
		end = totalRecords
	}
	paginatedOrders := activeOrders[offset:end]

	// 创建返回的 key 列表
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

	for _, order := range paginatedOrders {
		var key model.Keys
		if result := dao.Db.Where("id = ?", order.KeyId).First(&key); result.Error != nil {
			log.Println(result.Error)
			continue
		}

		var plan model.Plan
		if result := dao.Db.Where("id = ?", order.PlanId).First(&plan); result.Error != nil {
			log.Println(result.Error)
			continue
		}

		responseData := struct {
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
		}{
			OrderId:        order.OrderId,
			KeyId:          order.KeyId,
			Key:            key.Key,
			PlanId:         order.PlanId,
			PlanName:       plan.Name,
			IsActive:       order.IsActive,
			IsUsed:         order.IsUsed,
			IsValid:        order.IsValid,
			ClientId:       key.ClientId,
			CreatedAt:      key.CreatedAt.Format("2006-01-02 15:04:05"),
			ExpirationDate: order.ExpirationDate.Format("2006-01-02 15:04:05"),
		}

		responseDataList = append(responseDataList, responseData)
	}

	log.Println(responseDataList)

	if myKeyListJson, err := json.Marshal(responseDataList); err != nil {
		log.Println(err)
		return &pb.GetAllMyKeysByUserIdAscResponse{
			Code: http.StatusInternalServerError,
			Msg:  "序列化错误" + err.Error(),
		}, nil
	} else {
		return &pb.GetAllMyKeysByUserIdAscResponse{
			Code:      http.StatusOK,
			Msg:       "获取成功",
			Keys:      myKeyListJson,
			PageCount: totalPages,
		}, nil
	}
}
