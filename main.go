package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*
	1. 用户调用登录接口，根据userName+password或者userID生成token
	2. 把token放在Header中，用户再调用其他接口时，都会取下Header中的token进行解析
	3. 解析成功则继续调接口，失败则返回错误信息
*/

var jwtSecret = []byte("jwtSecretKey")

type CustomClaims struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func genToken() (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(10 * time.Hour)

	claims := CustomClaims{
		UserName: "pengjj",
		Password: "123456",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	fmt.Println(token)
	//Output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c
	return token, nil
}

func parseToken() (*CustomClaims, error) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"

	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
