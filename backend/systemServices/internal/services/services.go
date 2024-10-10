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
	//val, err := valueToJSONRawMessage(request.Value)
	//if err != nil {
	//	return &pb.UpdateSingleOptionResponse{
	//		Code: http.StatusInternalServerError,
	//		Msg:  "转换json格式失败" + err.Error(),
	//	}, nil
	//}
	// 保存或更新设置
	if err := saveSettingToDB(request.Category, request.Key, request.Value); err != nil {
		//context.JSON(http.StatusInternalServerError, gin.H{
		//	"error":   "Failed to update setting",
		//	"details": err.Error(),
		//})
		//return
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
			log.Println("Error unmarshaling setting value:", err)
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

//func anyToProtoValue(val any) (*pb.Value, error) {
//	switch v := val.(type) {
//	case string:
//		return &pb.Value{Kind: &pb.Value_StringValue{StringValue: v}}, nil
//	case int:
//		return &pb.Value{Kind: &pb.Value_IntValue{IntValue: int64(v)}}, nil
//	case bool:
//		return &pb.Value{Kind: &pb.Value_BoolValue{BoolValue: v}}, nil
//	case float64:
//		return &pb.Value{Kind: &pb.Value_DoubleValue{DoubleValue: v}}, nil
//	default:
//		return nil, fmt.Errorf("unsupported type: %T", v)
//	}
//}

//// handleUpdateOptions 将单个键值保存或更新
//func handleUpdateSingleOptions(context *gin.Context) {
//	var req struct {
//		Category string          `json:"category"`
//		Key      string          `json:"key"`
//		Value    json.RawMessage `json:"value"`
//	}
//
//	// 解析请求体
//	if err := context.ShouldBindJSON(&req); err != nil {
//		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data", "details": err.Error()})
//		return
//	}
//
//	// 如果 key 不需要进行格式拆分，可以直接使用 req.Key 作为 key
//	//category := "frontend" // 假设 category 是固定的 "frontend"，你可以根据需要进行调整
//	category := req.Category
//	key := req.Key
//
//	// 保存或更新设置
//	if err := saveSettingToDB(category, key, req.Value); err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"error":   "Failed to update setting",
//			"details": err.Error(),
//		})
//		return
//	}
//
//	context.JSON(http.StatusOK, gin.H{"message": "Setting updated successfully"})
//}

// pass
//func handleUpdateSystemSettings(context *gin.Context) {
//	var options = settings.SystemSettingOptions{}
//	if err := context.ShouldBind(&options); err != nil {
//		log.Println("获取设置失败")
//	}
//	log.Println(options)
//
//	// 使用反射保存所有字段
//	saveSettingsWithReflection("site", options.Site)
//	saveSettingsWithReflection("security", options.Security)
//	saveSettingsWithReflection("frontend", options.Frontend)
//	saveSettingsWithReflection("subscription", options.Subscribe)
//	saveSettingsWithReflection("server", options.Server)
//	saveSettingsWithReflection("sendmail", options.Sendmail)
//	saveSettingsWithReflection("notice", options.Notice)
//	saveSettingsWithReflection("myapp", options.Myapp)
//	context.JSON(http.StatusOK, gin.H{"message": "Settings updated successfully"})
//}
// passed

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
			log.Println("Error unmarshaling setting value:", err)
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
