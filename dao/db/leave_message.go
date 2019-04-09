package db

import (
	"fmt"
	"github.com/mesment/personblog/models"
)

func AddMessage(username string, msg string)  (err error){
	if len(username) <= 0 || len(msg) <= 0 {
		err = fmt.Errorf("名或评论不能为空")
	}
	str := `insert into message(username,content) values(?,?)`
	_,err = DB.Exec(str,username,msg)
	return
}

func GetMessage() (messageList []*models.Message, err error) {
	queryStr := `select id,username,content,create_time from message order by create_time desc,id desc limit 20`
	err = DB.Select(&messageList,queryStr)

	return
}
