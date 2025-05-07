package product

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/handler/structure"
	"github.com/lfyr/go-api/app/admin/service/product"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils"
	"time"
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
	param := structure.AddGoodsReq{}
	if err := c.ShouldBindJSON(&param); err != nil {
		utils.FailWithDetailed(c, map[string]interface{}{}, "参数错误")
		return
	}

	goodsData := model.AppGoods{
		GoodsName:      param.GoodsName,
		CatId:          param.CatId,
		BrandId:        param.BrandId,
		ShopPrice:      param.ShopPrice,
		Logo:           param.Logo,
		SmLogo:         param.SmLogo,
		IsHot:          param.IsHot,
		IsNew:          param.IsNew,
		IsBest:         param.IsBest,
		IsOnSale:       param.IsOnSale,
		SeoKeyword:     param.SeoKeyword,
		SeoDescription: param.SeoDescription,
		TypeId:         param.TypeId,
		SortNum:        param.SortNum,
		IsDelete:       0,
		GoodsDesc:      param.GoodsDesc,
		Addtime:        int(time.Now().Unix()),
	}

	if err := product.NewGoodsService().CreateGoods(goodsData, param.GoodsPicsData); err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}

	utils.OkWithDetailed(c, map[string]interface{}{}, "创建成功")
	return
}
