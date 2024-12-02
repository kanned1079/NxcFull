package payment

import (
	"context"
	"encoding/json"
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
func DoAlipayV3PreCreateTopUpOrder(subject string, orderId string, amount float32) (err error, code int, qrCode string, outTradeNo string) {
	// 检查金额是否合法
	if amount <= 0 {
		return errors.New("amount must be greater than 0"), http.StatusBadRequest, "", ""
	}

	// 从缓存中获取折扣信息
	confDiscount := model.AlipayConfCache.Discount
	// 规则如下
	// discount为前端传递的优惠金额 model.AlipayConfCache.Discount为从数据库中取的优惠金额 这两个值应当相等
	// 但是注意下面的几条规则
	// 1.如果用户的充值金额小于model.AlipayConfCache.Discount 那么则不使用优惠
	// 2.如果用户的充值金额等于model.AlipayConfCache.Discount 直接返回错误409，不允许充值0元
	// 3.只有当用户充值的金额大于model.AlipayConfCache.Discount时（如model.AlipayConfCache.Discount为10，充值金额amount为10.01）可以使用优惠
	// 4.充值金额不可为0或负数

	// 计算最终总金额
	var finalAmount float32
	if confDiscount > 0 {
		if amount < confDiscount {
			// 如果充值金额小于优惠金额，不使用优惠
			finalAmount = amount
		} else if amount == confDiscount {
			// 如果充值金额等于优惠金额，返回错误
			return errors.New("amount equals to discount is not allowed"), http.StatusConflict, "", ""
		} else {
			// 充值金额大于优惠金额，减去优惠后为最终金额
			finalAmount = amount - confDiscount
		}
	} else {
		// 如果没有任何优惠，直接使用原始金额
		finalAmount = amount
	}

	// 保留两位小数格式化金额
	totalAmount := fmt.Sprintf("%.2f", finalAmount)

	// 构建请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", subject).
		Set("out_trade_no", orderId).
		Set("total_amount", totalAmount) // 使用计算后的最终金额

	// 响应结果结构体
	rsp := new(struct {
		OutTradeNo string `json:"out_trade_no"`
		QrCode     string `json:"qr_code"`
	})

	// 调用支付 API
	res, err := PaymentInstanceRoot.AlipayInstance.DoAliPayAPISelfV3(context.Background(), "POST", "/v3/alipay/trade/precreate", bm, rsp)
	if err != nil {
		xlog.Errorf("client.TradePrecreate(), err:%v", err)
		return errors.New("TradePrecreate err"), http.StatusInternalServerError, "", ""
	}

	xlog.Debugf("aliRsp:%s", js.Marshal(rsp))
	if res.StatusCode != http.StatusOK {
		xlog.Errorf("aliRsp.StatusCode:%d", res.StatusCode)
		return errors.New("alipay response err"), res.StatusCode, "", ""
	}

	return nil, res.StatusCode, rsp.QrCode, rsp.OutTradeNo
}

func DoAlipayV3QueryPreCreateTopUpOrderStatus(orderId string) (code int32, msg string, err error, tradeStatus string, tradeNo string, outTradeNo string, totalAmount string, buyerPayAmount string, sendPayDate string) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", orderId)
	// 调用查询接口
	res, err := PaymentInstanceRoot.AlipayInstance.TradeQuery(context.Background(), bm)
	jsonResp, _ := json.Marshal(res)
	log.Println("res", string(jsonResp))
	if res == nil { // 如果没有返回值
		return 404, "支付结果获取失败", errors.New("missing valid response"), "NULL_RESPONSE", "", "", "", "", ""
	}

	if res.StatusCode == 400 {
		log.Println("请先扫码")
		return int32(WaitUserScanQrCode), "请先扫描二维码以创建订单", nil, "NULL_RESPONSE", "", "", "", "", ""
	}

	var resultCode int
	var resultMsg string
	switch res.TradeStatus {
	case "WAIT_BUYER_PAY":
		{
			resultCode = ScannedQrCodeWaitUserPay
			resultMsg = "交易创建成功_等待买家付款"
			//return int32(ScannedQrCodeWaitUserPay), nil, res.TradeStatus, res.TradeNo, res.OutTradeNo, res.TotalAmount, res.BuyerPayAmount, res.SendPayDate
		}

	case "TRADE_CLOSED":
		{
			resultCode = TradeClosed
			resultMsg = "未付款交易超时关闭_或支付完成后全额退款"
		}

	case "TRADE_SUCCESS":
		{
			resultCode = TradeSuccess
			resultMsg = "交易支付成功"
		}
	case "TRADE_FINISHED":
		{
			resultCode = TradeFinished
			resultMsg = "交易结束_不可退款"
		}

	}

	return int32(resultCode), resultMsg, nil, res.TradeStatus, res.TradeNo, res.OutTradeNo, res.TotalAmount, res.BuyerPayAmount, res.SendPayDate

	//return http.StatusOK, nil, res.TradeStatus, res.TradeNo, res.OutTradeNo, res.TotalAmount, res.BuyerPayAmount, res.SendPayDate
}
