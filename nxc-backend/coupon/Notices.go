package coupon

import (
	"gorm.io/gorm"
	"time"
)

type Notices struct {
	Id        int64          `json:"id" gorm:"primary_key"` // 公告唯一标识
	Title     string         `json:"title"`                 // 公告标题
	Content   string         `json:"content"`               // 公告内容
	Show      bool           `json:"show"`                  // 是否展示
	ImgUrl    string         `json:"img_url"`               // 背景图片地址
	Tags      string         `json:"tags"`                  // 标签
	CreatedAt time.Time      `json:"createdAt"`             // 创建日期
	UpdatedAt time.Time      `json:"updatedAt"`             // 更新日期
	DeletedAt gorm.DeletedAt `json:"deletedAt"`             // 删除日期
}

func (Notices) TableName() string {
	return "notices"
}
