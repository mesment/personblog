package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/dao/redis"
	"github.com/mesment/personblog/dao/db"
	"github.com/mesment/personblog/models"
	"log"
	"net/http"
	"strconv"
)

//Get 评论页面
func Comment(c *gin.Context)  {

	//获取登录状态和用户名
	var login,username = GetDefultUserName(c)

	var user = models.User{
		UserName:username,
	}
	articleIdStr := c.Query("article_id")
	articleId,err:= strconv.Atoi(articleIdStr)
	if err != nil {
		log.Printf("获取文章评论列表失败%v",err)
		ServerError(c,err)
		return
	}

	articleDetail,err := db.GetArticleDetailByArticleId(articleId)
	if err != nil {
		log.Printf("获取文章详情信息失败",err)
		ServerError(c,err)
		return
	}
	//更新详情中的文章ID
	articleDetail.ArticleInfo.Id = articleId

	//设置返回信息
	data:= make(map[string]interface{})
	data["detial"] = articleDetail
	data["islogin"] = login
	data["user"] = user

	c.HTML(http.StatusOK,	"views/htmls/comment.tmpl",data)
}


//提交评论
func PostComment(c *gin.Context)  {
	//获取用户名
	var _, user = GetDefultUserName(c)

	//获取文章ID，评论内容
	articleIdStr := c.PostForm("article_id")
	content := c.PostForm("content")
	articleId,err:= strconv.Atoi(articleIdStr)
	if err != nil {
		log.Printf("获取文章ID失败%v",err)
		ServerError(c, err)
		return
	}

	//对评论内容校验
	if len(content)== 0 || content =="" {
		err := errors.New("评论不能为空！！")
		ClientError(c,err.Error())
		return
	}

	comment := models.Comment{
		UserName:user,
		Content:content,
	}
	err = redis.AddArticleComment(articleIdStr,comment)
	if err != nil {
		log.Println(err)
	}

	//将评论入库
	err = db.AddComment(user,content,articleId)
	if err != nil {
		log.Printf("保存评论失败 %v",err)
		ServerError(c,err)
		return
	}

	//拼接跳转回文章详情页面
	detailURL := "/article/detail?article_id=" + articleIdStr
	c.Redirect(http.StatusFound,detailURL)

}
