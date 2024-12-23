package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	pb "systemServices/api/proto"
	"time"
)

//// GetWelcomePageConfig 获取欢迎页面的配置信息
//func (s *SettingServices) GetWelcomePageConfig(ctx context.Context, request *pb.GetWelcomePageConfigRequest) (*pb.GetWelcomePageConfigResponse, error) {
//	// 定义欢迎页面需要读取的配置项
//	var configKeys = []string{
//		"app_sub_description",
//		"why_choose_us_hint",
//		"bilibili_official_link",
//		"youtube_official_link",
//		"instagram_link",
//		"wechat_official_link",
//		"filing_number",
//		"page_suffix",
//	}
//
//	// 用于存储最终的配置信息
//	config := make(map[string]interface{})
//	var failedKeys []string
//
//	// 遍历配置项并从 Redis 获取值
//	for _, key := range configKeys {
//		value, err := readSettingFromRedisCache("welcome", key)
//		if err != nil {
//			log.Printf("Failed to read key %s from Redis: %v", key, err)
//			failedKeys = append(failedKeys, key)
//			continue
//		}
//
//		// 假设 Redis 中的值是 JSON 格式，需要解析
//		var parsedValue interface{}
//		if err := json.Unmarshal(value, &parsedValue); err != nil {
//			log.Printf("Failed to unmarshal key %s value: %v", key, err)
//			failedKeys = append(failedKeys, key)
//			continue
//		}
//
//		// 存储到配置 map 中
//		config[key] = parsedValue
//	}
//
//	// 如果有失败的键，记录日志或处理
//	if len(failedKeys) > 0 {
//		log.Printf("Some keys failed to load: %v", failedKeys)
//	}
//
//	// 序列化最终的配置为 JSON
//	data, err := json.Marshal(config)
//	if err != nil {
//		return nil, fmt.Errorf("failed to marshal config: %w", err)
//	}
//
//	// 返回响应
//	return &pb.GetWelcomePageConfigResponse{
//		Code:   http.StatusOK,
//		Msg:    "success",
//		Config: data,
//	}, nil
//}

func (s *SettingServices) GetWelcomePageConfig(ctx context.Context, request *pb.GetWelcomePageConfigRequest) (*pb.GetWelcomePageConfigResponse, error) {
	// 加锁保护缓存
	s.welcomePageConfigCacheMux.Lock()
	defer s.welcomePageConfigCacheMux.Unlock()

	// 检查缓存是否存在且未过期
	if s.welcomePageConfigCache != nil && time.Since(s.welcomePageConfigCache.LastSaved) < s.welcomePageConfigCacheTTL {
		// 返回缓存中的配置
		data, err := json.Marshal(s.welcomePageConfigCache.Config)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal cached config: %w", err)
		}
		return &pb.GetWelcomePageConfigResponse{
			Code:   http.StatusOK,
			Msg:    "success (from cache)",
			Config: data,
		}, nil
	}

	// 定义欢迎页面需要读取的配置项
	var configKeys = []string{
		"app_sub_description",
		"why_choose_us_hint",
		"bilibili_official_link",
		"youtube_official_link",
		"instagram_link",
		"wechat_official_link",
		"filing_number",
		"page_suffix",
	}

	// 用于存储最终的配置信息
	config := make(map[string]interface{})
	var failedKeys []string

	// 遍历配置项并从 Redis 获取值
	for _, key := range configKeys {
		value, err := readSettingFromRedisCache("welcome", key)
		if err != nil {
			log.Printf("Failed to read key %s from Redis: %v", key, err)
			failedKeys = append(failedKeys, key)
			continue
		}

		// 假设 Redis 中的值是 JSON 格式，需要解析
		var parsedValue interface{}
		if err := json.Unmarshal(value, &parsedValue); err != nil {
			log.Printf("Failed to unmarshal key %s value: %v", key, err)
			failedKeys = append(failedKeys, key)
			continue
		}

		// 存储到配置 map 中
		config[key] = parsedValue
	}

	// 如果有失败的键，记录日志或处理
	if len(failedKeys) > 0 {
		log.Printf("Some keys failed to load: %v", failedKeys)
	}

	// 更新缓存
	s.welcomePageConfigCache = &BasicSetting{
		Config:    config,
		LastSaved: time.Now(),
	}

	// 序列化最终的配置为 JSON
	data, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal config: %w", err)
	}

	// 返回响应
	return &pb.GetWelcomePageConfigResponse{
		Code:   http.StatusOK,
		Msg:    "success (from Redis)",
		Config: data,
	}, nil
}
