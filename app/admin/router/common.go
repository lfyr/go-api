package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/handler/common"
	"github.com/lfyr/go-api/middleware"
)

func commonRouter(routers *gin.RouterGroup) {
	commonRouters := routers.Group("common", middleware.LoginAuth())
	{
		fileApi := common.NewFileRoute()
		fileRouters := commonRouters.Group("file")
		{
			fileRouters.POST("fileUpload", fileApi.FileUpload)
		}
	}
}
