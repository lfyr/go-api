package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/middleware"
	"github.com/lfyr/go-api/utils"
)

func Router() (router *gin.Engine) {
	router = gin.New()
	router.Use(middleware.Cors()) // 直接放行全部跨域请求

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
