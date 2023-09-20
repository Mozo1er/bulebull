package routes

import (
	"Web_App/controller"
	"Web_App/logger"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 其他路由在这里注册
	r.POST("/signup", controller.SignUp)
	return r
}
