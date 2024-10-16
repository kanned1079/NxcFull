package config

import (
	"encoding/xml"
	"gateway/internal/etcd"
	"gateway/internal/model"
	"log"
)

var MyServerConfig model.ServerConfig

func GetServerEnv() error {
	var err error
	str, err := etcd.GetConfigFromEtcd("apiServer")
	if err != nil {
		return err
	}
	if err := xml.Unmarshal([]byte(str), &MyServerConfig); err != nil {
		return err
	}
	log.Println(MyServerConfig)
	return nil
}
