package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/dao/db"
	"github.com/mesment/personblog/models"
	"github.com/mesment/personblog/utils"
	"net/http"
)

//IndexHandler:显示文章列表首页
//url隐含第一页: http://localhost:8080/
// 第二页: http://localhost:8080/?page=2
func IndexHandler(c *gin.Context) {

	var page = models.Page{}  	//分页对象，用来存放分页信息
	var currentPage int			//当前页码

	//根据url获取当前页码
	currentPage,err := utils.GetPageNo(c)
	if err != nil {
		ServerError(c,err)
	}
	//文章总记录数
	totalRecords,err := db.CountArticles()
	if err != nil {
		ServerError(c,err)
	}
	fmt.Printf("文章总条数：%d\n",totalRecords)

	//设置分页对象
	page.SetPage(totalRecords,currentPage, pageLimit)
	//打印page信息
	fmt.Printf("page info：%v\n",page)

	//查询返回文章列表
	articleInfoList,err :=db.GetArticleInfosByPageAndLimit(currentPage,pageLimit)
	if err != nil {
		ServerError(c,err)
	}
	var data map[string]interface{} = make(map[string]interface{}, 10)
	data["article_list"] = articleInfoList
	data["pageinfo"] = page

	c.HTML(http.StatusOK,"views/htmls/index.tmpl",data)
}
