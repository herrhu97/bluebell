package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/model"
)

const secret = "bluebell"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64

	// 数据库查询错误
	err = db.Get(&count, sqlStr, username)
	if err != nil {
		return
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *model.User) (err error) {
	sqlStr := `insert into user(username, user_id, password) values (?, ?, ?)`
	user.Password = encryptPassword(user.Password)
	if _, err := db.Exec(sqlStr, user.Username, user.UserID, user.Password); err != nil {
		return err
	}
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *model.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	if encryptPassword(oPassword) != user.Password {
		return ErrorInvalidPassword
	}
	return
}
