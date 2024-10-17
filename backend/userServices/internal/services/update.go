package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	pb "userServices/api/proto"
	"userServices/internal/dao"
	"userServices/internal/model"
	"userServices/internal/utils"
)

//func (s *UserService) CheckPreviousPassword(ctx context.Context, req *pb.CheckPreviousPasswordRequest) (*pb.CheckPreviousPasswordResponse, error) {
//	// 查找用户
//	var auth model.Auth
//	if err := dao.Db.Model(&model.Auth{}).Where("email = ?", req.Email).First(&auth).Error; err != nil {
//		// 如果用户不存在
//		return &pb.CheckPreviousPasswordResponse{
//			Code:     http.StatusNotFound,
//			Verified: false,
//			Msg:      "用户不存在",
//		}, nil
//	}
//
//	// 直接比较哈希过的旧密码
//	if req.OldPassword != auth.Password {
//		// 如果哈希后的密码不匹配
//		log.Println("密码不匹配")
//		return &pb.CheckPreviousPasswordResponse{
//			Code:     http.StatusForbidden,
//			Verified: false,
//			Msg:      "旧密码不匹配",
//		}, nil
//	}
//
//	return &pb.CheckPreviousPasswordResponse{
//		Code:     http.StatusOK,
//		Verified: true,
//		Msg:      "旧密码验证通过",
//	}, nil
//}

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
	//var decodedPwd string
	//var err error
	//if decodedPwd, err = utils.DecodeBase64(req.OldPassword); err != nil {
	//	log.Println(err)
	//	return &pb.CheckPreviousPasswordResponse{
	//		Code:     http.StatusForbidden,
	//		Verified: false,
	//		Msg:      "解码失败",
	//	}, nil
	//}
	//log.Println("previous password:", decodedPwd)

	if code := utils.AuthUserInfo(req.Email, req.OldPassword); code != model.Auth_Pass {
		return &pb.CheckPreviousPasswordResponse{
			Code:     http.StatusForbidden,
			Verified: false,
			Msg:      "旧密码不匹配",
		}, nil
	}

	//if utils.ComparePasswordHash(decodedPwd, userAuth.Password) {
	//	code = model.Auth_Pass
	//} else {
	//	code = model.Auth_Failure
	//}

	//if req.OldPassword != auth.Password {
	//	// 如果哈希后的密码不匹配
	//	log.Println("密码不匹配")
	//	return &pb.CheckPreviousPasswordResponse{
	//		Code:     http.StatusForbidden,
	//		Verified: false,
	//		Msg:      "旧密码不匹配",
	//	}, nil
	//}

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
		//ctx, cancel := context.Context(context.Background(), 5*time.Second)
		//defer cancel()
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
