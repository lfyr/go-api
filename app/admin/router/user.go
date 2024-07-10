package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/handler/user"
)

func userRouter(routers *gin.RouterGroup) {
	routers.GET("/user/index", user.GetUser)
	routers.GET("/user/list", user.GetUser)
	routers.POST("/user/add", user.Add)
}
