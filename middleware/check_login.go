package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/app/web/service"
	"github.com/lfyr/go-api/utils"
)

func ParseToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getTokenByContext(c)
		if token == "" {
			c.Next()
			return
		}
		c.Set("token", token)
		user, _ := service.GetUserInfoByToken(token)

		c.Set("user_id", user.ID)
		c.Set("user_name", user.UserName)
		c.Set("token", token)
		c.Next()
		return
	}
}

func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := service.GetUid(c)
		if uid > 0 {
			c.Next()
			return
		} else {
			utils.FailWithMessage(c, "用户未登录")
			c.Abort()
			return
		}
	}
}

func getTokenByContext(c *gin.Context) (token string) {
	token = service.GetTokenFromHeader(c)
	if token == "" {
		token, _ = c.Cookie("token")
	}
	return
}
