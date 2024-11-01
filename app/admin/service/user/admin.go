package user

import (
	"github.com/lfyr/go-api/database/masterdb"
	"github.com/lfyr/go-api/model"
)

type AdminService struct{}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (this *AdminService) List(whereMap map[string]interface{}, fieldSlice []string, page int, size int, withSlice []string) (list []model.ListRsp, count int64) {
	list, count = model.NewAppAdmin().List(whereMap, fieldSlice, page, size, withSlice)
	return
}

func (this *AdminService) Add(data model.AppAdmin) (err error) {
	err = model.NewAppAdmin().Create(&data)
	return
}

func (this *AdminService) Update(data model.AppAdmin) (err error) {
	err = model.NewAppAdmin().Update(data.Id, map[string]interface{}{"is_use": data.IsUse})
	return
}

func (this *AdminService) GetUserById(id int, withSlice []string) (user model.AppAdmin) {
	user = model.NewAppAdmin().First(map[string]interface{}{"id": id}, withSlice)
	return
}

func (this *AdminService) Delete(ids []int) (err error) {
	err = masterdb.DB.Model(model.AppAdmin{}).Where("id in (?)", ids).Delete(model.AppAdmin{}).Error
	return err
}
