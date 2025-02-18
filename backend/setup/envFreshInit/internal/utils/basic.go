package utils

import (
	"envFreshInit/internal/dao"
	"envFreshInit/internal/model"
	"fmt"
)

// ClearTable 清除某一个表
func ClearTable(tableName string) (err error) {
	err = dao.Db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName)).Error
	if err == nil {
		ShowOK()
		fmt.Printf("Successfully cleared table %s.\n", tableName)
	}
	return
}

// InsertDataToDB 将数据插入到数据库
func InsertDataToDB(settings []model.SiteSetting) error {
	for _, setting := range settings {
		// 将数据插入到数据库

		result := dao.Db.Create(&setting)
		if result.Error != nil {
			ShowFailure()
			return fmt.Errorf("cannot insert rows to db: %w", result.Error)
		}
		ShowOK()
		fmt.Printf("key[%s.%s]: %s\n", setting.Category, setting.Key, setting.Value)

	}
	return nil
}
