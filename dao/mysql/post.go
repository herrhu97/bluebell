package mysql

import "web_app/model"

func CreatePost(p *model.Post) (err error) {
	sqlStr := `insert into post (post_id, title, content, author_id,
			community_id)
			value(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	if err != nil {
		return err
	}
	return
}

func GetPostById(id int64) (p *model.Post, err error) {
	p = new(model.Post)
	sqlStr := `select 
			post_id, title, content, author_id, community_id
			from post
			where post_id = ?`
	err = db.Get(p, sqlStr, id)
	return
}

func GetPostList(page, size int64) (p []*model.Post, err error) {
	sqlStr := `select 
			post_id, title, content, author_id, community_id
			from post
			limit ?, ?`
	p = make([]*model.Post, 0, size)
	err = db.Select(&p, sqlStr, (page-1)*size, size)
	return
}
