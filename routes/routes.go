package routes

import (
	"net/http"
	"web_app/controller"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	//if mode == gin.ReleaseMode {
	//	gin.SetMode(gin.ReleaseMode)
	//}
	gin.SetMode(mode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//用户注册功能
	r.POST("/signup", controller.SignUpHandler)

	r.GET("/version", func(c *gin.Context) {
		//c.String(http.StatusOK, settings.Conf.Version)
		c.JSON(http.StatusOK, map[string]interface{}{
			"Hello": "Hi",
		})
	})
	return r
}
