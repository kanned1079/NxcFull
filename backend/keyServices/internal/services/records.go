package services

import (
	"context"
	"encoding/json"
	pb "keyServices/api/proto"
	"keyServices/internal/dao"
	"keyServices/internal/model"
	"math"
	"net/http"
)

// GetActivateLogByUserId 根據用戶的id來檢索所有的激活紀錄
func (s *KeyServices) GetActivateLogByUserId(ctx context.Context, request *pb.GetActivateLogByUserIdRequest) (*pb.GetActivateLogByUserIdResponse, error) {
	// 获取分页参数
	page := request.GetPage()
	size := request.GetSize()

	// 如果分页参数无效，默认设置为第1页，每页10条记录
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	var err error
	var records []model.ActivateRecord
	var totalCount int64

	// 获取总记录数
	if result := dao.Db.Model(&model.ActivateRecord{}).Where("user_id = ?", request.GetUserId()).Count(&totalCount); result.Error != nil {
		return &pb.GetActivateLogByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "error: " + result.Error.Error(),
		}, nil
	}

	// 根据分页参数查询数据
	if result := dao.Db.Model(&model.ActivateRecord{}).
		Where("user_id = ?", request.GetUserId()).
		Offset(int((page - 1) * size)). // 设置偏移量
		Limit(int(size)).               // 设置每页数量
		Find(&records); result.Error != nil {
		return &pb.GetActivateLogByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "error: " + result.Error.Error(),
		}, nil
	}

	// 将查询到的记录序列化成JSON
	recordsJson, err := json.Marshal(records)
	if err != nil {
		return &pb.GetActivateLogByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "error: " + err.Error(),
		}, nil
	}

	// 计算总页数
	pageCount := int64(math.Ceil(float64(totalCount) / float64(size)))

	return &pb.GetActivateLogByUserIdResponse{
		Code:      http.StatusOK,
		Msg:       "success",
		Log:       recordsJson,
		PageCount: pageCount, // 返回总页数
	}, nil
}

func (s *KeyServices) GetKeyInfoById(ctx context.Context, request *pb.GetKeyInfoByIdRequest) (*pb.GetKeyInfoByIdResponse, error) {
	var key model.Keys

	if result := dao.Db.Model(&model.Keys{}).Where("id = ?", request.KeyId).First(&key); result.Error != nil {
		return &pb.GetKeyInfoByIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "error: " + result.Error.Error(),
		}, nil
	}
	jsonData, err := json.Marshal(key)
	if err != nil {
		return &pb.GetKeyInfoByIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "error: " + err.Error(),
		}, nil
	}

	return &pb.GetKeyInfoByIdResponse{
		Code:    http.StatusOK,
		Msg:     "success",
		Details: jsonData,
	}, nil
}
