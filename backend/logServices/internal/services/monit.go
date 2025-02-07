package services

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	pb "logsServices/api/proto"
	"logsServices/internal/dao"
	"logsServices/internal/model"
	"net/http"
	"time"
)

//func (s *LogService) GetServerLiveStatus(ctx context.Context, request *pb.GetServerLiveStatusRequest) (*pb.GetServerLiveStatusResponse, error) {
//	// 获取分页参数
//	page := request.Page
//	size := request.Size
//
//	// 计算分页的偏移量和每页大小
//	offset := (page - 1) * size
//	limit := size
//
//	var status200count, status404count, status500count, loginReq, regReq int64
//
//	// 查询一周内不同状态的日志数量
//	oneWeekAgo := time.Now().AddDate(0, 0, -7).Format("2006-01-02 15:04:05") // 获取一周前的时间
//
//	if err := dao.Db.Model(&model.ApiLog{}).Where("status_code = ? AND created_at >= ?", http.StatusOK, oneWeekAgo).Count(&status200count).Error; err != nil {
//		return &pb.GetServerLiveStatusResponse{
//			Code: http.StatusInternalServerError,
//		}, fmt.Errorf("failed to count status 200 logs: %v", err)
//	}
//
//	if err := dao.Db.Model(&model.ApiLog{}).Where("status_code = ? AND created_at >= ?", http.StatusNotFound, oneWeekAgo).Count(&status404count).Error; err != nil {
//		return &pb.GetServerLiveStatusResponse{
//			Code: http.StatusInternalServerError,
//		}, fmt.Errorf("failed to count status 404 logs: %v", err)
//	}
//
//	if err := dao.Db.Model(&model.ApiLog{}).Where("status_code = ? AND created_at >= ?", http.StatusInternalServerError, oneWeekAgo).Count(&status500count).Error; err != nil {
//		return &pb.GetServerLiveStatusResponse{
//			Code: http.StatusInternalServerError,
//		}, fmt.Errorf("failed to count status 500 logs: %v", err)
//	}
//
//	// 查询一周内的登录和注册请求数量
//	if err := dao.Db.Model(&model.ApiLog{}).Where("path IN (?, ?) AND created_at >= ?", "/api/admin/v1/login", "/api/user/v1/login", oneWeekAgo).Count(&loginReq).Error; err != nil {
//		return &pb.GetServerLiveStatusResponse{
//			Code: http.StatusInternalServerError,
//		}, fmt.Errorf("failed to count login requests: %v", err)
//	}
//
//	if err := dao.Db.Model(&model.ApiLog{}).Where("path = ? AND created_at >= ?", "/api/user/v1/register/register", oneWeekAgo).Count(&regReq).Error; err != nil {
//		return &pb.GetServerLiveStatusResponse{
//			Code: http.StatusInternalServerError,
//		}, fmt.Errorf("failed to count register requests: %v", err)
//	}
//
//	// 获取日志表大小（单位MB）
//	var tableSize float64
//	if err := dao.Db.Raw(`
//		SELECT ROUND(data_length / 1024 / 1024, 2) AS size_mb
//		FROM information_schema.tables
//		WHERE table_name = ?`, model.ApiLog{}.TableName()).Scan(&tableSize).Error; err != nil {
//		return &pb.GetServerLiveStatusResponse{
//			Code: http.StatusInternalServerError,
//		}, fmt.Errorf("failed to fetch table size: %v", err)
//	}
//
//	// 获取日志表的总记录数
//	var logTableRowsCount int64
//	if err := dao.Db.Model(&model.ApiLog{}).Count(&logTableRowsCount).Error; err != nil {
//		return &pb.GetServerLiveStatusResponse{
//			Code: http.StatusInternalServerError,
//		}, fmt.Errorf("failed to count total logs: %v", err)
//	}
//
//	// 获取符合条件的总记录数，用于分页计算
//	var totalCount int64
//	if err := dao.Db.Model(&model.ApiLog{}).Count(&totalCount).Error; err != nil {
//		return &pb.GetServerLiveStatusResponse{
//			Code: http.StatusInternalServerError,
//		}, fmt.Errorf("failed to count total logs: %v", err)
//	}
//
//	// 计算总页数（PageSize）
//	pageSize := (totalCount + int64(size) - 1) / int64(size) // 这里用向上取整公式计算总页数
//
//	// 分页查询日志数据
//	var apiList []model.ApiLog
//	if err := dao.Db.Model(&model.ApiLog{}).
//		Order("`request_at` DESC").
//		Offset(int(offset)).
//		Limit(int(limit)).
//		Find(&apiList).Error; err != nil {
//		return &pb.GetServerLiveStatusResponse{
//			Code: http.StatusInternalServerError,
//		}, fmt.Errorf("failed to fetch API logs: %v", err)
//	}
//
//	// 序列化日志数据
//	apiListJson, err := json.Marshal(apiList)
//	if err != nil {
//		return &pb.GetServerLiveStatusResponse{
//			Code: http.StatusInternalServerError,
//		}, fmt.Errorf("failed to marshal API logs: %v", err)
//	}
//
//	// 返回响应，包括总页数信息
//	return &pb.GetServerLiveStatusResponse{
//		Code:              http.StatusOK,
//		Status200:         status200count,
//		Status404:         status404count,
//		Status500:         status500count,
//		LoginReq:          loginReq,
//		RegisterReq:       regReq,
//		PageSize:          pageSize, // 返回计算出来的总页数
//		ApiLogList:        apiListJson,
//		TableSize:         float32(tableSize), // 返回查询到的日志表大小（单位MB）
//		LogTableRowsCount: logTableRowsCount,  // 返回查询到的日志表总记录数
//	}, nil
//}

