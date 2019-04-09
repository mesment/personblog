package models

import "time"

type Message struct {
	MessageId	int			`db:"id"`			//留言id
	Content 	string 		`db:"content"`		//留言内容
	UserName	string		`db:"username"`		//留言者用户名
	CreateTime  time.Time	`db:"create_time"`	//留言时间
}
