package services

import (
	"context"
	"encoding/base64"
	"errors"
	"gorm.io/gorm"
	"net/http"
	pb "userServices/api/proto"
	"userServices/internal/dao"
	"userServices/internal/model"
	"userServices/internal/utils"
)

//func (s *UserService) UploadUserAvatar(ctx context.Context, request *pb.UploadUserAvatarRequest) (*pb.UploadUserAvatarResponse, error) {
//	var existingAvatar model.Avatar
//	if err := dao.Db.Where("user_id = ?", request.UserId).First(&existingAvatar).Error; err == nil { // 查找是否已经存在相同的 UserId
//		existingAvatar.FileName = request.FileName
//		existingAvatar.FileData = request.FileData
//		if err := dao.Db.Save(&existingAvatar).Error; err != nil { // 如果记录存在，更新现有记录
//			return &pb.UploadUserAvatarResponse{
//				Code: http.StatusInternalServerError,
//				Msg:  err.Error(),
//			}, nil
//		}
//		return &pb.UploadUserAvatarResponse{
//			Code: http.StatusOK,
//			Msg:  "Avatar updated successfully",
//		}, nil
//	}
//
//	if result := dao.Db.Create(&model.Avatar{ // 如果记录不存在，创建新的头像记录
//		UserId:   request.UserId,
//		FileName: request.FileName,
//		FileData: request.FileData,
//	}); result.Error != nil {
//		return &pb.UploadUserAvatarResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  result.Error.Error(),
//		}, nil
//	}
//
//	return &pb.UploadUserAvatarResponse{
//		Code: http.StatusOK,
//		Msg:  "Avatar uploaded successfully",
//	}, nil
//}

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

		// 更新 Redis 缓存
		if err := utils.SetAvatarCache(ctx, request.UserId, existingAvatar.FileName, existingAvatar.FileData); err != nil {
			return &pb.UploadUserAvatarResponse{
				Code: http.StatusInternalServerError,
				Msg:  "Failed to update Redis cache: " + err.Error(),
			}, nil
		}

		return &pb.UploadUserAvatarResponse{
			Code: http.StatusOK,
			Msg:  "Avatar updated successfully (cache ok)",
		}, nil
	}

	// 如果记录不存在，创建新的头像记录
	if result := dao.Db.Create(&model.Avatar{
		UserId:   request.UserId,
		FileName: request.FileName,
		FileData: request.FileData,
	}); result.Error != nil {
		return &pb.UploadUserAvatarResponse{
			Code: http.StatusInternalServerError,
			Msg:  result.Error.Error(),
		}, nil
	}

	// 将头像数据缓存到 Redis
	if err := utils.SetAvatarCache(ctx, request.UserId, request.FileName, request.FileData); err != nil {
		return &pb.UploadUserAvatarResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to cache avatar in Redis: " + err.Error(),
		}, nil
	}

	return &pb.UploadUserAvatarResponse{
		Code: http.StatusOK,
		Msg:  "Avatar uploaded successfully (cache ok)",
	}, nil
}

//func (s *UserService) GetUserAvatar(ctx context.Context, request *pb.GetUserAvatarRequest) (*pb.GetUserAvatarResponse, error) {
//	// 查询数据库获取用户头像
//	var avatar model.Avatar
//	if err := dao.Db.Where("user_id = ?", request.UserId).First(&avatar).Error; err != nil {
//		// 如果没有找到头像
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return &pb.GetUserAvatarResponse{
//				Code: http.StatusNotFound,
//				Msg:  "User doesn't have avatar.",
//			}, nil
//		}
//		// 如果出现其他错误
//		return &pb.GetUserAvatarResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  err.Error(),
//		}, nil
//	}
//
//	// 返回头像数据
//	return &pb.GetUserAvatarResponse{
//		Code:     http.StatusOK,
//		Msg:      "Success",
//		FileName: avatar.FileName,
//		FileData: avatar.FileData, // 头像的二进制数据
//	}, nil
//}

func (s *UserService) GetUserAvatar(ctx context.Context, request *pb.GetUserAvatarRequest) (*pb.GetUserAvatarResponse, error) {
	// 检查 Redis 缓存中是否存在该用户的头像
	cachedAvatarData, fileName, err := utils.GetAvatarCache(ctx, request.UserId)
	if err != nil {
		// 如果发生错误，返回 500 错误
		return &pb.GetUserAvatarResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to read from Redis: " + err.Error(),
		}, nil
	}

	if cachedAvatarData != nil {
		// 如果缓存中有头像，直接返回
		decodedFileData, err := base64.StdEncoding.DecodeString(string(cachedAvatarData))
		if err != nil {
			return &pb.GetUserAvatarResponse{
				Code: http.StatusInternalServerError,
				Msg:  "Failed to decode cached file data",
			}, nil
		}

		return &pb.GetUserAvatarResponse{
			Code:     http.StatusOK,
			Msg:      "success (cache)",
			FileName: fileName,
			FileData: decodedFileData, // Redis 缓存中的头像数据
		}, nil
	}

	// 如果缓存没有，查询数据库获取用户头像
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

	// 将头像数据缓存到 Redis
	if err := utils.SetAvatarCache(ctx, request.UserId, avatar.FileName, avatar.FileData); err != nil {
		return &pb.GetUserAvatarResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to cache avatar in Redis: " + err.Error(),
		}, nil
	}

	// 返回头像数据
	return &pb.GetUserAvatarResponse{
		Code:     http.StatusOK,
		Msg:      "success (direct)",
		FileName: avatar.FileName,
		FileData: avatar.FileData, // 数据库中查询到的头像数据
	}, nil
}

func (s *UserService) AlterUserSecondaryName(ctx context.Context, request *pb.AlterUserSecondaryNameRequest) (*pb.AlterUserSecondaryNameResponse, error) {
	if result := dao.Db.Model(&model.User{}).Where("id = ?", request.UserId).Update("name", request.NewName); result.Error != nil {
		return &pb.AlterUserSecondaryNameResponse{
			Code: http.StatusInternalServerError,
			Msg:  result.Error.Error(),
		}, nil
	}
	return &pb.AlterUserSecondaryNameResponse{
		Code: http.StatusOK,
		Msg:  "success",
	}, nil
}
