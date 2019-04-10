package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/mesment/personblog/dao/db"
)



//获取发表文章页面
func NewArticle(c *gin.Context)  {
	categoryList,err := db.GetAllCategory()
	if err != nil {
		fmt.Printf("获取发表文章页面失败: %v",err)
		ServerError(c,err)
		return
	}
	fmt.Printf("category:%v",categoryList)
	c.HTML(http.StatusOK,	"views/htmls/newarticle.tmpl",categoryList)
}

//发布文章
func PostNewArticle(c *gin.Context)  {
	categoryId := c.PostForm("category_id")		//文章分类
	title := c.PostForm("title")    			//文章标题
	content := c.PostForm("content") 			//文章内容

	cateid,err := strconv.Atoi(categoryId)

	if err != nil {
		fmt.Printf("解析分类Id失败.%v",err)
		ServerError(c,err)
		return
	}
	err  = db.AddArticleDetail(title,cateid,content)
	if err != nil {
		fmt.Printf("新增文章失败.%v",err)
		c.String(http.StatusInternalServerError,err.Error())
		return
	}

	//跳转回首页
	c.Redirect(http.StatusFound,"/")
}

//文章详情
//url :http://localhost:8080/article/detail?article_id=1
func ArticleDetail(c *gin.Context)  {
	//获取文章ID
	articleIdStr := c.Query("article_id")
	articleId,err := strconv.Atoi(articleIdStr)
	if err != nil {
		fmt.Printf("解析文章ID失败 %v",err)
		ServerError(c, err)
		return
	}

	articleDetail,err := db.GetArticleDetailByArticleId(articleId)
	if err != nil {
		fmt.Printf("获取文章详情失败 %v",err)
		ServerError(c,err)
		return
	}

	articleDetail.ArticleInfo.Id = articleId
	commentList,err := db.GetCommentByArticleId(articleId)
	if err != nil {
		fmt.Printf("获取文章评论失败 %v",err)
	}

	//更新文章阅读次数
	err = db.UpdateViewCount(articleId)
	if err != nil {
		fmt.Printf("更新阅读次数失败")
		//忽略更新错误，继续往下执行
	}
	var m map[string]interface{} = make(map[string]interface{},10)
	m["comments"] = commentList  			//返回评论列表
	m["article_detail"] = articleDetail		//返回文章详情

	c.HTML(http.StatusOK,"views/htmls/details.tmpl", m)

}
