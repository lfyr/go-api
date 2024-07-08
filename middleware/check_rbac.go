package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/utils"
	"github.com/lfyr/go-api/utils/token"
)

func CheckPrivilege() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.CheckPrivilege(c)
		if err != nil {
			utils.FailWithMessage(c, "暂无权限")
			c.Abort()
		}
		c.Next()
		return
	}
}
