package utils

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
	"userServices/internal/dao"
	"userServices/internal/model"
)

// IsUserExist 指定邮箱的用户是否存在
//func IsUserExist(email string) (code int) {
//	user := model.Auth{
//		Email: email,
//	}
//	result := dao.Db.Model(&model.Auth{}).Where("email = ?", user.Email).First(&user)
//	if result.RowsAffected > 0 {
//		code = model.User_Exist
//	} else {
//		code = model.User_Not_Exist
//	}
//	return
//}

func IsUserExist(email string) (code int) {
	user := model.Auth{
		Email: email,
	}
	result := dao.Db.Model(&model.Auth{}).Where("email = ?", user.Email).First(&user)
	if result.RowsAffected > 0 {
		code = model.User_Exist
	} else {
		code = model.User_Not_Exist
	}
	return
}

// AuthUserInfo 校验用户凭据
//func AuthUserInfo(email string, password string) (code int) {
//	var userAuth model.Auth
//	result := dao.Db.Model(&model.Auth{}).Where("email = ?", email).First(&userAuth)
//	if result.Error != nil {
//		log.Println("未知错误")
//	}
//	log.Println("查询到的用户凭证", userAuth)
//	if password == userAuth.Password {
//		code = model.Auth_Pass
//	} else {
//		code = model.Auth_Failure
//	}
//	return
//}

func AuthUserInfo(email string, password string) (code int) {
	var userAuth model.Auth

	// 尝试从 Redis 获取用户凭据信息
	authKey := "auth:" + email

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	authData, err := dao.Rdb.Get(ctx, authKey).Result()
	if errors.Is(err, redis.Nil) {
		// Redis 中不存在，从数据库查询
		log.Println("该用户Auth信息Redis中不存在 从Mysql查询")
		result := dao.Db.Model(&model.Auth{}).Where("email = ?", email).First(&userAuth)
		if result.Error != nil {
			log.Println("数据库查询错误:", result.Error)
			return model.Auth_Failure
		}

		// 将用户凭据序列化为 JSON 字符串存入 Redis
		userJSON, jsonErr := json.Marshal(userAuth)
		if jsonErr != nil {
			log.Println("JSON 序列化用户凭据失败: ", jsonErr)
		} else {
			redisErr := dao.Rdb.Set(ctx, authKey, userJSON, 24*time.Hour).Err()
			if redisErr != nil {
				log.Println("Redis 存储用户凭据失败: ", redisErr)
			}
		}
	} else if err != nil {
		log.Println("Redis 获取用户凭据失败: ", err)
		return model.Auth_Failure
	} else {
		// 从 Redis 获取的数据反序列化为结构体
		log.Println("该用户Auth信息在Redis中存在")
		jsonErr := json.Unmarshal([]byte(authData), &userAuth)
		if jsonErr != nil {
			log.Println("JSON 反序列化用户凭据失败: ", jsonErr)
			return model.Auth_Failure
		}
	}

	// 验证用户密码
	if password == userAuth.Password {
		code = model.Auth_Pass
	} else {
		code = model.Auth_Failure
	}
	return
}

//// IsUserExist 指定邮箱的用户是否存在
//// IsUserExist 指定邮箱的用户是否存在
//// IsUserExist 指定邮箱的用户是否存在
//func IsUserExist(email string) (code int) {
//	redisKey := fmt.Sprintf("user_session:%s", email)
//
//	// 检查 Redis 键的类型
//	keyType, err := dao.Rdb.Type(context.Background(), redisKey).Result()
//	if err != nil {
//		log.Println("Redis 错误:", err)
//		code = model.User_Not_Exist
//		return
//	}
//
//	// 如果键是字符串类型，则说明以前可能用的是字符串存储
//	if keyType == "string" {
//		log.Println("Redis 错误: 键的类型不匹配，期望是哈希类型")
//		dao.Rdb.Del(context.Background(), redisKey) // 删除错误的键，避免后续冲突
//	}
//
//	// 正常处理 Redis 键值
//	_, err = dao.Rdb.HGetAll(context.Background(), redisKey).Result()
//	if err == redis.Nil {
//		// Redis 中不存在，查数据库
//		user := model.Auth{
//			Email: email,
//		}
//		result := dao.Db.Model(&model.Auth{}).Where("email = ?", user.Email).First(&user)
//		if result.RowsAffected > 0 {
//			// 将用户信息存入 Redis
//			userData := map[string]interface{}{
//				"id":    user.Id,
//				"email": user.Email,
//			}
//			dao.Rdb.HSet(context.Background(), redisKey, userData)
//			dao.Rdb.Expire(context.Background(), redisKey, 24*time.Hour)
//			code = model.User_Exist
//		} else {
//			code = model.User_Not_Exist
//		}
//	} else if err != nil {
//		log.Println("Redis 错误:", err)
//		code = model.User_Not_Exist
//	} else {
//		// Redis 中存在用户数据，直接返回
//		code = model.User_Exist
//		dao.Rdb.Expire(context.Background(), redisKey, 24*time.Hour) // 刷新过期时间
//	}
//	return
//}
//
//// AuthUserInfo 校验用户凭据
//func AuthUserInfo(email string, password string) (code int) {
//	redisKey := fmt.Sprintf("user_session:%s", email)
//	// 尝试从 Redis 获取用户凭据
//	userData, err := dao.Rdb.HGetAll(context.Background(), redisKey).Result()
//	if errors.Is(err, redis.Nil) {
//		// Redis 中不存在该用户，查询数据库
//		var userAuth model.Auth
//		result := dao.Db.Model(&model.Auth{}).Where("email = ?", email).First(&userAuth)
//		if result.Error != nil {
//			log.Println("数据库查询错误:", result.Error)
//			code = model.Auth_Failure
//			return
//		}
//		// 校验密码
//		if password == userAuth.Password {
//			// 将用户凭据存储到 Redis 并设置过期时间
//			userData := map[string]interface{}{
//				"id":       userAuth.Id,
//				"email":    userAuth.Email,
//				"password": userAuth.Password,
//			}
//			dao.Rdb.HSet(context.Background(), redisKey, userData)
//			dao.Rdb.Expire(context.Background(), redisKey, 24*time.Hour)
//			code = model.Auth_Pass
//		} else {
//			code = model.Auth_Failure
//		}
//	} else if err != nil {
//		log.Println("Redis 错误:", err)
//		code = model.Auth_Failure
//	} else {
//		// Redis 中存在用户数据，校验密码
//		if password == userData["password"] {
//			code = model.Auth_Pass
//			// 刷新过期时间
//			dao.Rdb.Expire(context.Background(), redisKey, 24*time.Hour)
//		} else {
//			code = model.Auth_Failure
//		}
//	}
//	return
//}
