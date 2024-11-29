package payment

import (
	"github.com/go-pay/gopay/alipay/v3"
	"log"
	model "orderServices/internal/payment/config"
	"orderServices/internal/payment/platform"
)

//const (
//	AlipayClient *alipay.ClientV3
//)

var (
	PaymentInstanceRoot PaymentInstance
)

type PaymentInstance struct {
	AlipayInstance *alipay.ClientV3
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
					log.Println("初始化支付宝客户端失败")
				}
				log.Println("初始化完成")
			}
		}
	}

}
