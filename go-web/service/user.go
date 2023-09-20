package service

import (
	"Web_App/dao/mysql"
	"Web_App/models"
	"Web_App/pkg/snowflake"
)

func SignUp(user *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	err = mysql.QueryUserByUsername(user.Username)
	if err != nil {
		return err
	}
	// 生成UID
	id, err := snowflake.GetID()
	if err != nil {
		return
	}
	// 构造实例
	u := models.User{
		UserID:   id,
		Username: user.Username,
		Password: user.Password,
	}
	// 保存进数据库
	return mysql.InsertUser(&u)
}
