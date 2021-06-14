package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CtxUserIDKey 如果引用controller包的CtxUserIDKey会造成循环引用问题
const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUser 获取当前登录的用户ID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func getPageInfo(c *gin.Context) (int64, int64) {
	var (
		page int64
		size int64
		err  error
	)
	pageStr := c.Query("page")
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}

	sizeStr := c.Query("size")
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}

	return page, size

}
