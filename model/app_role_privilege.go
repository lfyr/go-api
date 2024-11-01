package model

import (
	"github.com/lfyr/go-api/config/global"
	"github.com/lfyr/go-api/database/masterdb"
	"gorm.io/gorm"
)

type AppRolePrivilege struct {
	global.Model
	PriId        int          `json:"pri_id"`
	RoleId       int          `json:"role_id"`
	AppPrivilege AppPrivilege `json:"privilege" gorm:"foreignKey:Id;references:PriId"`
}

func NewAppRolePrivilege() *AppRolePrivilege {
	return &AppRolePrivilege{}
}

func (u *AppRolePrivilege) TableName() string {
	return "app_role_privilege"
}

func (this *AppRolePrivilege) List(whereMap map[string]interface{}, fieldSlice []string, page, size int, withSlice []string) (list []AppRolePrivilege, count int64) {
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

func (this *AppRolePrivilege) First(whereMap map[string]interface{}) (detail AppRolePrivilege) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	tx.Order("id desc").First(&detail)
	return
}

func (this *AppRolePrivilege) Many(whereMap map[string]interface{}) (list []AppRolePrivilege) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	tx.Find(&list)
	return
}

func (this *AppRolePrivilege) CreateInBatches(data []AppRolePrivilege, tx *gorm.DB) (err error) {
	err = tx.Model(this).CreateInBatches(&data, 50).Error
	return
}
