package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
