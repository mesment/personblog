package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/dao/db"
	"log"
	"net/http"
	"strconv"
	"github.com/mesment/personblog/models"
)

//Get 评论页面
func Comment(c *gin.Context)  {
	//继续获取文章id
	//获取登录状态用户名,用户名
	var login,username = GetDefultUserName(c)

	var user = models.User{
		UserName:username,
	}
	articleIdStr := c.Query("article_id")
	articleId,err:= strconv.Atoi(articleIdStr)
	if err != nil {
		fmt.Printf("获取文章评论失败%v",err)
		return
	}

	articleDetail,err := db.GetArticleDetailByArticleId(articleId)
	if err != nil {
		fmt.Printf("获取文章信息失败",err)
		ServerError(c,err)
		return
	}
	//更新详情中的文章ID
	articleDetail.ArticleInfo.Id = articleId
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

	//获取文章ID
	articleIdStr := c.PostForm("article_id")
	content := c.PostForm("content")
	articleId,err:= strconv.Atoi(articleIdStr)

	//对评论内容校验
	if len(content)== 0 || content =="" {
		err:= fmt.Errorf("评论不能为空！！")
		ClientError(c,err.Error())
		return
	}

	//将评论添加到数据库
	err = db.AddComment(user,content,articleId)
	if err != nil {
		fmt.Printf("PostComment：提交评论失败 %v",err)
		ServerError(c,err)
		return
	}
	//更新文章评论次数
	err = db.UpdateCommentCount(articleId)
	if err != nil {
		log.Printf("更新文章评论次数失败")
	}
	err = db.UpdateViewCount(articleId)
	if err != nil {
		log.Printf("更新文章阅读次数失败")
	}

	//拼接跳转回文章详情页面
	detailURL := "/article/detail?article_id=" + articleIdStr
	c.Redirect(http.StatusFound,detailURL)

}
