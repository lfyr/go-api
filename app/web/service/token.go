package service

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils/redis"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"
)

const sevenDaySecond = 7 * 24 * 60 * 60

func SetToken(user model.User) (token string, err error) {
	token = getUniquenessToken()
	err = setLoginInfoByToken(token, user)
	return
}

func getUniquenessToken() (token string) {
	token = generatorToken(randomSalt())
	return
}

func setLoginInfoByToken(token string, user model.User) (err error) {

	whereMap := map[string]interface{}{
		"id =?": user.ID,
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
	err = model.NewUser().Update(user.ID, data)
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
	key = key + strconv.Itoa(int(time.Now().Unix())) + randomSalt()
	md5 := crypto.MD5.New()
	md5.Write([]byte(key))
	token = hex.EncodeToString(md5.Sum(nil))
	return
}

func randomSalt() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRedisUserID(token string) (uid int) {
	info, err2 := GetUserInfoByToken(token)
	if err2 != nil {
		return
	}
	uid = info.ID
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
	client := redis.NewDefaultRedisStore(sevenDaySecond)
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
