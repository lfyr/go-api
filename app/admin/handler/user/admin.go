package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/service/user"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils"
	"github.com/lfyr/go-api/utils/token"
	"github.com/samber/lo"
	"strings"
)

type AdminRoute struct{}

func NewAdminRoute() *AdminRoute {
	return &AdminRoute{}
}

func (this *AdminRoute) List(c *gin.Context) {
	param := GetAdminListReq{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	list, count := user.NewAdminService().List(map[string]interface{}{"user_name like ?": "%" + param.UserName + "%"}, []string{"app_admin.*", "user.user_name,user.nickname,user.phone"}, param.Page, param.PageSize, []string{"Role.AppRolePrivilege.AppPrivilege"})
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	rData := []GetAdminListRsp{}
	for _, admin := range list {
		privilegeStr := ""
		for _, role := range admin.Role {
			for _, rolePrivilege := range role.AppRolePrivilege {
				privilegeStr += "," + rolePrivilege.AppPrivilege.PriName
			}
		}
		tmp := GetAdminListRsp{
			Id:           admin.Id,
			UserId:       admin.UserId,
			IsUse:        admin.IsUse,
			UserName:     admin.UserName,
			Nickname:     admin.Nickname,
			Phone:        admin.Phone,
			PrivilegeStr: strings.Trim(privilegeStr, ","),
		}
		rData = append(rData, tmp)
	}
	utils.OkWithDetailed(c, map[string]interface{}{
		"list":  rData,
		"count": count,
	}, "获取成功")
	return
}

func (this *AdminRoute) Info(c *gin.Context) {
	userId := token.GetUid(c)
	data := user.NewAdminService().GetUserById(userId, []string{"User"})
	rData := InfoReq{
		Id:       data.Id,
		UserId:   data.UserId,
		IsUse:    data.IsUse,
		UserName: data.User.UserName,
		Nickname: data.User.Nickname,
		Password: data.User.Password,
		Email:    data.User.Email,
		Phone:    data.User.Phone,
		Ip:       data.User.Ip,
		Token:    data.User.Token,
		Avatar:   data.User.Avatar,
	}

	utils.OkWithData(c, rData)
	return
}

func (this *AdminRoute) ToAssign(c *gin.Context) {
	param := ToAssignReq{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	list := user.NewRoleService().FindAdminRole(map[string]interface{}{"admin_id = ?": param.Id})
	allRole := user.NewRoleService().Many(map[string]interface{}{})

	roleId := lo.Map(list, func(role model.AppAdminRole, _ int) int {
		return role.RoleId
	})
	utils.OkWithDetailed(c, map[string]interface{}{
		"allRolesList": allRole,
		"assignRoles":  roleId,
	}, "获取成功")
	return
}

func (this *AdminRoute) DoAssign(c *gin.Context) {
	param := AddAdminRoleReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	data := []model.AppAdminRole{}
	for _, v := range param.RoleId {
		tmp := model.AppAdminRole{
			AdminId: param.AdminId,
			RoleId:  v,
		}
		data = append(data, tmp)
	}
	err = user.NewRoleService().AddAdminRole(param.AdminId, data)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithMessage(c, "添加成功")
	return
}

func (this *AdminRoute) Add(c *gin.Context) {
	param := AddAdminReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}

	userInfo := user.NewUserService().First(map[string]interface{}{"phone = ?": param.Phone})
	if userInfo.Id <= 0 {
		utils.FailWithMessage(c, "不存在用户请先添加用户")
		return
	}

	data := model.AppAdmin{
		UserId: userInfo.Id,
	}
	err = user.NewAdminService().Add(data)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithData(c, true)
	return
}

func (this *AdminRoute) Update(c *gin.Context) {
	param := UpdateAdminReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}

	data := model.AppAdmin{
		IsUse: param.IsUse,
	}
	data.Id = param.Id

	err = user.NewAdminService().Update(data)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithData(c, true)
	return
}

func (this *AdminRoute) Delete(c *gin.Context) {
	param := DelUserReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}

	err = user.NewAdminService().Delete(param.Ids)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithData(c, true)
	return
}
