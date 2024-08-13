package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/service/user"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils"
)

type User struct{}

func NewUserRoute() *User {
	return &User{}
}

func (this *User) Login(c *gin.Context) {
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

func (this *User) List(c *gin.Context) {
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

func (this *User) Add(c *gin.Context) {
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
