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
	}
}
