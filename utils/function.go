package utils

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"time"
	"unicode"
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

func CheckJsonParam(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(len(data))
	return
}

func IsPhoneNumber(phoneNumber string) bool {
	if len(phoneNumber) != 11 {
		return false
	}
	if phoneNumber[0] != '1' {
		return false
	}
	for _, r := range phoneNumber[1:] {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
