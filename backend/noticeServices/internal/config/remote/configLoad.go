package remote

import (
	"encoding/xml"
	"noticeServices/internal/etcd"
)

var (
	MyDbConfig    MysqlConfig
	MyRedisConfig RedisConfig
	err           error
)

func (m *MysqlConfig) Get() error {
	var tempData []byte
	//log.Println(5)

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
