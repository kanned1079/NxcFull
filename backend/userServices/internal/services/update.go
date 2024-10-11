package services

import (
	"context"
	"log"
	"net/http"
	pb "userServices/api/proto"
	"userServices/internal/dao"
	"userServices/internal/model"
)

func (s *UserService) CheckPreviousPassword(ctx context.Context, req *pb.CheckPreviousPasswordRequest) (*pb.CheckPreviousPasswordResponse, error) {
	// 查找用户
	var auth model.Auth
	if err := dao.Db.Model(&model.Auth{}).Where("email = ?", req.Email).First(&auth).Error; err != nil {
		// 如果用户不存在
		return &pb.CheckPreviousPasswordResponse{
			Code:     http.StatusNotFound,
			Verified: false,
			Msg:      "用户不存在",
		}, nil
	}

	// 直接比较哈希过的旧密码
	if req.OldPassword != auth.Password {
		// 如果哈希后的密码不匹配
		log.Println("密码不匹配")
		return &pb.CheckPreviousPasswordResponse{
			Code:     http.StatusForbidden,
			Verified: false,
			Msg:      "旧密码不匹配",
		}, nil
	}

	return &pb.CheckPreviousPasswordResponse{
		Code:     http.StatusOK,
		Verified: true,
		Msg:      "旧密码验证通过",
	}, nil
}

func (s *UserService) ApplyNewPassword(cxt context.Context, req *pb.ApplyNewPasswordRequest) (*pb.ApplyNewPasswordResponse, error) {
	// 查找用户
	var auth model.Auth
	if err := dao.Db.Model(&model.Auth{}).Where("email = ?", req.Email).First(&auth).Error; err != nil {
		// 如果用户不存在
		return &pb.ApplyNewPasswordResponse{
			Code:    http.StatusNotFound,
			Updated: false,
			Msg:     "用户不存在",
		}, nil
	}

	// 更新密码，直接存储哈希过的新密码
	auth.Password = req.NewPassword
	if err := dao.Db.Save(&auth).Error; err != nil {
		return &pb.ApplyNewPasswordResponse{
			Code:    http.StatusInternalServerError,
			Updated: false,
			Msg:     "更新密码失败",
		}, nil
	}

	log.Println("修改密码成功")
	// 密码更新成功
	return &pb.ApplyNewPasswordResponse{
		Code:    http.StatusOK,
		Updated: true,
		Msg:     "更新密码成功",
	}, nil
}
