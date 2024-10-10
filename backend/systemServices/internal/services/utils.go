package services

import (
	"NxcFull/backend/systemServices/internal/dao"
	settings "NxcFull/backend/systemServices/internal/model"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"log"
	"reflect"
)

func UnmarshalSingle(val json.RawMessage) (result string) {
	if err := json.Unmarshal(val, &result); err != nil {
		log.Println(err)
	}
	return
}

// saveSettingToDB 将单个设置保存到数据库
func saveSettingToDB(category, key string, value json.RawMessage) error {
	setting := settings.SiteSetting{
		Category: category,
		Key:      key,
	}

	// 查找现有记录
	if err := dao.Db.Where("`category` = ? AND `key` = ?", category, key).First(&setting).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有找到记录，则创建新记录
			setting.Value = value
			if err := dao.Db.Create(&setting).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		// 如果找到记录，则更新值
		setting.Value = value
		if err := dao.Db.Save(&setting).Error; err != nil {
			return err
		}
	}

	log.Println("Setting created or updated successfully:", key)
	return nil
}

// readSettingFromDB 从数据库中取出一条记录
func readSettingFromDB(category, key string) (json.RawMessage, error) {
	var setting settings.SiteSetting

	// 查找设置记录
	if err := dao.Db.Where("`category` = ? AND `key` = ?", category, key).First(&setting).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Setting not found for:", key)
			return nil, nil // 返回空值和nil表示找不到记录
		}
		return nil, err // 其他错误直接返回
	}

	log.Println("Setting retrieved successfully:", key)
	return setting.Value, nil
}

// saveSettingsWithReflection 使用反射来遍历结构体并保存字段
func saveSettingsWithReflection(category string, settingsStruct interface{}) {
	v := reflect.ValueOf(settingsStruct)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		key := t.Field(i).Tag.Get("json")
		if key == "" {
			key = t.Field(i).Name // 如果没有json tag，则使用字段名
		}

		// 将字段值转为 JSON
		valueJSON, err := json.Marshal(field.Interface())
		if err != nil {
			log.Println("Failed to marshal value:", err)
			continue
		}

		// 调用保存到数据库的函数
		if err := saveSettingToDB(category, key, valueJSON); err != nil {
			log.Println("Error saving settings to DB:", err)
		}
	}
}

func readMultipleSettingsFromDB() (*AppSettings, error) {
	var appSettings AppSettings

	// 查询 "app_name"
	if value, err := readSettingFromDB("site", "app_name"); err != nil {
		return nil, err
	} else if value != nil {
		//value = json.Unmarshal(value)
		appSettings.AppName = UnmarshalSingle(value)
	}

	// 查询 "app_description"
	if value, err := readSettingFromDB("site", "app_description"); err != nil {
		return nil, err
	} else if value != nil {
		appSettings.AppDescription = UnmarshalSingle(value)
	}

	// 查询 "tos_url"
	if value, err := readSettingFromDB("site", "tos_url"); err != nil {
		return nil, err
	} else if value != nil {
		appSettings.TOSURL = UnmarshalSingle(value)
	}

	// 查询 "frontend_background_url"
	if value, err := readSettingFromDB("frontend", "frontend_background_url"); err != nil {
		return nil, err
	} else if value != nil {
		appSettings.FrontendBackgroundURL = UnmarshalSingle(value)
	}

	// 查询 "frontend_theme"
	if value, err := readSettingFromDB("frontend", "frontend_theme"); err != nil {
		return nil, err
	} else if value != nil {
		appSettings.FrontendTheme = UnmarshalSingle(value)
	}

	return &appSettings, nil
}
