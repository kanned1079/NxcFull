package services

import (
	"context"
	"log"
	pb "logsServices/api/proto"
	"logsServices/internal/dao"
	"logsServices/internal/model"
	"net/http"
	"time"
)

func (s *LogService) SaveApiAccessLog2Db(ctx context.Context, request *pb.SaveApiAccessLog2DbRequest) (*pb.SaveApiAccessLog2DbResponse, error) {
	// 日志存储容器
	var apiLogs []model.ApiLog

	// 遍历请求日志数组并转换为数据库模型
	for _, logEntry := range request.Logs {
		apiLog := model.ApiLog{
			Level:        logEntry.Level,
			StatusCode:   logEntry.StatusCode,
			RequestAt:    time.Unix(logEntry.RequestAt, 0).UTC(),
			ResponseTime: logEntry.ResponseTime,
			ClientIp:     logEntry.ClientIp,
			Method:       logEntry.Method,
			Path:         logEntry.Path,
			CreatedAt:    time.Now(), // 设置当前时间为创建时间
			UpdatedAt:    time.Now(), // 设置当前时间为更新时间
		}
		apiLogs = append(apiLogs, apiLog)
	}

	// 将日志批量写入数据库
	err := dao.Db.Model(&model.ApiLog{}).Create(&apiLogs).Error
	if err != nil {
		log.Printf("Failed to save logs to database: %v", err)
		return &pb.SaveApiAccessLog2DbResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to save logs to database",
		}, err
	}

	log.Printf("Successfully saved %d logs to the database.", len(apiLogs))

	// 返回成功响应
	return &pb.SaveApiAccessLog2DbResponse{
		Code: http.StatusOK,
		Msg:  "Logs saved successfully",
	}, nil
}
