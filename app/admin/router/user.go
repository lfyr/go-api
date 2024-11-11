package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/handler/user"
	"github.com/lfyr/go-api/middleware"
)

func userRouter(routers *gin.RouterGroup) {

	userApi := user.NewUserRoute()
	adminApi := user.NewAdminRoute()
	routers.POST("/user/login", userApi.Login)
	adminRouter := routers.Group("", middleware.LoginAuth(), middleware.CheckPrivilege())
	{
		userRouters := adminRouter.Group("user")
		{
			userRouters.GET("/list", adminApi.List)
			userRouters.POST("/add", adminApi.Add)
			userRouters.POST("/update", adminApi.Update)
			userRouters.POST("/delete", adminApi.Delete)
			userRouters.GET("/info", adminApi.Info)
			userRouters.POST("/logout", userApi.Logout)
			userRouters.GET("/toAssign", adminApi.ToAssign)
			userRouters.POST("/doAssign", adminApi.DoAssign)
		}

		roleRouters := adminRouter.Group("role")
		roleApi := user.NewRoleRoute()
		{
			roleRouters.GET("/list", roleApi.List)
			roleRouters.POST("/add", roleApi.Add)
			roleRouters.POST("/update", roleApi.Update)
			roleRouters.GET("/delete", roleApi.Del)
			roleRouters.GET("/toAssign", roleApi.ToAssign)
			roleRouters.POST("/doAssign", roleApi.DoAssign)
		}

		privilegeRouters := adminRouter.Group("privilege")
		privilegeApi := user.NewPrivilegeRoute()
		{
			privilegeRouters.GET("/list", privilegeApi.List)
			privilegeRouters.POST("/add", privilegeApi.Add)
			privilegeRouters.POST("/update", privilegeApi.Update)
			privilegeRouters.GET("/delete", privilegeApi.Del)
		}

	}

}
