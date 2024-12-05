package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	pb "systemServices/api/proto"
	"systemServices/internal/dao"
	"systemServices/internal/model"
)

/*
  // pass 添加新的付款方式
  rpc AddPaymentMethod(AddPaymentMethodRequest) returns (AddPaymentMethodResponse);

  // 获取所有支付方式的键值对 名称 是否被启用
  rpc GetAllPaymentSettingsKv(GetAllPaymentSettingsKvRequest) returns (GetAllPaymentSettingsKvResponse);

  // 根据传入的支付方式名称来获取其详细信息
  rpc GetPaymentMethodDetailsByName(GetPaymentMethodDetailsByNameRequest) returns (GetPaymentMethodDetailsByNameResponse);

  // 根据名称来新增或修改支付方式的信息
  rpc EditPaymentSettingsBySystemName(EditPaymentSettingsBySystemNameRequest) returns (EditPaymentSettingsBySystemNameResponse);

  // 是否启用该支付防止 这个调用后取反enabled字段 之后需要刷新Redis缓存和页面
  rpc EnablePaymentSettingBySystemName(EnablePaymentSettingBySystemNameRequest) returns (EnablePaymentSettingBySystemNameResponse);

  // pass 删除付款方式
  rpc DeletePaymentSettingBySystemName(DeletePaymentSettingBySystemNameRequest) returns (DeletePaymentSettingBySystemNameResponse);
*/

func (s *SettingServices) EditPaymentSettingsBySystemName(ctx context.Context, request *pb.EditPaymentSettingsBySystemNameRequest) (*pb.EditPaymentSettingsBySystemNameResponse, error) {
	log.Println(string(request.Config))

	if err := SaveOrUpdatePaymentMethodBySystemName(request.System, request.Config); err != nil {
		return &pb.EditPaymentSettingsBySystemNameResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failure " + err.Error(),
		}, nil
	}

	return &pb.EditPaymentSettingsBySystemNameResponse{
		Code: http.StatusOK,
		Msg:  "Success",
	}, nil
}

//	func (s *SettingServices) GetAllPaymentSettingsKv(ctx context.Context, request *pb.GetAllPaymentSettingsKvRequest) (*pb.GetAllPaymentSettingsKvResponse, error) {
//		// singleConf 单个支付方式的结构体
//		singleConf := &struct {
//			System   string  `json:"system"`
//			Enabled  bool    `json:"enabled"`
//			Discount float32 `json:"discount"`
//		}{}
//
//		// 查询方式如下
//		// 1.查询System列中所有的内容并去重 得到总共有多少种支付方式
//		// 2.查询System为支付方式的名称下的Key为enabled得到其Value 这个代表此支付方式是否被启用 反序列化后得到bool类型
//		// 3.查询System为支付方式的名称下的Key为discount得到其Value 这个代表可以优惠的金额 反序列化后得到float32类型
//		// 最后返回的数据如下[{system: "alipay", enabled: true, discount: 0.71}, {...}, {...}] 但是注意这个需要被序列化后再返回
//
//		/*
//			数据库定义如下
//			type PaymentSettings struct {
//				ID        uint            `gorm:"primaryKey" json:"id"`
//				System    string          `gorm:"size:50;not null;index:idx_system_key" json:"system"` // 支付系统名称
//				Key       string          `gorm:"size:100;not null;index:idx_system_key" json:"key"`   // 配置键
//				Value     json.RawMessage `gorm:"type:longtext;not null" json:"value"`                 // 配置值
//				CreatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
//				UpdatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
//				DeletedAt gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
//			}
//		*/
//
//		// 数据库操作
//		dao.Db.
//
//		return &pb.GetAllPaymentSettingsKvResponse{
//			Code: http.StatusOK,
//			Msg:  "Success",
//			Config: <返回的数据>,
//		}, nil
//	}
var err error

