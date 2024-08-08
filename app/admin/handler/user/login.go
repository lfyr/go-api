package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/service/user"
	"github.com/lfyr/go-api/utils"
)

func Login(c *gin.Context) {
	param := LoginReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}

	// 登陆逻辑
	u, err := user.NewUserService().Login(param.Phone, param.Password)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithData(c, u)
	return
}
