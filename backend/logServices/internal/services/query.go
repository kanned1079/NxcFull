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

// SaveApiAccessLog2Db 用于保存Gin的用户访问日志 不含其他错误日志
func (s *LogService) SaveApiAccessLog2Db(ctx context.Context, request *pb.SaveApiAccessLog2DbRequest) (*pb.SaveApiAccessLog2DbResponse, error) {
	var apiLogs []model.ApiLog // 日志存储容器
	// 遍历请求日志数组并转换为数据库模型
	for _, logEntry := range request.Logs {
		apiLog := model.ApiLog{
			Level:        logEntry.Level,                         // 日志等级
			StatusCode:   logEntry.StatusCode,                    // 回应状态码
			RequestAt:    time.Unix(logEntry.RequestAt, 0).UTC(), // FIX 注意这里是秒级别时间戳
			ResponseTime: logEntry.ResponseTime,                  // 相应时间
			ClientIp:     logEntry.ClientIp,                      // 远程客户端IP
			Method:       logEntry.Method,                        // 请求方式 GET POST等
			Path:         logEntry.Path,                          // 请求路径
			CreatedAt:    time.Now(),                             // 设置当前时间为创建时间
			UpdatedAt:    time.Now(),                             // 设置当前时间为更新时间
		}
		apiLogs = append(apiLogs, apiLog) // 存储到日志容器中
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