func (s *SettingServices) GetAllPaymentSettingsKv(ctx context.Context, request *pb.GetAllPaymentSettingsKvRequest) (*pb.GetAllPaymentSettingsKvResponse, error) {
	// 定义结果集合
	var result []map[string]interface{}

	// 查询所有支付方式的名称（System列去重）
	var systems []string
	if err := dao.Db.Model(&model.PaymentSettings{}).
		Distinct("system").
		Pluck("system", &systems).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch payment systems: %v", err)
	}

	// 遍历每种支付方式
	for _, system := range systems {
		singleConf := map[string]interface{}{
			"system":   system,
			"enabled":  false,      // 默认值
			"discount": float32(0), // 默认值
		}

		// 查询 Key 为 "enabled" 的值
		var enabledSetting model.PaymentSettings
		if err := dao.Db.Where("`system` = ? AND `key` = ?", system, "enabled").
			First(&enabledSetting).Error; err == nil {
			var enabled bool
			if err := json.Unmarshal(enabledSetting.Value, &enabled); err == nil {
				singleConf["enabled"] = enabled
			}
		}

		// 查询 Key 为 "discount" 的值
		var discountSetting model.PaymentSettings
		if err := dao.Db.Where("`system` = ? AND `key` = ?", system, "discount").
			First(&discountSetting).Error; err == nil {
			var discount float32
			if err := json.Unmarshal(discountSetting.Value, &discount); err == nil {
				singleConf["discount"] = discount
			}
		}

		// 将当前支付方式的配置信息加入结果集合
		result = append(result, singleConf)
	}

	var resultJson, discountMsgJson json.RawMessage
	var discountMsgString string
	discountMsgJson, err = readSettingFromDB("invite", "discount_info")

	if err = json.Unmarshal(discountMsgJson, &discountMsgString); err != nil {
		log.Println("get discount msg err: ", err)
	}

	if resultJson, err = json.Marshal(result); err != nil {
		return &pb.GetAllPaymentSettingsKvResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failure " + err.Error(),
		}, nil
	}

	// 返回结果
	return &pb.GetAllPaymentSettingsKvResponse{
		Code:        http.StatusOK,
		Config:      resultJson,
		DiscountMsg: discountMsgString,
		Msg:         "Success",
	}, nil

}

func (s *SettingServices) EnablePaymentSettingBySystemName(ctx context.Context, request *pb.EnablePaymentSettingBySystemNameRequest) (*pb.EnablePaymentSettingBySystemNameResponse, error) {
	// 查找数据库中指定系统和键为 "enabled" 的记录
	var setting model.PaymentSettings
	if err := dao.Db.Where("`system` = ? AND `key` = ?", request.System, "enabled").First(&setting).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.EnablePaymentSettingBySystemNameResponse{
				Code: http.StatusNotFound,
				Msg:  "Payment setting not found",
			}, nil
		}
		return &pb.EnablePaymentSettingBySystemNameResponse{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("Failed to query database: %v", err),
		}, nil
	}

	// 反序列化 Value 字段，取反其值
	var enabled bool
	if err := json.Unmarshal(setting.Value, &enabled); err != nil {
		return &pb.EnablePaymentSettingBySystemNameResponse{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("Failed to parse value: %v", err),
		}, nil
	}

	// 取反逻辑
	enabled = !enabled

	// 更新数据库中对应记录的 Value 值
	newValue, err := json.Marshal(enabled)
	if err != nil {
		return &pb.EnablePaymentSettingBySystemNameResponse{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("Failed to marshal new value: %v", err),
		}, nil
	}

	setting.Value = json.RawMessage(newValue)
	if err := dao.Db.Save(&setting).Error; err != nil {
		return &pb.EnablePaymentSettingBySystemNameResponse{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("Failed to update value in database: %v", err),
		}, nil
	}

	// 返回成功响应
	return &pb.EnablePaymentSettingBySystemNameResponse{
		Code: http.StatusOK,
		Msg:  "Success",
	}, nil
}

//func (s *SettingServices) GetPaymentMethodDetailsByName(ctx context.Context, request *pb.GetPaymentMethodDetailsByNameRequest) (*pb.GetPaymentMethodDetailsByNameResponse, error) {
//	request.System
//
//	var details map[string]interface{}
//
//	//	/*
//	//		数据库定义如下
//	//		type PaymentSettings struct {
//	//			ID        uint            `gorm:"primaryKey" json:"id"`
//	//			System    string          `gorm:"size:50;not null;index:idx_system_key" json:"system"` // 支付系统名称
//	//			Key       string          `gorm:"size:100;not null;index:idx_system_key" json:"key"`   // 配置键
//	//			Value     json.RawMessage `gorm:"type:longtext;not null" json:"value"`                 // 配置值
//	//			CreatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
//	//			UpdatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
//	//			DeletedAt gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
//	//		}
//	//	*/
//	// 按照request.System 取出整个行的数据
//
//	return &pb.GetPaymentMethodDetailsByNameResponse{
//		Code: http.StatusOK,
//		Msg:  "Success",
//		Details:
//	}, nil
//}

