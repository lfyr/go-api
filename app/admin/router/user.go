package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/handler/user"
	"github.com/lfyr/go-api/middleware"
)

func userRouter(routers *gin.RouterGroup) {

	userApi := user.NewUserRoute()
	routers.POST("/user/login", userApi.Login)
	userRouters := routers.Group("user", middleware.LoginAuth())
	{
		userRouters.POST("/add", userApi.Add)
		userRouters.GET("/info", userApi.Info)
		userRouters.POST("/logout", userApi.Logout)
	}

	roleRouters := routers.Group("role", middleware.LoginAuth())
	roleApi := user.NewRoleRoute()
	{
		roleRouters.GET("/list", roleApi.List)
		roleRouters.POST("/add", roleApi.Add)
	}
}
