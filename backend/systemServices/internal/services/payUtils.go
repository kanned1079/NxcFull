package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"systemServices/internal/dao"
	"systemServices/internal/model"
)

// system 为systemName
// Key为结构体中每一个字段的字段名 json格式的
// Value为结构体中每一个字段的字段名的值
// 例子
// id system key value
// 1 alipay app_id 234424325275
// 2 alipay app_cert_public "dh3784qgfq38g74ofqf37q4f"
// 3 alipay app_private_key "df34g45g34gfh"
// 4 alipay discount 12

func SaveOrUpdatePaymentMethodBySystemName(systemName string, conf json.RawMessage) error {
	var data map[string]interface{}

	// 将 JSON 数据解析为通用 map
	if err := json.Unmarshal(conf, &data); err != nil {
		return fmt.Errorf("failed to parse configuration: %v", err)
	}

	// 遍历配置字段
	for key, value := range data {
		// 跳过空的和已设置的字段
		if strValue, ok := value.(string); ok && (strValue == "" || strValue == "---") {
			continue
		}

		// 将值序列化为 JSON 以存储
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return fmt.Errorf("failed to marshal value for key %s: %v", key, err)
		}

		// 检查数据库中是否已有相同的系统和键
		var existing model.PaymentSettings
		if err := dao.Db.Where("`system` = ? AND `key` = ?", systemName, key).First(&existing).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("failed to query database for key %s: %v", key, err)
			}

			// 如果记录不存在，则插入新记录
			newSetting := model.PaymentSettings{
				System: systemName,
				Key:    key,
				Value:  json.RawMessage(jsonValue),
			}
			if err := dao.Db.Create(&newSetting).Error; err != nil {
				return fmt.Errorf("failed to insert new setting for key %s: %v", key, err)
			}
		} else {
			// 如果记录已存在，更新其值
			existing.Value = json.RawMessage(jsonValue)
			if err := dao.Db.Save(&existing).Error; err != nil {
				return fmt.Errorf("failed to update setting for key %s: %v", key, err)
			}
		}
	}

	return nil
}
