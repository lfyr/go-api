package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/utils"
	"github.com/lfyr/go-api/utils/token"
)

func ParseToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokens := getTokenByContext(c)
		if tokens == "" {
			c.Next()
			return
		}
		c.Set("token", tokens)
		user, _ := token.GetUserInfoByToken(tokens)
		c.Set("user_id", user.Id)
		c.Set("user_name", user.UserName)
		c.Set("user", user)
		c.Next()
		return
	}
}

func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := token.GetUid(c)
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

func getTokenByContext(c *gin.Context) (t string) {
	t = token.GetTokenFromHeader(c)
	if t == "" {
		t, _ = c.Cookie("token")
	}
	return
}
