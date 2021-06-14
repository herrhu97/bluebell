package controller

import (
	"strconv"
	"web_app/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	//	1.参数校验
	//	2.logic处理
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//	3.返回
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	//	1.参数校验
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt()", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//	2.logic处理
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//	3.返回
	ResponseSuccess(c, data)
}
