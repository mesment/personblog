package controllers

import (
	"fmt"
	"github.com/mesment/personblog/models"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/dao/db"
	"net/http"
)

//Get 获取留言页面
func LeaveMessage(c *gin.Context)  {
	//获取登录状态用户名,用户名
	var login,username = GetDefultUserName(c)

	var user = models.User{
		UserName:username,
	}

	msgs, err := db.GetMessage(10) 	//返回前10条留言
	if err != nil {
		fmt.Println("LeaveMessage: 获取留言列表失败%v",err)
		ServerError(c,err)
	}
	data:= make(map[string]interface{})
	data["islogin"] = login
	data["user"] = user
	data["msgs"]  = msgs
	c.HTML(http.StatusOK,"views/htmls/message.tmpl",data)

}

//Post 提交留言
func SubmitMessage(c *gin.Context) {
	//获取用户名
	var _,username = GetDefultUserName(c)

	//获取留言内容
	msgContent := c.PostForm("content")
	if len(msgContent)== 0 || msgContent =="" {
		err := fmt.Errorf("留言不能为空！！")
		ClientError(c,err.Error())
		return
	}
	err := db.AddMessage(username, msgContent)
	if err != nil {
		fmt.Printf("提交留言失败,%v", err)
		ServerError(c,err)
		return
	}

	//刷新页面
	c.Redirect(http.StatusFound,"/leave/message")
}
