package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Login(ctx *gin.Context) {
	var u User
	u.Username = ctx.PostForm("user")
	u.Password = ctx.PostForm("pass")

	//判断用户是否存在
	if MyQuery(u) == false {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//判断密码是否正确
	if CheckPassword(u) == false {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不正确"})
		return
	}

	//发放token给前端
	token := "11"

	// c.JSON：返回JSON格式的数据
	ctx.JSON(200, gin.H{
		"code": 200,
		"token": token,
		"msg":  "登录成功",
	})
}
