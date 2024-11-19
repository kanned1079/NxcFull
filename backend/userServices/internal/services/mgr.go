package services

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
	pb "userServices/api/proto"
	"userServices/internal/dao"
	"userServices/internal/model"
	"userServices/internal/utils"
)

//adminAuthorized.GET("users", handler.HandleGetAllUsers)
//adminAuthorized.POST("users", handler.HandleAddUserManuallyFromAdmin)
//adminAuthorized.PUT("users", handler.HandleUpdateUserInfoByIdFromAdmin)
//adminAuthorized.PATCH("users", handler.HandleBlock2UnblockUserByIdFromAdmin)

//rpc GetAllUsers(GetAllUsersRequest) returns (GetAllUsersResponse);
//rpc AddUserManually(AddUserManuallyRequest) returns (AddUserManuallyResponse);
//rpc UpdateUserInfoAdmin(UpdateUserInfoAdminRequest) returns (UpdateUserInfoAdminResponse);
//rpc Block2UnblockUserById(Block2UnblockUserByIdRequest) returns (Block2UnblockUserByIdResponse);

//func (s *UserService) GetAllUsers(context context.Context, request *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
//	request.Page
//	request.Size // 分页参数
//	request.SearchEmail // 模糊搜索邮箱
//
//	var user []model.User
//	if result := dao.Db.Model(&model.User{}).
//
//
//	return &pb.GetAllUsersResponse{
//		Code:      http.StatusOK,
//		PageCount: 0, // 计算这里的总页面数
//		Users: //这个是[]byte类型 需要将用户的列表序列化,如果找不到就到返回空的数组即可
//	}, nil
//}

//// 以下是前端需要的数据， 下面的后端代码中返回的是用户的信息而不包含以下两个需要的数据
//// 修改代码让下面两个数据也包含进users信息并返回
//// plan_count 需要查询 &model.ActiveOrders{}
//// order_count 需要查询 &model.Orders{}
///*
//interface User {
//  id?: number
//  invite_user_id?: number
//  email?: string
//  is_admin?: boolean
//  is_staff?: boolean
//  status?: boolean
//  plan_count?: number
//  order_count?: number
//  balance?: number
//  created_at?: string
//}
//*/
//
//func (s *UserService) GetAllUsers(ctx context.Context, request *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
//	// 从请求中获取分页参数
//	page := int(request.Page)
//	size := int(request.Size)
//	searchEmail := request.SearchEmail
//
//	var users []model.User
//
//	//// 得到 plan_count
//	//dao.Db.Model(&model.ActiveOrders{}).Where("user_id = ?", 12).Count()
//	//
//	//// 得到 order_count
//	//dao.Db.Model(&model.Orders{}).Where("user_id = ?", 12).Count()
//
//	// 查询数据库，支持分页和模糊搜索邮箱
//	query := dao.Db.Model(&model.User{})
//	if searchEmail != "" {
//		query = query.Where("email LIKE ?", "%"+searchEmail+"%")
//	}
//
//	var total int64
//	if err := query.Count(&total).Error; err != nil {
//		return nil, err
//	}
//
//	// 分页查询
//	if err := query.Offset((page - 1) * size).Limit(size).Find(&users).Error; err != nil {
//		return nil, err
//	}
//
//	// 计算总页数
//	pageCount := int((total + int64(size) - 1) / int64(size))
//
//	// 将用户列表序列化为 []byte
//	userBytes, err := json.Marshal(users)
//	if err != nil {
//		return nil, err
//	}
//
//	// 构建返回响应
//	return &pb.GetAllUsersResponse{
//		Code:      http.StatusOK,
//		PageCount: int64(pageCount),
//		Users:     userBytes,
//	}, nil
//}

