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

func (this *User) Info(c *gin.Context) {
	userId := token.GetUid(c)
	data := user.NewUserService().GetUserById(userId, []string{})
	utils.OkWithData(c, data)
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

func (this *User) List(c *gin.Context) {
	param := GetUserReq{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	list, count := user.NewUserService().List(map[string]interface{}{}, []string{}, param.Page, param.PageSize, []string{"Admin"})
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

func (this *User) ToAssign(c *gin.Context) {
	param := ToAssignReq{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	list := user.NewRoleService().FindAdminRole(map[string]interface{}{"user_id = ?": param.Id})
	role := user.NewRoleService().Many(map[string]interface{}{})

	utils.OkWithDetailed(c, map[string]interface{}{
		"allRolesList": list,
		"assignRoles":  role,
	}, "获取成功")
	return
}
