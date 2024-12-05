package platform

import (
	"github.com/go-pay/gopay/apple"
	"github.com/go-pay/xlog"
	"orderServices/internal/payment/model"
)

//var (
//	err error
//)

// InitApplePayClient 初始化ApplePay客户端
func InitApplePayClient() (error, *apple.Client) {
	client, err := apple.NewClient(model.AppleConfCache.Iss, model.AppleConfCache.Bid, model.AppleConfCache.Kid, model.AppleConfCache.PrivateKey, false)
	if err != nil {
		xlog.Error(err)
		return err, nil
	}

	return nil, client

}
