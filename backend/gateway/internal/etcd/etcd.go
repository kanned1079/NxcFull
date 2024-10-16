package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sync"
	"time"
)

var etcdServersAddr = []string{"api.ikanned.com:22379"}

var (
	cli      *clientv3.Client
	srvCache = make(map[string]string)
	cacheMux sync.RWMutex
)

// InitEtcdClient 初始化 etcd 客户端
func InitEtcdClient() {
	var err error
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   etcdServersAddr,
		DialTimeout: 5 * time.Second,
		Username:    "services",
		Password:    "Passcode1!",
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	log.Println("Etcd初始化完成")
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
	if cli == nil {
		InitEtcdClient()
	}

	resp, err := cli.Get(ctx, "/services/"+serviceName)
	if err != nil || len(resp.Kvs) == 0 {
		//log.Printf("获取服务 %s 地址错误: %v", serviceName, err)
		panic(fmt.Sprintf("获取服务 %s 地址错误: %v", serviceName, err))
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
