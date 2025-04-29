package product

import (
	"github.com/gin-gonic/gin"
)

type GoodsService struct{}

func NewGoodsService() *GoodsService {
	return &GoodsService{}
}

func (this *GoodsService) Create(c *gin.Context) {

	return
}
