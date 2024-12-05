package platform

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"orderServices/internal/payment/model"
)

//var (
//	err error
//)

// InitWechatClient 初始化微信
func InitWechatClient() (error, *wechat.ClientV3) {
	client, err := wechat.NewClientV3(model.WechatConfCache.MchId, model.WechatConfCache.SerialNo, model.WechatConfCache.ApiV3Key, model.WechatConfCache.PrivateKey)
	if err != nil {
		return err, nil
	}

	// 设置微信平台API证书和序列号（推荐开启自动验签，无需手动设置证书公钥等信息）
	//client.SetPlatformCert([]byte(""), "")

	// 启用自动同步返回验签，并定时更新微信平台API证书（开启自动验签时，无需单独设置微信平台API证书和序列号）
	err = client.AutoVerifySign()
	if err != nil {
		return err, nil
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	client.SetBodySize(10) // 没有特殊需求，可忽略此配置

	// 设置自定义RequestId生成方法，非必须
	//client.SetRequestIdFunc()

	// 打开Debug开关，输出日志，默认是关闭的
	client.DebugSwitch = gopay.DebugOn

	return nil, client

}
