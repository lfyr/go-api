package model

import (
	"github.com/lfyr/go-api/config/global"
	"gorm.io/gorm"
)

type AppGoods struct {
	global.Model
	GoodsName      string  `json:"goods_name"`
	CatId          int     `json:"cat_id"`
	BrandId        int     `json:"brand_id"`
	ShopPrice      float64 `json:"shop_price"`
	Logo           string  `json:"logo"`
	SmLogo         string  `json:"sm_logo"`
	IsHot          int     `json:"is_hot"`
	IsNew          int     `json:"is_new"`
	IsBest         int     `json:"is_best"`
	IsOnSale       int     `json:"is_on_sale"`
	SeoKeyword     string  `json:"seo_key_word"`
	SeoDescription string  `json:"seo_description"`
	TypeId         int     `json:"type_id"`
	SortNum        int     `json:"sort_num"`
	IsDelete       int     `json:"is_delete"`
	GoodsDesc      string  `json:"goods_desc"`
	Addtime        int     `json:"addtime"`
}

func (AppGoods) TableName() string {
	return "app_goods"
}
func NewAppGoods() *AppGoods {
	return &AppGoods{}
}

func (this *AppGoods) Create(data *AppBrand, tx *gorm.DB) (err error) {
	err = tx.Model(this).Create(&data).Error
	return
}
