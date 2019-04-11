package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mesment/personblog/models"
)

//关于
func AboutMe(c *gin.Context) {
	var login,username = GetDefultUserName(c)

	var user = models.User{
		UserName:username,
	}
	data:= make(map[string]interface{})
	data["islogin"] = login
	data["user"] = user

	c.HTML(http.StatusOK,"views/htmls/about.tmpl",data)
}
