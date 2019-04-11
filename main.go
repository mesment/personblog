package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/auth"
	"github.com/mesment/personblog/controllers"
	"github.com/mesment/personblog/dao/db"
	"log"
	"net/http"
)

//token 鉴权 middleware
func TestAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//通过从cookie中取出token来进行验证用户是否已登录
		tokenStr, err := c.Cookie("token")
		//t := c.Request.Cookie("token")
		//str:=t.Value

		if err != nil {
			if err == http.ErrNoCookie {
				//如果没有设置cookie，返回未授权
				log.Printf("没有权限，请先登录,%v",err)
				c.HTML(http.StatusUnauthorized,"views/htmls/404.tmpl",err)
				return
			}
		}

		log.Printf("token 字符串:%s\n",tokenStr)

		// Initialize a new instance of `Claims`
		claims := &auth.CustomClaims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return auth.JWTKey , nil
		})
		newErr := errors.New("请先登录")
		if !tkn.Valid {

			c.HTML(http.StatusUnauthorized,"views/htmls/404.tmpl",newErr)
			return
		}
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.HTML(http.StatusUnauthorized,"views/htmls/404.tmpl",newErr)
				return
			}
			c.HTML(http.StatusBadRequest,"views/htmls/404.tmpl",newErr)
			return
		}
		c.Next()
	}
}

func main() {

	//设置数据库连接信息(mysql:3306),其中mysql是docker里mysql的容器名
	//在本地主机上运行mysql时改为localhost
	dns := "root:personblog@(localhost:3306)/myblog?charset=utf8mb4&parseTime=true"

	err := db.Init("mysql", dns)
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	router := gin.Default()

	//测试中间件
	//router.Use(TestAuth())

	router.Static("/views", "./views")
	router.LoadHTMLGlob("views/htmls/*")

	//首页
	router.GET("/", controllers.IndexHandler)

	//登录
	router.GET("/user/login", controllers.ShowLogin)
	router.POST("/user/login", controllers.LoginHandler)

	//用户注册
	router.POST("/user/register", controllers.UserRegister)
	router.GET("/user/register", controllers.ShowRegister)

	//发表文章
	router.GET("/article/new", TestAuth(),controllers.NewArticle)
	router.POST("/article/new", TestAuth(),controllers.PostNewArticle)

	//文章详情
	router.GET("/article/detail", controllers.ArticleDetail)

	//文章评论
	router.GET("/article/comment", TestAuth(),controllers.Comment)
	router.POST("/article/comment", TestAuth(),controllers.PostComment)

	//留言
	router.GET("/leave/message", TestAuth(),controllers.LeaveMessage)
	router.POST("/leave/message", TestAuth(),controllers.SubmitMessage)

	//关于
	router.GET("/about/me", controllers.AboutMe)

	router.Run() // listen and serve on 0.0.0.0:8080
}
