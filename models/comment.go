package models

import "time"

type Comment struct {
	Id         int64     `db:"id" json:"id"`			//评论id
	Content    string    `db:"content" json:"content"`		//评论内容
	UserName   string    `db:"username" json:"user_name"`	//评论者名称
	CreateTime time.Time `db:"create_time" json:"create_time"`	//评论时间
	Status     int       `db:"status" json:"status"`		//评论状态，用语标记删除
	ArticleId  int64     `db:"article_id" json:"article_id"`	//评论所属文章id
}
