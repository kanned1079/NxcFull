package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	pb "systemServices/api/proto"
	"systemServices/internal/dao"
	"systemServices/internal/model"
	"time"
)

type AppSettings struct {
	AppName               string `json:"app_name"`
	AppDescription        string `json:"app_description"`
	TOSURL                string `json:"tos_url"`
	FrontendBackgroundURL string `json:"frontend_background_url"`
	FrontendTheme         string `json:"frontend_theme"`
}

type SettingServices struct {
	pb.UnimplementedSettingsServiceServer
}

func NewSettingServices() *SettingServices {
	return &SettingServices{}
}

// UpdateSingleOption 更新单个设置的键值
func (s *SettingServices) UpdateSingleOption(context context.Context, request *pb.UpdateSingleOptionRequest) (*pb.UpdateSingleOptionResponse, error) {
	// 保存或更新设置
	if err := saveSettingToDB(request.Category, request.Key, request.Value); err != nil {
		return &pb.UpdateSingleOptionResponse{
			Code: http.StatusInternalServerError,
			Msg:  "保存或更新设置失败" + err.Error(),
		}, nil
	}

	//context.JSON(http.StatusOK, gin.H{"message": "Setting updated successfully"})
	return &pb.UpdateSingleOptionResponse{
		Code: http.StatusOK,
		Msg:  "配置项更新成功",
	}, nil
}

//func (s *SettingServices) GetSystemSettings(context context.Context, request *pb.GetSystemSettingsRequest) (*pb.GetSystemSettingsResponse, error) {
//	var settingOptions []model.SiteSetting
//
//	// 从数据库中读取所有的系统设置
//	if err := dao.Db.Find(&settingOptions).Error; err != nil {
//		log.Println("Error fetching settings:", err)
//		//context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch settings"})
//		return &pb.GetSystemSettingsResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "Failed to fetch settings" + err.Error(),
//		}, nil
//	}
//
//	// 创建一个map来组织数据
//	settingsMap := make(map[string]map[string]any)
//	for _, setting := range settingOptions {
//		if _, exists := settingsMap[setting.Category]; !exists {
//			settingsMap[setting.Category] = make(map[string]any)
//		}
//
//		var value any
//		if err := json.Unmarshal(setting.Value, &value); err != nil {
//			log.Println("Error unmarshaling settings value:", err)
//			continue
//		}
//		settingsMap[setting.Category][setting.Key] = value
//	}
//
//	// 返回组织好的数据
//	jsonSettings, err := json.Marshal(settingsMap)
//	if err != nil {
//		return &pb.GetSystemSettingsResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "序列化失败: " + err.Error(),
//		}, nil
//	}
//	log.Println(string(jsonSettings))
//
//	return &pb.GetSystemSettingsResponse{
//		Code: http.StatusOK,
//		Msg:  "Query ok",
//		//Settings: settingsMap,
//		Settings: jsonSettings,
//	}, nil
//}

//func (s *SettingServices) GetSystemSettings(rpcCtx context.Context, request *pb.GetSystemSettingsRequest) (*pb.GetSystemSettingsResponse, error) {
//	redisKey := "system:settings"
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	// 尝试从 Redis 获取缓存的系统设置
//	cachedSettings, err := dao.Rdb.Get(ctx, redisKey).Result()
//	if err == nil {
//		// 如果缓存存在，直接解析返回
//		log.Println("Returning system settings from Redis cache.")
//		return &pb.GetSystemSettingsResponse{
//			Code:     http.StatusOK,
//			Msg:      "Query ok",
//			Settings: []byte(cachedSettings),
//		}, nil
//	} else if err != redis.Nil {
//		// 如果 Redis 出现其他错误，记录日志但继续从数据库中读取
//		log.Printf("Error fetching settings from Redis: %v\n", err)
//	}
//
//	// Redis 缓存未命中，读取数据库
//	var settingOptions []model.SiteSetting
//	if err := dao.Db.Find(&settingOptions).Error; err != nil {
//		log.Println("Error fetching settings from DB:", err)
//		return &pb.GetSystemSettingsResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "Failed to fetch settings: " + err.Error(),
//		}, nil
//	}
//
//	// 创建一个 map 来组织数据
//	settingsMap := make(map[string]map[string]any)
//	for _, setting := range settingOptions {
//		if _, exists := settingsMap[setting.Category]; !exists {
//			settingsMap[setting.Category] = make(map[string]any)
//		}
//
//		var value any
//		if err := json.Unmarshal(setting.Value, &value); err != nil {
//			log.Println("Error unmarshaling setting value:", err)
//			continue
//		}
//		settingsMap[setting.Category][setting.Key] = value
//	}
//
//	// 将组织好的数据序列化
//	jsonSettings, err := json.Marshal(settingsMap)
//	if err != nil {
//		return &pb.GetSystemSettingsResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "Serialization failed: " + err.Error(),
//		}, nil
//	}
//
//	// 将序列化后的数据存入 Redis
//	if err := dao.Rdb.Set(ctx, redisKey, jsonSettings, 10*time.Minute).Err(); err != nil {
//		log.Printf("Failed to cache settings in Redis: %v\n", err)
//	}
//
//	// 返回结果
//	log.Println("Returning system settings from database.")
//	return &pb.GetSystemSettingsResponse{
//		Code:     http.StatusOK,
//		Msg:      "Query ok",
//		Settings: jsonSettings,
//	}, nil
//}

