package controller

import (
	"Web_App/models"
	"Web_App/service"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUp(c *gin.Context) {
	// 获取参数并校验
	var user models.ParamSignUp
	err := c.ShouldBindJSON(&user)
	if err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
		}

		return
	}
	// 处理业务
	if err := service.SignUp(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
