package services

import (
	"context"
	pb "keyServices/api/proto"
	"keyServices/internal/client"
	userPb "keyServices/internal/client/api/user/proto"
	"keyServices/internal/dao"
	"keyServices/internal/model"
	"keyServices/internal/utils"
	"log"
	"net/http"
	"time"
)

//func (s *KeyServices) BindOrActiveMyKey2App(ctx context.Context, request *pb.BindOrActiveMyKey2AppRequest) (*pb.BindOrActiveMyKey2AppResponse, error) {
//	// 调用用户微服务来验证账户信息
//	authResponse, err := client.GrpcClient.UserServiceClient.CheckUserAuthInternal(context.Background(), &userPb.CheckUserAuthInternalRequest{
//		Email:    request.Email,
//		Password: request.Password,
//	})
//	if err != nil { // 如果微服务调用出现错误
//		return &pb.BindOrActiveMyKey2AppResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "rpc error:" + err.Error(),
//		}, nil
//	}
//	if authResponse.Code != http.StatusOK || authResponse.UserId == -1 { // 如果返回的信息有错误则退出
//		return &pb.BindOrActiveMyKey2AppResponse{
//			Code: http.StatusConflict,
//			Msg:  "username or password incorrect",
//		}, nil
//	}
//	log.Println("验证账户通过 开始验证密钥")
//	log.Println(request.Key)
//
//	// 查找密钥
//	var key model.Keys
//	dao.Db.Model(&model.Keys{}).Where("`key` = ?", request.Key).First(&key)
//
//	// 查询该密钥的使用细节
//	var activeOrder model.ActiveOrders
//	dao.Db.Model(&model.ActiveOrders{}).Where("key_id = ? and email = ?", key.Id, request.Email).First(&activeOrder)
//
//	// 检查该密钥
//	// activeOrder.IsValid 该密钥是否被管理员禁用 false为禁用状态
//	// activeOrder.IsUsed false为未被使用 true则为使用过了
//	// 要求activeOrder.IsValid为true并且activeOrder.IsUsed为false才能继续下一步
//
//	// 开始操作
//	// 1.新增一条激活记录
//	reqAt := time.Now()
//	var NewActivateRecord model.ActivateRecord = model.ActivateRecord{
//		UserId:        authResponse.UserId,
//		Email:         request.Email,
//		OrderId:       activeOrder.OrderId,
//		KeyId:         key.Id,
//		RequestAt:     &reqAt,
//		ClientVersion: request.ClientVersion,
//		OsType:        request.OsType,
//		Remark:        request.Remark,
//		IsBind:        true,
//	}
//	dao.Db.Model(&model.ActivateRecord{}).Create(&NewActivateRecord)
//
//	// 2.修改密钥绑定的ClientId
//	key.ClientId = request.ClientId
//	dao.Db.Model(&model.Keys{}).Save(&key)
//
//	// 3.修改密钥的激活信息
//	activeOrder.IsUsed = true
//	dao.Db.Model(&model.ActiveOrders{}).Save(&activeOrder)
//
//	return &pb.BindOrActiveMyKey2AppResponse{
//		Code:              http.StatusOK,
//		Msg:               "success",
//		SignificantNumber: "23456ytgfdd",
//	}, nil
//
//}

