package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/auth"
	"github.com/mesment/personblog/dao/db"
)

//Get 评论页面
func Comment(c *gin.Context)  {

	//通过从cookie中取出token来进行验证用户是否已登录
	tokenStr, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			//如果没有设置cookie，返回未授权
			log.Printf("没有权限，请先登录,%v",err)
			c.HTML(http.StatusUnauthorized,"views/htmls/500.tmpl",err)
			return
		}
	}

	log.Printf("tokenstr:%s\n",tokenStr)

	// Initialize a new instance of `Claims`
	claims := &auth.CustomClaims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		//return auth.VerifyKey, nil
		return auth.JWTKey, nil
	})
	if !tkn.Valid {
		c.HTML(http.StatusUnauthorized,"views/htmls/500.tmpl",err)
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.HTML(http.StatusUnauthorized,"views/htmls/500.tmpl",err)
			return
		}
		c.HTML(http.StatusBadRequest,"views/htmls/500.tmpl",err)
		return
	}

	//终于通过验证了
	//继续获取文章id
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

	c.HTML(http.StatusOK,	"views/htmls/comment.tmpl",articleDetail)
}
//提交评论
func PostComment(c *gin.Context)  {
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

	err = db.AddComment("mesment",content,articleId)
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
