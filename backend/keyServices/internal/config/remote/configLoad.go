package remote

import (
	"encoding/xml"
	"keyServices/internal/etcd"
)

var (
	MyDbConfig    MysqlConfig
	MyRedisConfig RedisConfig
	MyMqConfig    MqConfig
	err           error
)

func (m *MysqlConfig) Get() error {
	var tempData []byte
	if tempData, err = etcd.GetConfigFromEtcdBytes("mysql"); err != nil {
		return err
	}
	if err := xml.Unmarshal(tempData, m); err != nil {
		return err
	}
	return nil
}

func (r *RedisConfig) Get() error {
	var tempData []byte
	if tempData, err = etcd.GetConfigFromEtcdBytes("redis"); err != nil {
		return err
	}
	if err := xml.Unmarshal(tempData, r); err != nil {
		return err
	}
	return nil
}

func (r *MqConfig) Get() error {
	var tempData []byte
	if tempData, err = etcd.GetConfigFromEtcdBytes("mq"); err != nil {
		return err
	}
	if err := xml.Unmarshal(tempData, r); err != nil {
		return err
	}
	return nil
}
