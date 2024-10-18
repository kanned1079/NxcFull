package services

import (
	"context"
	"log"
	"net/http"
	"time"
	pb "userServices/api/proto"
	"userServices/internal/auth"
	"userServices/internal/dao"
	"userServices/internal/model"
	"userServices/internal/utils"
)

// Setup2FA 设置新的2FA 生成key 存入redis
func (s *UserService) Setup2FA(context context.Context, request *pb.Setup2FARequest) (*pb.Setup2FAResponse, error) {
	var secretKey string
	var url string
	var err error

	// 生成2FA密钥和URL
	if secretKey, url, err = auth.Generate2FASecret("Nxc Networks", request.Email); err != nil {
		return &pb.Setup2FAResponse{
			Code: http.StatusInternalServerError,
			Msg:  "生成2FA密钥失败" + err.Error(),
		}, nil
	}

	// 将生成的2FA数据存储到Redis中，等待后续验证
	redisKey := "2fa_new:" + request.Email
	redisValue := map[string]interface{}{
		"secret_key": secretKey,
		"url":        url,
		"user_id":    request.Id,
		"user_email": request.Email,
	}

	// 假设 redis 的客户端是 dao.Rdb
	if err := dao.Rdb.HMSet(context, redisKey, redisValue).Err(); err != nil {
		log.Println("存储2FA数据到Redis失败:", err)
		return &pb.Setup2FAResponse{
			Code: http.StatusInternalServerError,
			Msg:  "存储2FA数据失败" + err.Error(),
		}, nil
	}

	// 设置Redis数据的过期时间，防止未验证的数据长期存在
	if err := dao.Rdb.Expire(context, redisKey, time.Minute).Err(); err != nil {
		log.Println("设置Redis过期时间失败:", err)
	}

	// 返回URL供用户扫码
	return &pb.Setup2FAResponse{
		Code: http.StatusOK,
		Msg:  "生成url成功，请测试2FA",
		Url:  url,
	}, nil
}

// Test2FA 测试2FA可用性 如果通过 存储进数据库 删除redis中的缓存
func (s *UserService) Test2FA(ctx context.Context, request *pb.Test2FARequest) (*pb.Test2FAResponse, error) {
	// 从 Redis 中获取用户的 2FA 秘钥
	log.Println(1)
	redisKey := "2fa_new:" + request.Email
	secretData, err := dao.Rdb.HMGet(ctx, redisKey, "secret_key").Result()
	if err != nil || len(secretData) == 0 || secretData[0] == nil {
		log.Println(secretData)
		return &pb.Test2FAResponse{
			Code: http.StatusUnauthorized,
			Msg:  "2FA 密钥无效或未找到",
		}, err
	}

	secret := secretData[0].(string)
	log.Println(secret)

	log.Println(2)
	// 验证用户输入的 TOTP 验证码
	log.Println(request.TwoFaCode)
	if !auth.Verify2FACode(secret, request.TwoFaCode) {
		return &pb.Test2FAResponse{
			Code: http.StatusConflict,
			Msg:  "验证码无效",
		}, nil
	}

	log.Println(3)
	// 将 2FA 密钥存储进数据库
	if err := dao.Db.Model(&model.TwoFA{}).Create(&model.TwoFA{
		Email:    request.Email,
		TwoFAKey: secret,
	}).Error; err != nil {
		return &pb.Test2FAResponse{
			Code: http.StatusInternalServerError,
			Msg:  "存储 2FA 密钥失败请重试",
		}, err
	}

	// 更新用户表中的 is_two_fa 字段，标记启用 2FA
	if err := dao.Db.Model(&model.User{}).Where("id = ?", request.Id).Update("is_two_fa", true).Error; err != nil {
		return &pb.Test2FAResponse{
			Code: http.StatusInternalServerError,
			Msg:  "更新用户 2FA 状态失败",
		}, err
	}

	// 删除 Redis 中的 2FA 数据
	// 此处的HDel
	if err := dao.Rdb.Del(ctx, redisKey).Err(); err != nil {
		return &pb.Test2FAResponse{
			Code: http.StatusInternalServerError,
			Msg:  "删除 Redis 数据失败",
		}, err
	}

	// 返回成功信息
	return &pb.Test2FAResponse{
		Code: http.StatusOK,
		Msg:  "注册 2FA 成功",
	}, nil
}

func (s *UserService) CancelSetup2FA(context context.Context, request *pb.CancelSetup2FARequest) (*pb.CancelSetup2FAResponse, error) {
	if request.Email == "" {
		return &pb.CancelSetup2FAResponse{
			Code: http.StatusBadRequest,
			Msg:  "缺少有效信息",
		}, nil
	}
	redisKey := "2fa_new:" + request.Email
	if err := dao.Rdb.Del(context, redisKey).Err(); err != nil {
		return &pb.CancelSetup2FAResponse{
			Code: http.StatusInternalServerError,
			Msg:  "删除 Redis 数据失败",
		}, nil
	} else {
		return &pb.CancelSetup2FAResponse{
			Code: http.StatusOK,
			Msg:  "已删除键值",
		}, nil
	}
}

func (s *UserService) Get2FAStatus(context context.Context, request *pb.Get2FAStatusRequest) (*pb.Get2FAStatusResponse, error) {
	var isTwoFa bool
	var err error
	if isTwoFa, err = utils.IsUser2FAEnabled(request.Id); err != nil {
		return &pb.Get2FAStatusResponse{
			Code:    http.StatusOK,
			Msg:     "查询失败" + err.Error(),
			Enabled: isTwoFa,
		}, nil
	}
	log.Println("isTwoFa: ", isTwoFa)
	return &pb.Get2FAStatusResponse{
		Code:    http.StatusOK,
		Msg:     "查询成功",
		Enabled: isTwoFa,
	}, nil
}

func (s *UserService) Disable2FA(context context.Context, request *pb.Disable2FARequest) (*pb.Disable2FAResponse, error) {
	if request.Email == "" {
		return &pb.Disable2FAResponse{
			Code: http.StatusBadRequest,
			Msg:  "缺少有效信息",
		}, nil
	}
	if result := dao.Db.Model(&model.User{}).Where("email = ?", request.Email).Update("is_two_fa", false); result.Error != nil {
		log.Println("关闭失败")
		return &pb.Disable2FAResponse{
			Code:     http.StatusInternalServerError,
			Msg:      "关闭失败" + result.Error.Error(),
			Disabled: false,
		}, nil
	}
	return &pb.Disable2FAResponse{
		Code:     http.StatusOK,
		Msg:      "已关闭2FA",
		Disabled: true,
	}, nil
}
