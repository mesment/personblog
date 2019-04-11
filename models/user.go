package models


type User struct {
	UserID 		string 	`db:"id"`		//用户id
	UserName 	string	`db:"name"`		//用户名
	PassWord	string  `db:"password"`	//密码
}
