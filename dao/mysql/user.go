package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
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

func GetUserById(uid int64) (user *model.User, err error) {
	user = new(model.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
