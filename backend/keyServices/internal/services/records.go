package services

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	pb "keyServices/api/proto"
	"keyServices/internal/dao"
	"keyServices/internal/model"
	"keyServices/internal/utils"
	"log"
	"math"
	"net/http"
	"strings"
	"time"
)

// GetActivateLogByAdmin 管理员获取所有的密钥列表
func (s *KeyServices) GetActivateLogByAdmin(ctx context.Context, request *pb.GetActivateLogByAdminRequest) (*pb.GetActivateLogByAdminResponse, error) {
	// 新增功能说明：
	// request.Email 模糊查询 email，对应数据库的 email 字段
	// request.Valid 如果为 true，则只查询字段 is_bind 为 true 的记录
	// request.Sort 默认为 DESC，如果为 ASC，则按 created_at 升序排列

	page := request.Page
	size := request.Size

	// 默认页数和每页大小检查
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10 // 默认每页显示10条
	}

	var records []model.ActivateRecord

	// 计算分页的偏移量和限制
	offset := (page - 1) * size

	// 构建查询条件
	query := dao.Db.Model(&model.ActivateRecord{})

	// 模糊查询 email
	if request.Email != "" {
		query = query.Where("email LIKE ?", "%"+request.Email+"%")
	}

	// 筛选 is_bind 字段
	if request.Valid {
		query = query.Where("is_bind = ?", true)
	}

	// 排序
	sortOrder := "DESC"
	if strings.ToUpper(request.Sort) == "ASC" {
		sortOrder = "ASC"
	}
	query = query.Order("created_at " + sortOrder)

	// 查询记录
	if result := query.Offset(int(offset)).Limit(int(size)).Find(&records); result.Error != nil {
		return &pb.GetActivateLogByAdminResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Error fetching records: " + result.Error.Error(),
		}, nil
	}

	// 查询总记录数
	var totalCount int64
	if result := query.Count(&totalCount); result.Error != nil {
		return &pb.GetActivateLogByAdminResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Error counting records: " + result.Error.Error(),
		}, nil
	}

	// 计算总页数
	pageCount := totalCount / int64(size)
	if totalCount%int64(size) != 0 {
		pageCount++
	}

	// 序列化记录
	recordsJson, err := json.Marshal(records)
	if err != nil {
		return &pb.GetActivateLogByAdminResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Error serializing records: " + err.Error(),
		}, nil
	}

	return &pb.GetActivateLogByAdminResponse{
		Code:      http.StatusOK,
		Log:       recordsJson,
		PageCount: pageCount,
	}, nil
}

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

	if err := utils.CloseActiveActivationCheckCacheByKeyId(ctx, activateRecord.KeyId); err != nil {
		log.Println("关闭激活检查缓存失败:", err)
		tx.Rollback() // 更新失败，回滚事务
		return &pb.UnbindKeyResponse{
			Code: http.StatusInternalServerError,
			Msg:  "关闭激活检查缓存失败",
		}, err
	}

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

func (s *KeyServices) AlterKeyIsValidByAdmin(ctx context.Context, request *pb.AlterKeyIsValidByAdminRequest) (*pb.AlterKeyIsValidByAdminResponse, error) {
	if err := utils.CloseActiveActivationCheckCacheByKeyId(ctx, request.KeyId); err != nil {
		return &pb.AlterKeyIsValidByAdminResponse{
			Code: http.StatusInternalServerError,
			Msg:  "error: " + err.Error(),
		}, nil
	}
	// 如果清楚缓存都没有成功 取消后面的操作
	if result := dao.Db.Model(&model.ActiveOrders{}).
		Where("`key_id` = ?", request.KeyId).
		Update("is_valid", gorm.Expr("NOT `is_valid`")); result.RowsAffected == 0 {
		return &pb.AlterKeyIsValidByAdminResponse{
			Code: http.StatusInternalServerError,
			Msg:  "error: " + result.Error.Error(),
		}, nil
	}
	utils.ClearUserKeyCache(request.UserId)
	return &pb.AlterKeyIsValidByAdminResponse{
		Code: http.StatusOK,
		Msg:  "success",
	}, nil
}

