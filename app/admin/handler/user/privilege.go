package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/admin/service/user"
	"github.com/lfyr/go-api/utils"
)

type Privilege struct{}

func NewPrivilegeRoute() *Privilege {
	return &Privilege{}
}

func (this *Privilege) List(c *gin.Context) {
	data := user.NewPrivilegeService().Many(map[string]interface{}{})
	utils.OkWithData(c, data)
	return
}

func (this *Privilege) Add(c *gin.Context) {

}

func (this *Privilege) Update(c *gin.Context) {

}

func (this *Privilege) Del(c *gin.Context) {

}
