package controller

import (
	"web_app/logic"
	"web_app/model"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostVoteController(c *gin.Context) {
	//	参数校验
	p := new(model.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	userId, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error(" getCurrentUserID(c)", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}

	if err := logic.VoteForPost(userId, p); err != nil {
		zap.L().Error(" logic.VoteForPost", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
