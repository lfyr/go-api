package user

import (
	"github.com/lfyr/go-api/database/masterdb"
	"github.com/lfyr/go-api/model"
)

type RoleService struct{}

func (r *RoleService) Add(data model.AppRole) (err error) {
	err = model.NewAppRole().Create(&data)
	return
}

func (r *RoleService) UpData(data model.AppRole) (err error) {
	upData := map[string]interface{}{
		"role_name": data.RoleName,
	}
	err = model.NewAppRole().Update(data.Id, upData)
	return
}

func (r *RoleService) delData(id int) (err error) {
	err = model.NewAppRole().Delete(id)
	return
}

func (r *RoleService) AddAdminRole(roleId int, data []model.AppAdminRole) (err error) {
	tx := masterdb.DB.Begin()
	err = model.NewAppAdminRole().DeleteByRoleId(roleId, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	err = model.NewAppAdminRole().CreateInBatches(data, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
