package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func Application(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Gin HTML Example",
	})
}
func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	r.StaticFS("/js", http.Dir("dist/js"))
	r.StaticFS("/css", http.Dir("dist/css"))
	r.StaticFS("/fonts", http.Dir("dist/fonts"))
	r.LoadHTMLGlob("dist/index.html")
	//r.Use(UserSrv.HandleRequest)
	r.GET("/", Application)
	r.GET("/api/goodslist", ProductSrc.SendAll)
	r.POST("/api/Login", UserSrv.Login)
	r.POST("/api/Register", UserSrv.Register)
	r.Run()
}
