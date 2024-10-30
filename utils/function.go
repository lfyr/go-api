package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"math/rand"
	"unicode"
)

func HashPassword(password string) string {
	// GenerateFromPassword函数用于生成哈希密码
	// 第二个参数是cost，取值范围在4 - 31之间，它决定了哈希计算的成本
	// 成本越高，哈希计算越慢，但安全性越高
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		logrus.Error("生成密码失败", err.Error())
		return ""
	}
	return string(hashedPassword)
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func RandomSalt(num int) string {
	buffer := make([]byte, num)
	_, err := rand.Read(buffer)
	if err != nil {
		panic("读取随机字节失败")
	}
	// 可以根据需要将字节切片转换为其他形式的字符串，如十六进制等
	return fmt.Sprintf("%x", buffer)
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
