package logic

import (
	"web_app/dao/mysql"
	"web_app/model"
)

func GetCommunityList() ([]*model.Community, error) {

	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*model.CommunityDetail, error) {

	return mysql.GetCommunityDetail(id)
}
