package payment

import (
	"github.com/go-pay/gopay/alipay/v3"
	"github.com/go-pay/gopay/apple"
	"github.com/go-pay/gopay/wechat/v3"
	"log"
	model "orderServices/internal/payment/model"
	"orderServices/internal/payment/platform"
)

//const (
//	AlipayClient *alipay.ClientV3
//)

var (
	PaymentInstanceRoot PaymentInstance
)

type PaymentInstance struct {
	AlipayInstance   *alipay.ClientV3
	WechatInstance   *wechat.ClientV3
	ApplePayInstance *apple.Client
}

func (p *PaymentInstance) Initial() {
	log.Println("总初始化函数")
	var err error
	for _, instance := range model.PaymentMethodsArr {
		switch instance {
		case "alipay":
			{
				log.Println("开始初始化支付宝")
				err, p.AlipayInstance = platform.InitAlipayClient()
				if err != nil {
					log.Println("初始化支付宝客户端失败", err)
				}
				log.Println("初始化完成")
			}
		case "wechat":
			{
				// 项目待完善...
				log.Println("开始初始化微信支付客户端")
				err, p.WechatInstance = platform.InitWechatClient()
				if err != nil {
					log.Println("初始化微信客户端失败", err)
				}
				log.Println("初始化完成")
			}
		case "apple":
			{
				// 项目待完善...
				log.Println("开始初始化ApplePay支付客户端")
				err, p.ApplePayInstance = platform.InitApplePayClient()
				if err != nil {
					log.Println("初始化ApplePay客户端失败", err)
				}
				log.Println("初始化完成")
			}
		default:
			{
				log.Println("没有启用的支付网关 跳过初始化")
			}
		}
	}

}
