package routers

import (
	"NxcFull/nxc-backend/dao"
	"NxcFull/nxc-backend/settings"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
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
		if err == gorm.ErrRecordNotFound {
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

// handleUpdateOptions 将单个键值保存或更新
func handleUpdateSingleOptions(context *gin.Context) {
	var req struct {
		Category string          `json:"category"`
		Key      string          `json:"key"`
		Value    json.RawMessage `json:"value"`
	}

	// 解析请求体
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data", "details": err.Error()})
		return
	}

	// 如果 key 不需要进行格式拆分，可以直接使用 req.Key 作为 key
	//category := "frontend" // 假设 category 是固定的 "frontend"，你可以根据需要进行调整
	category := req.Category
	key := req.Key

	// 保存或更新设置
	if err := saveSettingToDB(category, key, req.Value); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update settings",
			"details": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Setting updated successfully"})
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

func handleUpdateSystemSettings(context *gin.Context) {
	var options = settings.SystemSettingOptions{}
	if err := context.ShouldBind(&options); err != nil {
		log.Println("获取设置失败")
	}
	log.Println(options)

	// 使用反射保存所有字段
	saveSettingsWithReflection("site", options.Site)
	saveSettingsWithReflection("security", options.Security)
	saveSettingsWithReflection("frontend", options.Frontend)
	saveSettingsWithReflection("subscription", options.Subscribe)
	saveSettingsWithReflection("server", options.Server)
	saveSettingsWithReflection("sendmail", options.Sendmail)
	saveSettingsWithReflection("notice", options.Notice)
	saveSettingsWithReflection("myapp", options.Myapp)
	context.JSON(http.StatusOK, gin.H{"message": "Settings updated successfully"})
}

// handleGetSystemSetting 取出所有设置项
func handleGetSystemSetting(context *gin.Context) {
	var settingOptions []settings.SiteSetting

	// 从数据库中读取所有的系统设置
	if err := dao.Db.Find(&settingOptions).Error; err != nil {
		log.Println("Error fetching settings:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch settings"})
		return
	}

	// 创建一个map来组织数据
	settingsMap := make(map[string]map[string]any)
	for _, setting := range settingOptions {
		if _, exists := settingsMap[setting.Category]; !exists {
			settingsMap[setting.Category] = make(map[string]any)
		}

		var value any
		if err := json.Unmarshal(setting.Value, &value); err != nil {
			log.Println("Error unmarshaling settings value:", err)
			continue
		}
		settingsMap[setting.Category][setting.Key] = value
	}

	// 返回组织好的数据
	context.JSON(http.StatusOK, settingsMap)
}

type AppSettings struct {
	AppName               string `json:"app_name"`
	AppDescription        string `json:"app_description"`
	TOSURL                string `json:"tos_url"`
	FrontendBackgroundURL string `json:"frontend_background_url"`
	FrontendTheme         string `json:"frontend_theme"`
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

	//log.Println(appSettings)

	//json.Unmarshal(appSettings, appSettings)

	return &appSettings, nil
}

// getStartTheme 启动页面需要一些配置
func getStartTheme(context *gin.Context) {
	appSettings, err := readMultipleSettingsFromDB()
	if err != nil {
		log.Println(err.Error())
	}
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": appSettings,
	})
}
