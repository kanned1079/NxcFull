package payment

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/js"
	"github.com/go-pay/xlog"
	"net/http"
)

// DoAliPayAPISelfV3 预生成订单并回传订单号 支付二维码
func DoAliPayAPISelfV3(subject string, orderId string, amount float32, discount float32) (code int32, err error) {
	// 格式化浮点数为字符串
	totalAmount := fmt.Sprintf("%.2f", amount) // 保留两位小数
	discountableAmount := fmt.Sprintf("%.2f", discount)

	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", subject).
		Set("out_trade_no", orderId).
		Set("total_amount", totalAmount). // 使用格式化后的字符串
		Set("discountable_amount", discountableAmount)

	rsp := new(struct {
		OutTradeNo string `json:"out_trade_no"`
		QrCode     string `json:"qr_code"`
	})
	// 创建订单
	res, err := PaymentInstanceRoot.AlipayInstance.DoAliPayAPISelfV3(context.Background(), "POST", "/v3/alipay/trade/precreate", bm, rsp)
	if err != nil {
		xlog.Errorf("client.TradePrecreate(), err:%v", err)
		return http.StatusInternalServerError, err
	}
	xlog.Debugf("aliRsp:%s", js.Marshal(rsp))
	if res.StatusCode != http.StatusOK {
		xlog.Errorf("aliRsp.StatusCode:%d", res.StatusCode)
		//return
		return http.StatusInternalServerError, err

	}
	return http.StatusOK, nil
}
