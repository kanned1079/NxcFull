package setup

import (
	"encoding/json"
	"envFreshInit/internal/dao"
	"envFreshInit/internal/model"
	"envFreshInit/internal/utils"
	"errors"
	"fmt"
)

// SetupPaymentSettings 添加支付方式列表
func SetupPaymentSettings() error {
	if err := utils.ClearTable(model.PaymentSettings{}.TableName()); err != nil {
		utils.ShowFailure()
		fmt.Println("Failed to clear table, please check the database configuration.")
		return err
	}

	var paymentSettings []model.PaymentSettings

	// Alipay 配置项
	var alipayKeys []string = []string{
		"alipay_cert_public_key",
		"alipay_root_cert",
		"app_cert_public",
		"app_id",
		"app_private_key",
		"discount",
		"enabled",
	}

	// WeChat 配置项
	var wechatKeys []string = []string{
		"mch_id",
		"private_key",
		"serial_no",
		"api_v3_key",
		"discount",
		"enabled",
	}

	// Apple 配置项
	var appleKeys []string = []string{
		"private_key",
		"bid",
		"discount",
		"enabled",
		"iss",
		"kid",
	}

	// 为 Alipay 配置项添加支付设置
	for _, key := range alipayKeys {
		paymentSettings = append(paymentSettings, model.PaymentSettings{
			System: "alipay",
			Key:    key,
			Value:  json.RawMessage(`""`),
		})
	}

	// 为 WeChat 配置项添加支付设置
	for _, key := range wechatKeys {
		paymentSettings = append(paymentSettings, model.PaymentSettings{
			System: "wechat",
			Key:    key,
			Value:  json.RawMessage(`""`),
		})
	}

	// 为 Apple 配置项添加支付设置
	for _, key := range appleKeys {
		paymentSettings = append(paymentSettings, model.PaymentSettings{
			System: "apple",
			Key:    key,
			Value:  json.RawMessage(`""`),
		})
	}

	// 插入支付方式设置到数据库
	if result := dao.Db.Model(&model.PaymentSettings{}).Create(&paymentSettings); result.RowsAffected == 0 {
		utils.ShowFailure()
		fmt.Println("Failed to setup payment settings")
		return errors.New("failed to setup payment settings" + result.Error.Error())
	}

	return nil
}
