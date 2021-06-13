package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"web_app/model"
)

const secret = "bluebell"

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64

	// 数据库查询错误
	err = db.Get(&count, sqlStr, username)
	if err != nil {
		return
	}
	if count > 0 {
		return errors.New("用户已存在")
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
