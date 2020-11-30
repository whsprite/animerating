package modules

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeUserExist
	CodeUserNotExist
	CodeInValidParam
	CodeServerBusy
)

var CodeMsg = map[ResCode]string{
	CodeSuccess:      "Success",
	CodeUserExist:    "用户已存在",
	CodeUserNotExist: "用户不存在",
	CodeInValidParam: "参数错误",
	CodeServerBusy:   "服务器繁忙",
}

func (r ResCode) GetMsg() string {
	msg, ok := CodeMsg[r]
	if !ok {
		return CodeMsg[CodeServerBusy]
	}
	return msg
}

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	msg := code.GetMsg()
	c.JSON(http.StatusOK, &ResponseData{
		code,
		msg,
		nil,
	})
}
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg string) {
	c.JSON(http.StatusOK, &ResponseData{
		code,
		msg,
		nil,
	})
}
func ResponseSuccess(c *gin.Context, data interface{}) {
	msg := CodeSuccess.GetMsg()
	c.JSON(http.StatusOK, &ResponseData{
		CodeSuccess,
		msg,
		data,
	})
}
