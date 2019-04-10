package models

type Category struct {
	Id 	int `db:"id"`							//分类id
	CategoryName string	`db:"category_name""`	//分类名称
	CategoryNo  int 	`db:"category_no"`		//分类编号

}
