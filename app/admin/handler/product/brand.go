package product

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/service/product"
	"github.com/lfyr/go-api/config/global"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils"
)

type Brand struct{}

func NewBrandRoute() *Brand {
	return &Brand{}
}

func (this *Brand) List(c *gin.Context) {
	param := GetBrandReq{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	list, count := product.NewBrandService().List(map[string]interface{}{}, []string{}, param.Page, param.PageSize, []string{})
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithDetailed(c, map[string]interface{}{
		"list":  list,
		"count": count,
	}, "获取成功")
	return
}

func (this *Brand) Add(c *gin.Context) {
	param := AddBrandReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	data := model.AppBrand{
		BrandName: param.BrandName,
		Logo:      param.Logo,
	}
	err = product.NewBrandService().Add(data)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithMessage(c, "添加成功")
	return
}

func (this *Brand) Update(c *gin.Context) {
	param := UpdateBrandReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	data := model.AppBrand{
		Model: global.Model{
			Id: param.Id,
		},
		BrandName: param.BrandName,
		Logo:      param.Logo,
	}
	err = product.NewBrandService().Update(data)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithMessage(c, "更新成功")
	return
}
