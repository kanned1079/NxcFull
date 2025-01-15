package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"log"
	"logsServices/internal/config/local"
	"time"
)

var (
	// EtcdCli etcd客户端
	EtcdCli *clientv3.Client
	err     error
)

// RegisterService2Etcd 在etcd中注册该服务
func RegisterService2Etcd(ttl int64) {
	//cli, err := clientv3.New(clientv3.Config{
	//	Endpoints:   []string{serverAddr},
	//	DialTimeout: 5 * time.Second,
	//	Username:    "services",
	//	Password:    "Passcode1!",
	//})
	//if err != nil {
	//	log.Println("Error creating etcd client", err)
	//	return
	//}

	//if EtcdCli == nil {
	//	if err := InitEtcdClient(); err != nil {
	//		log.Println("Error init etcd client", err)
	//		panic(err)
	//	}
	//}

	defer func() {
		if err := EtcdCli.Close(); err != nil {
			//return err
			log.Println("Error closing etcd client", err)
		}
	}()

	// 创建租约
	leaseResp, err := EtcdCli.Grant(context.TODO(), ttl)
	if err != nil {
		log.Println("Error creating etcd lease", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//put, err := cli.Put(ctx, serviceName, serviceAddr)
	put, err := EtcdCli.Put(ctx,
		"/services/"+local.MyLocalConfig.RpcConfig.RpcName,
		fmt.Sprintf("%s:%d",
			local.MyLocalConfig.RpcConfig.ListenAddr,
			local.MyLocalConfig.RpcConfig.ListenPort),
		clientv3.WithLease(leaseResp.ID),
	)
	if err != nil {
		log.Println("Error registering services", err)
	}
	log.Println("Put operation completed. Revision:", put.Header.Revision)

	// 启动续约操作
	ch, err := EtcdCli.KeepAlive(context.TODO(), leaseResp.ID)
	if err != nil {
		log.Println("Error keeping lease alive", err)
	}
	// 监听续约相应
	go func() {
		for {
			kaResp, ok := <-ch
			if !ok {
				log.Println("Lease keep-alive channel closed, services may be expired")
				return
			}
			log.Printf("Lease renewed successfully for leaseID: %v\n", kaResp.ID)
		}
	}()
	// 保持服务运行
	select {}
}
