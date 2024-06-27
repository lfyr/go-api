package config

import "time"

type Model struct {
	ID        int       `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time `json:"createdAt"`            // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`            // 更新时间
}

type Server struct {
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
