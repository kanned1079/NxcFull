package model

var MyEtcdConfig EtcdConf

type EtcdConf struct {
	Endpoints string `json:"endpoints"`
	Port      int32  `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
