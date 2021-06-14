package logic

import (
	"web_app/dao/mysql"
	"web_app/model"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *model.Post) (err error) {
	//	1.生成post_id
	p.ID = snowflake.GenID()
	//	2.写入数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return
	}
	//	3.返回
	return
}

func GetPostById(id int64) (p *model.ApiPostDetail, err error) {
	// 组合post属性
	post, err := mysql.GetPostById(id)
	if err != nil {
		zap.L().Error("mysql.GetPostById(id)", zap.Error(err))
		return
	}

	// 组合authorName属性
	uid := post.AuthorID
	user, err := mysql.GetUserById(uid)
	if err != nil {
		zap.L().Error("mysql.GetUserById(uid)", zap.Error(err))
		return
	}
	authorName := user.Username

	// 组合community属性
	communityId := post.CommunityID
	community, err := mysql.GetCommunityDetail(communityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetail(communityId) failed", zap.Error(err))
		return
	}

	p = &model.ApiPostDetail{
		AuthorName:      authorName,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

func GetPostList(page, size int64) (data []*model.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList(page, size) failed", zap.Error(err))
		return nil, err
	}
	data = make([]*model.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		// 组合authorName属性
		uid := post.AuthorID
		user, err := mysql.GetUserById(uid)
		if err != nil {
			zap.L().Error("mysql.GetUserById(uid)", zap.Error(err))
			continue
		}
		authorName := user.Username

		// 组合community属性
		communityId := post.CommunityID
		community, err := mysql.GetCommunityDetail(communityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetail(communityId) failed", zap.Error(err))
			continue
		}

		p := &model.ApiPostDetail{
			AuthorName:      authorName,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, p)
	}
	return
}
