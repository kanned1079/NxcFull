package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"systemServices/internal/dao"
	settings "systemServices/internal/model"
	"time"
)

//func UnmarshalSingle(val json.RawMessage) (result string) {
//	if err := json.Unmarshal(val, &result); err != nil {
//		log.Println(err)
//	}
//	return
//}

func UnmarshalSingle(val json.RawMessage) (string, error) {
	var result string
	if err := json.Unmarshal(val, &result); err != nil {
		log.Println("Failed to unmarshal JSON:", err)
		return "", err
	}
	return result, nil
}

// saveSettingToDB 将单个设置保存到数据库
//func saveSettingToDB(category, key string, value json.RawMessage) error {
//	setting := settings.SiteSetting{
//		Category: category,
//		Key:      key,
//	}
//
//	// 查找现有记录
//	if err := dao.Db.Where("`category` = ? AND `key` = ?", category, key).First(&setting).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			// 如果没有找到记录，则创建新记录
//			setting.Value = value
//			if err := dao.Db.Create(&setting).Error; err != nil {
//				return err
//			}
//		} else {
//			return err
//		}
//	} else {
//		// 如果找到记录，则更新值
//		setting.Value = value
//		if err := dao.Db.Save(&setting).Error; err != nil {
//			return err
//		}
//	}
//
//	log.Println("Setting created or updated successfully:", key)
//	return nil
//}

// saveSettingToDB 保存单个键值到数据库 *并刷新redis缓存
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

	// 数据库保存成功后，刷新 Redis 缓存
	log.Println("刷新redis缓存")
	if err := FlushSettingsCache(category, key, value); err != nil {
		log.Printf("Failed to refresh Redis cache for category '%s' and key '%s': %v", category, key, err)
		return err
	}

	log.Println("Setting created or updated successfully:", key)
	return nil
}

// readSettingFromDB 从数据库中取出一条记录
//func readSettingFromDB(category, key string) (json.RawMessage, error) {
//	var setting settings.SiteSetting
//
//	// 查找设置记录
//	if err := dao.Db.Where("`category` = ? AND `key` = ?", category, key).First(&setting).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			log.Println("Setting not found for:", key)
//			return nil, nil // 返回空值和nil表示找不到记录
//		}
//		return nil, err // 其他错误直接返回
//	}
//
//	log.Println("Setting retrieved successfully:", key)
//	return setting.Value, nil
//}

// readSettingFromDB 从数据库中取出一条记录 *先检索redis
func readSettingFromDB(category, key string) (json.RawMessage, error) {
	// 先从 Redis 缓存中读取数据
	value, err := readSettingFromRedisCache(category, key)
	if err != nil {
		return nil, err
	}

	// 如果 Redis 中找到缓存，则直接返回
	if value != nil {
		log.Println("Setting retrieved from Redis cache:", key, ", skipped query db.")
		return value, nil
	}
	log.Println("Setting not existed in redis, query db.")

	// Redis 中未找到，查询数据库
	var setting settings.SiteSetting
	if err := dao.Db.Where("`category` = ? AND `key` = ?", category, key).First(&setting).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Setting not found in database for:", key)
			return nil, nil // 返回空值和nil表示找不到记录
		}
		return nil, err // 其他错误直接返回
	}

	// 将数据库中查到的数据缓存到 Redis
	if err := FlushSettingsCache(category, key, setting.Value); err != nil {
		log.Printf("Failed to cache setting in Redis for category '%s' and key '%s': %v", category, key, err)
	}

	log.Println("Setting retrieved from database and cached in Redis:", key)
	return setting.Value, nil
}

// MakeSettingsCache 在init()的时候将数据库x_site_settings中所有字段保存进redis以提供缓存查询
func MakeSettingsCache() error {
	var items []settings.SiteSetting

	// 从数据库中读取所有设置项
	if err := dao.Db.Find(&items).Error; err != nil {
		return err
	}

	// 将设置项写入 Redis
	for _, setting := range items {
		redisKey := fmt.Sprintf("settings:%s", setting.Category)

		// 将 json.RawMessage 转换为 string
		valueStr := string(setting.Value)

		// 使用哈希类型存储，字段名为设置项的 Key，值为 JSON 格式的 Value
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		//defer cancel()
		if err := dao.Rdb.HSet(ctx, redisKey, setting.Key, valueStr).Err(); err != nil {
			cancel() // 直接结束
			return err
		}
		cancel() // 每次for循环结束后直接关闭ctx
	}

	log.Println("All settings cached successfully in Redis.")
	return nil
}

