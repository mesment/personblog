package models

type Category struct {
	Id 	int `db:"id"`
	CategoryName string	`db:"category_name""`
	CategoryNo  int 	`db:"category_no"`

}
