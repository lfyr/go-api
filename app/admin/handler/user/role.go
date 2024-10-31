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

func (this *Role) Update(c *gin.Context) {
	param := UpdateRoleReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	data := model.AppRole{RoleName: param.RoleName}
	data.Id = param.Id
	err = user.NewRoleService().Update(data)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithMessage(c, "修改成功")
	return
}

func (this *Role) Del(c *gin.Context) {
	param := DelRoleReq{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	data := user.NewRoleService().FirstAdminRole(map[string]interface{}{"role_id = ?": param.Id})
	if data.Id > 0 {
		utils.FailWithMessage(c, "该角色存在关联管理员无法删除")
		return
	}
	err = user.NewRoleService().Del(param.Id)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithMessage(c, "删除成功")
	return
}
