package Controller

import (
	"Goweb_cli/Logic"
	"Goweb_cli/dao/mysql"
	"Goweb_cli/modules"
	"errors"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//获取参数
	var p modules.SignUpParam
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("c.ShouldBindJSON failed", zap.Error(err))
		_, ok := err.(validator.FieldError)
		if !ok {
			modules.ResponseError(c, modules.CodeInValidParam)
		}
		modules.ResponseErrorWithMsg(c, modules.CodeInValidParam, err.Error())
		return
	}

	if err := Logic.SignUp(&p); err != nil {
		zap.L().Error("Logic.SignUp() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			modules.ResponseError(c, modules.CodeUserExist)
			return
		}
		modules.ResponseError(c, modules.CodeServerBusy)
		return
	}
	modules.ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	var p modules.LoginParam
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("c.ShouldBindJSON failed", zap.Error(err))
		_, ok := err.(validator.FieldError)
		if !ok {
			modules.ResponseError(c, modules.CodeInValidParam)
		}
		modules.ResponseErrorWithMsg(c, modules.CodeInValidParam, err.Error())
		return
	}
	token, err := Logic.Login(&p)
	if err != nil {
		zap.L().Error("Logic.Login failed", zap.Error(err))
		modules.ResponseError(c, modules.CodeInValidParam)
		return
	}

	modules.ResponseSuccess(c, token)
}
