package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	pb "userServices/api/proto"
	"userServices/internal/dao"
	"userServices/internal/model"
)

func (s *UserService) UpdateUserInfo(context context.Context, request *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
	var user model.User
	if result := dao.Db.Model(&model.User{}).Where("id = ?", request.UserId).First(&user); result.RowsAffected == 0 {
		log.Println("无该记录 id：", request.UserId)
		return &pb.UpdateUserInfoResponse{
			Code: http.StatusNotFound,
			Msg:  fmt.Sprintf("User not found. id: %d", request.UserId),
		}, nil
	}
	if userJson, err := json.Marshal(user); err != nil {
		return &pb.UpdateUserInfoResponse{
			Code: http.StatusNotFound,
			Msg:  fmt.Sprintf("User marshal fail. id: %d", request.UserId),
		}, nil
	} else {
		return &pb.UpdateUserInfoResponse{
			Code:     http.StatusOK,
			Msg:      "Query ok",
			UserInfo: userJson,
		}, nil
	}
}
