package etcd

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

// RegisterService2Etcd 在etcd中注册该服务
func RegisterService2Etcd(serverAddr string, ttl int64, serviceName, serviceAddr string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{serverAddr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println("Error creating etcd client", err)
		return
	}
	defer func() {
		if err := cli.Close(); err != nil {
			//return err
			log.Println("Error closing etcd client", err)
		}
	}()
	// 创建租约
	leaseResp, err := cli.Grant(context.TODO(), ttl)
	if err != nil {
		log.Println("Error creating etcd lease", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//put, err := cli.Put(ctx, serviceName, serviceAddr)
	put, err := cli.Put(ctx, "/services/"+serviceName, serviceAddr, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		log.Println("Error registering service", err)
	}
	log.Println("Put operation completed. Revision:", put.Header.Revision)

	// 启动续约操作
	ch, err := cli.KeepAlive(context.TODO(), leaseResp.ID)
	if err != nil {
		log.Println("Error keeping lease alive", err)
	}
	// 监听续约相应
	go func() {
		for {
			kaResp, ok := <-ch
			if !ok {
				log.Println("Lease keep-alive channel closed, service may be expired")
				return
			}
			log.Printf("Lease renewed successfully for leaseID: %v\n", kaResp.ID)
		}
	}()
	// 保持服务运行
	select {}
}
