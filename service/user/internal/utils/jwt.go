package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var JwyKey = []byte("hdunb")

type MyClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

//生成token
func SetToken(uid int) string {
	//设置到期时间
	expireTime := time.Now().Add(24 * time.Hour) //24小时
	//创建Jwt声明
	SetClaims := MyClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	//使用用于签名的算法和令牌
	reqClain := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	//创建jwt字符串
	token, err := reqClain.SignedString(JwyKey)
	//如果出错，则返回服务器内部错误
	if err != nil {
		return ""
	}
	return token
}

//验证token
func CheckToken(token string) (*MyClaims, error) {

	//解析jwt字符串并将结果存储
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwyKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, nil
	}
	return nil, nil
}

//验证token
func GetUserIdByToken(token string) (int) {

	//解析jwt字符串并将结果存储
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwyKey, nil
	})
	return setToken.Claims.(*MyClaims).UserId
}