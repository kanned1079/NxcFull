package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"time"
	pb "userServices/api/proto"
	"userServices/internal/auth"
	"userServices/internal/dao"
	"userServices/internal/model"
	"userServices/internal/utils"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

//func (s *UserService) Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
//	log.Println(req.Email, req.Password, req.Role, req.TwoFaCode)
//
//	var thisUser model.User
//	// 尝试从 Redis 获取用户信息
//	userKey := "user:" + req.Email
//	userData, err := dao.Rdb.Get(ctx, userKey).Result()
//
//	if errors.Is(err, redis.Nil) {
//		// Redis 缓存中不存在用户信息，从数据库查询
//		log.Println("用户信息在Redis中不存在")
//		if utils.IsUserExist(req.Email) == model.User_Exist { // 数据库查询
//			// 验证用户密码
//			if utils.AuthUserInfo(req.Email, req.Password) == model.Auth_Pass {
//				// 从数据库获取用户信息
//				log.Println("用户信息验证成功")
//
//				if result := dao.Db.Model(&model.User{}).Where("email = ?", req.Email).First(&thisUser); result.Error != nil {
//					return &pb.UserLoginResponse{
//						Code:     http.StatusInternalServerError,
//						IsAuthed: false,
//						Msg:      "查询用户信息失败" + result.Error.Error(),
//					}, nil
//				}
//
//				if thisUser.IsTwoFA {
//					if passed, err := auth.GetSecretAndVerify2FACode(req.Email, req.TwoFaCode); !passed || err != nil {
//						return &pb.UserLoginResponse{
//							Code:     http.StatusUnauthorized,
//							IsAuthed: false,
//							Msg:      "2fa验证未通过",
//						}, nil
//					}
//				}
//
//				// 将用户信息序列化为 JSON 字符串
//				userJSON, jsonErr := json.Marshal(thisUser)
//				if jsonErr != nil {
//					log.Println("JSON 序列化用户信息失败: ", jsonErr)
//				}
//
//				// 登录成功后，存储用户信息到 Redis
//				redisErr := dao.Rdb.Set(ctx, userKey, userJSON, 24*time.Hour).Err()
//				if redisErr != nil {
//					log.Println("Redis 存储用户信息失败: ", redisErr)
//				}
//			} else {
//				log.Println("密码不正确")
//				return &pb.UserLoginResponse{
//					Code:     http.StatusInternalServerError,
//					IsAuthed: false,
//					Msg:      "incorrect_password",
//				}, nil
//			}
//		} else {
//			return &pb.UserLoginResponse{
//				Code:     http.StatusNotFound,
//				IsAuthed: false,
//				Msg:      "user_not_exist",
//			}, nil
//		}
//	} else if err != nil {
//		log.Println("Redis 获取用户信息失败: ", err)
//		return &pb.UserLoginResponse{
//			Code:     http.StatusInternalServerError,
//			IsAuthed: false,
//			Msg:      err.Error(),
//		}, nil
//	} else {
//		// 反序列化用户信息
//		log.Println("用户信息在Redis中存在")
//		err := json.Unmarshal([]byte(userData), &thisUser)
//		if err != nil {
//			log.Println("JSON 反序列化用户信息失败: ", err)
//			return &pb.UserLoginResponse{
//				Code:     http.StatusInternalServerError,
//				IsAuthed: false,
//				Msg:      "failed_to_parse_user_data",
//			}, nil
//		}
//
//		// 即使 Redis 中有用户数据，仍然需要验证密码
//		if utils.AuthUserInfo(req.Email, req.Password) != model.Auth_Pass {
//			return &pb.UserLoginResponse{
//				Code:     http.StatusUnauthorized,
//				IsAuthed: false,
//				Msg:      "incorrect_password",
//			}, nil
//		}
//
//		if thisUser.IsTwoFA {
//			if passed, err := auth.GetSecretAndVerify2FACode(req.Email, req.TwoFaCode); !passed || err != nil {
//				return &pb.UserLoginResponse{
//					Code:     http.StatusUnauthorized,
//					IsAuthed: false,
//					Msg:      "2fa验证未通过",
//				}, nil
//			}
//		}
//
//		log.Println("用户是否启用2FA", thisUser.IsTwoFA)
//
//	}
//
//	// 生成 token
//	token, err := auth.GenerateToken(req.Email, req.Role)
//	if err != nil {
//		return &pb.UserLoginResponse{
//			Code:     http.StatusInternalServerError,
//			IsAuthed: false,
//			Msg:      err.Error(),
//		}, nil
//	}
//
//	if jsonUser, err := json.Marshal(thisUser); err != nil {
//		return &pb.UserLoginResponse{
//			Code:     http.StatusInternalServerError,
//			IsAuthed: false,
//			Msg:      err.Error(),
//		}, nil
//	} else {
//		return &pb.UserLoginResponse{
//			Code:     http.StatusOK,
//			IsAuthed: true,
//			Msg:      "auth_passed",
//			Token:    token,
//			UserData: jsonUser,
//		}, nil
//	}
//
//	//return &pb.UserLoginResponse{
//	//	Code:     http.StatusOK,
//	//	IsAuthed: true,
//	//	Msg:      "auth_passed",
//	//	Token:    token,
//	//	UserData: convertModelUserToProtoUser(&thisUser),
//	//}, nil
//
//}

