package model

type ServerConfig struct {
	ListenAddr string `json:"listen_addr"`
	ListenPort int32  `json:"listen_port"`
	SSL        bool   `json:"ssl"`
	Crt        string `json:"crt"`
	Key        string `json:"key"`
}
