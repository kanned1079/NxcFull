package services

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"
	pb "userServices/api/proto"
	"userServices/internal/dao"
	"userServices/internal/model"
)

func (s *UserService) UploadUserAvatar(ctx context.Context, request *pb.UploadUserAvatarRequest) (*pb.UploadUserAvatarResponse, error) {
	var existingAvatar model.Avatar
	if err := dao.Db.Where("user_id = ?", request.UserId).First(&existingAvatar).Error; err == nil { // 查找是否已经存在相同的 UserId
		existingAvatar.FileName = request.FileName
		existingAvatar.FileData = request.FileData
		if err := dao.Db.Save(&existingAvatar).Error; err != nil { // 如果记录存在，更新现有记录
			return &pb.UploadUserAvatarResponse{
				Code: http.StatusInternalServerError,
				Msg:  err.Error(),
			}, nil
		}
		return &pb.UploadUserAvatarResponse{
			Code: http.StatusOK,
			Msg:  "Avatar updated successfully",
		}, nil
	}

	if result := dao.Db.Create(&model.Avatar{ // 如果记录不存在，创建新的头像记录
		UserId:   request.UserId,
		FileName: request.FileName,
		FileData: request.FileData,
	}); result.Error != nil {
		return &pb.UploadUserAvatarResponse{
			Code: http.StatusInternalServerError,
			Msg:  result.Error.Error(),
		}, nil
	}

	return &pb.UploadUserAvatarResponse{
		Code: http.StatusOK,
		Msg:  "Avatar uploaded successfully",
	}, nil
}

func (s *UserService) GetUserAvatar(ctx context.Context, request *pb.GetUserAvatarRequest) (*pb.GetUserAvatarResponse, error) {
	// 查询数据库获取用户头像
	var avatar model.Avatar
	if err := dao.Db.Where("user_id = ?", request.UserId).First(&avatar).Error; err != nil {
		// 如果没有找到头像
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.GetUserAvatarResponse{
				Code: http.StatusNotFound,
				Msg:  "User doesn't have avatar.",
			}, nil
		}
		// 如果出现其他错误
		return &pb.GetUserAvatarResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}, nil
	}

	// 返回头像数据
	return &pb.GetUserAvatarResponse{
		Code:     http.StatusOK,
		Msg:      "Success",
		FileName: avatar.FileName,
		FileData: avatar.FileData, // 头像的二进制数据
	}, nil
}
