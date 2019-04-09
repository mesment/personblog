package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/dao/db"
	"github.com/mesment/personblog/models"
	"github.com/mesment/personblog/utils"
	"log"
	"net/http"
	"strconv"
)


var pageLimit = 5

func ServerError(c *gin.Context,code int,tplname string, err error)  {
	fmt.Printf("ERROR:%v",err)
	c.HTML(code,tplname,err)
	return
}




//第一页:隐含 http://localhost:8080/
//第一页: http://localhost:8080/?page=1
//第10页:http://localhost:8080/?page=10
func IndexHandler(c *gin.Context) {
	var page = models.Page{}  //分页对象
	var currentPage int		//当前页码

	//获取当前页码
	currentPage,err := utils.GetPage(c)
	if err != nil {
		ServerError(c,http.StatusInternalServerError,"views/htmls/404.tmpl",err)
	}
	//总记录数
	totalRecords,err := db.CountArticles()
	if err != nil {
		ServerError(c,http.StatusInternalServerError,"views/htmls/404.tmpl",err)
	}
	fmt.Printf("文章总条数：%d\n",totalRecords)

	//设置分页对象
	page.SetPage(totalRecords,currentPage, pageLimit)
	//打印page信息
	fmt.Printf("page info：%v\n",page)

	//查询返回文章列表
	articleInfoList,err :=db.GetArticleInfosByPageAndLimit(currentPage,pageLimit)
	if err != nil {
		ServerError(c,http.StatusInternalServerError,"views/htmls/404.tmpl",err)
	}
	var data map[string]interface{} = make(map[string]interface{}, 10)
	data["article_list"] = articleInfoList
	data["pageinfo"] = page

	c.HTML(http.StatusOK,"views/htmls/index.tmpl",data)
}

//用户注册
func UserRegister(c *gin.Context)  {
	var username string = c.PostForm("loginName")
	password := c.PostForm("password")
	//检查用户名是否已存在
	exist := db.UserExist(username)

	if exist {
		errmsg :=fmt.Sprintf("用户名%s已存在，请更换一个用户名\n",username)
		fmt.Printf(errmsg)
		c.HTML(http.StatusBadRequest,"views/htmls/404.tmpl",errmsg)
		//c.AbortWithError(http.StatusBadRequest,errors.New(errmsg))
		return
	}

	fmt.Printf("username:%s, passwrod:%s\n",username,password)
	err :=db.AddUser(username,password)
	if err != nil {
		fmt.Printf("ERROR:%v",err)
		c.HTML(http.StatusInternalServerError,"views/htmls/404.tmpl",err.Error())
		return
	}
	if err != nil {
		c.HTML(http.StatusInternalServerError,"views/htmls/404.tmpl",err.Error())
		return
	}
	//跳转回首页
	c.Redirect(http.StatusFound ,"/")
}


//返回用户登录页面
func ShowLogin(c *gin.Context)  {
	//跳转到登录页面
	c.HTML(http.StatusOK,"views/htmls/login.tmpl",nil)
}

//用户登录
func LoginHandler(c *gin.Context) {
	var errmsg string
	var username string = c.PostForm("loginName")
	password := c.PostForm("password")
	if username == "" ||password =="" {
		errmsg = "用户名或密码不能为空"
		c.HTML(http.StatusInternalServerError,"views/htmls/404.tmpl",errmsg)
	}

	_,err := db.GetUser(username,password)
	if err != nil {
		err := errors.New("登录失败:" + err.Error())
		c.HTML(http.StatusInternalServerError,"views/htmls/404.tmpl",err)
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
		fmt.Printf("ArticleDetail：获取文章详情失败 %v",err)
		c.HTML(http.StatusInternalServerError,	"views/htmls/404.tmpl",err)
		return
	}

	articleDetail,err := db.GetArticleDetailByArticleId(articleId)
	articleDetail.ArticleInfo.Id = articleId
	if err != nil {
		fmt.Printf("ArticleDetail：获取文章详情失败 %v",err)
		c.HTML(http.StatusInternalServerError,	"views/htmls/404.tmpl",err)
		return
	}

	commentList,err := db.GetCommentByArticleId(articleId)
	for _,v :=range commentList {
		fmt.Println("COMMENT LIST:%v",v)
	}
	if err != nil {
		fmt.Printf("ArticleDetail：获取文章评论失败 %v",err)
	}
	err = db.UpdateViewCount(articleId)
	if err != nil {
		fmt.Printf("更新阅读次数失败")
	}
	var m map[string]interface{} = make(map[string]interface{},10)
	m["comments"] = commentList
	m["article_detail"] = articleDetail
	//fmt.Printf("map:%v\n",m)
	c.HTML(http.StatusOK,"views/htmls/details.tmpl", m)

}