// FlushSettingsCache 刷新一个redis缓存
func FlushSettingsCache(category, key string, value json.RawMessage) error {
	redisKey := fmt.Sprintf("settings:%s", category)

	// 将 json.RawMessage 转换为 string
	valueStr := string(value)

	// 更新 Redis 中对应的 Key 值
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := dao.Rdb.HSet(ctx, redisKey, key, valueStr).Err(); err != nil {
		return err
	}

	log.Printf("Setting for category '%s' and key '%s' has been refreshed in Redis.", category, key)
	return nil
}

// readSettingFromRedisCache 从redis中读取一个配置项目
func readSettingFromRedisCache(category, key string) (json.RawMessage, error) {
	redisKey := fmt.Sprintf("settings:%s", category)

	// 从 Redis 中获取指定 Key 的值
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	valueStr, err := dao.Rdb.HGet(ctx, redisKey, key).Result()
	if errors.Is(err, redis.Nil) {
		log.Printf("Setting not found in Redis cache for category '%s' and key '%s'.", category, key)
		return nil, nil // 未找到记录
	} else if err != nil {
		return nil, err // 其他错误
	}

	// 将字符串转换为 json.RawMessage
	// 直接使用 `json.RawMessage(valueStr)` 将 string 转换为 RawMessage
	return json.RawMessage(valueStr), nil
}

// saveSettingsWithReflection 使用反射来遍历结构体并保存字段
//func saveSettingsWithReflection(category string, settingsStruct interface{}) {
//	v := reflect.ValueOf(settingsStruct)
//	t := v.Type()
//
//	for i := 0; i < v.NumField(); i++ {
//		field := v.Field(i)
//		key := t.Field(i).Tag.Get("json")
//		if key == "" {
//			key = t.Field(i).Name // 如果没有json tag，则使用字段名
//		}
//
//		// 将字段值转为 JSON
//		valueJSON, err := json.Marshal(field.Interface())
//		if err != nil {
//			log.Println("Failed to marshal value:", err)
//			continue
//		}
//
//		// 调用保存到数据库的函数
//		if err := saveSettingToDB(category, key, valueJSON); err != nil {
//			log.Println("Error saving settings to DB:", err)
//		}
//	}
//}

func readMultipleSettingsFromDB() (*AppSettings, error) {
	var appSettings AppSettings
	var err error
	var value json.RawMessage
	// 查询 "app_name"
	if value, err = readSettingFromDB("site", "app_name"); err != nil {
		return nil, err
	} else if value != nil {
		//value = json.Unmarshal(value)
		if appSettings.AppName, err = UnmarshalSingle(value); err != nil {
			log.Println(err.Error())
		}
	}

	// 查询 "app_description"
	if value, err := readSettingFromDB("site", "app_description"); err != nil {
		return nil, err
	} else if value != nil {
		if appSettings.AppDescription, err = UnmarshalSingle(value); err != nil {
			log.Println(err.Error())
		}
	}

	// 查询 "tos_url"
	if value, err := readSettingFromDB("site", "tos_url"); err != nil {
		return nil, err
	} else if value != nil {
		if appSettings.TOSURL, err = UnmarshalSingle(value); err != nil {
			log.Println(err.Error())
		}
	}

	// 查询 "frontend_background_url"
	if value, err := readSettingFromDB("frontend", "frontend_background_url"); err != nil {
		return nil, err
	} else if value != nil {
		if appSettings.FrontendBackgroundURL, err = UnmarshalSingle(value); err != nil {
			log.Println(err.Error())
		}
	}

	// 查询 "frontend_theme"
	if value, err := readSettingFromDB("frontend", "frontend_theme"); err != nil {
		return nil, err
	} else if value != nil {
		if appSettings.FrontendTheme, err = UnmarshalSingle(value); err != nil {
			log.Println(err.Error())
		}
	}

	return &appSettings, nil
}
