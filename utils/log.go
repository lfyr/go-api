package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/middleware"
	"github.com/sirupsen/logrus"
)

func LogInit() {
	if GVA_CONFIG.System.Env == "test" { //开发者
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetLevel(logrus.DebugLevel)
	} else { //非开发者
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func GetNewGinEngine(logBody bool) (router *gin.Engine) {
	LogInit()
	router = gin.New()
	router.Use(middleware.LoggerWithWriter(logBody), middleware.Recovery())
	return
}
