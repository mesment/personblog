package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/dao/db"

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




//返回注册页面
func ShowRegister(c * gin.Context)  {
	c.HTML(http.StatusOK,"views/htmls/register.tmpl",nil)
}
