package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"system/model"
	"time"
)

var jwtKey = []byte("secret_key") // 密钥

// GenerateToken 生成JWT
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 设置过期时间为1天

	// 创建声明
	claims := &model.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "issuer",
			Subject:   "user-token",
		},
	}

	// 创建令牌并签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken 验证JWT
func VerifyToken(tokenString string) (*jwt.Token, *model.Claims, error) {
	// 解析JWT并验证签名
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return token, nil, err
	}

	// 验证令牌是否有效
	claims, ok := token.Claims.(*model.Claims)
	if !ok || !token.Valid {
		return token, nil, fmt.Errorf("invalid token")
	}

	return token, claims, nil
}

// 更新token
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*model.Claims)
	if !ok {
		panic("token is valid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	Tk, _ := GenerateToken(claims.Username)
	return Tk
}
