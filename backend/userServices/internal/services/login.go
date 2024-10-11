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
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

// Login 实现用户登录的逻辑
func (s *UserService) Login(cxt context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	if model.IsUserExist(req.Email) == model.User_Exist {
		//log.Println("验证密码")
		// 验证用户密码
		if model.AuthUserInfo(req.Email, req.Password) == model.Auth_Pass {
			var thisUser model.User
			if result := dao.Db.Model(&model.User{}).Where("email = ?", req.Email).First(&thisUser); result.Error != nil {
				return &pb.UserLoginResponse{
					Code:     http.StatusInternalServerError,
					IsAuthed: false,
					Msg:      result.Error.Error(),
				}, nil
			}
			token, err := auth.GenerateToken(req.Email, req.Role)
			log.Println("Token: ", token)
			if err != nil {
				return &pb.UserLoginResponse{
					Code:     http.StatusInternalServerError,
					IsAuthed: false,
					Msg:      err.Error(),
				}, nil
			}
			return &pb.UserLoginResponse{
				Code:     http.StatusOK,
				IsAuthed: true,
				Msg:      "auth_passed",
				Token:    token,
				UserData: convertModelUserToProtoUser(&thisUser),
			}, nil
		} else {
			return &pb.UserLoginResponse{
				Code:     http.StatusInternalServerError,
				IsAuthed: false,
				Msg:      "incorrect_password",
			}, nil
		}
	} else {
		return &pb.UserLoginResponse{
			Code:     http.StatusNotFound,
			IsAuthed: false,
			Msg:      "user_not_exist",
		}, nil
	}
}

func convertModelUserToProtoUser(modelUser *model.User) *pb.User {
	return &pb.User{
		Id:           modelUser.Id,
		InviteUserId: modelUser.InviteUserID,
		GroupId:      modelUser.GroupId,
		Name:         modelUser.Name,
		Email:        modelUser.Email,
		IsAdmin:      modelUser.IsAdmin,
		IsStaff:      modelUser.IsStaff,
		Balance:      modelUser.Balance,
		LastLogin:    convertPointerTimeToCustomTimestamp(modelUser.LastLogin), // 将 Go 的时间转换为 protobuf 的时间
		LastLoginIp:  modelUser.LastLoginIp,
		CreatedAt:    convertTimeToCustomTimestamp(modelUser.CreatedAt),
		UpdatedAt:    convertTimeToCustomTimestamp(modelUser.UpdatedAt),
		DeletedAt:    convertTimeToCustomTimestamp(modelUser.DeletedAt.Time), // 处理软删除时间
	}
}

// 将 time.Time 转换为自定义的 Timestamp
func convertTimeToCustomTimestamp(t time.Time) *pb.Timestamp {
	return &pb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

// 处理指针类型的 time.Time 转换为自定义的 Timestamp
func convertPointerTimeToCustomTimestamp(t *time.Time) *pb.Timestamp {
	if t == nil {
		return nil // 如果指针为 nil，返回 nil
	}
	return &pb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
