package db

import (
	"fmt"
	"github.com/mesment/personblog/models"
)

//新增留言
func AddMessage(username string, msg string)  (err error){
	if len(username) <= 0 || len(msg) <= 0 {
		err = fmt.Errorf("名或评论不能为空")
	}
	str := `insert into message(username,content) values(?,?)`
	_,err = DB.Exec(str,username,msg)
	return
}

//返回前 limitNum 条留言
func GetMessage(limitNum int) (messageList []*models.Message, err error) {
	queryStr := `select id,username,content,create_time from message order by create_time desc,id desc limit ?`
	err = DB.Select(&messageList,queryStr,limitNum)

	return
}
