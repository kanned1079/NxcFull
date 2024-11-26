package config

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type PaymentSettings struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	System    string          `gorm:"size:50;not null;index" json:"system"` // 支付系统名称，如 "alipay", "wechat"
	Key       string          `gorm:"size:100;not null;index" json:"key"`   // 配置键
	Value     json.RawMessage `gorm:"type:longtext;not null" json:"value"`  // 配置值，JSON 存储
	CreatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
}

func (PaymentSettings) TableName() string {
	return "x_payment_settings"
}

type Alipay struct {
	AppId               string `json:"app_id"`                 // 应用id
	AppCertPublicKey    string `json:"app_cert_public"`        // 应用公钥证书内容
	AppPrivateKey       string `json:"app_private_key"`        // 应用私钥
	AlipayRootCert      string `json:"alipay_root_cert"`       // 支付宝根证书
	AlipayCertPublicKey string `json:"alipay_cert_public_key"` // 支付宝公钥证书内容
}

type Wechat struct {
	MchId      string `json:"mch_id"`      // 商户ID 或者服务商模式的 sp_mchid
	SerialNo   string `json:"serial_no"`   // 商户证书的证书序列号
	ApiV3Key   string `json:"api_v3_key"`  // apiV3Key商户平台获取
	PrivateKey string `json:"private_key"` // 私钥 apiclient_key.pem 读取后的内容
}

type Apple struct {
	Iss        string `json:"iss"`         // issuer ID
	Bid        string `json:"bid"`         // bundle ID
	Kid        string `json:"kid"`         // private key ID
	PrivateKey string `json:"private_key"` // 私钥文件读取后的字符串内容
}
