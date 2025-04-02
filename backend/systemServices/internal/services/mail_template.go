package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	pb "systemServices/api/proto"
)

func (s *SettingServices) GetSendMailTemplateFillContent(ctx context.Context, request *pb.GetSendMailTemplateFillContentRequest) (*pb.GetSendMailTemplateFillContentResponse, error) {
	// 读取 site.app_name
	var appName string
	appNameRaw, err := readSettingFromDB("site", "app_name")
	if err != nil {
		log.Printf("Failed to fetch app_name: %v", err)
		return &pb.GetSendMailTemplateFillContentResponse{
			Code: http.StatusInternalServerError,
		}, err
	}
	if appNameRaw != nil {
		if err := json.Unmarshal(appNameRaw, &appName); err != nil {
			log.Printf("Failed to unmarshal app_name: %v", err)
			appName = "App" // 解析失败时使用默认值
		}
	} else {
		appName = "App" // 默认值
	}

	// 读取 site.app_url
	var appUrl string
	appUrlRaw, err := readSettingFromDB("site", "app_url")
	if err != nil {
		log.Printf("Failed to fetch app_url: %v", err)
		return &pb.GetSendMailTemplateFillContentResponse{
			Code: http.StatusInternalServerError,
		}, err
	}
	if appUrlRaw != nil {
		if err := json.Unmarshal(appUrlRaw, &appUrl); err != nil {
			log.Printf("Failed to unmarshal app_url: %v", err)
			appUrl = "https://example.com" // 解析失败时使用默认值
		}
	} else {
		appUrl = "https://example.com" // 默认值
	}

	return &pb.GetSendMailTemplateFillContentResponse{
		Code:    http.StatusOK,
		AppName: appName,
		AppUrl:  appUrl,
	}, nil
}

// 金坛金江苑
