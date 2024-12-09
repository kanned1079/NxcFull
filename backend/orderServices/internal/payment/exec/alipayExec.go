package exec

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/js"
	"github.com/go-pay/xlog"
	"gorm.io/gorm"
	"log"
	"math"
	"net/http"
	"orderServices/internal/dao"
	externalModel "orderServices/internal/model"
	"orderServices/internal/payment"
	"orderServices/internal/payment/config"
	"orderServices/internal/payment/count"
	"orderServices/internal/payment/model"
	"orderServices/internal/payment/utils"
)

// DoAlipayV3PreCreateTopUpOrder 预生成订单并回传订单号 支付二维码
func DoAlipayV3PreCreateTopUpOrder(subject string, orderId string, amount float32) (err error, code int, qrCode string, outTradeNo string) {
	// 检查金额是否合法
	if amount <= 0 {
		return errors.New("amount must be greater than 0"), http.StatusBadRequest, "", ""
	}

	// 从缓存中获取折扣信息
	confDiscount := model.AlipayConfCache.Discount
	var discountCopy float32
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
			discountCopy = confDiscount
		}
	} else {
		// 如果没有任何优惠，直接使用原始金额
		finalAmount = amount
	}

	// 保留两位小数格式化金额
	totalAmount := fmt.Sprintf("%.2f", finalAmount)

	//log.Println(totalAmount)
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
	res, err := payment.PaymentInstanceRoot.AlipayInstance.DoAliPayAPISelfV3(context.Background(), "POST", "/v3/alipay/trade/precreate", bm, rsp)
	if err != nil {
		xlog.Errorf("client.TradePrecreate(), err:%v", err)
		return errors.New("TradePrecreate err"), http.StatusInternalServerError, "", ""
	}

	xlog.Debugf("aliRsp:%s", js.Marshal(rsp))
	if res.StatusCode != http.StatusOK {
		xlog.Errorf("aliRsp.StatusCode:%d", res.StatusCode)
		return errors.New("alipay response err"), res.StatusCode, "", ""
	}

	// 保存订单进redis缓存 有效时间2h
	if err = utils.SavePreCreateOrder2Redis(rsp.OutTradeNo, amount, finalAmount, discountCopy); err != nil {
		log.Println("set pre create order to redis err:", err)

		return err, http.StatusInternalServerError, "", ""
	}

	return nil, res.StatusCode, rsp.QrCode, rsp.OutTradeNo
}

func DoAlipayV3QueryPreCreateTopUpOrderStatus(orderId string, userid int64, inviteUserId int64) (code int32, msg string, err error, tradeStatus string, tradeNo string, outTradeNo string, totalAmount string, buyerPayAmount string, sendPayDate string) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", orderId)
	// 调用查询接口
	res, err := payment.PaymentInstanceRoot.AlipayInstance.TradeQuery(context.Background(), bm)
	//jsonResp, _ := json.Marshal(res)
	//log.Println("res", string(jsonResp))
	if res == nil { // 如果没有返回值
		return 404, "支付结果获取失败", errors.New("missing valid response"), "NULL_RESPONSE", "", "", "", "", ""
	}

	if res.StatusCode == 400 {
		log.Println("请先扫码")
		return int32(payment.WaitUserScanQrCode), "请先扫描二维码以创建订单", nil, "NULL_RESPONSE", "", "", "", "", ""
	}

	var resultCode int
	var resultMsg string
	switch res.TradeStatus {
	case "WAIT_BUYER_PAY":
		{
			resultCode = payment.ScannedQrCodeWaitUserPay
			resultMsg = "交易创建成功_等待买家付款"
			//return int32(ScannedQrCodeWaitUserPay), nil, res.TradeStatus, res.TradeNo, res.OutTradeNo, res.TotalAmount, res.BuyerPayAmount, res.SendPayDate
		}

	case "TRADE_CLOSED":
		{
			resultCode = payment.TradeClosed
			resultMsg = "未付款交易超时关闭_或支付完成后全额退款"
		}

	case "TRADE_SUCCESS":
		{
			resultCode = payment.TradeSuccess
			resultMsg = "交易支付成功"
			// 执行更新
			resultCode = int(UpdateUserBalanceAfterPaymentSuccess(orderId, userid, inviteUserId))

		}
	case "TRADE_FINISHED":
		{
			resultCode = payment.TradeFinished
			resultMsg = "交易结束_不可退款"
		}

	}

	return int32(resultCode), resultMsg, nil, res.TradeStatus, res.TradeNo, res.OutTradeNo, res.TotalAmount, res.BuyerPayAmount, res.SendPayDate

	//return http.StatusOK, nil, res.TradeStatus, res.TradeNo, res.OutTradeNo, res.TotalAmount, res.BuyerPayAmount, res.SendPayDate
}

