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
	var decodedPwd string
	if decodedPwd, err = DecodeBase64(password); err != nil {
		log.Println(err)
		return model.Unknow_Error
	}

	if ComparePasswordHash(decodedPwd, userAuth.Password) {
		code = model.Auth_Pass
	} else {
		code = model.Auth_Failure
	}

	//if password == userAuth.Password {
	//} else {
	//}
	//return
	return
}
