package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/handler/product"
	"github.com/lfyr/go-api/middleware"
)

func productRouter(routers *gin.RouterGroup) {
	productRouters := routers.Group("product", middleware.LoginAuth())
	{
		brandApi := product.NewBrandRoute()
		brandRouters := productRouters.Group("brand")
		{
			brandRouters.GET("list", brandApi.List)
			brandRouters.POST("create", brandApi.Add)
			brandRouters.POST("update", brandApi.Update)
			brandRouters.GET("del", brandApi.Del)
		}

		categoryApi := product.NewCategoryRoute()
		categoryRouters := productRouters.Group("category")
		{
			categoryRouters.GET("list", categoryApi.List)
			categoryRouters.POST("create", categoryApi.Add)
			categoryRouters.POST("update", categoryApi.Update)
			categoryRouters.GET("del", categoryApi.Del)
		}

		goodsApi := product.NewGoodsRoute()
		goodsRouters := productRouters.Group("category")
		{
			goodsRouters.GET("list", goodsApi.List)
			goodsRouters.POST("create", goodsApi.Add)
		}
	}
}
