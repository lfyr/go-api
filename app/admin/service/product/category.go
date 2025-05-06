package product

import (
	"github.com/lfyr/go-api/model"
)

type CategoryService struct{}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (this *CategoryService) List(whereMap map[string]interface{}, fieldSlice []string, page, size int, withSlice []string) (list []model.AppCategory, count int64) {
	list, count = model.NewAppCategory().List(whereMap, fieldSlice, page, size, withSlice)
	return
}

func (this *CategoryService) Create(data model.AppCategory) (err error) {
	err = model.NewAppCategory().Create(&data)
	return
}

func (this *CategoryService) Update(data model.AppCategory) (err error) {
	upData := map[string]interface{}{
		"cat_name":  data.CatName,
		"parent_id": data.ParentId,
	}
	err = model.NewAppCategory().Update(map[string]interface{}{"id=?": data.Id}, upData)
	return
}

func (this *CategoryService) Delete(id int) (err error) {
	upData := map[string]interface{}{
		"delete_status": 1,
	}
	err = model.NewAppCategory().Update(map[string]interface{}{"id = ?": id}, upData)
	return
}
