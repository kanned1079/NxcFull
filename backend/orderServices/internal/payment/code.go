package payment

import "net/http"

const (
	// WaitUserScanQrCode 预生成订单成功后的等待用户扫码
	WaitUserScanQrCode = http.StatusAccepted // 202 - 请求已接受，但尚未处理，等待用户操作

	// ScannedQrCodeWaitUserPay 用户扫描二维码后生成订单等待支付
	ScannedQrCodeWaitUserPay = http.StatusProcessing // 102 - 请求已接收，正在处理中，但尚未完成

	// TradeClosed 未付款交易超时关闭，或支付完成后全额退款
	TradeClosed = http.StatusRequestTimeout // 408 - 请求超时，通常表示支付超时未完成

	// TradeSuccess 交易成功
	TradeSuccess = http.StatusOK // 200 - 请求成功并且已处理完毕

	// TradeFinished TRADE_FINISHED，交易完成，表示交易结束
	TradeFinished = http.StatusGone // 410 - 资源不再可用，交易已经结束

	// UpdateUserBalanceSuccess 更新用户和邀请人的信息都成功
	UpdateUserBalanceSuccess = http.StatusOK

	// UpdateUserBalanceAccept 更新用户的成功 但是邀请人可能没有成功
	UpdateUserBalanceAccept = http.StatusPartialContent

	// UpdateUserBalanceFailed 都更新失败
	UpdateUserBalanceFailed = http.StatusInternalServerError
)
