package handler

import (
	pb "NxcFull/backend/gateway/internal/grpc/api/settings/proto"
	sysContext "context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 辅助函数：将 Go 的 `any` 类型转换为 gRPC 的 `*pb.Value`
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

func HandleUpdateSingleOptions(context *gin.Context) {
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
	//category := req.Category
	//key := req.Key

	//val, err := anyToProtoValue(req.Value)
	//log.Println(req)
	//if err != nil {
	//	context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data", "details": err.Error()})
	//}
	jsonVal, err := json.Marshal(req.Value)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "序列化失败" + err.Error(),
		})
	}
	log.Println(string(jsonVal))
	resp, err := grpcClient.SettingServiceClient.UpdateSingleOption(sysContext.Background(), &pb.UpdateSingleOptionRequest{
		Category: req.Category,
		Key:      req.Key,
		Value:    jsonVal,
	})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

func HandleGetSystemSetting(context *gin.Context) {
	resp, err := grpcClient.SettingServiceClient.GetSystemSettings(sysContext.Background(), &pb.GetSystemSettingsRequest{})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	//log.Println(string(resp.Settings)) // 此处json格式正确
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	// 反序列化 resp.Settings 到 Go 的 map 或其他结构
	var settings map[string]interface{}
	if err := json.Unmarshal(resp.Settings, &settings); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to parse settings",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"msg":      resp.Msg,
		"settings": settings, // 此处在浏览器返回不正确
	})
}
