package db

import (
	"fmt"
	"github.com/mesment/personblog/models"
)

//根据文章id获取文章的评论
func GetCommentByArticleId(articleId int) (commentList []*models.Comment, err error){
	str :=`select
					id,content,username,create_time,status,article_id 
				from 
					comment
				where 
					article_id = ?
				order by create_time desc`
	err = DB.Select(&commentList,str,articleId)
	return
}

//添加评论
func AddComment(username,content string,articleId int) (err error)  {
	if len(content) == 0 || articleId < 0 {
		err = fmt.Errorf("参数非法：username=%s,articleId=%d\n",username,articleId)
	}
	str := `insert into comment(username,content,article_id) values(?,?,?)`

	_, err = DB.Exec(str,username,content,articleId)


	return
}
