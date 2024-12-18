package services

import (
	"context"
	"encoding/json"
	pb "keyServices/api/proto"
	"keyServices/internal/dao"
	"keyServices/internal/model"
	"keyServices/internal/utils"
	"log"
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
		Order("request_at desc").
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

//func (s *KeyServices) UnbindKey(ctx context.Context, request *pb.UnbindKeyRequest) (*pb.UnbindKeyResponse, error) {
//	//var err error
//	// 1.修改激活记录的表
//	var activateRecord model.ActivateRecord
//	if result := dao.Db.Model(&model.ActivateRecord{}).Where("id = ? and user_id = ?", request.RecordId, request.UserId).First(&activateRecord); result.Error != nil {
//
//	}
//	activateRecord.IsBind = false
//	// 保存
//	dao.Db.Model(&model.ActivateRecord{}).Where("id = ? and user_id = ?", request.RecordId, request.UserId).Updates(&activateRecord)
//
//	// 2.将key绑定的client_id重新置为空
//	dao.Db.Model(&model.Keys{}).Where("id = ?", activateRecord.KeyId).Updates(map[string]any{"client_id": nil})
//
//	// 3.将key的是否使用设置为否
//	dao.Db.Model(&model.ActiveOrders{}).Where("order_id = ?", activateRecord.OrderId).Updates(map[string]any{"is_used": false})
//
//	// 4.清除用户的key缓存
//	go utils.ClearUserKeyCache(request.UserId)
//
//	return &pb.UnbindKeyResponse{
//		Code: http.StatusOK,
//		Msg:  "success",
//	}, nil
//}

func (s *KeyServices) UnbindKey(ctx context.Context, request *pb.UnbindKeyRequest) (*pb.UnbindKeyResponse, error) {
	// 开始一个事务
	tx := dao.Db.Begin()
	if tx.Error != nil {
		log.Println("事务开始失败:", tx.Error)
		return &pb.UnbindKeyResponse{
			Code: http.StatusInternalServerError,
			Msg:  "事务开始失败",
		}, tx.Error
	}

	// 1. 修改激活记录的表
	var activateRecord model.ActivateRecord
	if result := tx.Model(&model.ActivateRecord{}).
		Where("`id` = ? and user_id = ?", request.RecordId, request.UserId).
		First(&activateRecord); result.Error != nil {
		tx.Rollback() // 查询失败，回滚事务
		log.Println("未找到激活记录:", result.Error)
		return &pb.UnbindKeyResponse{
			Code: http.StatusNotFound,
			Msg:  "未找到激活记录",
		}, result.Error
	}
	log.Println("查找ok", activateRecord)

	// 设置 IsBind 为 false
	//activateRecord.IsBind = false
	//activateRecord.Remark = ""
	if err := tx.Model(&model.ActivateRecord{}).
		Where("id = ? and user_id = ?", request.RecordId, request.UserId).
		Updates(map[string]any{
			"is_bind": false,
			"remark":  nil,
		}).Error; err != nil {
		tx.Rollback() // 更新失败，回滚事务
		log.Println("更新激活记录失败:", err)
		return &pb.UnbindKeyResponse{
			Code: http.StatusInternalServerError,
			Msg:  "更新激活记录失败",
		}, err
	}
	log.Println("修改激活记录ok")

	// 2. 将 key 绑定的 client_id 重新置为空
	if err := tx.Model(&model.Keys{}).
		Where("id = ?", activateRecord.KeyId).
		Updates(map[string]any{"client_id": nil}).Error; err != nil {
		tx.Rollback() // 更新失败，回滚事务
		log.Println("更新密钥的 client_id 失败:", err)
		return &pb.UnbindKeyResponse{
			Code: http.StatusInternalServerError,
			Msg:  "更新密钥的 client_id 失败",
		}, err
	}
	log.Println("修改key ok")

	// 3. 将 key 的是否使用设置为 false
	if err := tx.Model(&model.ActiveOrders{}).
		Where("order_id = ?", activateRecord.OrderId).
		Updates(map[string]any{"is_used": false}).Error; err != nil {
		tx.Rollback() // 更新失败，回滚事务
		log.Println("更新订单的 is_used 失败:", err)
		return &pb.UnbindKeyResponse{
			Code: http.StatusInternalServerError,
			Msg:  "更新订单的 is_used 失败",
		}, err
	}
	log.Println("修改使用记录 ok")

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Println("事务提交失败:", err)
		return &pb.UnbindKeyResponse{
			Code: http.StatusInternalServerError,
			Msg:  "事务提交失败",
		}, err
	}

	// 4. 清除用户的 key 缓存（异步操作，不影响事务）
	go utils.ClearUserKeyCache(request.UserId)

	// 返回成功响应
	return &pb.UnbindKeyResponse{
		Code: http.StatusOK,
		Msg:  "success",
	}, nil
}

func (s *KeyServices) AlterKeyBindRemark(ctx context.Context, request *pb.AlterKeyBindRemarkRequest) (*pb.AlterKeyBindRemarkResponse, error) {
	var code int32 = http.StatusOK
	var msg string = "success"
	if result := dao.Db.Model(&model.ActivateRecord{}).Where("id = ? and user_id = ?", request.RecordId, request.UserId).Updates(map[string]any{
		"remark": request.Remark,
	}); result.Error != nil {
		code = http.StatusInternalServerError
		msg = result.Error.Error()
	}
	return &pb.AlterKeyBindRemarkResponse{
		Code: code,
		Msg:  msg,
	}, nil
}
