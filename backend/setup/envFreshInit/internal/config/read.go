package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var LocalCfg Config

type Config struct {
	Db MysqlConfig `yaml:"db" json:"db"`
}

// Get 方法：从 ./config/db_credential.yaml 文件中读取数据并反序列化到 Config 结构体
func (c *Config) Get(filePath string) error {
	// 读取 YAML 文件内容
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("cannot read file %s: %v", filePath, err)
	}

	// 反序列化 YAML 数据到 Config 结构体
	if err := yaml.Unmarshal(data, c); err != nil {
		return fmt.Errorf("cannot resolve yaml file %v", err)
	}

	return nil
}

// Set 方法：将 Config 结构体序列化并写入 ./config/db_credential.yaml 文件
func (c *Config) Set(filePath string) error {
	// 序列化 Config 结构体为 YAML 数据
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("cannot marshal data: %v", err)
	}

	// 确保配置文件所在的目录存在，如果不存在则创建
	if err := os.MkdirAll("./config", os.ModePerm); err != nil {
		return fmt.Errorf("cannot create dir ./config: %v", err)
	}

	// 写入数据到文件
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("cannot write file to %s: %v", filePath, err)
	}

	return nil
}
