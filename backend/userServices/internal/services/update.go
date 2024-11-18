package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
	pb "userServices/api/proto"
	"userServices/internal/dao"
	"userServices/internal/model"
	"userServices/internal/utils"
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

	if code := utils.AuthUserInfo(req.Email, req.OldPassword); code != model.Auth_Pass {
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

	// 将用户凭据序列化为 JSON 字符串存入 Redis
	userJSON, jsonErr := json.Marshal(auth)
	if jsonErr != nil {
		log.Println("JSON 序列化用户凭据失败: ", jsonErr)
	} else {
		redisErr := dao.Rdb.Set(cxt, "auth:"+auth.Email, userJSON, 24*time.Hour).Err()
		if redisErr != nil {
			log.Println("Redis 存储用户凭据失败: ", redisErr)
		}
	}

	return &pb.ApplyNewPasswordResponse{
		Code:    http.StatusOK,
		Updated: true,
		Msg:     "更新密码成功",
	}, nil
}

func (s *UserService) DeleteMyAccount(ctx context.Context, request *pb.DeleteMyAccountRequest) (*pb.DeleteMyAccountResponse, error) {
	tx := dao.Db.Begin() // 开启事务
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v", tx.Error)
		return nil, fmt.Errorf("failed to start transaction")
	}

	// 查找用户
	var user model.User
	if err := tx.Model(&model.User{}).Where("id = ?", request.UserId).First(&user).Error; err != nil {
		tx.Rollback() // 事务回滚
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("User not found for ID: %d", request.UserId)
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("Error querying user: %v", err)
		return nil, fmt.Errorf("failed to query user")
	}

	// 删除用户认证信息
	if err := tx.Model(&model.Auth{}).Where("email = ?", user.Email).Delete(&model.Auth{}).Error; err != nil {
		tx.Rollback() // 事务回滚
		log.Printf("Error deleting auth info for email %s: %v", user.Email, err)
		return nil, fmt.Errorf("failed to delete auth info")
	}

	// 删除用户信息
	if err := tx.Model(&model.User{}).Where("id = ?", request.UserId).Delete(&model.User{}).Error; err != nil {
		tx.Rollback() // 事务回滚
		log.Printf("Error deleting user with ID %d: %v", request.UserId, err)
		return nil, fmt.Errorf("failed to delete user")
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, fmt.Errorf("failed to commit transaction")
	}

	// 返回成功响应
	return &pb.DeleteMyAccountResponse{
		Code:    http.StatusOK,
		Deleted: true,
		Msg:     "Deleted successfully",
	}, nil
}
