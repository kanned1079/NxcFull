package payment

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/js"
	"github.com/go-pay/xlog"
	"log"
	"net/http"
	"orderServices/internal/payment/model"
)

// DoAlipayV3PreCreateTopUpOrder 预生成订单并回传订单号 支付二维码
func DoAlipayV3PreCreateTopUpOrder(subject string, orderId string, amount float32, discount float32) (err error, code int, qrCode string, outTradeNo string) {
	// 格式化浮点数为字符串
	totalAmount := fmt.Sprintf("%.2f", amount)                                                 // 保留两位小数
	if discount != model.AlipayConfCache.Discount && amount > model.AlipayConfCache.Discount { // 如果折扣金额不匹配
		return errors.New("discount amount not match"), http.StatusConflict, "", ""
	}
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
		//return http.StatusInternalServerError, err
		return errors.New("TradePrecreate err"), http.StatusInternalServerError, "", ""
	}
	xlog.Debugf("aliRsp:%s", js.Marshal(rsp))
	if res.StatusCode != http.StatusOK {
		xlog.Errorf("aliRsp.StatusCode:%d", res.StatusCode)
		//return
		//return res.StatusCode, "", ""
		return errors.New("alipay response err"), res.StatusCode, "", ""
	}
	//return http.StatusOK, nil

	return nil, res.StatusCode, rsp.QrCode, rsp.OutTradeNo
}

func DoAlipayV3QueryPreCreateTopUpOrderStatus(orderId string) (err error, tradeStatus string, tradeNo string, outTradeNo string, totalAmount string) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", orderId)

	//rsp := new(struct {
	//	OutTradeNo string `json:"out_trade_no"`
	//	QrCode     string `json:"qr_code"`
	//})
	res, err := PaymentInstanceRoot.AlipayInstance.TradeQuery(context.Background(), bm)
	log.Println(res)
	if res == nil { // 如果没有返回值
		return errors.New("missing valid response"), "NULL_RESPONSE", "", "", ""
	}
	switch res.TradeStatus {
	case "WAIT_BUYER_PAY":
		{
			log.Println("交易创建，等待买家付款")
		}

	case "TRADE_CLOSED":
		{
			log.Println("未付款交易超时关闭，或支付完成后全额退款")
		}

	case "RADE_SUCCESS":
		{
			log.Println("交易支付成功")
		}
	case "TRADE_FINISHED":
		{
			log.Println("交易结束，不可退款")
		}

	}
	return nil, res.TradeStatus, res.TradeNo, res.OutTradeNo, res.TotalAmount
}
