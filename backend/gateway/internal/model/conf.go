package model

type ServerConfig struct {
	ListenAddr string `json:"listen_addr"`
	ListenPort int32  `json:"listen_port"`
}
