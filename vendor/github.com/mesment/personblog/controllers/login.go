package controllers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/auth"
	"github.com/mesment/personblog/dao/db"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Text string `json:"text"`
}


//返回用户登录页面
func ShowLogin(c *gin.Context) {

	c.HTML(http.StatusOK, "views/htmls/login.tmpl", nil)
}

//读取用户提交的表单信息进行校验，登录成功创建JWT token
func LoginHandler(c *gin.Context) {
	var errmsg string

	username := c.PostForm("loginName")
	password := c.PostForm("password")

	if username == "" || password == "" {
		errmsg = "用户名或密码不能为空"
		ClientError(c, errmsg)
	}
	log.Printf("username:%s,password:%s\n",username,password)

	_, err := db.GetUser(username, password)
	if err != nil {
		err := errors.New("登录失败:" + err.Error())
		ServerError(c, err)
	}

	//登录成功,设置token
	authjwt:= auth.JWT{}

	//声明JWT token有效时间1个小时
	expirationTime := time.Now().Add(3600 * time.Second)
	//创建JWT claims,包含用户名和超时时间
	claims := auth.CustomClaims{
		Username:username,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt:expirationTime.Unix(),
		},
	}
	log.Printf("claims: %v",claims)

	//创建token
	tokenString, err := authjwt.CreateToken(claims)
	if err != nil {
		log.Printf("Token签名失败:%v", err)
		ServerError(c, err)
	}
	log.Printf("Token:%s",tokenString)

	//将新的token写入到cookie，超时时间等于token超时时间
	c.SetCookie("token",tokenString,int(expirationTime.Unix()),
		"/","",false,true)

	//跳转回首页
	//c.String(http.StatusOK,"登录成功")
	c.Redirect(http.StatusOK, "/")
}




