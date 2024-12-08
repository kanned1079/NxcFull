package services

import (
	"context"
	"net/http"
	pb "systemServices/api/proto"
	"systemServices/internal/utils"
)

func (s *SettingServices) GetAdminDashboardData(ctx context.Context, request *pb.GetAdminDashboardDataRequest) (*pb.GetAdminDashboardDataResponse, error) {
	apiAccessCount, err := utils.GetLast7DaysAccessApiCounts()
	if err != nil {

	}

	return &pb.GetAdminDashboardDataResponse{
		Code:                  http.StatusOK,
		Msg:                   "success",
		ApiAccessCountHistory: apiAccessCount,
	}, nil
}