func (s *UserService) GetAllUsers(ctx context.Context, request *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	// 从请求中获取分页参数
	page := int(request.Page)
	size := int(request.Size)
	searchEmail := request.SearchEmail

	// 用于存储用户数据，包含扩展信息
	var usersWithCounts []struct {
		model.User
		PlanCount  int64 `json:"plan_count"`
		OrderCount int64 `json:"order_count"`
	}

	// 构建查询，确保 Unscoped() 始终生效
	query := dao.Db.Unscoped().Model(&model.User{})
	if searchEmail != "" {
		query = query.Where("email LIKE ?", "%"+searchEmail+"%")
	}

	// 获取总记录数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// 分页查询用户基础信息
	var users []model.User
	if err := query.Offset((page - 1) * size).Limit(size).Find(&users).Error; err != nil {
		return nil, err
	}

	// 为每个用户查询扩展信息（plan_count 和 order_count）
	for _, user := range users {
		var planCount int64
		var orderCount int64

		// 计算 plan_count
		if err := dao.Db.Model(&model.ActiveOrders{}).Where("user_id = ?", user.Id).Count(&planCount).Error; err != nil {
			return nil, err
		}

		// 计算 order_count
		if err := dao.Db.Model(&model.Orders{}).Where("user_id = ?", user.Id).Count(&orderCount).Error; err != nil {
			return nil, err
		}

		// 将用户信息和附加信息添加到结果中
		usersWithCounts = append(usersWithCounts, struct {
			model.User
			PlanCount  int64 `json:"plan_count"`
			OrderCount int64 `json:"order_count"`
		}{
			User:       user,
			PlanCount:  planCount,
			OrderCount: orderCount,
		})
	}

	// 序列化用户列表数据
	userBytes, err := json.Marshal(usersWithCounts)
	if err != nil {
		return nil, err
	}

	// 计算总页数
	pageCount := int((total + int64(size) - 1) / int64(size))

	// 构建返回响应
	return &pb.GetAllUsersResponse{
		Code:      http.StatusOK,
		PageCount: int64(pageCount),
		Users:     userBytes,
	}, nil
}

func (s *UserService) AddUserManually(context context.Context, request *pb.AddUserManuallyRequest) (*pb.AddUserManuallyResponse, error) {

	return &pb.AddUserManuallyResponse{
		Code: http.StatusOK,
	}, nil
}

// 原始版本 无事务
func (s *UserService) UpdateUserInfoAdmin(context context.Context, request *pb.UpdateUserInfoAdminRequest) (*pb.UpdateUserInfoAdminResponse, error) {
	var user model.User

	if result := dao.Db.Model(&model.User{}).Where("id = ?", request.Id).First(&user); result.RowsAffected == 0 {
		return &pb.UpdateUserInfoAdminResponse{
			Code:    http.StatusNotFound,
			Success: false,
			Msg:     "user not found" + result.Error.Error(),
		}, nil
	}
	//user.Email = request.Email
	log.Println("previous", user)
	user.InviteCode = request.InviteCode
	user.IsAdmin = request.IsAdmin
	user.IsStaff = request.IsStaff
	user.Balance = request.Balance

	log.Println("--------previous", user)

	if dao.Db.Model(&model.User{}).Where("id = ?", request.Id).Save(&user).Error != nil {
		return &pb.UpdateUserInfoAdminResponse{
			Code:    http.StatusInternalServerError,
			Success: false,
			Msg:     "Update user data failure",
		}, nil
	}

	if request.Password != "" {
		var userAuth model.Auth

		if result := dao.Db.Model(&model.Auth{}).Where("id = ?", request.Id).First(&userAuth); result.RowsAffected == 0 {
			return &pb.UpdateUserInfoAdminResponse{
				Code:    http.StatusNotFound,
				Success: false,
				Msg:     "user not found" + result.Error.Error(),
			}, nil
		}

		//userAuth.Email = request.Email
		userAuth.Password = request.Password

		if err := dao.Db.Model(&model.Auth{}).Where("id = ?", request.Id).Save(&userAuth).Error; err != nil {
			return &pb.UpdateUserInfoAdminResponse{
				Code:    http.StatusInternalServerError,
				Success: false,
				Msg:     "Update user auth data failure" + err.Error(),
			}, nil
		}
	}

	if utils.ClearUserCacheByEmail(user.Email) != nil {
		return &pb.UpdateUserInfoAdminResponse{
			Code:    http.StatusInternalServerError,
			Success: false,
			Msg:     "Flush user cache failure",
		}, nil
	}

	return &pb.UpdateUserInfoAdminResponse{
		Code:    http.StatusOK,
		Success: true,
		Msg:     "Updated user successfully",
	}, nil
}