func (s *KeyServices) BindOrActiveMyKey2App(ctx context.Context, request *pb.BindOrActiveMyKey2AppRequest) (*pb.BindOrActiveMyKey2AppResponse, error) {
	// 调用用户微服务来验证账户信息
	authResponse, err := client.GrpcClient.UserServiceClient.CheckUserAuthInternal(context.Background(), &userPb.CheckUserAuthInternalRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		// 如果微服务调用出现错误
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusInternalServerError,
			Msg:  "rpc error: " + err.Error(),
		}, nil
	}
	if authResponse.Code != http.StatusOK || authResponse.UserId == -1 {
		// 如果返回的信息有错误则退出
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusConflict,
			Msg:  "username or password incorrect",
		}, nil
	}

	log.Println("请求的密钥：", request.Key)

	// 查找密钥
	var key model.Keys
	if err := dao.Db.Model(&model.Keys{}).Where("`key` = ?", request.Key).First(&key).Error; err != nil {
		log.Printf("找不到密钥：%v\n", err)
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusNotFound,
			Msg:  "key not found",
		}, nil
	}

	// 查询该密钥的使用细节
	var activeOrder model.ActiveOrders
	if err := dao.Db.Model(&model.ActiveOrders{}).Where("`key_id` = ? and `email` = ?", key.Id, request.Email).First(&activeOrder).Error; err != nil {
		log.Printf("找不到激活订单：%v\n", err)
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusNotFound,
			Msg:  "activation order not found",
		}, nil
	}

	// 检查该密钥
	if !activeOrder.IsValid {
		// 密钥被禁用
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusForbidden,
			Msg:  "this key is disabled by the admin",
		}, nil
	}
	if activeOrder.IsUsed {
		// 密钥已经使用
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusConflict,
			Msg:  "this key has already been used",
		}, nil
	}

	// 开始操作：使用数据库事务来确保操作的一致性
	tx := dao.Db.Begin()
	if tx.Error != nil {
		log.Printf("开始事务失败：%v\n", tx.Error)
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusInternalServerError,
			Msg:  "failed to begin transaction",
		}, nil
	}

	// 1. 新增一条激活记录
	reqAt := time.Now()
	NewActivateRecord := model.ActivateRecord{
		UserId:    authResponse.UserId,
		Email:     request.Email,
		OrderId:   activeOrder.OrderId,
		KeyId:     key.Id,
		RequestAt: &reqAt,

		ClientVersion: request.ClientVersion,
		OsType:        request.OsType,
		Remark:        request.Remark,
		IsBind:        true,
	}

	if err := tx.Model(&model.ActivateRecord{}).Create(&NewActivateRecord).Error; err != nil {
		tx.Rollback()
		log.Printf("创建激活记录失败：%v\n", err)
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusInternalServerError,
			Msg:  "failed to create activation record",
		}, nil
	}

	// 2. 修改密钥绑定的ClientId
	//key.ClientId = request.ClientId
	newClientId, err := utils.GenerateClientId()
	if err != nil {
		tx.Rollback()
		log.Printf("err on generate client_id: %v\n", err)
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusInternalServerError,
			Msg:  "err on generate client_id: " + err.Error(),
		}, nil
	}

	key.ClientId = newClientId
	if err := tx.Model(&model.Keys{}).Where("`id` = ?", key.Id).Updates(&key).Error; err != nil {
		tx.Rollback()
		log.Printf("修改密钥绑定信息失败：%v\n", err)
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusInternalServerError,
			Msg:  "failed to update key binding",
		}, nil
	}

	// 3. 修改密钥的激活信息
	activeOrder.IsUsed = true
	if err := tx.Model(&model.ActiveOrders{}).Where("`id` = ?", activeOrder.ID).Updates(&activeOrder).Error; err != nil {
		tx.Rollback()
		log.Printf("修改激活订单失败：%v\n", err)
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusInternalServerError,
			Msg:  "failed to update activation order",
		}, nil
	}

	if err := utils.SetActivationRecordCache2Redis(ctx, newClientId, key.Key, activeOrder.ExpirationDate); err != nil {
		tx.Rollback()
		log.Printf(err.Error())
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusInternalServerError,
			Msg:  "failed to set cache: " + err.Error(),
		}, nil
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Printf("提交事务失败：%v\n", err)
		return &pb.BindOrActiveMyKey2AppResponse{
			Code: http.StatusInternalServerError,
			Msg:  "failed to commit transaction",
		}, nil
	}

	// 清除用户的缓存
	go utils.ClearUserKeyCache(authResponse.UserId)

	log.Println("密钥绑定并激活成功")
	return &pb.BindOrActiveMyKey2AppResponse{
		Code:     http.StatusOK,
		Msg:      "success",
		ClientId: key.ClientId,
	}, nil
}
