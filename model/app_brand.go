package model

import (
	"github.com/lfyr/go-api/config/global"
	"github.com/lfyr/go-api/database/masterdb"
)

type AppBrand struct {
	global.Model
	BrandName    string `json:"brandName"`
	Logo         string `json:"logo"`
	DeleteStatus int    `json:"deleteStatus"`
}

func (AppBrand) TableName() string {
	return "app_brand"
}
func NewAppBrand() *AppBrand {
	return &AppBrand{}
}

func (this *AppBrand) List(whereMap map[string]interface{}, fieldSlice []string, page, size int, withSlice []string) (list []AppBrand, count int64) {
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

func (this *AppBrand) First(whereMap map[string]interface{}) (detail AppBrand) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	tx.Order("id desc").First(&detail)
	return
}

func (this *AppBrand) Create(data *AppBrand) (err error) {
	err = masterdb.DB.Model(this).Create(&data).Error
	return
}

func (this *AppBrand) Update(whereMap map[string]interface{}, data map[string]interface{}) (err error) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	err = tx.Updates(&data).Error
	return
}
