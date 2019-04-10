package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)


//例如url: http://localhost:8080/?page=2
//GetPageNo: 根据URL获取当前的分页页码，返回页码2
func GetPageNo(c *gin.Context) (currentPageNo int, err error) {
	current := c.Query("page")
	if current == "" {
		currentPageNo = 1 //隐含显示第一页
		return
	}
	currentPageNo, err = strconv.Atoi(current)
	if err != nil {
		currentPageNo = 1 //默认第一页
		return
	}
	return
}
