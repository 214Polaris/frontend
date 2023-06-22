package main

import (
	"github.com/gin-contrib/cors"
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
	//r.StaticFS("/js", http.Dir("dist/js"))
	//r.StaticFS("/css", http.Dir("dist/css"))
	//r.StaticFS("/fonts", http.Dir("dist/fonts"))
	//r.LoadHTMLGlob("dist/index.html")
	//r.Use(UserSrv.HandleRequest)
	// 配置跨域请求
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"} // 允许的前端地址
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "token", "Authorization",
		"Access-Control-Allow-Origin"} // 允许的请求头
	config.AllowMethods = []string{"GET", "POST"}
	r.Use(cors.New(config))
	r.GET("/", Application)
	r.GET("/api/goodslist", ProductSrc.SendAll)
	r.POST("/api/Search", ProductSrc.Search)
	r.POST("/api/details", ProductSrc.GetByID)
	r.POST("/api/details/info", ProductSrc.GetDetail)

	r.GET("/pay/alipay/return", OrderSrv.AliPayReturn)
	r.POST("/pay/alipay/notify", OrderSrv.AliPayNotify)
	r.POST("/api/alipay", OrderSrv.Alipay)
	r.POST("/api/getuserorders", OrderSrv.GetUserOrders)
	r.POST("/api/createorder", OrderSrv.CreateOrder)
	r.POST("/api/cancelorder", OrderSrv.CancelOrder)

	r.POST("/api/Login", UserSrv.Login)
	r.POST("/api/Register", UserSrv.Register)
	r.POST("/api/Edit", UserSrv.Edit)
	r.POST("/api/showUser", UserSrv.GetByID)
	r.POST("/api/EditPass", UserSrv.EditPassword)

	r.POST("/api/cart", CartSrv.GetTotal)
	r.POST("/api/add", CartSrv.Add)
	r.POST("/api/delete", CartSrv.Delete)
	r.Run(":8081")
	//
}
