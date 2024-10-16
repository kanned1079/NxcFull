package etcd

import (
	"couponServices/internal/config/local"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

// InitEtcdClient 初始化Etcd服务器
func InitEtcdClient() error {
	addr := fmt.Sprintf("%s:%d", local.MyLocalConfig.EtcdConfig.Host, local.MyLocalConfig.EtcdConfig.Port)
	EtcdCli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
		Username:    local.MyLocalConfig.EtcdConfig.Username,
		Password:    local.MyLocalConfig.EtcdConfig.Password,
	})
	if err != nil {
		log.Println("Error creating etcd client", err)
		return err
	}
	return nil
}
