package mysql

import (
	"Web_App/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

// secret 用于当作密码的加密
const secret string = "go is best language"

// InsertUser 向数据库中插入一条新纪录
func InsertUser(user *models.User) (err error) {
	// 密码加密
	password := encryptPassword(user.Password)
	// 执行sql语句
	sqlStr := `insert into user(user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, password)
	if err != nil {
		return err
	}
	return
}

// QueryUserByUsername 查找用户名是否已出现
func QueryUserByUsername(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	err = db.Get(&count, sqlStr, username)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