func UpdateUserBalanceAfterPaymentSuccess(outTradeNo string, userId int64, inviteUserId int64) (code int32) {
	// 获取订单的相关信息
	_, amount, finalAmount, discount, err := utils.GetPreCreateOrderFromRedis(outTradeNo)
	// amount 用户充值金额
	// finalAmount 用户实际支付的金额
	// discount 优惠金额
	if err != nil {
		log.Println("GetPreCreateOrderFromRedis error:", err)
		return payment.UpdateUserBalanceFailed
	}
	log.Println("redis订单信息", finalAmount, discount, amount)

	// 用于管理员统计页面的图表 开启协程不阻塞当前线程
	go count.SaveIncomeCount2Redis(finalAmount)

	// 启动事务
	tx := dao.Db.Begin()

	// 确保在操作失败时回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Recovered from panic in transaction:", r)
		}
	}()

	// 更新用户余额
	var user externalModel.User
	if result := tx.Model(&externalModel.User{}).Where("id =?", userId).First(&user); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		tx.Rollback()
		log.Println("找不到用户", userId)
		return payment.UpdateUserBalanceFailed
	}
	user.Balance += amount

	if result := tx.Model(&externalModel.User{}).Where("id =?", userId).Save(&user); result.RowsAffected == 0 {
		tx.Rollback()
		log.Println("没有修改用户余额", userId)
		return payment.UpdateUserBalanceFailed
	}

	// 如果有邀请人，处理佣金
	// 1.计算佣金金额 增加的佣金金额为实际付款金额*比例 如finalAmount * (config.InviteConfCache.InviteRebateRate / 100)
	// 如果佣金金额保留两位小数后为0 则不增加邀请人的余额了 如 0.2(实付金额) * 0.02(佣金比例2%) = 0.004
	// 2.满足上一步的条件后再根据inviteUserId来查询用户并增加其余额
	if inviteUserId > 0 {
		// 获取最新的佣金比例
		if err := config.InviteConfCache.FetchConfFromRedis(); err != nil {
			log.Println("查询最新佣金比例失败")
		} else if config.InviteConfCache.InviteRebateEnable {
			// 计算佣金
			commission := finalAmount * (float32(config.InviteConfCache.InviteRebateRate) / 100)
			// 保留两位小数
			commission = float32(math.Round(float64(commission)*100) / 100)

			if commission > 0 {
				log.Println("佣金金额:", commission)

				// 更新邀请人余额
				var inviteUser externalModel.User
				if result := tx.Model(&externalModel.User{}).Where("id =?", inviteUserId).First(&inviteUser); errors.Is(result.Error, gorm.ErrRecordNotFound) {
					tx.Rollback()
					log.Println("找不到邀请人", inviteUserId)
					return payment.UpdateUserBalanceFailed
				}
				inviteUser.Balance += commission

				if result := tx.Model(&externalModel.User{}).Where("id =?", inviteUserId).Save(&inviteUser); result.RowsAffected == 0 {
					tx.Rollback()
					log.Println("没有修改邀请人余额", inviteUserId)
					return payment.UpdateUserBalanceFailed
				}

				log.Println("邀请人佣金已更新")
			}
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Println("事务提交失败:", err)
		return payment.UpdateUserBalanceFailed
	} else {
		log.Println("用户余额和佣金更新成功")
		return payment.UpdateUserBalanceSuccess
	}
}
