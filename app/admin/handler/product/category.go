package product

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/handler/structure"
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
	data := product.NewCategoryService().Many(map[string]interface{}{})
	rData := []structure.CategoryTree{}
	if len(data) > 0 {
		rData = getTree(data, 0)
	}
	utils.OkWithData(c, rData)
	return
}

func (this *Category) Add(c *gin.Context) {
	param := structure.AddCategoryReq{}
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
	param := structure.UpdateCategoryReq{}
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
		return
	}
	utils.OkWithMessage(c, "更新成功")
	return
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

func getTree(data []model.AppCategory, pid int) (dataTree []structure.CategoryTree) {
	for _, item := range data {
		if item.ParentId == pid {
			pri := structure.CategoryTree{
				Id:       item.Id,
				CatName:  item.CatName,
				ParentId: item.ParentId,
			}
			child := getTree(data, item.Id)
			pri.Children = child
			dataTree = append(dataTree, pri)
		}
	}
	return dataTree
}
