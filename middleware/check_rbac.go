package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/service/user"
	"github.com/lfyr/go-api/utils"
	"github.com/lfyr/go-api/utils/token"
)

func CheckPrivilege() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !checkPrivilege(c) {
			utils.FailWithMessage(c, "暂无权限")
			c.Abort()
		}
		c.Next()
		return
	}
}

func checkPrivilege(c *gin.Context) bool {

	// 如果是超级管理员直接返回所有权限
	userId := token.GetUid(c)
	if token.GetUid(c) == 1 {
		return true
	}

	// 获取用户信息
	//adminUser := model.NewAppAdmin().First(map[string]interface{}{"user_id = ?": userId}, []string{"Role"})

	// 获取地址
	path := c.Request.URL.Path
	if path == "/admin/user/login" || path == "/admin/user/logout" {
		return true
	}

	// 获取用户所有权限  先获取用户id->角色->权限 最后通过比对判断是否具有访问权限
	menuList := []string{}
	data := user.NewAdminService().GetUserById(userId, []string{"Role.AppRolePrivilege.AppPrivilege"})
	for _, role := range data.Role {
		for _, privilege := range role.AppRolePrivilege {
			if len(privilege.AppPrivilege.MenuName) > 0 {
				menuList = append(menuList, privilege.AppPrivilege.MenuName)
			}
		}
	}
	c.Set("menuList", menuList)
	for _, role := range data.Role {
		for _, privilege := range role.AppRolePrivilege {
			if privilege.AppPrivilege.ActionName == path {
				return true
			}
		}
	}
	return false
}
