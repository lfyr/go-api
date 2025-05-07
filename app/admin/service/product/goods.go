package product

import (
	"github.com/lfyr/go-api/app/admin/handler/structure"
	"github.com/lfyr/go-api/database/masterdb"
	"github.com/lfyr/go-api/model"
)

type GoodsService struct{}

func NewGoodsService() *GoodsService {
	return &GoodsService{}
}

func (this *GoodsService) CreateGoods(data model.AppGoods, pics []structure.GoodsPics) (err error) {
	tx := masterdb.DB.Begin()
	// 创建商品
	gId, err := model.NewAppGoods().Create(&data, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	// 创建商品图片
	goodsPicsData := []model.AppGoodsPics{}
	for _, v := range pics {
		v.GoodsId = gId
		tmp := model.AppGoodsPics{
			Pic:      v.Pic,
			SmPic:    v.SmPic,
			GoodsId:  gId,
			IsDelete: 0,
		}
		goodsPicsData = append(goodsPicsData, tmp)
	}
	err = model.NewAppGoodsPics().BatchCreate(&goodsPicsData, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	// 创建商品属性

	tx.Commit()
	return
}
