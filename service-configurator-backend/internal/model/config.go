package model

type MysqlConfig struct {
	Protocol string `yaml:"protocol" json:"protocol"`
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Database string `yaml:"database" json:"database"`
}

type RedisConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     int32  `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Database int    `yaml:"database" json:"database"`
}

type EtcdConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type MqConfig struct {
	Host       string `yaml:"host" json:"host"`
	Port       string `yaml:"port" json:"port"`
	Username   string `yaml:"username" json:"username"`
	Password   string `yaml:"password" json:"password"`
	SwitchName string `yaml:"switch_name" json:"switch_name"`
}

type RpcConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	MySql MysqlConfig `yaml:"mysql"`
	Redis RedisConfig `yaml:"redis"`
	Etcd  EtcdConfig  `yaml:"etcd"`
	Rpc   RpcConfig   `yaml:"rpc"`
}
