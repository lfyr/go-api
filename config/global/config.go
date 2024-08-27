package global

import "time"

type Model struct {
	Id        int       `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time `json:"createdAt"`            // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`            // 更新时间
}
