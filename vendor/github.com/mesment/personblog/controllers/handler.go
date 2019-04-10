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


//每页显示条数
var pageLimit = 5

//服务器出错
func ServerError(c *gin.Context, err error)  {
	fmt.Printf("Error:%v",err)
	c.HTML(http.StatusInternalServerError,"views/htmls/500.tmpl",err)
	return
}
//Client出错
func ClientError(c *gin.Context, errmsg string)  {
	fmt.Printf("%s",errmsg)
	c.HTML(http.StatusBadRequest,"views/htmls/404.tmpl",errmsg)
	return
}

//IndexHandler:显示文章列表首页
//url第一页:隐含 http://localhost:8080/ ,第二页: http://localhost:8080/?page=2
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

//用户注册
func UserRegister(c *gin.Context)  {
	var username string = c.PostForm("loginName")
	password := c.PostForm("password")
	//检查用户名是否已存在
	exist := db.UserExist(username)
	if exist {
		errmsg :=fmt.Sprintf("用户名%s已存在，请更换一个用户名\n",username)
		ClientError(c , errmsg )
		//c.AbortWithError(http.StatusBadRequest,errors.New(errmsg))
		return
	}
	//打印用户信息
	fmt.Printf("username:%s, passwrod:%s\n",username,password)
	err :=db.AddUser(username,password)
	if err != nil {
		ServerError(c,err)
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
		ClientError(c,errmsg)
	}

	_,err := db.GetUser(username,password)
	if err != nil {
		err := errors.New("登录失败:" + err.Error())
		ServerError(c,err)
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

//返回注册页面
func ShowRegister(c * gin.Context)  {
	c.HTML(http.StatusOK,"views/htmls/register.tmpl",nil)
}


//Get 评论页面
func Comment(c *gin.Context)  {
	//获取文章id
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




//Get 获取留言页面
func LeaveMessage(c *gin.Context)  {

	msgs, err := db.GetMessage(10) 	//返回前50条留言
	if err != nil {
		fmt.Println("LeaveMessage: 获取留言列表失败%v",err)
		ServerError(c,err)
	}
	c.HTML(http.StatusOK,"views/htmls/message.tmpl",msgs)

}

//Post 提交留言
func SubmitMessage(c *gin.Context) {
	username := "mesment"

	msgContent := c.PostForm("content")
	if len(msgContent)== 0 || msgContent =="" {
		err := fmt.Errorf("留言不能为空！！")
		ClientError(c,err.Error())
		return
	}
	err := db.AddMessage(username, msgContent)
	if err != nil {
			fmt.Printf("提交留言失败,%v", err)
			ServerError(c,err)
			return
	}

	//刷新页面
	c.Redirect(http.StatusFound,"/leave/message")
}

//关于
func AboutMe(c *gin.Context) {

	c.HTML(http.StatusOK,"views/htmls/about.tmpl",nil)
}