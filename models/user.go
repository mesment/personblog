package models


type  Article struct {
	Username string
	Avatar string
	Praise string
	Content string
}


type User struct {
	UserID 		string 	`db:"id"`
	UserName 	string	`db:name`
	PassWord	string  `db:password`
}
