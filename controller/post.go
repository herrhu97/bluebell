package controller

import (
	"strconv"
	"web_app/logic"
	"web_app/model"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(c *gin.Context) {
	//1.参数校验
	p := new(model.Post)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("c.ShouldBindJSON", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从 c 取到当前发请求的用户的ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	//	2. logic
	if err := logic.CreatePost(p); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	//	3.返回
	ResponseSuccess(c, nil)
}

func GetPostHandler(c *gin.Context) {
	//	1. 参数校验
	idStr := c.Param("id")
	pid, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//	2. logic
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//	3. 返回
	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {
	page, size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
