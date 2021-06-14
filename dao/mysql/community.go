package mysql

import (
	"database/sql"
	"web_app/model"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*model.Community, err error) {
	sqlStr := `select community_id, community_name from community`
	err = db.Select(&communityList, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Info("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetail(id int64) (communityDetail *model.CommunityDetail, err error) {
	communityDetail = new(model.CommunityDetail)
	sqlStr := `select 
       community_id, community_name, introduction, create_time 
		from community
		where community_id = ?`
	err = db.Get(communityDetail, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return communityDetail, err
}
