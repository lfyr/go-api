package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/service/user"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils"
)

func GetUser(c *gin.Context) {
	param := GetUserReq{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		fmt.Println(err)
		utils.FailWithMessage(c, err.Error())
		return
	}
	data := user.NewUserService().GetUserById(param.Id)
	utils.OkWithData(c, data)
	return
}

func Add(c *gin.Context) {
	param := AddUserReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	data := model.User{
		UserName: param.UserName,
		Password: param.Password,
		Email:    param.Email,
		Phone:    param.Phone,
	}
	err = user.NewUserService().Add(data)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithData(c, true)
	return
}
