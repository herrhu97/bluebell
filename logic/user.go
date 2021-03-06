package logic

import (
	"web_app/dao/mysql"
	"web_app/model"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *model.ParamSignUp) (err error) {
	// 1. 查看用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2. 创建user
	userID := snowflake.GenID()
	user := model.User{
		Username: p.Username,
		Password: p.Password,
		UserID:   userID,
	}

	// 3. 插入数据库
	if err := mysql.InsertUser(&user); err != nil {
		return err
	}
	return
}

func Login(p *model.ParamLogin) (user *model.User, err error) {
	user = &model.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	token, err := jwt.GenToken(user.Username, user.UserID)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return
}