//返回注册页面
func ShowRegister(c * gin.Context)  {
	//跳转回首页
	c.HTML(http.StatusOK,"views/htmls/register.tmpl",nil)
}


//评论
func Comment(c *gin.Context)  {
	//获取文章id参数
	articleIdStr := c.Query("article_id")
	articleId,err:= strconv.Atoi(articleIdStr)
	if err != nil {
		fmt.Printf("Comment：获取文章评论失败%v",err)
		return
	}

	articleDetail,err := db.GetArticleDetailByArticleId(articleId)
	if err != nil {
		fmt.Printf("Comment：获取文章信息失败",err)
		c.HTML(http.StatusInternalServerError,	"views/htmls/404.tmpl",err)
		return
	}
	articleDetail.ArticleInfo.Id = articleId
	fmt.Printf("ArticleDetial:categoryInfo:%v\n",articleDetail.Category)
	fmt.Printf("ArticleDetial:ArticleInfo:Id:%d\n",articleDetail.ArticleInfo.Id)
	c.HTML(http.StatusOK,	"views/htmls/comment.tmpl",articleDetail)
}
//提交评论
func PostComment(c *gin.Context)  {
	//获取文章id
	articleIdStr := c.PostForm("article_id")
	content := c.PostForm("content")
	articleId,err:= strconv.Atoi(articleIdStr)

	if len(content)== 0 || content =="" {
		err:= fmt.Errorf("评论不能为空！！")
		c.HTML(http.StatusBadRequest,	"views/htmls/404.tmpl",err)
		return
	}

	err = db.AddComment("mesment",content,articleId)
	if err != nil {
		fmt.Printf("PostComment：提交评论失败 %v",err)
		c.HTML(http.StatusInternalServerError,	"views/htmls/500.tmpl",err)
		return
	}
	//更新文章评论次数
	err = db.UpdateCommentCount(articleId)
	if err != nil {
		log.Printf("更新文章评论次数失败")
	}

	detailURL := "/article/detail?article_id=" + articleIdStr
	c.Redirect(http.StatusFound,detailURL)
}

//获取发表文章页面
func NewArticle(c *gin.Context)  {
	categoryList,err := db.GetAllCategory()
	if err != nil {
		fmt.Printf("NewArticle Failed %v",err)
		c.HTML(http.StatusInternalServerError,	"views/htmls/404.tmpl",err)
		return
	}
	fmt.Printf("category:%v",categoryList)
	c.HTML(http.StatusOK,	"views/htmls/newarticle.tmpl",categoryList)
}

//发布文章
func PostNewArticle(c *gin.Context)  {
	categoryId := c.PostForm("category_id")
	title := c.PostForm("title")
	content := c.PostForm("content")

	cateid,err := strconv.Atoi(categoryId)

	if err != nil {
		fmt.Printf("新增文章失败,解析文章Id失败.%v",err)
		c.String(http.StatusInternalServerError,err.Error())
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




//Get 留言
func LeaveMessage(c *gin.Context)  {

	msgs, err := db.GetMessage()
	if err != nil {
		fmt.Println("LeaveMessage: 获取留言列表失败%v",err)
	}
	c.HTML(http.StatusOK,"views/htmls/message.tmpl",msgs)

}

//Post 提交留言
func SubmitMessage(c *gin.Context) {
	username := "mesment"

	msgContent := c.PostForm("content")
	if len(msgContent)== 0 || msgContent =="" {
		err := fmt.Errorf("留言不能为空！！")
		c.HTML(http.StatusInternalServerError,	"views/htmls/404.tmpl",err)
		return
	}
	log.Printf("SubmitMessage: msg:%s\n", msgContent)
	err := db.AddMessage(username, msgContent)
	if err != nil {
			fmt.Printf("添加留言失败,%v", err)
			c.HTML(http.StatusInternalServerError, "views/htmls/404.tmpl", err)
			return
	}
	c.Redirect(http.StatusFound,"/leave/message")
}

//关于
func AboutMe(c *gin.Context) {

	c.HTML(http.StatusOK,"views/htmls/about.tmpl",nil)
}