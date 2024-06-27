package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(c *gin.Context, message string) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(c *gin.Context, data interface{}, message string) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(c *gin.Context, message string) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(c *gin.Context, data interface{}, message string) {
	Result(ERROR, data, message, c)
}
