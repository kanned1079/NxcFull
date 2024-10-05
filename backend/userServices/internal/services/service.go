package services

import (
	pb "NxcFull/backend/userServices/api/proto"
	"NxcFull/backend/userServices/internal/auth"
	"NxcFull/backend/userServices/internal/dao"
	"NxcFull/backend/userServices/internal/model"
	"NxcFull/nxc-backend/user"
	"context"
	"log"
	"net/http"
	"time"
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

func (s *UserService) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {

	var newUserInfo user.User = user.User{
		Email:        req.Email,
		InviteUserID: req.InviteUserId,
		//LicenseExpiration: nil,
	}
	var newUserAuth user.Auth = user.Auth{
		Email:    req.Email,
		Password: req.Password,
	}
	result1 := dao.Db.Model(&user.User{}).Create(&newUserInfo)
	if result1.Error != nil {
		//context.JSON(http.StatusInternalServerError, gin.H{
		//	"code":  http.StatusInternalServerError,
		//	"msg":   "插入数据错误",
		//	"error": result1.Error.Error(),
		//})
		return &pb.UserRegisterResponse{
			Code:         http.StatusInternalServerError,
			Msg:          "插入数据错误",
			IsRegistered: false,
		}, nil
	}
	result2 := dao.Db.Model(&user.Auth{}).Create(&newUserAuth)
	if result2.Error != nil {
		//context.JSON(http.StatusInternalServerError, gin.H{
		//	"code":  http.StatusInternalServerError,
		//	"msg":   "插入验证数据错误",
		//	"error": result2.Error.Error(),
		//})
		return &pb.UserRegisterResponse{
			Code:         http.StatusInternalServerError,
			Msg:          "插入验证数据错误",
			IsRegistered: false,
		}, nil
	}
	if result1.Error == nil && result2.Error == nil {
		//context.JSON(http.StatusOK, gin.H{
		//	"code": http.StatusOK,
		//	"msg":  "新建用户成功",
		//})
		return &pb.UserRegisterResponse{
			Code:         http.StatusOK,
			Msg:          "插入数据成功",
			IsRegistered: true,
		}, nil
	} else {
		return &pb.UserRegisterResponse{
			Code:         http.StatusInternalServerError,
			Msg:          "未知错误",
			IsRegistered: false,
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
