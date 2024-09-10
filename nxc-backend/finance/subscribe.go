package finance

import (
	"gorm.io/gorm"
	"time"
)

type Subscribes struct {
	Id              int64          `json:"id"`               // 订阅ID
	Name            string         `json:"name"`             // 名称
	Sale            bool           `json:"sale"`             // 是否启用销售
	Renew           bool           `json:"renew"`            // 是否允许续费
	MonthlyPrice    float64        `json:"monthly_price"`    // 月付价格
	SeasonalPrice   float64        `json:"seasonal_price"`   // 季付价格
	SemiannualPrice float64        `json:"semiannual_price"` // 半年付价格
	AnnualPrice     float64        `json:"annual_price"`     // 年付价格
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

func (Subscribes) TableName() string {
	return "x_subscribes"
}
