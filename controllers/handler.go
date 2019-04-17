package controllers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/middleware"
	"github.com/mesment/personblog/models"
	"log"
	"net/http"
)


//服务器出错
func ServerError(c *gin.Context, err error)  {
	log.Printf("Error:%v",err)
	c.HTML(http.StatusInternalServerError,"views/htmls/500.tmpl",err)
	return
}
//Client出错
func ClientError(c *gin.Context, errmsg string)  {
	log.Printf("%s",errmsg)
	c.HTML(http.StatusBadRequest,"views/htmls/404.tmpl",errmsg)
	return
}

//根据用户的登录状态来进行跳转
func UserHandler(c *gin.Context)  {
	username,islogin := IsLogin(c)

	//已登录，跳转到欢迎界面
	if islogin {
		user := models.User{
			UserName:username,
		}
		m := make(map[string] interface{})
		m["user"] = user
		m["islogin"] = true
		c.HTML(http.StatusFound,"views/htmls/welcome.tmpl",m)
	}

	//未登录，跳转到登录界面
	c.Redirect(http.StatusFound,"/user/login")

}


//从cookie中找到用户名则认为已登录，否则未登录
func IsLogin(c *gin.Context) (string,bool) {

	//通过从cookie中取出token来查找用户名
	token, err := c.Cookie("token")
	if err != nil {
		return "",false //状态未登录
	}
	name, err := GetUserNameFromToken(token)
	if err != nil {
		return "",false  //状态未登录
	}

	//查找成功，已登录
	return name,true
}

//从token中获取用户名，获取失败返回默认Guest用户
func GetDefultUserName(c *gin.Context) (login bool, name string ){
	//设置默认用户名
	var defaultName = "Guest"

	name,login = IsLogin(c)
	if !login {
		name = defaultName
	}
	return
}

//创建token
func CreateToken(authjwt *middleware.JWT, username string , exptime int64 ) (token string, err error){
	//设置 token有效时间
	expirationTime := exptime

	//创建JWT claims,包含用户名和超时时间
	claims := middleware.CustomClaims{
		Username:username,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt:expirationTime,
		},
	}
	log.Printf("claims: %v",claims)

	//创建token
	token, err = authjwt.CreateToken(claims)
	if err != nil {
		log.Printf("Token签名失败%v",err)
		return "", err
	}
	log.Printf("Token:%s",token)

	return token,nil
}

//从token中获取用户名，成功返回用户名，失败返回错误信息
func GetUserNameFromToken(tokenStr string) (username string ,err error) {
	// 解析token
	authjwt := middleware.JWT{}

	claims,err := authjwt.ParseToken(tokenStr)
	if err != nil {
		return
	}
	//遍历claims找出用户名
	for key,value := range claims {
		log.Printf("key:%v,value:%v",key, value)
		if key == "username" {
			username =  value.(string) //token中的用户名
			return
		}
	}

	err = errors.New("Token里没有用户名字段")
	return
}
