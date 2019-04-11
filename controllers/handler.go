package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/auth"
	"log"
	"net/http"
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

//从token中获取用户名，如果获取失败返回默认Guest用户
func GetDefultUserName(c *gin.Context) (login bool, name string ){
	var user = "Guest"	//默认用户名

	//通过从cookie中取出token来查找用户名
	tokenStr, err := c.Cookie("token")

	// 解析token获取用户名
	username, err := GetUserNameFromToken(tokenStr)
	if err != nil {
		log.Println(err)  	//如果获取用户名失败，使用默认用户名
		login = false 	 	//登录状态未登录
	} else {
		user = username   	//更新用户名
		login = true	  	//登录状态已登录
	}

	return login ,user
}

func GetUserNameFromToken(tokenStr string) (username string ,err error) {
	var find = false  //是否找到
	// 解析token
	authjwt := auth.JWT{}
	claims,err := authjwt.ParseToken(tokenStr)
	if err != nil {
		return
	}
	//遍历claims找出用户名
	for key,value := range claims {
		log.Printf("key:%v,value:%v",key, value)
		if key == "username" {
			username =  value.(string) //token中的用户名
			find = true
			break
		}
	}
	if find {
		return
	}
	err = errors.New("Token里没有用户名字段")
	return
}
