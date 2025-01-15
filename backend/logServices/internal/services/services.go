package services

import pb "logsServices/api/proto"

type LogService struct {
	pb.UnimplementedLogServiceServer
}

func NewLogService() *LogService {
	return &LogService{}
}
