package services

import (
	"context"
	"gorm.io/gorm"
	"net/http"
	pb "userServices/api/proto"
	"userServices/internal/dao"
	"userServices/internal/model"
)

// v1 没有开启事务
//func (s *UserService) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
//	var newUserInfo model.User = model.User{
//		Email:        req.Email,
//		InviteUserID: req.InviteUserId,
//	}
//	var newUserAuth model.Auth = model.Auth{
//		Email:    req.Email,
//		Password: req.Password,
//	}
//	result1 := dao.Db.Model(&model.User{}).Create(&newUserInfo)
//	if result1.Error != nil {
//		return &pb.UserRegisterResponse{
//			Code:         http.StatusInternalServerError,
//			Msg:          "插入数据错误",
//			IsRegistered: false,
//		}, nil
//	}
//	result2 := dao.Db.Model(&model.Auth{}).Create(&newUserAuth)
//	if result2.Error != nil {
//		return &pb.UserRegisterResponse{
//			Code:         http.StatusInternalServerError,
//			Msg:          "插入验证数据错误",
//			IsRegistered: false,
//		}, nil
//	}
//	if result1.Error == nil && result2.Error == nil {
//		return &pb.UserRegisterResponse{
//			Code:         http.StatusOK,
//			Msg:          "插入数据成功",
//			IsRegistered: true,
//		}, nil
//	} else {
//		return &pb.UserRegisterResponse{
//			Code:         http.StatusInternalServerError,
//			Msg:          "未知错误",
//			IsRegistered: false,
//		}, nil
//	}
//}

func (s *UserService) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	var newUserInfo model.User = model.User{
		Email:        req.Email,
		InviteUserID: req.InviteUserId,
	}
	var newUserAuth model.Auth = model.Auth{
		Email:    req.Email,
		Password: req.Password,
	}

	// 开启事务
	err := dao.Db.Transaction(func(tx *gorm.DB) error {
		// 插入用户数据
		if result := tx.Create(&newUserInfo); result.Error != nil {
			return result.Error // 回滚事务
		}

		// 插入认证数据
		if result := tx.Create(&newUserAuth); result.Error != nil {
			return result.Error // 回滚事务
		}

		return nil // 提交事务
	})

	// 检查事务执行结果
	if err != nil {
		return &pb.UserRegisterResponse{
			Code:         http.StatusInternalServerError,
			Msg:          "注册失败，事务回滚",
			IsRegistered: false,
		}, nil
	}

	return &pb.UserRegisterResponse{
		Code:         http.StatusOK,
		Msg:          "注册成功",
		IsRegistered: true,
	}, nil
}
