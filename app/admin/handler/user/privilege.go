package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/service/user"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils"
)

type Privilege struct{}

func NewPrivilegeRoute() *Privilege {
	return &Privilege{}
}

func (this *Privilege) List(c *gin.Context) {
	data := user.NewPrivilegeService().Many(map[string]interface{}{})
	rData := []PrivilegeTree{}
	if len(data) > 0 {
		rData = getTree(data, 0)
	}
	utils.OkWithData(c, rData)
	return
}

func (this *Privilege) Add(c *gin.Context) {
	param := AddPrivilegeReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	data := model.AppPrivilege{PriName: param.PriName, ActionName: param.ActionName, ParentId: param.Pid}
	err = user.NewPrivilegeService().Add(data)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithMessage(c, "添加成功")
	return
}

func (this *Privilege) Update(c *gin.Context) {
	param := UpDatePrivilegeReq{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	data := model.AppPrivilege{PriName: param.PriName, ActionName: param.ActionName, ParentId: param.ParentId}
	data.Id = param.Id
	err = user.NewPrivilegeService().UpData(data)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithMessage(c, "修改成功")
	return
}

func (this *Privilege) Del(c *gin.Context) {
	param := DelPrivilegeReq{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	data := user.NewPrivilegeService().GetPriRoleByCond(map[string]interface{}{"pri_id = ?": param.Id}, []string{"pri_id"})
	if len(data) > 0 {
		utils.FailWithMessage(c, "该权限存在关联角色无法删除")
		return
	}
	err = user.NewPrivilegeService().DelData(param.Id)
	if err != nil {
		utils.FailWithMessage(c, err.Error())
		return
	}
	utils.OkWithMessage(c, "删除成功")
	return
}

//func getTree(data []model.AppPrivilege, pid int) (dataTree []PrivilegeTree) {
//	for _, item := range data {
//		if item.ParentId == pid {
//			pri := PrivilegeTree{
//				Id:         item.Id,
//				PriName:    item.PriName,
//				ActionName: item.ActionName,
//				ParentId:   item.ParentId,
//			}
//			child := getTree(data, item.Id)
//			pri.Children = child
//			dataTree = append(dataTree, pri)
//		}
//	}
//	return dataTree
//}
