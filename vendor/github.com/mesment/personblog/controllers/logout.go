package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//处理用户登录和退出登录
func Logout(c *gin.Context) {
	_, login := IsLogin(c)
	if login {
		//清理cookie
		c.SetCookie("token","",0,"/","",false,true)
	}
	m := make(map[string] interface{})
	m["islogin"] = false

	//重定向到首页
	c.Redirect(http.StatusFound,"/")
}


