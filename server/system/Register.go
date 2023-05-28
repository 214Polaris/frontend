package main

import (
	"net/http"
	_ "net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var u User
	u.Username = ctx.PostForm("user")
	u.Password = ctx.PostForm("pass")
	//密码加密
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	//将加密后的密码赋给u.Password
	u.Password = string(hasePassword)

	//判断用户是否存在
	if MyQuery(u) == true {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}

	//添加用户
	MyInsert(u)

	// c.JSON：返回JSON格式的数据
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}
