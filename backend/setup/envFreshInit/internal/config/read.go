package config

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strconv"
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

func (c *Config) GetFromArgs() error {
	reader := bufio.NewReader(os.Stdin)

	// 获取 Db.Host，若为空使用默认值
	fmt.Print("Enter Database Host (default 'localhost'): ")
	hostInput, _ := reader.ReadString('\n')
	c.Db.Host = hostInput[:len(hostInput)-1] // 去掉末尾的换行符
	if c.Db.Host == "" {
		c.Db.Host = "localhost" // 设置默认值
	}

	// 获取 Db.Username
	fmt.Print("Enter Database Username: ")
	usernameInput, _ := reader.ReadString('\n')
	c.Db.Username = usernameInput[:len(usernameInput)-1] // 去掉末尾的换行符

	// 获取 Db.Password
	fmt.Print("Enter Database Password: ")
	passwordInput, _ := reader.ReadString('\n')
	c.Db.Password = passwordInput[:len(passwordInput)-1] // 去掉末尾的换行符

	// 获取 Db.Port，若为空使用默认值
	fmt.Print("Enter Database Port (default '3306'): ")
	portInput, _ := reader.ReadString('\n')
	portInput = portInput[:len(portInput)-1] // 去掉末尾的换行符

	// 如果输入为空，使用默认值
	if portInput == "" {
		c.Db.Port = 3306
	} else {
		// 如果输入了内容，尝试将其转换为 int
		port, err := strconv.Atoi(portInput)
		if err != nil {
			return fmt.Errorf("invalid port number: %v", err)
		}
		c.Db.Port = port
	}

	return nil
}
