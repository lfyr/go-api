package model

import (
	"github.com/lfyr/go-api/config/global"
	"gorm.io/gorm"
)

type AppGoodsPics struct {
	global.Model
	Pic      string `json:"pic"`
	SmPic    string `json:"sm_pic"`
	GoodsId  int    `json:"goods_id"`
	IsDelete int    `json:"isDelete"`
}

func (AppGoodsPics) TableName() string {
	return "app_goods_pics"
}
func NewAppGoodsPics() *AppGoodsPics {
	return &AppGoodsPics{}
}

func (this *AppGoodsPics) BatchCreate(data *[]AppGoodsPics, tx *gorm.DB) (err error) {
	err = tx.Model(this).CreateInBatches(&data, 500).Error
	return
}
