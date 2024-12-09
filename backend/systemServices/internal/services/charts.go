package services

import (
	"context"
	"log"
	"net/http"
	pb "systemServices/api/proto"
	"systemServices/internal/utils"
)

func (s *SettingServices) GetAdminDashboardData(ctx context.Context, request *pb.GetAdminDashboardDataRequest) (*pb.GetAdminDashboardDataResponse, error) {
	apiAccessCount, err := utils.GetLast7DaysAccessApiCounts()
	if err != nil {
		log.Println("get api access count err:", err)
	}
	incomeCount, yesterday, month, err := utils.GetLast7DaysAndMonthIncome()
	if err != nil {
		log.Println("get income count err:", err)
	}

	// TODO 待实现用户占比操作

	return &pb.GetAdminDashboardDataResponse{
		Code:                  http.StatusOK,
		Msg:                   "success",
		ApiAccessCountHistory: apiAccessCount,
		IncomeCountHistory:    incomeCount,
		IncomeYesterday:       yesterday,
		IncomeThisMonth:       month,
	}, nil
}
