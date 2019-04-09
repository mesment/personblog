package models

import "time"

type Message struct {
	MessageId	int			`db:"id"`
	Content 	string 		`db:"content"`
	UserName	string		`db:"username"`
	CreateTime  time.Time	`db:"create_time"`
}
