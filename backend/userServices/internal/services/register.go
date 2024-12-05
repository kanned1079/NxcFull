package services

import (
	"context"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

//func (s *UserService) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
//	var inviteCode = req.InviteUserId // 这个是另一个用户的邀请码
//	// 获取该邀请码对应的用户id
//	var inviteUserId int64
//	// select id from nxc_networks.x_users where invite_code = "<inviteCode>"
//
//	var newUserInfo model.User = model.User{
//		Email: req.Email,
//		//InviteUserID: req.InviteUserId,
//		InviteUserID: strconv.FormatInt(inviteUserId, 10), // 设置邀请用户的id
//	}
//	var newUserAuth model.Auth = model.Auth{
//		Email:    req.Email,
//		Password: req.Password,
//	}
//
//	// 开启事务
//	err := dao.Db.Transaction(func(tx *gorm.DB) error {
//		// 插入用户数据
//		if result := tx.Create(&newUserInfo); result.Error != nil {
//			return result.Error // 回滚事务
//		}
//
//		// 插入认证数据
//		if result := tx.Create(&newUserAuth); result.Error != nil {
//			return result.Error // 回滚事务
//		}
//
//		return nil // 提交事务
//	})
//
//	// 检查事务执行结果
//	if err != nil {
//		return &pb.UserRegisterResponse{
//			Code:         http.StatusInternalServerError,
//			Msg:          "注册失败，事务回滚",
//			IsRegistered: false,
//		}, nil
//	}
//
//	return &pb.UserRegisterResponse{
//		Code:         http.StatusOK,
//		Msg:          "注册成功",
//		IsRegistered: true,
//	}, nil
//}

func (s *UserService) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	// 获取邀请用户的邀请码
	var inviteCode = req.InviteUserId // 这是另一个用户的邀请码

	// 根据邀请码查询对应的用户 ID
	var inviteUserId int64
	err := dao.Db.Model(model.User{}). // 指定查询的表
						Where("invite_code = ?", inviteCode). // 条件：邀请码匹配
						Select("id").                         // 选择需要的字段
						Scan(&inviteUserId).Error             // 执行查询并绑定结果到 inviteUserId

	if err != nil || inviteUserId == 0 {
		// 如果查询失败或者没有找到记录，将 inviteUserId 设置为空字符串
		inviteUserId = 0
	}

	// 初始化用户信息模型
	var newUserInfo model.User = model.User{
		Email:        req.Email,                           // 设置用户的邮箱
		InviteUserID: strconv.FormatInt(inviteUserId, 10), // 如果 inviteUserId 为 0，则转换为空字符串
		Status:       true,
	}

	// 初始化认证信息模型
	var newUserAuth model.Auth = model.Auth{
		Email:    req.Email,    // 设置认证邮箱
		Password: req.Password, // 设置认证密码
	}

	// 开启事务，确保用户注册和认证数据的一致性
	err = dao.Db.Transaction(func(tx *gorm.DB) error {
		// 插入用户信息到数据库
		if result := tx.Create(&newUserInfo); result.Error != nil {
			return result.Error // 如果插入失败，则回滚事务
		}

		// 插入认证信息到数据库
		if result := tx.Create(&newUserAuth); result.Error != nil {
			return result.Error // 如果插入失败，则回滚事务
		}

		// 如果所有插入操作成功，则提交事务
		return nil
	})

	// 检查事务执行结果
	if err != nil {
		// 如果事务失败，返回错误信息
		return &pb.UserRegisterResponse{
			Code:         http.StatusInternalServerError, // HTTP 状态码：500
			Msg:          "注册失败，事务回滚",                    // 错误提示信息
			IsRegistered: false,                          // 标记注册状态为失败
		}, nil
	}

	// 如果事务成功，返回成功响应
	return &pb.UserRegisterResponse{
		Code:         http.StatusOK, // HTTP 状态码：200
		Msg:          "注册成功",        // 成功提示信息
		IsRegistered: true,          // 标记注册状态为成功
	}, nil
}