func (s *SettingServices) GetSystemSettings(rpcCtx context.Context, request *pb.GetSystemSettingsRequest) (*pb.GetSystemSettingsResponse, error) {
	settingsMap := make(map[string]map[string]any)
	categories := []string{"frontend", "my_app", "notice", "security", "sendmail", "server", "site", "invite"}

	// 遍历每个类别从 Redis 缓存中获取设置
	for _, category := range categories {
		redisKey := fmt.Sprintf("settings:%s", category)
		// 使用 HGetALL 获取整个类别的键值对
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		entries, err := dao.Rdb.HGetAll(ctx, redisKey).Result()
		if err == nil && len(entries) > 0 {
			// 如果 Redis 存在缓存数据，解析并填充 settingsMap
			settingsMap[category] = make(map[string]any)
			for key, value := range entries {
				var parsedValue any
				if err := json.Unmarshal([]byte(value), &parsedValue); err != nil {
					log.Printf("Error unmarshalling value for key '%s': %v\n", key, err)
					continue
				}
				settingsMap[category][key] = parsedValue
			}
			log.Printf("Category '%s' loaded from Redis cache.\n", category)
		} else if err != nil && err != redis.Nil {
			// 如果发生 Redis 错误（非空值错误），记录日志并继续数据库读取
			log.Printf("Redis error for category '%s': %v\n", category, err)
		}
	}

	// 检查是否所有类别都在缓存中
	if len(settingsMap) == len(categories) {
		// 全部类别均已从 Redis 缓存中获取，直接返回
		jsonSettings, err := json.Marshal(settingsMap)
		if err != nil {
			return &pb.GetSystemSettingsResponse{
				Code: http.StatusInternalServerError,
				Msg:  "Serialization failed: " + err.Error(),
			}, nil
		}

		log.Println("Returning settings from Redis cache.")
		return &pb.GetSystemSettingsResponse{
			Code:     http.StatusOK,
			Msg:      "Query ok",
			Settings: jsonSettings,
		}, nil
	}

	// Redis 未完全命中，读取数据库
	var settingOptions []model.SiteSetting
	if err := dao.Db.Find(&settingOptions).Error; err != nil {
		log.Println("Error fetching settings from DB:", err)
		return &pb.GetSystemSettingsResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to fetch settings: " + err.Error(),
		}, nil
	}

	go func() {
		if err := MakeSettingsCache(); err != nil {
			log.Println("Error making settings cache:", err)
		}
	}()

	// 返回组织好的数据
	jsonSettings, err := json.Marshal(settingsMap)
	if err != nil {
		return &pb.GetSystemSettingsResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Serialization failed: " + err.Error(),
		}, nil
	}

	log.Println("Returning settings from database and cached in Redis.")
	return &pb.GetSystemSettingsResponse{
		Code:     http.StatusOK,
		Msg:      "Query ok",
		Settings: jsonSettings,
	}, nil
}

// handleGetSystemSetting 取出所有设置项
func handleGetSystemSetting(context *gin.Context) {
	var settingOptions []model.SiteSetting

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
