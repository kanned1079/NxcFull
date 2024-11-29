package platform

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay/v3"
	"github.com/go-pay/xlog"
	"log"
	"orderServices/internal/payment/config"
)

var (
	err error
)

// InitAlipayClient 初始化支付宝客户端
func InitAlipayClient() (error, *alipay.ClientV3) {
	xlog.SetLevel(xlog.DebugLevel)
	log.Println(model.AlipayConfCache.AppId)
	log.Println(model.AlipayConfCache.AppPrivateKey)
	client, err := alipay.NewClientV3(model.AlipayConfCache.AppId, model.AlipayConfCache.AppPrivateKey, true)
	if err != nil {
		xlog.Error(err)
		return err, nil
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	client.SetBodySize(10) // 没有特殊需求，可忽略此配置

	// Debug开关，输出/关闭日志
	client.DebugSwitch = gopay.DebugOn

	// 设置自定义RequestId生成方法
	//client.SetRequestIdFunc()

	// 设置biz_content加密KEY，设置此参数默认开启加密（目前不可用）
	//client.SetAESKey("KvKUTqSVZX2fUgmxnFyMaQ==")

	// 传入证书内容
	err = client.SetCert([]byte(model.AlipayConfCache.AppCertPublicKey), []byte(model.AlipayConfCache.AlipayRootCert), []byte(model.AlipayConfCache.AlipayCertPublicKey))
	if err != nil {
		xlog.Debug("SetCert:", err)
		return err, nil
	}

	return nil, client

}
