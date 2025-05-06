package product

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/utils"
)

type Goods struct {
}

func NewGoodsRoute() *Goods {
	return &Goods{}
}

func (this *Goods) List(c *gin.Context) {
	utils.OkWithDetailed(c, map[string]interface{}{}, "获取成功")
	return
}

func (this *Goods) Add(c *gin.Context) {
	param := AddGoodsReq{}
	if err := c.ShouldBindJSON(&param); err != nil {
		utils.FailWithDetailed(c, map[string]interface{}{}, "参数错误")
		return
	}
	utils.OkWithDetailed(c, map[string]interface{}{}, "创建成功")
	return
}
