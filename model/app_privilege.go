package model

import (
	"github.com/lfyr/go-api/config"
	"github.com/lfyr/go-api/database/masterdb"
)

type AppPrivilege struct {
	config.Model
	PriName    string `json:"pri_name"`
	ActionName string `json:"action_name"`
	ParentId   int    `json:"parent_id"`
}

func (u *AppPrivilege) TableName() string {
	return "app_privilege"
}

func NewAppPrivilege() *AppPrivilege {
	return &AppPrivilege{}
}

func (this *AppPrivilege) List(whereMap map[string]interface{}, fieldSlice []string, page, size int, withSlice []string) (list []AppPrivilege, count int64) {
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
	tx.Count(&count).Order("id asc").Offset((page - 1) * size).Limit(size).Find(&list)
	return
}

func (this *AppPrivilege) First(whereMap map[string]interface{}) (detail AppPrivilege) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	tx.Order("id desc").First(&detail)
	return
}

func (this *AppPrivilege) Many(whereMap map[string]interface{}) (list []AppPrivilege) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	tx.Preload("Category").Find(&list)
	return
}

func (this *AppPrivilege) Create(data *AppPrivilege) (err error) {
	tx := masterdb.DB.Model(this)
	err = tx.Create(&data).Error
	return
}

func (this *AppPrivilege) Update(Id int, user map[string]interface{}) (err error) {
	err = masterdb.DB.Model(this).Where("id = ?", Id).Updates(&user).Error
	return
}
