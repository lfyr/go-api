package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/middleware"
	"github.com/lfyr/go-api/utils"
	"github.com/sirupsen/logrus"
)

func Router() (router *gin.Engine) {

	if utils.GVA_CONFIG.System.Env == "test" { //开发者
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetLevel(logrus.DebugLevel)
	} else { //非开发者
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.InfoLevel)
	}
	router = gin.New()
	router.Use(middleware.Cors()) // 直接放行全部跨域请求

	// 服务健康检测
	router.GET("/admin/health", func(c *gin.Context) {
		utils.Ok(c)
	})
	router.Use(middleware.ParseToken())
	router.Use(middleware.LoggerWithWriter(true), middleware.Recovery())

	// 内部功能
	adminRouter := router.Group("/admin")
	{
		userRouter(adminRouter)
		productRouter(adminRouter)
		commonRouter(adminRouter)
	}
	return
}
