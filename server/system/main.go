package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Application(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Gin HTML Example",
	})
}
func main() {
	//建立数据库
	db := InitDB()
	db.Begin()
	// 创建一个默认的路由引擎
	r := gin.Default()
	r.StaticFS("/js", http.Dir("dist/js"))
	r.StaticFS("/css", http.Dir("dist/css"))
	r.LoadHTMLGlob("dist/index.html")
	r.GET("/", Application)
	r.POST("/api/Login", Login)
	r.POST("/api/Register", Register)
	r.Run()
	db.Close()
}
