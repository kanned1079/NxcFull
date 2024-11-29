package payment

import (
	"encoding/json"
	"fmt"
	"log"
	"orderServices/internal/dao"
	"orderServices/internal/payment/model"
	"strconv"
)

func InitPaymentConf() {
	var paymentSettings []model.PaymentSettings

	// 查询数据库
	err := dao.Db.Model(&model.PaymentSettings{}).Find(&paymentSettings).Error
	if err != nil {
		fmt.Printf("Failed to load payment settings: %v\n", err)
		return
	}

	// 分类存储配置信息
	alipayConfig := make(map[string]interface{})
	wechatConfig := make(map[string]interface{})
	appleConfig := make(map[string]interface{})

	// 遍历配置项
	for _, setting := range paymentSettings {
		value := normalizeValue(setting.Key, setting.Value)
		switch setting.System {
		case "alipay":
			alipayConfig[setting.Key] = value
		case "wechat":
			wechatConfig[setting.Key] = value
		case "apple":
			appleConfig[setting.Key] = value
		}
	}

	// 解析为结构体并缓存
	if err := mapToStruct(alipayConfig, &model.AlipayConfCache); err != nil {
		log.Panicf("Failed to parse Alipay config: %v\n", err)
	} else if model.AlipayConfCache.Enabled {
		model.PaymentMethodsArr = append(model.PaymentMethodsArr, "alipay")
	}

	if err := mapToStruct(wechatConfig, &model.WechatConfCache); err != nil {
		log.Panicf("Failed to parse WeChat config: %v\n", err)
	} else if model.WechatConfCache.Enabled {
		model.PaymentMethodsArr = append(model.PaymentMethodsArr, "wechat")
	}

	if err := mapToStruct(appleConfig, &model.AppleConfCache); err != nil {
		log.Panicf("Failed to parse Apple config: %v\n", err)
	} else if model.AppleConfCache.Enabled {
		model.PaymentMethodsArr = append(model.PaymentMethodsArr, "apple")
	}

	fmt.Printf("Payment configurations initialized successfully. Enabled methods: %v\n", model.PaymentMethodsArr)

}

// normalizeValue 对 JSON 数据进行解析，处理具体字段
func normalizeValue(key string, value json.RawMessage) interface{} {
	switch key {
	case "discount":
		var discount float32
		// 解析为 float32
		if err := json.Unmarshal(value, &discount); err == nil {
			return discount
		}

		// 尝试解析为字符串再转换为 float32
		var strValue string
		if err := json.Unmarshal(value, &strValue); err == nil {
			floatVal, _ := strconv.ParseFloat(strValue, 32)
			return float32(floatVal)
		}

		// 默认返回 0
		return float32(0)

	case "enabled":
		var enabled bool
		// 尝试解析为 bool
		if err := json.Unmarshal(value, &enabled); err == nil {
			return enabled
		}

		// 尝试解析为字符串再转换为 bool
		var strValue string
		if err := json.Unmarshal(value, &strValue); err == nil {
			return strValue == "true"
		}

		// 默认返回 false
		return false

	default:
		// 默认返回字符串
		var str string
		if err := json.Unmarshal(value, &str); err == nil {
			return str
		}
		return string(value)
	}
}

// mapToStruct 将 map 转换为结构体
func mapToStruct(data map[string]interface{}, out interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	if err := json.Unmarshal(jsonData, out); err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}
	return nil
}

// 序列化有效的支付配置
func serializeConfig(config interface{}, name string) {
	serializedData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Printf("Failed to serialize %s config: %v\n", name, err)
	} else {
		fmt.Printf("%s config serialized: %s\n", name, string(serializedData))
	}
}
