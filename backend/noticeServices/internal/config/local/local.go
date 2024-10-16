package local

// 这里的配置来自于本地的xml配置文件

type LocalConfig struct {
	RpcConfig  RpcConfig  `yaml:"rpc_config" xml:"rpc_config"`
	EtcdConfig EtcdConfig `yaml:"etcd_config" xml:"etcd_config"`
}

type EtcdConfig struct {
	Host     string `yaml:"host" xml:"host"`
	Port     int    `yaml:"port" xml:"port"`
	Username string `yaml:"username" xml:"username"`
	Password string `yaml:"password" xml:"password"`
}

type RpcConfig struct {
	RpcName    string `yaml:"rpc_name" xml:"rpc_name"`
	ListenAddr string `yaml:"listen_addr" xml:"listen_addr"`
	ListenPort int    `yaml:"listen_port" xml:"listen_port"`
}
