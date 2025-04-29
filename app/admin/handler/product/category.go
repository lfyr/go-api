package product

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/service/product"
	"github.com/lfyr/go-api/config/global"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils"
	"strconv"
)

type Category struct {
}

func NewCategoryRoute() *Category {
	return &Category{}
}

func (this *Category) List(c *gin.Context) {
	utils.OkWithDetailed(c, map[string]interface{}{}, "获取成功")
	return
}

func (this *Category) Add(c *gin.Context) {
	param := AddCategoryReq{}
	if err := c.ShouldBind(&param); err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}

	data := model.AppCategory{
		CatName:  param.CatName,
		ParentId: param.ParentId,
	}
	if err := product.NewCategoryService().Create(data); err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithDetailed(c, map[string]interface{}{}, "创建成功")
	return
}

func (this *Category) Update(c *gin.Context) {
	param := UpdateCategoryReq{}
	if err := c.ShouldBind(&param); err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}

	data := model.AppCategory{
		Model: global.Model{
			Id: param.Id,
		},
		CatName:  param.CatName,
		ParentId: param.ParentId,
	}
	if err := product.NewCategoryService().Update(data); err != nil {
		utils.FailWithMessage(c, err.Error())
	}
}

func (this *Category) Del(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	if err = product.NewCategoryService().Delete(idInt); err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithMessage(c, "删除成功")
	return
}
