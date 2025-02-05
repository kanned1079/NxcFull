package services

import (
	"encoding/json"
	"golang.org/x/net/context"
	pb "logsServices/api/proto"
	"logsServices/internal/dao"
	"logsServices/internal/model"
	"net/http"
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

	var status200count, status404count, status500count, loginReq, regReq int64
	dao.Db.Model(&model.ApiLog{}).Where("status_code = ?", http.StatusOK).Count(&status200count)
	dao.Db.Model(&model.ApiLog{}).Where("status_code = ?", http.StatusNotFound).Count(&status404count)
	dao.Db.Model(&model.ApiLog{}).Where("status_code = ?", http.StatusInternalServerError).Count(&status500count)

	dao.Db.Model(&model.ApiLog{}).Where("path = ? or ?", "/api/admin/v1/login", "/api/user/v1/login").Count(&loginReq)
	dao.Db.Model(&model.ApiLog{}).Where("path = ? or ?", "/api/user/v1/register/register").Count(&regReq)

	var apiList []model.ApiLog
	dao.Db.Model(&model.ApiLog{}).Order("`request_at` DESC").Limit(1000).Find(&apiList)

	apiListJson, err := json.Marshal(apiList)
	if err != nil {
		return &pb.GetServerLiveStatusResponse{
			Code: http.StatusInternalServerError,
		}, err
	}

	return &pb.GetServerLiveStatusResponse{
		Code:        http.StatusOK,
		Status200:   status200count,
		Status404:   status404count,
		Status500:   status500count,
		LoginReq:    loginReq,
		RegisterReq: regReq,
		ApiLogList:  apiListJson,
	}, nil
}
