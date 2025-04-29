package model

import (
	"github.com/lfyr/go-api/config/global"
	"github.com/lfyr/go-api/database/masterdb"
)

type AppCategory struct {
	global.Model
	CatName  string `json:"cat_name"`
	ParentId int    `json:"parent_id"`
}

func (AppCategory) TableName() string {
	return "app_category"
}

func NewAppCategory() *AppCategory {
	return &AppCategory{}
}

func (this *AppCategory) Create(data *AppCategory) (err error) {
	err = masterdb.DB.Model(this).Create(&data).Error
	return
}
