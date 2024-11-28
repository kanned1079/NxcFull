package handler

import (
	sysContext "context"
	"encoding/json"
	pb "gateway/internal/grpc/api/settings/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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
		"settings": settings,
	})
}

func HandleEditOrSavePaymentMethodBySystemName(context *gin.Context) {
	// 获取 system 参数
	systemName := context.Query("system")
	if systemName == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "system 参数不能为空",
		})
		return
	}

	// 解析请求体中的 JSON
	var body map[string]interface{}
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "无效的 JSON 请求体",
		})
		return
	}

	// 将 JSON 转换为字节数组
	configBytes, err := json.Marshal(body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "JSON 编码失败",
		})
		return
	}

	// 调用 gRPC 服务
	resp, err := grpcClient.SettingServiceClient.EditPaymentSettingsBySystemName(sysContext.Background(), &pb.EditPaymentSettingsBySystemNameRequest{
		System: systemName,
		Config: configBytes,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	// 返回响应
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

func HandleGetAllPaymentMethodKv(context *gin.Context) {
	resp, err := grpcClient.SettingServiceClient.GetAllPaymentSettingsKv(sysContext.Background(), &pb.GetAllPaymentSettingsKvRequest{})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	var confMap []map[string]interface{}
	if err := json.Unmarshal(resp.Config, &confMap); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
		"conf": confMap,
	})
}

func HandleGetPaymentMethodDetailsBySystemName(context *gin.Context) {
	system := context.Query("system")
	if system == "" {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Bad Request",
		})
		return
	}
	resp, err := grpcClient.SettingServiceClient.GetPaymentMethodDetailsByName(sysContext.Background(), &pb.GetPaymentMethodDetailsByNameRequest{
		System: system,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	var detailsMap map[string]interface{}
	if err := json.Unmarshal(resp.Details, &detailsMap); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"msg":     resp.Msg,
		"details": detailsMap,
	})
}

func HandleSwitchPaymentMethodEnableBySystemName(context *gin.Context) {
	postData := &struct {
		System string `json:"system"`
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := grpcClient.SettingServiceClient.EnablePaymentSettingBySystemName(sysContext.Background(), &pb.EnablePaymentSettingBySystemNameRequest{
		System: postData.System,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}
