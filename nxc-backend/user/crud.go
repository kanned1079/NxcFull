package user

import (
	"NxcFull/nxc-backend/dao"
	"log"
)

// IsUserExist 指定邮箱的用户是否存在
func IsUserExist(email string) (code int) {
	user := Auth{
		Email: email,
	}
	result := dao.Db.Model(&Auth{}).Where("email = ?", user.Email).First(&user)
	if result.RowsAffected > 0 {
		code = User_Exist
	} else {
		code = User_Not_Exist
	}
	return
}

// AuthUserInfo 校验用户凭据
func AuthUserInfo(email string, password string) (code int) {
	var userAuth Auth
	result := dao.Db.Model(&Auth{}).Where("email = ?", email).First(&userAuth)
	if result.Error != nil {
		log.Println("未知错误")
	}
	if password == userAuth.Password {
		code = Auth_Pass
	} else {
		code = Auth_Failure
	}
	return
}
