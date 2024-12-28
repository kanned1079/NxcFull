package services

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"net/http"
	pb "userServices/api/proto"
	"userServices/internal/dao"
	"userServices/internal/model"
	"userServices/internal/utils"
)

// GetUserIdFromEmail 根据用户邮箱获取用户id，如果查询出错或找不到返回-1
func GetUserIdFromEmail(email string) int64 {
	var userId int64
	result := dao.Db.Model(&model.User{}).Where("email = ?", email).Select("id").Scan(&userId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return -1
		}
		log.Println("Error querying user:", result.Error)
		return -1
	}

	return userId
}

func (s *UserService) CheckUserAuthInternal(ctx context.Context, request *pb.CheckUserAuthInternalRequest) (*pb.CheckUserAuthInternalResponse, error) {
	authResultCode := utils.AuthUserInfo(request.Email, request.Password)
	userId := GetUserIdFromEmail(request.Email)
	if userId == -1 {
		return &pb.CheckUserAuthInternalResponse{
			Code: http.StatusInternalServerError,
			Msg:  "unknown error, please try again",
		}, nil
	}
	switch authResultCode {
	case model.Auth_Pass:
		return &pb.CheckUserAuthInternalResponse{
			Code:   http.StatusOK,
			Msg:    "success",
			UserId: userId,
		}, nil
	case model.Auth_Failure:
		return &pb.CheckUserAuthInternalResponse{
			Code:   http.StatusConflict,
			Msg:    "email or password is incorrect",
			UserId: userId,
		}, nil
	case model.User_Not_Exist:
		return &pb.CheckUserAuthInternalResponse{
			Code:   http.StatusNotFound,
			Msg:    "account may not exist, please check your email",
			UserId: userId,
		}, nil
	default:
		return &pb.CheckUserAuthInternalResponse{
			Code:   http.StatusInternalServerError,
			Msg:    "unknown error, please try again",
			UserId: userId,
		}, nil
	}
}
