package user

import (
	"github.com/lfyr/go-api/database/masterdb"
	"github.com/lfyr/go-api/model"
)

type RoleService struct{}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (r *RoleService) List(whereMap map[string]interface{}, fieldSlice []string, page, size int, withSlice []string) (list []model.AppRole, count int64) {
	list, count = model.NewAppRole().List(whereMap, fieldSlice, page, size, withSlice)
	return
}

func (r *RoleService) Many(whereMap map[string]interface{}) (list []model.AppRole) {
	list = model.NewAppRole().Many(whereMap)
	return
}

func (r *RoleService) Add(data model.AppRole) (err error) {
	err = model.NewAppRole().Create(&data)
	return
}

func (r *RoleService) Update(data model.AppRole) (err error) {
	upData := map[string]interface{}{
		"role_name": data.RoleName,
	}
	err = model.NewAppRole().Update(data.Id, upData)
	return
}

func (r *RoleService) Del(id int) (err error) {
	err = model.NewAppRole().Delete(id)
	return
}

func (r *RoleService) FirstAdminRole(whereMap map[string]interface{}) (data model.AppAdminRole) {
	data = model.NewAppAdminRole().First(whereMap)
	return
}

func (r *RoleService) FindAdminRole(whereMap map[string]interface{}) (data []model.AppAdminRole) {
	data = model.NewAppAdminRole().Many(whereMap)
	return
}

func (r *RoleService) AddAdminRole(userId int, data []model.AppAdminRole) (err error) {

	tx := masterdb.DB.Begin()
	err = model.NewAppAdminRole().DeleteByAdminId(userId, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	if len(data) > 0 {
		err = model.NewAppAdminRole().CreateInBatches(data, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}
