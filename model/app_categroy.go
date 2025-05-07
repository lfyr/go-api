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

func (this *AppCategory) List(whereMap map[string]interface{}, fieldSlice []string, page, size int, withSlice []string) (list []AppCategory, count int64) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}

	if len(fieldSlice) > 0 {
		tx = tx.Select(fieldSlice)
	}

	if len(withSlice) > 0 {
		for _, v := range withSlice {
			tx = tx.Preload(v)
		}
	}
	tx.Where("delete_status", 0).Count(&count).Order("id asc").Offset((page - 1) * size).Limit(size).Find(&list)
	return
}

func (this *AppCategory) Many(whereMap map[string]interface{}) (list []AppCategory) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	tx.Find(&list)
	return
}

func (this *AppCategory) Create(data *AppCategory) (err error) {
	err = masterdb.DB.Model(this).Create(&data).Error
	return
}

func (this *AppCategory) Update(whereMap map[string]interface{}, data map[string]interface{}) (err error) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	err = tx.Updates(&data).Error
	return
}
