package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/controllers"
	"github.com/mesment/personblog/dao/db"
)

func main() {

	dns := "root:personblog@110@(localhost:3306)/myblog?charset=utf8mb4&parseTime=true"
	//dns := "root:root@tcp(localhost:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()
	router := gin.Default()

	router.Static("/views", "./views")
	router.LoadHTMLGlob("views/htmls/*")

	router.GET("/", controllers.IndexHandler)

	//登录
	router.GET("/user/login", controllers.ShowLogin)
	router.POST("/user/login", controllers.LoginHandler)

	//注册
	router.POST("/user/register", controllers.UserRegister)
	router.GET("/user/register", controllers.ShowRegister)

	//发表文章
	router.GET("/article/new", controllers.NewArticle)
	router.POST("/article/new", controllers.PostNewArticle)

	//文章详情
	router.GET("/article/detail", controllers.ArticleDetail)

	//文章评论
	router.GET("/article/comment", controllers.Comment)
	router.POST("/article/comment", controllers.PostComment)


	router.GET("/leave/message", controllers.LeaveMessage)
	router.POST("/leave/message", controllers.SubmitMessage)

	router.GET("/about/me", controllers.AboutMe)

	router.Run() // listen and serve on 0.0.0.0:8080
}
