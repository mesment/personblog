package models


type ArticleInfo struct {
	Id           int     `db:"id"`				//文章id
	CategoryId   int     `db:"category_id"`		//分类id
	Summary      string    `db:"summary"`		//摘要
	Title        string    `db:"title"`			//文章标题
	ViewCount    int    `db:"view_count"`	//阅读次数
	CreateTime   string `db:"create_time"`	//创建时间
	CommentCount int32    `db:"comment_count"`	//评论次数
	Username     string    `db:"username"`		//用户名(作者)
}

type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`				//文章内容
	Category
}

type ArticleRecord struct {
	ArticleInfo
	Category
}