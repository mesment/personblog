package controllers

import (
	"fmt"
	"github.com/mesment/personblog/middleware"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/dao/db"
	"net/http"
	"time"
	"github.com/mesment/personblog/models"
)

//用户注册
func UserRegister(c *gin.Context)  {

	var username string = c.PostForm("loginName")
	password := c.PostForm("password")
	//检查用户名是否已存在
	exist := db.UserExist(username)
	if exist {
		errmsg :=fmt.Sprintf("用户名%s已存在，请更换一个用户名\n",username)
		ClientError(c , errmsg )
		return
	}
	//打印用户信息
	log.Printf("username:%s, passwrod:%s\n",username,password)

	//将用户添加到数据库
	err :=db.AddUser(username,password)
	if err != nil {
		ServerError(c,err)
		return
	}

	//设置token
	authjwt:= middleware.JWT{}
	//声明JWT token有效时间单位为1小时(3600秒)
	expirationTime := time.Now().Add(3600 * time.Second).Unix()
	//创建token
	tokenStr,err := CreateToken(&authjwt,username,expirationTime)

	//如果创建token失败，则跳转到登录界面让用户登录
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound ,"/user/login")
		return;
	}

	//创建成功将token写入到cookie，cookie超时时间等于token超时时间
	c.SetCookie("token",tokenStr,int(expirationTime),
		"/","",false,true)

	var user = models.User{
		UserName:username,
	}
	m := make(map[string] interface{})
	m["user"] = user
	m["islogin"] = true
	//跳转回首页
	c.Redirect(http.StatusFound ,"/user")
}




//返回注册页面
func ShowRegister(c * gin.Context)  {
	c.HTML(http.StatusOK,"views/htmls/register.tmpl",nil)
}
