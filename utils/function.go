package utils

import (
	"crypto"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"time"
)

func HashPassword(password, salt string) string {
	md5 := crypto.MD5.New()
	md5.Write([]byte(password))
	password = hex.EncodeToString(md5.Sum(nil))
	// 添加 salt
	password = password + salt
	md5.Reset()
	md5.Write([]byte(password))
	return hex.EncodeToString(md5.Sum(nil))
}

func RandomSalt(num int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func CheckJsonParam(c *gin.Context, param interface{}) (err error) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		err = errors.New("参数错误")
		return
	}
	fmt.Println(param)
	err = c.ShouldBindJSON(&param)
	fmt.Println("CheckJsonParam", err)
	return
}