func (s *LogService) GetServerLiveStatus(ctx context.Context, request *pb.GetServerLiveStatusRequest) (*pb.GetServerLiveStatusResponse, error) {
	// 获取分页参数
	page := request.Page
	size := request.Size

	// 计算分页的偏移量和每页大小
	offset := (page - 1) * size
	limit := size

	// 获取一周前的时间
	oneWeekAgo := time.Now().AddDate(0, 0, -7).Format("2006-01-02 15:04:05")

	// 定义计数变量
	var status200count, status404count, status500count, loginReq, regReq, logTableRowsCount, totalCount int64
	var tableSize float64

	// 定义查询条件和对应的计数变量
	queries := []struct {
		whereClause string
		args        []interface{}
		countPtr    *int64
	}{
		{"status_code = ? AND created_at >= ?", []interface{}{http.StatusOK, oneWeekAgo}, &status200count},
		{"status_code = ? AND created_at >= ?", []interface{}{http.StatusNotFound, oneWeekAgo}, &status404count},
		{"status_code = ? AND created_at >= ?", []interface{}{http.StatusInternalServerError, oneWeekAgo}, &status500count},
		{"path IN (?, ?) AND created_at >= ?", []interface{}{"/api/admin/v1/login", "/api/user/v1/login", oneWeekAgo}, &loginReq},
		{"path = ? AND created_at >= ?", []interface{}{"/api/user/v1/register/register", oneWeekAgo}, &regReq},
		{"1 = 1", nil, &totalCount},
	}

	// 批量执行查询计数
	for _, query := range queries {
		if err := dao.Db.Model(&model.ApiLog{}).Where(query.whereClause, query.args...).Count(query.countPtr).Error; err != nil {
			return &pb.GetServerLiveStatusResponse{
				Code: http.StatusInternalServerError,
			}, fmt.Errorf("failed to count logs: %v", err)
		}
	}

	// 获取日志表大小（单位MB）
	if err := dao.Db.Raw(`
       SELECT ROUND(data_length / 1024 / 1024, 4) AS size_mb
       FROM information_schema.tables
       WHERE table_name = ?`, model.ApiLog{}.TableName()).Scan(&tableSize).Error; err != nil {
		return &pb.GetServerLiveStatusResponse{
			Code: http.StatusInternalServerError,
		}, fmt.Errorf("failed to fetch table size: %v", err)
	}

	// 获取日志表的总记录数
	if err := dao.Db.Model(&model.ApiLog{}).Count(&logTableRowsCount).Error; err != nil {
		return &pb.GetServerLiveStatusResponse{
			Code: http.StatusInternalServerError,
		}, fmt.Errorf("failed to count total logs: %v", err)
	}

	// 计算总页数（PageSize）
	pageSize := (totalCount + int64(size) - 1) / int64(size)

	// 分页查询日志数据
	var apiList []model.ApiLog
	if err := dao.Db.Model(&model.ApiLog{}).
		Order("`request_at` DESC").
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&apiList).Error; err != nil {
		return &pb.GetServerLiveStatusResponse{
			Code: http.StatusInternalServerError,
		}, fmt.Errorf("failed to fetch API logs: %v", err)
	}

	// 序列化日志数据
	apiListJson, err := json.Marshal(apiList)
	if err != nil {
		return &pb.GetServerLiveStatusResponse{
			Code: http.StatusInternalServerError,
		}, fmt.Errorf("failed to marshal API logs: %v", err)
	}

	// 返回响应，包括总页数信息
	return &pb.GetServerLiveStatusResponse{
		Code:              http.StatusOK,
		Status200:         status200count,
		Status404:         status404count,
		Status500:         status500count,
		LoginReq:          loginReq,
		RegisterReq:       regReq,
		PageSize:          pageSize,
		ApiLogList:        apiListJson,
		TableSize:         float32(tableSize),
		LogTableRowsCount: logTableRowsCount,
	}, nil
}

func (s *LogService) ClearPreviousApiLog(ctx context.Context, request *pb.ClearPreviousApiLogRequest) (*pb.ClearPreviousApiLogResponse, error) {
	// 获取当前时间
	now := time.Now()

	// 计算两天前的时间
	oneWeekAgo := now.AddDate(0, 0, -7)

	// 直接删除一周前的日志
	result := dao.Db.Model(&model.ApiLog{}).
		Where("created_at < ?", oneWeekAgo).
		Unscoped(). // 禁用软删除
		Delete(&model.ApiLog{})

	// 检查错误
	if result.Error != nil {
		// 错误处理
		return &pb.ClearPreviousApiLogResponse{
			Code: http.StatusInternalServerError,
			Msg:  "failure to clear previous api log: " + result.Error.Error(),
		}, nil
	}

	// 返回成功，RowDeleted 为删除的记录数
	return &pb.ClearPreviousApiLogResponse{
		Code:       http.StatusOK,
		RowDeleted: result.RowsAffected, // 返回删除的行数
		Msg:        "success",
	}, nil
}
