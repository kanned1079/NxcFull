package etcd

import (
	"context"
	"fmt"
	"log"
	"time"
)

// GetConfigFromEtcd 用于从 etcd 中获取一个配置
func GetConfigFromEtcd(configName string) (string, error) {
	// 设置 5 秒超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 从 etcd 中获取指定的配置
	getResp, err := EtcdCli.Get(ctx, "/config/"+configName) // 此处已经有前缀
	if err != nil {
		log.Printf("获取设置 %s 失败: %v\n", configName, err)
		return "", err
	}

	// 检查是否找到值
	if len(getResp.Kvs) > 0 {
		configValue := string(getResp.Kvs[0].Value)
		//log.Printf("成功获取设置 %s, 值: %s\n", configName, configValue)
		return configValue, nil
	}

	log.Printf("未找到设置 %s\n", configName)
	return "", fmt.Errorf("config %s not found", configName)
}

func GetConfigFromEtcdBytes(configName string) ([]byte, error) {
	// 设置 5 秒超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 从 etcd 中获取指定的配置
	getResp, err := EtcdCli.Get(ctx, "/config/"+configName) // 此处已经有前缀
	if err != nil {
		log.Printf("获取设置 %s 失败: %v\n", configName, err)
		return []byte(""), err
	}

	// 检查是否找到值
	if len(getResp.Kvs) > 0 {
		//configValue := string(getResp.Kvs[0].Value)
		//log.Printf("成功获取设置 %s, 值: %s\n", configName, configValue)
		return getResp.Kvs[0].Value, nil
	}

	log.Printf("未找到设置 %s\n", configName)
	return []byte(""), fmt.Errorf("config %s not found", configName)
}
