package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/handler/user"
	"github.com/lfyr/go-api/middleware"
)

func userRouter(routers *gin.RouterGroup) {

	userApi := user.NewUserRoute()
	routers.POST("/user/login", userApi.Login)
	adminRouter := routers.Group("", middleware.LoginAuth(), middleware.CheckPrivilege())
	{
		userRouters := adminRouter.Group("user")
		{
			userRouters.GET("/list", userApi.List)
			userRouters.POST("/add", userApi.Add)
			userRouters.GET("/info", userApi.Info)
			userRouters.POST("/logout", userApi.Logout)
			userRouters.GET("/toAssign", userApi.ToAssign)
		}

		roleRouters := adminRouter.Group("role")
		roleApi := user.NewRoleRoute()
		{
			roleRouters.GET("/list", roleApi.List)
			roleRouters.POST("/add", roleApi.Add)
			roleRouters.POST("/update", roleApi.Update)
			roleRouters.GET("/delete", roleApi.Del)
		}

		privilegeRouters := adminRouter.Group("privilege")
		privilegeApi := user.NewRoleRoute()
		{
			privilegeRouters.GET("/list", privilegeApi.List)
			privilegeRouters.POST("/add", privilegeApi.Add)
			privilegeRouters.POST("/update", privilegeApi.Update)
			privilegeRouters.GET("/delete", privilegeApi.Del)
		}

	}

}
