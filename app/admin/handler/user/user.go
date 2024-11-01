package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/service/user"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils"
	"github.com/lfyr/go-api/utils/token"
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
	u, err := user.NewUserService().Login(param.UserName, param.Password)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	c.SetCookie("token", u.Token, utils.GVA_CONFIG.System.TokenExpire, "/", c.Request.Host, false, true)
	utils.OkWithData(c, u)
	return
}

func (this *User) Logout(c *gin.Context) {
	id := token.GetUid(c)
	tk := token.GetTokenFromHeader(c)
	err := token.DelRedisToken(id, tk)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.Ok(c)
	return
}

func (this *User) Update(c *gin.Context) {
	param := UpdateUserReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	data := model.User{}
	data.Id = param.Id
	data.UserName = param.UserName
	data.Phone = param.Phone
	err = user.NewUserService().Update(data)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithMessage(c, "修改成功")
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
		Phone: param.Phone,
	}
	err = user.NewUserService().Add(data)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithData(c, true)
	return
}

func (this *User) Delete(c *gin.Context) {
	param := DelUserReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}

	err = user.NewUserService().Delete(param.Ids)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithData(c, true)
	return
}

func (this *User) List(c *gin.Context) {
	param := GetUserReq{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	list, count := user.NewUserService().List(map[string]interface{}{}, []string{}, param.Page, param.PageSize, []string{"Admin", "Role.AppRolePrivilege.AppPrivilege"})
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithDetailed(c, map[string]interface{}{
		"list":  list,
		"count": count,
	}, "获取成功")
	return
}
