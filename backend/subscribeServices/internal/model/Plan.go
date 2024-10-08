package model

import (
	"gorm.io/gorm"
	"time"
)

type Plan struct {
	Id            int64   `json:"id" gorm:"primary_key"`     // 订阅计划ID
	GroupId       int64   `json:"group_id"`                  // 群组Id
	Name          string  `json:"name"`                      // 订阅名称
	IsSale        bool    `json:"is_sale"`                   // 启用销售
	IsRenew       bool    `json:"is_renew"`                  // 是否允许续费
	CapacityLimit int64   `json:"capacity_limit"`            // 最大用户数量限制
	Residue       int64   `json:"residue"`                   // 剩余订阅数量
	Describe      string  `json:"describe" gorm:"type:TEXT"` // 订阅计划的备注
	MonthPrice    float64 `json:"month_price"`               // 月付价格
	QuarterPrice  float64 `json:"quarter_price"`             // 季付价格
	HalfYearPrice float64 `json:"half_year_price"`           // 半年付价格
	YearPrice     float64 `json:"year_price"`                // 年付价格
	Sort          int64   `json:"sort"`                      // 排序顺序

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Plan) TableName() string {
	return "x_plan"
}
