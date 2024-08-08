package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/handler/user"
	"github.com/lfyr/go-api/middleware"
)

func userRouter(routers *gin.RouterGroup) {
	routers.POST("/login", user.Login)
	userRouters := routers.Group("user", middleware.LoginAuth())
	{
		userRouters.GET("/index", user.GetUser)
		userRouters.GET("/list", user.GetUser)
		userRouters.POST("/add", user.Add)
	}

}
