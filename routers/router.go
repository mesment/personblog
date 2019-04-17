package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/controllers"
	"github.com/mesment/personblog/middleware"
)

//初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	//测试中间件
	//router.Use(TestAuth())

	router.Static("/views", "./views")
	router.LoadHTMLGlob("views/htmls/*")

	//首页
	router.GET("/", controllers.IndexHandler)

	//用户注册
	router.POST("/user/register", controllers.UserRegister)
	router.GET("/user/register", controllers.ShowRegister)

	//登录
	router.GET("/user/login", controllers.ShowLogin)
	router.POST("/user/login", controllers.LoginHandler)

	//退出登录
	router.GET("/user/logout", controllers.Logout)

	//登录退出管理
	router.GET("/user", controllers.UserHandler)


	//发表文章
	router.GET("/article/new", middleware.TestAuth(),controllers.NewArticle)
	router.POST("/article/new", middleware.TestAuth(),controllers.PostNewArticle)

	//文章详情
	router.GET("/article/detail", controllers.ArticleDetail)

	//文章评论
	router.GET("/article/comment", middleware.TestAuth(),controllers.Comment)
	router.POST("/article/comment", middleware.TestAuth(),controllers.PostComment)

	//留言
	router.GET("/leave/message", middleware.TestAuth(),controllers.LeaveMessage)
	router.POST("/leave/message", middleware.TestAuth(),controllers.SubmitMessage)

	//关于
	router.GET("/about/me", controllers.AboutMe)

	return router
}


