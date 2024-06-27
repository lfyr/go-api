package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/utils"
)

func Router() (router *gin.Engine) {
	router = gin.New()

	// 服务健康检测
	router.GET("/admin/health", func(c *gin.Context) {
		utils.Ok(c)
	})
	// 内部功能
	adminRouter := router.Group("/admin")
	{
		UserRouter(adminRouter)
	}
	return
}
