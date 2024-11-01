package user

import (
	"github.com/lfyr/go-api/database/masterdb"
	"github.com/lfyr/go-api/model"
)

type PrivilegeService struct{}

func (r *PrivilegeService) Add(data model.AppPrivilege) (err error) {
	err = model.NewAppPrivilege().Create(&data)
	return
}

func (r *PrivilegeService) UpData(data model.AppPrivilege) (err error) {
	upData := map[string]interface{}{
		"role_name":   data.PriName,
		"action_name": data.ActionName,
		"parent_id":   data.ParentId,
	}
	err = model.NewAppPrivilege().Update(data.Id, upData)
	return
}

func (r *PrivilegeService) delData(id int) (err error) {
	err = model.NewAppPrivilege().Delete(id)
	return
}

func (r *PrivilegeService) AddRolePrivilege(roleId int, data []model.AppRolePrivilege) (err error) {
	tx := masterdb.DB.Begin()
	err = model.NewAppAdminRole().DeleteByRoleId(roleId, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	err = model.NewAppRolePrivilege().CreateInBatches(data, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
