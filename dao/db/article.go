package db

import (
	"fmt"
	"github.com/mesment/personblog/models"
	"math"
)

func GetArticleDetailByArticleId( articleId int) (articleDetail *models.ArticleDetail, err error ) {
	if articleId < 0 {
		err = fmt.Errorf("文章ID非法，id=%d\n",articleId)
		return
	}
	querStr := `select 
					  	  id,summary,category_id,title,view_count,create_time,
					  	  comment_count,username,content
					from 
						  article 
					where 
						  status=1
					and 
						  id =?`

	articleDetail = &models.ArticleDetail{}

	err = DB.Get(articleDetail, querStr, articleId)
	//fmt.Printf("ArticleDetial:%v\n",articleDetail)
	return
}

//GetArticleInfosByPageAndLimit:通过页码和每页条数查询

func GetArticleInfosByPageAndLimit(page int, limit int) (list []*models.ArticleInfo, err error) {
	if page <= 0 || limit <= 0 {
		err = fmt.Errorf("页码或每页条数非法,页码:%d,每页条数:%d",page,limit)
	}
	offset := (page - 1) * limit
	list, err = GetArticleInfoList(offset,limit)
	return
}

func GetArticleInfoList(offset int, limit int ) (list []*models.ArticleInfo, err error){
	if offset <0 || limit < 0 {
		fmt.Printf("ArticleRecordList: argument can't be negativd")
		return
	}

	queryStr  := `select id,category_id,summary,title,view_count,
	create_time,comment_count,username from article where status=1 order by create_time desc limit ?,?`

	err = DB.Select(&list,queryStr,offset,limit)

	return
}


func AddArticleDetail(title string,categoryid int, content string) error {
	articleDetail := models.ArticleDetail{}
	articleDetail.Title = title
	articleDetail.Content = content
	articleDetail.ArticleInfo.CategoryId = categoryid

	contentutf8 := []rune(content)
	minLength := int(math.Min(float64(len(contentutf8)),128.0))
	articleDetail.Summary = string(contentutf8[:minLength])
	id, err := InsertArticleDetail(&articleDetail)
	if err != nil {
		return err
	}
	fmt.Printf("insert article success,article id:%d",id)
	return nil
}

func InsertArticleDetail(detail *models.ArticleDetail) (id int, err error) {
	stmtstr := `insert into article(category_id,content,title,summary) values(?,?,?,?)`
	if err != nil {
		fmt.Printf("errinfo:%v",err)
		return -1, err
	}

	_, err = DB.Exec(stmtstr, detail.ArticleInfo.CategoryId, detail.Content, detail.Title,detail.Summary)
	if err != nil {
		return -1 ,err
	}
	return id,nil
}

func CountArticles()  (int, error) {
	var count int
	query :=`select count(*) from article`
	err := DB.Get(&count,query)
	if err != nil  {
		return 0,err
	}
	return count, nil
}


func GetArticleCount() (total int, err error){
	str := `select count(*) from article`

	err = DB.Get(&total,str)
	return
}

func UpdateViewCount(articleId int) (err error) {

	sqlstr := ` update 
						article 
					set 
						view_count = view_count + 1
					where
						id = ?`

	_, err = DB.Exec(sqlstr, articleId)
	if err != nil {
		return
	}

	return
}

func UpdateCommentCount(articleId int) (err error) {

	sqlstr := ` update 
						article 
					set 
						comment_count = comment_count + 1
					where
						id = ?`

	_, err = DB.Exec(sqlstr, articleId)
	if err != nil {
		return
	}

	return
}



