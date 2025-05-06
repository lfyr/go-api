package product

import (
	"fmt"
	"github.com/lfyr/go-api/database/masterdb"
	"github.com/lfyr/go-api/model"
)

type GoodsService struct{}

func NewGoodsService() *GoodsService {
	return &GoodsService{}
}

func (this *GoodsService) CreateGoods(data model.AppGoods) (err error) {
	tx := masterdb.DB.Begin()
	// 创建商品
	gId, err := model.NewAppGoods().Create(&data, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	fmt.Println(gId)

	// 创建商品图片

	// 创建商品属性

	tx.Commit()
	return
}
