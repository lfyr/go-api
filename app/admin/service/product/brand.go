package product

import "github.com/lfyr/go-api/model"

type BrandService struct{}

func NewBrandService() *BrandService {
	return &BrandService{}
}

func (this *BrandService) List(whereMap map[string]interface{}, fieldSlice []string, page, size int, withSlice []string) (list []model.AppBrand, count int64) {
	list, count = model.NewAppBrand().List(whereMap, fieldSlice, page, size, withSlice)
	return
}

func (this *BrandService) Add(data model.AppBrand) (err error) {
	err = model.NewAppBrand().Create(&data)
	return
}

func (this *BrandService) Update(data model.AppBrand) (err error) {
	upData := map[string]interface{}{
		"brand_name": data.BrandName,
		"logo":       data.Logo,
	}
	err = model.NewAppBrand().Update(map[string]interface{}{"id=?": data.Id}, upData)
	return
}

func (this *BrandService) Delete(data model.AppBrand) (err error) {
	upData := map[string]interface{}{
		"delete_status": 1,
	}
	err = model.NewAppBrand().Update(map[string]interface{}{"id = ?": data.Id}, upData)
	return
}
