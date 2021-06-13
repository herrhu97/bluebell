package controller

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/model"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(model.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			zap.L().Error("gin bind json", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic sign up failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(model.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error("gin bind json", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2. 业务处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic login failed", zap.Error(err))

		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, token)
}