//func (s *KeyServices) CheckActivationStatusIntervalReq(ctx context.Context, request *pb.CheckActivationStatusIntervalReqRequest) (*pb.CheckActivationStatusIntervalReqResponse, error) {
//
//	code, expiredAt := utils.CheckActivationFromCache(ctx, request.ClientId, request.KeyContent)
//	log.Println("检查: ", code, expiredAt)
//	// 200 成功 直接返回 取消下面查询数据库
//	// 404 在redis中查询不到该密钥
//	// 409 密钥和客户端ID不对应
//	// 500 其他错误
//	// 除了200 其他状态码后都需要查询数据库
//
//	if code != 200 {
//		// 1 先查询密钥
//		var key model.Keys
//		result := dao.Db.Model(&model.Keys{}).Where("`key` = ?", request.KeyContent).First(&key)
//		if result.RowsAffected == 0 {
//			code = http.StatusNotFound // 404
//		} else if result.Error != nil {
//			code = http.StatusInternalServerError // 500
//		}
//		// 2 再根据密钥的id查询有效订单表中的is_valid字段来判断该密钥是否被管理员禁用 如果是false则是禁用状态
//		var activeOrder model.ActiveOrders
//		// 添加下面代码的错误处理
//		result = dao.Db.Model(&model.ActiveOrders{}).Where("`key_id` = ?", key.Id).First(&activeOrder)
//		if activeOrder.IsValid == false {
//			code = http.StatusForbidden // 403 密钥被管理员禁用且不可使用
//		}
//		// 上面如果有一个出现错误则不执行下面的
//		// 3 检查密钥的内容和客户端id
//		if request.KeyContent == key.Key && request.ClientId == key.ClientId {
//
//			code = http.StatusOK                                                                                                 // 200
//			if err := utils.SetActivationRecordCache2Redis(ctx, key.ClientId, key.Key, activeOrder.ExpirationDate); err != nil { //保存到缓存
//
//			}
//
//		}
//	}
//
//	// 200 匹配成功
//	// 206 密钥可用但是即将过期
//	// 409 密钥和客户端id不匹配
//	// 403 密钥被管理员禁用
//	// 404 密钥不存在
//	// 500 其他运行错误
//	// 503 密钥已过期
//
//	// 检查密钥是否过期
//	var msg string = ""
//	if expiredAt.Sub(time.Now()) <= time.Hour*24*7 {
//		// 密钥即将过期 返回code 为206
//		code = http.StatusPartialContent // 206
//		msg = "key is comming to expire"
//	} else if expiredAt.Sub(time.Now()) < 0 {
//		// 密钥即将过期 返回code 为207
//		code = http.StatusServiceUnavailable // 503 密钥过期
//		msg = "key is expired"
//	}
//
//	return &pb.CheckActivationStatusIntervalReqResponse{
//		Code: code,
//		Msg:  msg,
//	}, nil
//}

func (s *KeyServices) CheckActivationStatusIntervalReq(ctx context.Context, request *pb.CheckActivationStatusIntervalReqRequest) (*pb.CheckActivationStatusIntervalReqResponse, error) {
	code, expiredAt := utils.CheckActivationFromCache(ctx, request.ClientId, request.KeyContent) // 从缓存中检查激活状态
	var msg string = "success"
	if code != http.StatusOK {
		// 1. 查询密钥是否存在
		var key model.Keys
		result := dao.Db.Model(&model.Keys{}).Where("`key` = ?", request.KeyContent).First(&key)
		if result.RowsAffected == 0 {
			code = http.StatusNotFound // 404
			msg = "cannot find key of " + request.KeyContent
		} else if result.Error != nil {
			code = http.StatusInternalServerError // 500
			msg = "err on query db: " + result.Error.Error()
		} else {
			code = http.StatusOK
		}
		// 2. 查询有效订单表，检查该密钥是否被禁用
		if code == http.StatusOK { // 如果没有出错，继续检查是否禁用
			var activeOrder model.ActiveOrders
			result = dao.Db.Model(&model.ActiveOrders{}).Where("`key_id` = ?", key.Id).First(&activeOrder)

			if result.Error != nil {
				code = http.StatusInternalServerError // 500 查询数据库出错
				msg = "err on query db: " + result.Error.Error()
			} else if activeOrder.IsValid == false {
				code = http.StatusForbidden // 403 密钥被管理员禁用
				msg = "the key has been blocked by admin, please contact us via ticket"
			}
			// 3. 检查密钥和客户端 ID 是否匹配
			if code == http.StatusOK && request.KeyContent == key.Key && request.ClientId == key.ClientId {
				code = http.StatusOK // 200 密钥有效，且匹配客户端
				// 如果成功，则将激活记录保存到缓存
				if err := utils.SetActivationRecordCache2Redis(ctx, key.ClientId, key.Key, activeOrder.ExpirationDate); err != nil {
					msg = "cache in redis failure, its not effected: " + err.Error() // 不影响返回状态
				}
			} else if request.KeyContent != key.Key || request.ClientId != key.ClientId {
				code = http.StatusConflict // 409 密钥和客户端 ID 不匹配
				msg = "the client_id and key_content is not match, please try again later"
			}
		}
	}
	// 检查密钥是否接近过期或已过期
	if expiredAt != nil {
		if expiredAt.Sub(time.Now()) <= time.Hour*24*7 {
			// 密钥即将过期，返回 code 206
			code = http.StatusPartialContent // 206
			// TODO
			msg = "key is coming to expire, please"
		} else if expiredAt.Sub(time.Now()) < 0 {
			// 密钥已经过期，返回 code 503
			code = http.StatusServiceUnavailable // 503 密钥过期
			msg = "key is expired"
		}
	}

	return &pb.CheckActivationStatusIntervalReqResponse{
		Code: code,
		Msg:  msg,
	}, nil
}
