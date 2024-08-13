package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/handler/goods"
	"github.com/lfyr/go-api/middleware"
)

func goodsRouter(routers *gin.RouterGroup) {
	userRouters := routers.Group("goods", middleware.LoginAuth())
	{
		userRouters.GET("/index", goods.List)
		userRouters.GET("/list", goods.Index)
		userRouters.POST("/edit", goods.Edit)
	}
}
