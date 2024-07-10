package model

import (
	"github.com/lfyr/go-api/config"
	"github.com/lfyr/go-api/database/masterdb"
)

type AppAdminRole struct {
	config.Model
	AdminId          int                `json:"admin_id"`
	RoleId           int                `json:"role_id" `
	AppRolePrivilege []AppRolePrivilege `gorm:"foreignKey:RoleId;references:RoleId"`
}

func (u *AppAdminRole) TableName() string {
	return "app_admin_role"
}

func NewAppAdminRole() *AppAdminRole {
	return &AppAdminRole{}
}

func (this *AppAdminRole) List(whereMap map[string]interface{}, fieldSlice []string, page, size int, withSlice []string) (list []AppAdminRole, count int64) {
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

func (this *AppAdminRole) First(whereMap map[string]interface{}) (detail AppAdminRole) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	tx.Order("id desc").First(&detail)
	return
}

func (this *AppAdminRole) Many(whereMap map[string]interface{}) (list []AppAdminRole) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	tx.Preload("Category").Find(&list)
	return
}

func (this *AppAdminRole) Create(data *AppAdminRole) (err error) {
	tx := masterdb.DB.Model(this)
	err = tx.Create(&data).Error
	return
}

func (this *AppAdminRole) Update(Id int, user map[string]interface{}) (err error) {
	err = masterdb.DB.Model(this).Where("id = ?", Id).Updates(&user).Error
	return
}