package models

import "time"

type Comment struct {
	Id         int64     `db:"id"`			//评论id
	Content    string    `db:"content"`		//评论内容
	UserName   string    `db:"username"`	//评论者名称
	CreateTime time.Time `db:"create_time"`	//评论时间
	Status     int       `db:"status"`		//评论状态，用语标记删除
	ArticleId  int64     `db:"article_id"`	//评论所属文章id
}
