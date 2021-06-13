package controller

import (
	"net/http"
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
			c.JSON(http.StatusOK, gin.H{
				"msg": "请求参数格式不正确",
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	//if len(p.RePassword) == 0 || len(p.Password) == 0 ||
	//	len(p.Username) == 0 || p.Password != p.RePassword {
	//	zap.L().Error("param sign up error")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数非法",
	//	})
	//	return
	//}

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic sign up failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}

	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