//func (s *SettingServices) GetPaymentMethodDetailsByName(ctx context.Context, request *pb.GetPaymentMethodDetailsByNameRequest) (*pb.GetPaymentMethodDetailsByNameResponse, error) {
//	// 检查请求参数
//	if request.System == "" {
//		return &pb.GetPaymentMethodDetailsByNameResponse{
//			Code: http.StatusBadRequest,
//			Msg:  "System name cannot be empty",
//		}, nil
//	}
//
//	var settings []model.PaymentSettings
//	// 查询数据库中所有符合条件的配置项
//	err := dao.Db.Where("`system` = ?", request.System).Find(&settings).Error
//	if err != nil {
//		return &pb.GetPaymentMethodDetailsByNameResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  fmt.Sprintf("Failed to query database: %v", err),
//		}, nil
//	}
//
//	// 如果查询结果为空
//	if len(settings) == 0 {
//		return &pb.GetPaymentMethodDetailsByNameResponse{
//			Code: http.StatusNotFound,
//			Msg:  fmt.Sprintf("No settings found for system: %s", request.System),
//		}, nil
//	}
//
//	// 将查询结果组织为 map
//	details := make(map[string]interface{})
//	for _, setting := range settings {
//
//		var value interface{}
//		// 尝试反序列化 Value 字段
//		if err := json.Unmarshal(setting.Value, &value); err != nil {
//			return &pb.GetPaymentMethodDetailsByNameResponse{
//				Code: http.StatusInternalServerError,
//				Msg:  fmt.Sprintf("Failed to unmarshal value for key %s: %v", setting.Key, err),
//			}, nil
//		}
//		details[setting.Key] = value
//	}
//
//	if detailsJson, err := json.Marshal(details); err != nil {
//		return &pb.GetPaymentMethodDetailsByNameResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "Failure " + err.Error(),
//		}, nil
//	} else {
//		return &pb.GetPaymentMethodDetailsByNameResponse{
//			Code:    http.StatusOK,
//			Msg:     "Success",
//			Details: detailsJson,
//		}, nil
//	}
//
//	//// 返回成功响应
//	//return &pb.GetPaymentMethodDetailsByNameResponse{
//	//	Code:    http.StatusOK,
//	//	Msg:     "Success",
//	//	Details: details,
//	//}, nil
//}

func (s *SettingServices) GetPaymentMethodDetailsByName(ctx context.Context, request *pb.GetPaymentMethodDetailsByNameRequest) (*pb.GetPaymentMethodDetailsByNameResponse, error) {
	// 检查请求参数
	if request.System == "" {
		return &pb.GetPaymentMethodDetailsByNameResponse{
			Code: http.StatusBadRequest,
			Msg:  "System name cannot be empty",
		}, nil
	}

	var settings []model.PaymentSettings
	// 查询数据库中所有符合条件的配置项
	err := dao.Db.Where("`system` = ?", request.System).Find(&settings).Error
	if err != nil {
		return &pb.GetPaymentMethodDetailsByNameResponse{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("Failed to query database: %v", err),
		}, nil
	}

	// 如果查询结果为空
	if len(settings) == 0 {
		return &pb.GetPaymentMethodDetailsByNameResponse{
			Code: http.StatusNotFound,
			Msg:  fmt.Sprintf("No settings found for system: %s", request.System),
		}, nil
	}

	// 将查询结果组织为 map
	details := make(map[string]interface{})
	for _, setting := range settings {
		// 跳过 key 为 "enabled" 和 "discount" 的记录
		if setting.Key == "enabled" || setting.Key == "discount" {
			continue
		}

		var value string
		// 尝试反序列化 Value 字段为字符串
		if err := json.Unmarshal(setting.Value, &value); err != nil {
			return &pb.GetPaymentMethodDetailsByNameResponse{
				Code: http.StatusInternalServerError,
				Msg:  fmt.Sprintf("Failed to unmarshal value for key %s: %v", setting.Key, err),
			}, nil
		}

		// 如果值不是空字符串，则设置为 "---"
		if value != "" {
			details[setting.Key] = "---"
		} else {
			details[setting.Key] = value
		}
	}

	if detailsJson, err := json.Marshal(details); err != nil {
		return &pb.GetPaymentMethodDetailsByNameResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Failure " + err.Error(),
		}, nil
	} else {
		return &pb.GetPaymentMethodDetailsByNameResponse{
			Code:    http.StatusOK,
			Msg:     "Success",
			Details: detailsJson,
		}, nil
	}
}
