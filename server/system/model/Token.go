package model

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims 结构体定义了JWT的声明
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
