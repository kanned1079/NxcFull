package services

import (
	pb "NxcFull/backend/userServices/api/proto"
	"NxcFull/backend/userServices/internal/dao"
	"NxcFull/backend/userServices/internal/model"
	"context"
	"net/http"
)

func (s *UserService) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {

	var newUserInfo model.User = model.User{
		Email:        req.Email,
		InviteUserID: req.InviteUserId,
		//LicenseExpiration: nil,
	}
	var newUserAuth model.Auth = model.Auth{
		Email:    req.Email,
		Password: req.Password,
	}
	result1 := dao.Db.Model(&model.User{}).Create(&newUserInfo)
	if result1.Error != nil {
		//context.JSON(http.StatusInternalServerError, gin.H{
		//	"code":  http.StatusInternalServerError,
		//	"msg":   "插入数据错误",
		//	"error": result1.Error.Error(),
		//})
		return &pb.UserRegisterResponse{
			Code:         http.StatusInternalServerError,
			Msg:          "插入数据错误",
			IsRegistered: false,
		}, nil
	}
	result2 := dao.Db.Model(&model.Auth{}).Create(&newUserAuth)
	if result2.Error != nil {
		//context.JSON(http.StatusInternalServerError, gin.H{
		//	"code":  http.StatusInternalServerError,
		//	"msg":   "插入验证数据错误",
		//	"error": result2.Error.Error(),
		//})
		return &pb.UserRegisterResponse{
			Code:         http.StatusInternalServerError,
			Msg:          "插入验证数据错误",
			IsRegistered: false,
		}, nil
	}
	if result1.Error == nil && result2.Error == nil {
		//context.JSON(http.StatusOK, gin.H{
		//	"code": http.StatusOK,
		//	"msg":  "新建用户成功",
		//})
		return &pb.UserRegisterResponse{
			Code:         http.StatusOK,
			Msg:          "插入数据成功",
			IsRegistered: true,
		}, nil
	} else {
		return &pb.UserRegisterResponse{
			Code:         http.StatusInternalServerError,
			Msg:          "未知错误",
			IsRegistered: false,
		}, nil
	}
}
