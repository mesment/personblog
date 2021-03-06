package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/dao/db"
	"github.com/mesment/personblog/models"
	"github.com/mesment/personblog/utils"
	"log"
	"net/http"
	"github.com/mesment/personblog/pkg/logger"
)

//IndexHandler:显示文章列表首页
//url隐含第一页: http://localhost:8080/
// 第二页: http://localhost:8080/?page=2
func IndexHandler(c *gin.Context) {

	//获取登录状态用户名,用户名
	var login,username = GetDefultUserName(c)

	var user = models.User{
		UserName:username,
	}
	logger.Debug("username:%s,来访了。",username)
	var page = models.NewPage()  //分页对象，用来存放分页信息
	var curPageNo int			//当前页码

	//根据url获取当前页码
	curPageNo,err := utils.GetPageNo(c)
	if err != nil {
		ServerError(c,err)
	}
	//文章总记录数
	totalRecords,err := db.CountArticles()
	if err != nil {
		ServerError(c,err)
	}
	log.Printf("文章总条数：%d\n",totalRecords)

	//设置分页对象
	page.SetPage(totalRecords,curPageNo, page.Limit)
	//打印page信息
	log.Printf("page info：%v\n",page)

	//查询返回文章列表
	articleInfoList,err :=db.GetArticleInfosByPageAndLimit(curPageNo,page.Limit)
	if err != nil {
		ServerError(c,err)
	}
	var data map[string]interface{} = make(map[string]interface{}, 10)
	data["article_list"] = articleInfoList
	data["pageinfo"] = page
	data["islogin"] = login
	data["user"] = user

	c.HTML(http.StatusOK,"views/htmls/index.tmpl",data)
}
