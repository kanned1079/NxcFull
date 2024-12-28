package etcd

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	srvCache = make(map[string]string)
	cacheMux sync.RWMutex
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

// GetServiceAddress 获取服务地址，带缓存
func GetServiceAddress(serviceName string) string {
	// 首先检查缓存
	log.Println("GetServiceAddress", serviceName)
	cacheMux.RLock()
	if addr, ok := srvCache[serviceName]; ok {
		cacheMux.RUnlock()
		log.Println("缓存中存有地址:", addr)
		return addr
	}
	cacheMux.RUnlock()

	// 如果缓存中没有，则从 etcd 获取
	log.Println("从 etcd 获取")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 如果cli没有被初始化则进行初始化
	if EtcdCli == nil {
		if InitEtcdClient() != nil {
			log.Panicln("KeyServices init other services err: ", err)
		}
	}
	log.Println("获取地址GET")
	resp, err := EtcdCli.Get(ctx, "/services/"+serviceName)
	if err != nil || len(resp.Kvs) == 0 {
		//log.Printf("获取服务 %s 地址错误: %v", serviceName, err)
		log.Printf("\033[31m **START** 获取服务 %s 地址错误: %v **END**\033[0m", serviceName, err)
		//panic(fmt.Sprintf("获取服务 %s 地址错误: %v", serviceName, err))
		return ""
	}

	serviceAddress := string(resp.Kvs[0].Value)
	log.Printf("获取服务 %s 地址成功 地址为： %s\n", serviceName, serviceAddress)

	// 更新缓存
	cacheMux.Lock()
	srvCache[serviceName] = serviceAddress
	cacheMux.Unlock()

	return serviceAddress
}
