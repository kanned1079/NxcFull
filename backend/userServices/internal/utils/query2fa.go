package utils

import (
	"log"
	"userServices/internal/dao"
	"userServices/internal/model"
)

// IsUser2FAEnabled 检查用户是否启用了2FA验证
//func IsUser2FAEnabled(userId int64) (enabled bool, err error) {
//	var data bool
//	if result := dao.Db.Model(&model.User{}).Where("id = ?", userId).Select("is_two_fa").Scan(&data); result.Error != nil {
//		log.Println("查询失败" + result.Error.Error())
//		return false, result.Error
//	}
//	log.Println(data)
//	return data, nil
//}

func IsUser2FAEnabled(userId int64) (enabled bool, err error) {
	var data bool
	// 使用 Row() 获取单个值
	log.Println("userId:", userId)
	err = dao.Db.Model(&model.User{}).Where("id = ?", userId).Select("is_two_fa").Row().Scan(&data)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return false, err
	}

	log.Printf("用户 %d 的 2FA 状态: %v", userId, data)
	return data, nil
}
