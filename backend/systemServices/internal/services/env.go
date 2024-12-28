package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	pb "systemServices/api/proto"
	"time"
)

type BasicSetting struct {
	Config    map[string]any
	LastSaved time.Time
}

func (s *SettingServices) GetBasicRuntimeEnvConfig(ctx context.Context, request *pb.GetBasicRuntimeEnvConfigRequest) (*pb.GetBasicRuntimeEnvConfigResponse, error) {
	// 加锁保护缓存
	s.basicRuntimeEnvCacheMux.Lock()
	defer s.basicRuntimeEnvCacheMux.Unlock()

	// 检查缓存是否存在且未过期
	if s.basicRuntimeEnvCache != nil && time.Since(s.basicRuntimeEnvCache.LastSaved) < s.basicRuntimeEnvCacheTTL {
		// 返回缓存中的配置
		data, err := json.Marshal(s.basicRuntimeEnvCache.Config)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal cached config: %w", err)
		}
		return &pb.GetBasicRuntimeEnvConfigResponse{
			Code:   http.StatusOK,
			Msg:    "success (cache)",
			Config: data,
		}, nil
	}

	// 定义要读取的配置项
	siteKeys := []string{
		"app_name",
		"app_sub_name",
		"app_description",
		"app_url",
		"logo_url",
		"currency",
		"currency_symbol",
		"stop_register",
	}
	frontendKeys := []string{
		"frontend_theme",
		"frontend_background_url",
	}
	securityKeys := []string{
		"secure_path",
		"safe_mode_enable",
	}

	// 用于存储最终的配置
	config := make(map[string]interface{})

	// 从 settings.site 获取配置
	for _, key := range siteKeys {
		value, err := readSettingFromRedisCache("site", key)
		if err != nil {
			return nil, fmt.Errorf("failed to read settings.site.%s: %w", key, err)
		}
		// 假设值是 JSON 序列化的，反序列化后存入 config
		var parsedValue interface{}
		if err := json.Unmarshal(value, &parsedValue); err != nil {
			return nil, fmt.Errorf("failed to unmarshal settings.site.%s: %w", key, err)
		}
		config[key] = parsedValue
	}

	// 从 settings.frontend 获取配置
	for _, key := range frontendKeys {
		value, err := readSettingFromRedisCache("frontend", key)
		if err != nil {
			return nil, fmt.Errorf("failed to read settings.frontend.%s: %w", key, err)
		}
		// 假设值是 JSON 序列化的，反序列化后存入 config
		var parsedValue interface{}
		if err := json.Unmarshal(value, &parsedValue); err != nil {
			return nil, fmt.Errorf("failed to unmarshal settings.frontend.%s: %w", key, err)
		}
		config[key] = parsedValue
	}

	// 从 settings.security 获取配置
	for _, key := range securityKeys {
		value, err := readSettingFromRedisCache("security", key)
		if err != nil {
			return nil, fmt.Errorf("failed to read settings.frontend.%s: %w", key, err)
		}
		// 假设值是 JSON 序列化的，反序列化后存入 config
		var parsedValue interface{}
		if err := json.Unmarshal(value, &parsedValue); err != nil {
			return nil, fmt.Errorf("failed to unmarshal settings.frontend.%s: %w", key, err)
		}
		config[key] = parsedValue
	}

	// 更新缓存
	s.basicRuntimeEnvCache = &BasicSetting{
		Config:    config,
		LastSaved: time.Now(),
	}

	// 将最终的配置序列化为 JSON
	data, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal final config: %w", err)
	}

	// 返回 RPC 响应
	return &pb.GetBasicRuntimeEnvConfigResponse{
		Code:   http.StatusOK,
		Msg:    "success (redis)",
		Config: data,
	}, nil
}

func (s *SettingServices) GetRegisterEnvConfig(ctx context.Context, request *pb.GetRegisterEnvConfigRequest) (*pb.GetRegisterEnvConfigResponse, error) {

	//  string app_name = 3;
	//  string app_sub_name = 4;
	//  string app_description = 5;
	//  string app_url = 6;
	//  bool email_whitelist_suffix = 7;
	//  bool is_email_verify = 8;
	//  bool email_gmail_limit_enable = 9;
	//  bool is_invite_force = 10;
	//  bool is_recaptcha = 11;
	//  string logo = 12;
	//  string recaptcha_site_key = 13;
	//  string tos_url = 14;
	//  bool stop_register = 15;

	return &pb.GetRegisterEnvConfigResponse{
		Code: http.StatusOK,
		Msg:  "success",
	}, nil
}