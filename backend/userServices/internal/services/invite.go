package services

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	pb "userServices/api/proto"
	"userServices/internal/dao"
	"userServices/internal/model"
	"userServices/internal/utils"
)

/*
  rpc GetUserInviteCodeByUserId(GetUserInviteCodeByUserIdRequest) returns (GetUserInviteCodeByUserIdResponse);
  rpc GetUserInvitedUserListByUserId(GetUserInvitedUserListByUserIdRequest) returns (GetUserInvitedUserListByUserIdResponse);
  rpc CreateUserInviteCodeByUserId(CreateUserInviteCodeByUserIdRequest) returns (CreateUserInviteCodeByUserIdResponse);
*/

func (s *UserService) GetUserInviteCodeByUserId(ctx context.Context, request *pb.GetUserInviteCodeByUserIdRequest) (*pb.GetUserInviteCodeByUserIdResponse, error) {
	var inviteCode string

	// 根据request.UserId查询其invite_code字段的值
	if result := dao.Db.Model(&model.User{}).Select("invite_code").Where("id = ?", request.UserId).Scan(&inviteCode); result.Error != nil {
		// 数据库查询错误
		return &pb.GetUserInviteCodeByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Database Error",
		}, result.Error
	}

	// 如果字段没有值
	if inviteCode == "" {
		return &pb.GetUserInviteCodeByUserIdResponse{
			Code: http.StatusNotFound,
			Msg:  "InviteCode Not Found",
		}, nil
	}

	// 如果字段有值
	return &pb.GetUserInviteCodeByUserIdResponse{
		Code:         http.StatusOK,
		Msg:          "ok",
		MyInviteCode: inviteCode,
	}, nil
}

func (s *UserService) CreateUserInviteCodeByUserId(ctx context.Context, request *pb.CreateUserInviteCodeByUserIdRequest) (*pb.CreateUserInviteCodeByUserIdResponse, error) {
	// 开启事务
	tx := dao.Db.Begin()

	// 确保事务最终被提交或回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 查询用户邮箱
	var email string
	if err := tx.Model(&model.User{}).Select("email").Where("id = ?", request.UserId).Scan(&email).Error; err != nil {
		// 查询邮箱失败，回滚事务并返回错误
		tx.Rollback()
		return &pb.CreateUserInviteCodeByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to retrieve user email",
		}, err
	}

	// 生成邀请码
	newInviteCode := generateInviteCode(request.UserId)

	// 更新用户的 invite_code 字段
	result := tx.Model(&model.User{}).Where("id = ?", request.UserId).Update("invite_code", newInviteCode)
	if result.Error != nil {
		// 数据库更新出错，回滚事务并返回错误
		tx.Rollback()
		return &pb.CreateUserInviteCodeByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to update invite code",
		}, result.Error
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		// 提交事务失败
		return &pb.CreateUserInviteCodeByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to commit transaction",
		}, err
	}

	// 清除用户缓存（基于邮箱）
	if err := utils.ClearUserCacheByEmail(email); err != nil {
		return &pb.CreateUserInviteCodeByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to clear user cache",
		}, err
	}

	// 返回成功响应
	return &pb.CreateUserInviteCodeByUserIdResponse{
		Code: http.StatusOK,
		Msg:  "Invite code created successfully",
	}, nil
}

func (s *UserService) GetUserInvitedUserListByUserId(ctx context.Context, request *pb.GetUserInvitedUserListByUserIdRequest) (*pb.GetUserInvitedUserListByUserIdResponse, error) {
	// 请求中的分页参数
	page := request.Page
	size := request.Size

	// 校验分页参数
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10 // 默认每页 10 条
	}

	var userList []struct {
		Id        int64      `json:"id"`
		Email     string     `json:"email"`
		CreatedAt *time.Time `json:"created_at"`
	}

	// 查询被邀请用户的总数
	var totalCount int64
	if err := dao.Db.Model(&model.User{}).
		Where("invite_user_id = ?", request.UserId).
		Count(&totalCount).Error; err != nil {
		return &pb.GetUserInvitedUserListByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to fetch total count: " + err.Error(),
		}, err
	}

	// 计算总页数
	pageCount := int((totalCount + int64(size) - 1) / int64(size)) // 向上取整

	// 查询被邀请用户列表（分页）
	offset := (page - 1) * size
	result := dao.Db.Model(&model.User{}).
		Select("id, email, created_at").
		Where("invite_user_id = ?", request.UserId).
		Order("created_at DESC"). // 根据创建时间倒序排列
		Offset(int(offset)).      // 跳过的记录数
		Limit(int(size)).         // 每页记录数
		Scan(&userList)

	// 检查数据库查询是否发生错误
	if result.Error != nil {
		return &pb.GetUserInvitedUserListByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to fetch invited user list: " + result.Error.Error(),
		}, result.Error
	}

	// 如果没有被邀请的用户，返回空列表
	//if len(userList) == 0 {
	//	return &pb.GetUserInvitedUserListByUserIdResponse{
	//		Code:      http.StatusNotFound,
	//		Msg:       "No invited users found",
	//		UserList:  nil,
	//		PageCount: int64(pageCount),
	//	}, nil
	//}

	// 将查询结果转换为 JSON
	jsonUserList, err := json.Marshal(userList)
	if err != nil {
		return &pb.GetUserInvitedUserListByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to marshal user list: " + err.Error(),
		}, err
	}

	// 返回成功响应
	return &pb.GetUserInvitedUserListByUserIdResponse{
		Code:      http.StatusOK,
		Msg:       "ok",
		UserList:  jsonUserList,
		PageCount: int64(pageCount),
	}, nil
}

func generateInviteCode(userId int64) (code string) {
	str := "abcdefghijklmnopqrstuvwxyz0123456789"
	// 生成的随机字符串长度为10
	// 组成为 <随机字符串><用户id>
	userIdStr := strconv.FormatInt(userId, 10)
	randomStrLen := 18 - len(userIdStr)
	// 注意随机字符串的长度为 10-用户id长度
	// 例子 用户id为4 整个输出为fj82jd3as4  用户id为378 整个输出为i20d2h1378
	rand.Seed(time.Now().UnixNano())

	randomStr := make([]byte, randomStrLen)
	for i := 0; i < randomStrLen; i++ {
		randomStr[i] = str[rand.Intn(len(str))]
	}
	code = string(randomStr) + userIdStr
	return
}
