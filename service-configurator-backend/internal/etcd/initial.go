package etcd

import (
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

var (
	EtcdCli *clientv3.Client
	err     error
)

func InitEtcdCli() error {
	EtcdCli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"api.ikanned.com:22379"},
		DialTimeout: 5 * time.Second,
		Username:    "config",
		Password:    "Passcode1!",
	})
	if err != nil {
		log.Println("Error creating etcd client", err)
		return err
	}
	//defer EtcdCli.Close()

	return nil
}