func (s *UserService) Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	// 调试日志：输入参数
	log.Printf("登录请求: Email=%s, Role=%s", req.Email, req.Role)

	var thisUser model.User
	userKey := "user:" + req.Email

	// 从 Redis 获取用户信息
	userData, err := dao.Rdb.Get(ctx, userKey).Result()
	if errors.Is(err, redis.Nil) {
		// Redis 缓存中不存在用户信息，从数据库查询
		log.Println("用户信息在Redis中不存在，尝试从数据库查询")
		if utils.IsUserExist(req.Email) == model.User_Exist { // 检查用户是否存在
			// 验证用户密码
			if utils.AuthUserInfo(req.Email, req.Password) == model.Auth_Pass {
				// 从数据库获取用户信息
				log.Println("用户验证通过，开始从数据库查询用户信息")
				result := dao.Db.Model(&model.User{}).Where("email = ?", req.Email).First(&thisUser)
				if result.Error != nil {
					log.Printf("数据库查询用户信息失败: %s", result.Error)
					return &pb.UserLoginResponse{
						Code:     http.StatusInternalServerError,
						IsAuthed: false,
						Msg:      "查询用户信息失败: " + result.Error.Error(),
					}, nil
				}

				// 将用户信息存储到 Redis
				userJSON, jsonErr := json.Marshal(thisUser)
				if jsonErr != nil {
					log.Printf("JSON 序列化用户信息失败: %v", jsonErr)
				} else {
					redisErr := dao.Rdb.Set(ctx, userKey, userJSON, 24*time.Hour).Err()
					if redisErr != nil {
						log.Printf("Redis 存储用户信息失败: %v", redisErr)
					} else {
						log.Println("用户信息已成功存储到 Redis")
					}
				}
			} else {
				log.Println("密码验证失败")
				return &pb.UserLoginResponse{
					Code:     http.StatusUnauthorized,
					IsAuthed: false,
					Msg:      "密码错误",
				}, nil
			}
		} else {
			log.Println("用户不存在")
			return &pb.UserLoginResponse{
				Code:     http.StatusNotFound,
				IsAuthed: false,
				Msg:      "用户不存在",
			}, nil
		}
	} else if err != nil {
		log.Printf("Redis 获取用户信息失败: %v", err)
		return &pb.UserLoginResponse{
			Code:     http.StatusInternalServerError,
			IsAuthed: false,
			Msg:      "Redis获取失败: " + err.Error(),
		}, nil
	} else {
		// 反序列化 Redis 中的用户信息
		log.Println("用户信息在Redis中存在，尝试反序列化")
		err := json.Unmarshal([]byte(userData), &thisUser)
		if err != nil {
			log.Printf("JSON 反序列化失败: %v", err)
			return &pb.UserLoginResponse{
				Code:     http.StatusInternalServerError,
				IsAuthed: false,
				Msg:      "反序列化用户信息失败",
			}, nil
		}

		// 验证密码（管理员无需2FA检查）
		if utils.AuthUserInfo(req.Email, req.Password) != model.Auth_Pass {
			log.Println("密码验证失败")
			return &pb.UserLoginResponse{
				Code:     http.StatusUnauthorized,
				IsAuthed: false,
				Msg:      "密码错误",
			}, nil
		}

		if thisUser.IsTwoFA && !thisUser.IsAdmin {
			if passed, err := auth.GetSecretAndVerify2FACode(req.Email, req.TwoFaCode); !passed || err != nil {
				return &pb.UserLoginResponse{
					Code:     http.StatusUnauthorized,
					IsAuthed: false,
					Msg:      "2fa验证未通过",
				}, nil
			}
		}
	}

	// 调试日志：用户信息检查
	log.Printf("用户信息加载完成: %+v", thisUser)

	// 生成 token
	token, err := auth.GenerateToken(req.Email, req.Role)
	if err != nil {
		log.Printf("生成Token失败: %v", err)
		return &pb.UserLoginResponse{
			Code:     http.StatusInternalServerError,
			IsAuthed: false,
			Msg:      "生成Token失败: " + err.Error(),
		}, nil
	}

	// 返回登录成功结果
	userJSON, err := json.Marshal(thisUser)
	if err != nil {
		log.Printf("JSON 序列化最终用户信息失败: %v", err)
		return &pb.UserLoginResponse{
			Code:     http.StatusInternalServerError,
			IsAuthed: false,
			Msg:      "序列化用户信息失败",
		}, nil
	}

	log.Println("管理员登录成功，返回Token和用户信息")

	log.Println(string(userJSON))
	return &pb.UserLoginResponse{
		Code:     http.StatusOK,
		IsAuthed: true,
		Msg:      "登录成功",
		Token:    token,
		UserData: userJSON,
	}, nil
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
