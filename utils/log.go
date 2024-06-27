package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogInit() {
	//if gin_utils.IsDev() { //开发者
	//	logrus.SetFormatter(&logrus.TextFormatter{})
	//	logrus.SetLevel(logrus.DebugLevel)
	//} else { //非开发者
	//	logrus.SetFormatter(&logrus.JSONFormatter{})
	//	logrus.SetLevel(logrus.InfoLevel)
	//}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}

func GetNewGinEngine(logBody bool) (router *gin.Engine) {
	LogInit()
	router = gin.New()
	return
}
