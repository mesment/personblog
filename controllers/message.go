package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/mesment/personblog/dao/db"
)

//Get 获取留言页面
func LeaveMessage(c *gin.Context)  {

	msgs, err := db.GetMessage(10) 	//返回前50条留言
	if err != nil {
		fmt.Println("LeaveMessage: 获取留言列表失败%v",err)
		ServerError(c,err)
	}
	c.HTML(http.StatusOK,"views/htmls/message.tmpl",msgs)

}

//Post 提交留言
func SubmitMessage(c *gin.Context) {
	username:= ""
	token,err := c.Cookie("token")
	if err != nil {
		username = "Guest"
	}
	log.Printf("token:%s\n",token)

	msgContent := c.PostForm("content")
	if len(msgContent)== 0 || msgContent =="" {
		err := fmt.Errorf("留言不能为空！！")
		ClientError(c,err.Error())
		return
	}
	err = db.AddMessage(username, msgContent)
	if err != nil {
		fmt.Printf("提交留言失败,%v", err)
		ServerError(c,err)
		return
	}

	//刷新页面
	c.Redirect(http.StatusFound,"/leave/message")
}
