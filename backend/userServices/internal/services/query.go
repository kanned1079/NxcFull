package services

import (
	"context"
	"net/http"
	pb "userServices/api/proto"
	"userServices/internal/dao"
	"userServices/internal/model"
)

//func (s *UserService) UpdateUserInfo(context context.Context, request *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
//	var user model.User
//	if result := dao.Db.Model(&model.User{}).Where("id = ?", request.UserId).First(&user); result.RowsAffected == 0 {
//		log.Println("无该记录 id：", request.UserId)
//		return &pb.UpdateUserInfoResponse{
//			Code: http.StatusNotFound,
//			Msg:  fmt.Sprintf("User not found. id: %d", request.UserId),
//		}, nil
//	}
//	if userJson, err := json.Marshal(user); err != nil {
//		return &pb.UpdateUserInfoResponse{
//			Code: http.StatusNotFound,
//			Msg:  fmt.Sprintf("User marshal fail. id: %d", request.UserId),
//		}, nil
//	} else {
//		return &pb.UpdateUserInfoResponse{
//			Code:     http.StatusOK,
//			Msg:      "Query ok",
//			UserInfo: userJson,
//		}, nil
//	}
//}

func (s *UserService) GetUsersLayout(ctx context.Context, request *pb.GetUsersLayoutRequest) (*pb.GetUsersLayoutResponse, error) {

	// UserAll
	var userAll int64
	if result := dao.Db.Model(&model.User{}).Unscoped().Count(&userAll); result.Error != nil {
		return &pb.GetUsersLayoutResponse{
			Code: http.StatusInternalServerError,
			Msg:  result.Error.Error(),
		}, nil
	}

	// UserActive: 活跃用户 — 在 model.Order{} 表中有订单的用户
	var userActive int64
	if result := dao.Db.Model(&model.User{}).
		Joins("JOIN `x_orders` o ON o.user_id = x_users.id").
		Where("o.deleted_at IS NULL").            // 排除软删除的订单
		Distinct("x_users.id").                   // 确保每个用户只计算一次
		Count(&userActive); result.Error != nil { // 计算唯一活跃用户数量
		return &pb.GetUsersLayoutResponse{
			Code: http.StatusInternalServerError,
			Msg:  result.Error.Error(),
		}, nil
	}

	// UserInactive: 不活跃用户 — 在 model.Order{} 表中没有订单的用户
	var userInactive int64
	if result := dao.Db.Model(&model.User{}).
		Where("NOT EXISTS (?)",
			dao.Db.Model(&model.Orders{}).
				Select("1").
				Where("`x_orders`.user_id = `x_users`.id").
				Where("`x_orders`.deleted_at IS NULL")).
		Where("`x_users`.deleted_at IS NULL").
		Count(&userInactive); result.Error != nil {
		return &pb.GetUsersLayoutResponse{
			Code: http.StatusInternalServerError,
			Msg:  result.Error.Error(),
		}, nil
	}
	// UserBlocked: 被禁用或已删除的用户
	var userBlocked int64
	if result := dao.Db.Model(&model.User{}).Unscoped().Where("status = ? OR deleted_at IS NOT NULL", false).Count(&userBlocked); result.Error != nil {
		return &pb.GetUsersLayoutResponse{
			Code: http.StatusInternalServerError,
			Msg:  result.Error.Error(),
		}, nil
	}

	return &pb.GetUsersLayoutResponse{
		Code:         http.StatusOK,
		Msg:          "success",
		UserAll:      userAll,
		UserActive:   userActive,
		UserInactive: userInactive,
		UserBlocked:  userBlocked,
	}, nil
}
