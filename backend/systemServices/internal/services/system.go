package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	pb "systemServices/api/proto"
	"systemServices/internal/dao"
	"systemServices/internal/model"
	"time"
)

func (s *SettingServices) GetSystemSettings(rpcCtx context.Context, request *pb.GetSystemSettingsRequest) (*pb.GetSystemSettingsResponse, error) {
	settingsMap := make(map[string]map[string]any)
	//categories := []string{"frontend", "my_app", "notice", "security", "sendmail", "server", "site", "invite"}
	categories := []string{"frontend", "my_app", "notice", "security", "sendmail", "welcome", "site", "invite"}

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

// UpdateSingleOption 更新单个设置的键值
func (s *SettingServices) UpdateSingleOption(context context.Context, request *pb.UpdateSingleOptionRequest) (*pb.UpdateSingleOptionResponse, error) {
	// 保存或更新设置
	if err := saveSettingToDB(request.Category, request.Key, request.Value); err != nil {
		return &pb.UpdateSingleOptionResponse{
			Code: http.StatusInternalServerError,
			Msg:  "保存或更新设置失败" + err.Error(),
		}, nil
	}
	if err = s.ClearAllCaches(context); err != nil {
		return &pb.UpdateSingleOptionResponse{
			Code: http.StatusAccepted,
			Msg:  "可能需要等待缓存失效" + err.Error(),
		}, nil
	}

	//context.JSON(http.StatusOK, gin.H{"message": "Setting updated successfully"})
	return &pb.UpdateSingleOptionResponse{
		Code: http.StatusOK,
		Msg:  "配置项更新成功",
	}, nil
}
