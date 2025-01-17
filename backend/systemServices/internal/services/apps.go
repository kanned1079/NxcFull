package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	pb "systemServices/api/proto"
)

func (s *SettingServices) GetAppDownloadLink(ctx context.Context, request *pb.GetAppDownloadLinkRequest) (*pb.GetAppDownloadLinkResponse, error) {

	// 配置项列表，包含各平台的下载链接配置项
	var myApp = []string{
		"download_enabled", // 是否启用下载
		"win_download",     // Windows 下载链接
		"osx_download",     // macOS 下载链接
		"android_download", // Android 下载链接
	}

	// 存储下载链接的结构体，AppItem 用于存放每个平台的下载链接
	var downloadLinks []*pb.AppItem

	var downloadEnabled bool
	var failedKeys []string // 用于记录读取失败的键

	// 遍历每个配置项，读取 Redis 中的配置
	for _, key := range myApp {
		value, err := readSettingFromRedisCache("my_app", key)
		if err != nil {
			log.Printf("Failed to read key %s from Redis: %v", key, err)
			failedKeys = append(failedKeys, key)
			continue
		}

		// 处理配置项
		switch key {
		case "download_enabled":
			// 解析是否启用下载的配置
			if err := json.Unmarshal(value, &downloadEnabled); err != nil {
				log.Printf("Failed to unmarshal key %s value: %v", key, err)
				failedKeys = append(failedKeys, key)
			}
		case "win_download", "osx_download", "android_download":
			// 处理下载链接配置项
			var platform string
			var downloadLink string

			if err := json.Unmarshal(value, &downloadLink); err != nil {
				log.Printf("Failed to unmarshal key %s value: %v", key, err)
				failedKeys = append(failedKeys, key)
				continue
			}

			// 根据当前 key 添加平台信息
			platformMap := map[string]string{
				"win_download":     "Windows",
				"osx_download":     "macOS",
				"android_download": "Android",
			}

			platform, ok := platformMap[key]
			if !ok {
				// 如果 key 不在 map 中，说明是一个无效的键
				log.Printf("Invalid platform key: %s", key)
				continue
			}

			// 将获取到的下载链接信息添加到下载链接列表中
			downloadLinks = append(downloadLinks, &pb.AppItem{
				Platform: platform,
				Link:     downloadLink,
			})
		}
	}

	// 如果有读取失败的键，记录日志并返回错误
	if len(failedKeys) > 0 {
		log.Printf("Some keys failed to load: %v", failedKeys)
	}

	// 返回成功响应
	return &pb.GetAppDownloadLinkResponse{
		Code:            http.StatusOK,
		Msg:             "success",
		DownloadEnabled: downloadEnabled, // 是否启用下载
		DownloadLinks:   downloadLinks,   // 下载链接列表
	}, nil
}
