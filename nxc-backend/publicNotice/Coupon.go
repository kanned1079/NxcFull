package publicNotice

import (
	"gorm.io/gorm"
	"time"
)

type PublicNotices struct {
	Id        int64          `json:"id" gorm:"primary_key"`    // 公告唯一标识
	Title     string         `json:"title"`                    // 公告标题
	Content   string         `json:"content" gorm:"type:text"` // 公告内容
	Show      bool           `json:"show"`                     // 是否展示
	ImgUrl    string         `json:"img_url"`                  // 背景图片地址
	Tags      string         `json:"tags"`                     // 标签
	CreatedAt time.Time      `json:"created_at"`               // 创建日期
	UpdatedAt time.Time      `json:"updated_at"`               // 更新日期
	DeletedAt gorm.DeletedAt `json:"deleted_at"`               // 删除日期
}

func (PublicNotices) TableName() string {
	return "x_public_notices"
}
