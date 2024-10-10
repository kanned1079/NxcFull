package services

import (
	pb "NxcFull/backend/systemServices/api/proto"
	"NxcFull/backend/systemServices/internal/dao"
	"NxcFull/nxc-backend/settings"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

func (s *SettingServices) GetSystemSettings(context context.Context, request *pb.GetSystemSettingsRequest) (*pb.GetSystemSettingsResponse, error) {
	var settingOptions []settings.SiteSetting

	// 从数据库中读取所有的系统设置
	if err := dao.Db.Find(&settingOptions).Error; err != nil {
		log.Println("Error fetching settings:", err)
		//context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch settings"})
		return &pb.GetSystemSettingsResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failed to fetch settings" + err.Error(),
		}, nil
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
	jsonSettings, err := json.Marshal(settingsMap)
	if err != nil {
		return &pb.GetSystemSettingsResponse{
			Code: http.StatusInternalServerError,
			Msg:  "序列化失败: " + err.Error(),
		}, nil
	}
	log.Println(string(jsonSettings))

	return &pb.GetSystemSettingsResponse{
		Code: http.StatusOK,
		Msg:  "Query ok",
		//Settings: settingsMap,
		Settings: jsonSettings,
	}, nil
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
