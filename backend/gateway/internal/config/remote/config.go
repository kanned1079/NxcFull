package remote

//// MysqlConfig 数据库的配置项目 来自于Etcd
//type MysqlConfig struct {
//	Protocol string `yaml:"protocol" json:"protocol"`
//	Host     string `yaml:"host" json:"host"`
//	Port     int    `yaml:"port" json:"port"`
//	Username string `yaml:"username" json:"username"`
//	Password string `yaml:"password" json:"password"`
//	Database string `yaml:"database" json:"database"`
//}

// RedisConfig Redis的配置项目 来自于Etcd
type RedisConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Database int    `yaml:"database" json:"database"`
}
