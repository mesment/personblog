package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//关于
func AboutMe(c *gin.Context) {

	c.HTML(http.StatusOK,"views/htmls/about.tmpl",nil)
}