//func (s *UserService) UpdateUserInfoAdmin(context context.Context, request *pb.UpdateUserInfoAdminRequest) (*pb.UpdateUserInfoAdminResponse, error) {
//	// 开启事务
//	tx := dao.Db.Begin()
//	defer func() {
//		if r := recover(); r != nil {
//			tx.Rollback()
//		}
//	}()
//
//	var user model.User
//	if result := tx.Model(&model.User{}).Where("id = ?", request.Id).First(&user); result.RowsAffected == 0 {
//		tx.Rollback()
//		return &pb.UpdateUserInfoAdminResponse{
//			Code:    http.StatusNotFound,
//			Success: false,
//			Msg:     "user not found: " + result.Error.Error(),
//		}, nil
//	}
//
//	// 更新 x_auth 表中的 email 字段
//	if err := tx.Model(&model.Auth{}).Where("id = ?", request.Id).Update("email", request.Email).Error; err != nil {
//		tx.Rollback()
//		return &pb.UpdateUserInfoAdminResponse{
//			Code:    http.StatusInternalServerError,
//			Success: false,
//			Msg:     "failed to update auth email: " + err.Error(),
//		}, nil
//	}
//
//	// 更新 x_users 表中的信息
//	user.Email = request.Email
//	if err := tx.Save(&user).Error; err != nil {
//		tx.Rollback()
//		return &pb.UpdateUserInfoAdminResponse{
//			Code:    http.StatusInternalServerError,
//			Success: false,
//			Msg:     "failed to update user: " + err.Error(),
//		}, nil
//	}
//
//	// 提交事务
//	if err := tx.Commit().Error; err != nil {
//		return &pb.UpdateUserInfoAdminResponse{
//			Code:    http.StatusInternalServerError,
//			Success: false,
//			Msg:     "failed to commit transaction: " + err.Error(),
//		}, nil
//	}
//
//	return &pb.UpdateUserInfoAdminResponse{
//		Code:    http.StatusOK,
//		Success: true,
//		Msg:     "Updated user successfully",
//	}, nil
//}

//func (s *UserService) Block2UnblockUserById(ctx context.Context, request *pb.Block2UnblockUserByIdRequest) (*pb.Block2UnblockUserByIdResponse, error) {
//	var user model.User
//	if err := dao.Db.Model(&model.User{}).Where("id = ?", request.UserId).First(&user).Error; err != nil {
//		return &pb.Block2UnblockUserByIdResponse{
//			Code: http.StatusNotFound,
//			Msg:  "User not found",
//		}, nil
//	}
//
//	if err := dao.Db.Model(&model.User{}).Where("id = ?", request.UserId).Update("status", !user.Status).Error; err != nil {
//		return &pb.Block2UnblockUserByIdResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "Failed to update user status",
//		}, nil
//	}
//
//	if err := utils.ClearUserCacheByEmail(user.Email); err != nil {
//		return &pb.Block2UnblockUserByIdResponse{
//			Code: http.StatusAccepted,
//			Msg:  "Clear user cache failure, data may not be latest",
//		}, nil
//	}
//
//	return &pb.Block2UnblockUserByIdResponse{
//		Code: http.StatusOK,
//		Msg:  "success",
//	}, nil
//}

func (s *UserService) Block2UnblockUserById(ctx context.Context, request *pb.Block2UnblockUserByIdRequest) (*pb.Block2UnblockUserByIdResponse, error) {
	var email string
	if err := dao.Db.Model(&model.User{}).Where("id = ?", request.UserId).Select("email").Scan(&email).Error; err != nil {
		return &pb.Block2UnblockUserByIdResponse{
			Code: http.StatusNotFound,
			Msg:  "User not found",
		}, nil
	}

	if err := dao.Db.Model(&model.User{}).Where("id = ?", request.UserId).
		Update("status", gorm.Expr("NOT status")).Error; err != nil {
		return &pb.Block2UnblockUserByIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to update user status",
		}, nil
	}

	if err := utils.ClearUserCacheByEmail(email); err != nil {
		return &pb.Block2UnblockUserByIdResponse{
			Code: http.StatusAccepted,
			Msg:  "Clear user cache failure, data may not be latest",
		}, nil
	}

	return &pb.Block2UnblockUserByIdResponse{
		Code: http.StatusOK,
		Msg:  "success",
	}, nil
}
