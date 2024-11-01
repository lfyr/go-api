package token

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils"
	"github.com/lfyr/go-api/utils/redis"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func SetToken(user model.User) (token string, err error) {
	token = getUniquenessToken()
	err = setLoginInfoByToken(token, user)
	return
}

func getUniquenessToken() (token string) {
	token = generatorToken(utils.RandomSalt(6))
	return
}

func setLoginInfoByToken(token string, user model.User) (err error) {

	whereMap := map[string]interface{}{
		"id =?": user.Id,
	}
	userInfo := model.NewUser().First(whereMap)
	if userInfo.Token != "" {
		// 删除上一次在该类型登录的 Token
		client := redis.NewDefaultRedisStore(0)
		err = client.Del(userInfo.Token)
		if err != nil {
			return err
		}
	}

	// 更新 loginInfo
	data := map[string]interface{}{
		"token": token,
	}
	err = model.NewUser().Update(user.Id, data)
	if err != nil {
		return err
	}

	// 保存到 Redis
	err = SaveRedisToken(token, user)
	if err != nil {
		return
	}
	return
}

func generatorToken(key string) (token string) {
	key = key + strconv.Itoa(int(time.Now().Unix())) + utils.RandomSalt(6)
	md5 := crypto.MD5.New()
	md5.Write([]byte(key))
	token = hex.EncodeToString(md5.Sum(nil))
	return
}

func GetRedisUserID(token string) (uid int) {
	info, err2 := GetUserInfoByToken(token)
	if err2 != nil {
		return
	}
	uid = info.Id
	return
}

func GetUserInfoByToken(token string) (user model.User, err error) {
	client := redis.NewDefaultRedisStore(0)
	result := client.Get(token, false)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		logrus.Error("getUserInfoByToken!", err.Error())
		return
	}
	return
}

func SaveRedisToken(token string, user model.User) (err error) {
	client := redis.NewDefaultRedisStore(utils.GVA_CONFIG.System.TokenExpire)
	userStr, err := json.Marshal(user)
	if err != nil {
		logrus.Error("SaveRedisToken", err.Error())
		return
	}
	err = client.SetNX(token, string(userStr))
	if err != nil {
		logrus.Error("SaveRedisToken", err.Error())
		return
	}
	return
}

func GetUid(c *gin.Context) int {
	if userId, exists := c.Get("user_id"); !exists {
		return 0
	} else {
		if value, ok := userId.(int); ok {
			return value
		} else {
			return 0
		}
	}
}

func GetTokenFromHeader(c *gin.Context) string {
	return c.Request.Header.Get("Authorization")
}

func GetUserInfo(c *gin.Context) model.User {
	if user, exists := c.Get("user"); !exists {
		return model.User{}
	} else {
		if value, ok := user.(model.User); ok {
			return value
		} else {
			return model.User{}
		}
	}
}

func DelRedisToken(id int, token string) (err error) {
	client := redis.NewDefaultRedisStore(0)
	err = client.Del(token)
	if err != nil {
		logrus.Error("DelRedisToken", err.Error())
		return
	}
	data := map[string]interface{}{
		"token": "",
	}
	err = model.NewUser().Update(id, data)
	if err != nil {
		logrus.Error("DelToken", err.Error())
		return
	}
	return
}
