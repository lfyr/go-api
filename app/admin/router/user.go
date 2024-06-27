package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/handler/user"
)

func UserRouter(routers *gin.RouterGroup) {
	routers.GET("/user/index", user.GetUser)
	routers.POST("/user/add", user.Add)
}
