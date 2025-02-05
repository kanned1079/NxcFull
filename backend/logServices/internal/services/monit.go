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

/**
message GetServerLiveStatusResponse {
  int32 code = 1;
  int64 status200 = 2;
  int64 status404 = 3;
  int64 status500 = 4;
  int64 login_req = 5;
  int64 register_req = 6;
  bytes api_log_list = 7;
  int64 page_size = 8;
}
*/

func (s *LogService) GetServerLiveStatus(ctx context.Context, request *pb.GetServerLiveStatusRequest) (*pb.GetServerLiveStatusResponse, error) {
	// 获取分页参数
	page := request.Page
	size := request.Size

	// 计算分页的偏移量和每页大小
	offset := (page - 1) * size
	limit := size

	var status200count, status404count, status500count, loginReq, regReq int64

	// 查询不同状态的日志数量
	if err := dao.Db.Model(&model.ApiLog{}).Where("status_code = ?", http.StatusOK).Count(&status200count).Error; err != nil {
		return &pb.GetServerLiveStatusResponse{
			Code: http.StatusInternalServerError,
		}, fmt.Errorf("failed to count status 200 logs: %v", err)
	}

	if err := dao.Db.Model(&model.ApiLog{}).Where("status_code = ?", http.StatusNotFound).Count(&status404count).Error; err != nil {
		return &pb.GetServerLiveStatusResponse{
			Code: http.StatusInternalServerError,
		}, fmt.Errorf("failed to count status 404 logs: %v", err)
	}

	if err := dao.Db.Model(&model.ApiLog{}).Where("status_code = ?", http.StatusInternalServerError).Count(&status500count).Error; err != nil {
		return &pb.GetServerLiveStatusResponse{
			Code: http.StatusInternalServerError,
		}, fmt.Errorf("failed to count status 500 logs: %v", err)
	}

	// 查询登录和注册请求数量
	if err := dao.Db.Model(&model.ApiLog{}).Where("path = ? or ?", "/api/admin/v1/login", "/api/user/v1/login").Count(&loginReq).Error; err != nil {
		return &pb.GetServerLiveStatusResponse{
			Code: http.StatusInternalServerError,
		}, fmt.Errorf("failed to count login requests: %v", err)
	}

	if err := dao.Db.Model(&model.ApiLog{}).Where("path = ?", "/api/user/v1/register/register").Count(&regReq).Error; err != nil {
		return &pb.GetServerLiveStatusResponse{
			Code: http.StatusInternalServerError,
		}, fmt.Errorf("failed to count register requests: %v", err)
	}

	// 获取符合条件的总记录数，用于分页计算
	var totalCount int64
	if err := dao.Db.Model(&model.ApiLog{}).Count(&totalCount).Error; err != nil {
		return &pb.GetServerLiveStatusResponse{
			Code: http.StatusInternalServerError,
		}, fmt.Errorf("failed to count total logs: %v", err)
	}

	// 计算总页数（PageSize）
	pageSize := (totalCount + int64(size) - 1) / int64(size) // 这里用向上取整公式计算总页数

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
		Code:        http.StatusOK,
		Status200:   status200count,
		Status404:   status404count,
		Status500:   status500count,
		LoginReq:    loginReq,
		RegisterReq: regReq,
		PageSize:    pageSize, // 返回计算出来的总页数
		ApiLogList:  apiListJson,
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
