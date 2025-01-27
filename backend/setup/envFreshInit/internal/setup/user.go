package setup

import (
	"envFreshInit/internal/dao"
	"envFreshInit/internal/model"
	"envFreshInit/internal/utils"
	"fmt"
)

// SetupNewAdminAccount 添加一个默认的管理员
func SetupNewAdminAccount(email, password string) error {
	// 启动事务
	tx := dao.Db.Begin()

	// 确保事务在最后回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			fmt.Println("Recovered in AddNewAdminAccount:", r)
		}
	}()

	// 对密码进行哈希处理
	hashedPwd, err := utils.HashPassword(password)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 创建管理员账户
	if result := tx.Model(&model.User{}).Create(&model.User{
		Email:   email,
		IsAdmin: true,
		IsStaff: false,
		Status:  true,
	}); result.Error != nil {
		tx.Rollback() // 如果用户创建失败，回滚事务
		return result.Error
	}

	// 创建身份验证记录
	if result := tx.Model(&model.Auth{}).Create(&model.Auth{
		Email:    email,
		Password: hashedPwd,
	}); result.Error != nil {
		tx.Rollback() // 如果身份验证记录创建失败，回滚事务
		return result.Error
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
