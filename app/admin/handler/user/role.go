package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/service/user"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils"
)

type Role struct{}

func NewRoleRoute() *Role {
	return &Role{}
}

func (this *Role) List(c *gin.Context) {
	param := GetRoleReq{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	list, count := user.NewRoleService().List(map[string]interface{}{}, []string{}, param.Page, param.PageSize, []string{})
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

func (this *Role) Add(c *gin.Context) {
	param := AddRoleReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	data := model.AppRole{RoleName: param.RoleName}
	err = user.NewRoleService().Add(data)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithMessage(c, "添加成功")
	return
}
