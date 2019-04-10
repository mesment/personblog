package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

var JWTKey = []byte("my_secret_key")

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

// Create a new token object, specifying signing method and the claims
// you would like it to contain.
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
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
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

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
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


//解析token例子
func AuthHandler(c *gin.Context)  {
	//检查token
	// sample token string taken from the New example
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return JWTKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

}



