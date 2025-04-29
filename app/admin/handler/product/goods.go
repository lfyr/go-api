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

func (this *Goods) Create(c *gin.Context) {
	utils.OkWithDetailed(c, map[string]interface{}{}, "创建成功")
}
