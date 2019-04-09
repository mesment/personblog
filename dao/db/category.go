package db

import (
	"github.com/mesment/personblog/models"
)

func GetAllCategory() (categoryList []*models.Category, err error ){

	queryStr := `select id, category_no, category_name from category order by id asc`

	err = DB.Select(&categoryList,queryStr)
	if err != nil {
		return nil,err
	}
	return categoryList,err

}

func GetCategory(categoryID int) (category *models.Category, err error) {
	queryStr := `select id,category_no, category_name from category where id=?`

	err = DB.Get(category,queryStr,categoryID)
	return
}

