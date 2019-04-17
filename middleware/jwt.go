package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mesment/personblog/pkg/setting"
	"log"
	"net/http"
	"time"
)

const (
	privateKeyPath = "" // openssl genrsa -out app.rsa 1024

	publicKeyPath = "" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

/*
var (
	VerifyKey []byte //验证的key
	SignKey  []byte //签名的key
)
*/

var JWTKey = []byte(setting.AppCfg.PagJWTSecret)

//错误声明
var (
	TokenExpired error = errors.New("Token 已过期")
	TokenNotValidYet error = errors.New("Token 尚未生效")
	TokenErrorformed error = errors.New("Token 格式错误")
	TokenInvalid error = errors.New("Token 无效")
)

/*
func init() {
	var err error
	//读取文件私钥
	SignKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal("读取私钥文件失败")
		return
	}

	//读取公钥文件
	VerifyKey, err = ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal("读取公钥文件失败")
		return
	}

}
*/


//token 鉴权 middleware
func TestAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//通过从cookie中取出token来进行验证用户是否已登录
		tokenStr, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				//如果没有设置cookie，提示需要登录
				errs := errors.New("请先登录")
				log.Printf("没有权限，请先登录,%v",err)
				c.HTML(http.StatusUnauthorized,"views/htmls/404.tmpl",errs)
				return
			}
		}

		log.Printf("token 字符串:%s\n",tokenStr)

		claims := &CustomClaims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return JWTKey , nil
		})
		newErr := errors.New("请先登录")
		if err != nil {
			c.HTML(http.StatusUnauthorized,"views/htmls/404.tmpl",newErr)
			return
		}
		if err != nil || !tkn.Valid  {
			if err == jwt.ErrSignatureInvalid {
				c.HTML(http.StatusUnauthorized,"views/htmls/404.tmpl",newErr)
				return
			}
			c.HTML(http.StatusBadRequest,"views/htmls/404.tmpl",newErr)
			return
		}
		c.Next()
	}
}



type JWT struct {
	Token string `json:"token"`
}

//载荷
type CustomClaims struct {
	Username string	`json:"username"`
	jwt.StandardClaims
}

func (j *JWT)SetToken(token string)  {
	j.Token = token
}

func (j *JWT)GetToken()(token string)  {
	return j.Token
}

//声明通过claims,和rs256签名算法创建Token 字符串
//成功返回签名后的JWT字符串,失败返回错误信息
func (j *JWT) CreateToken(claims CustomClaims) (encodedTokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encodedTokenStr,err = token.SignedString(JWTKey)
	j.Token = encodedTokenStr
	return
}

func (j *JWT) NewToken(username string) (t *JWT, err error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodRS512)
	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token.Claims = claims
	encodedTokenString,err := token.SignedString(JWTKey)
	if err != nil {
		return
	}
	t = &JWT{encodedTokenString}

	return
}

//解析Token
func (j *JWT) ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})
	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError:   //Token校验错误
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorMalformed:
				return nil, TokenErrorformed
			case jwt.ValidationErrorExpired:
				return nil, TokenExpired
			case jwt.ValidationErrorNotValidYet:
				return nil,TokenNotValidYet
			default:
				return nil, TokenInvalid
			}
		default:  //其他错误

		return nil, errors.New("Token 校验失败")
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

//刷新Token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}


