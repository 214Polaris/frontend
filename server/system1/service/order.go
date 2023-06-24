package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
	"system/config"
	"system/data"
	"system/model"
	"time"
)

type OrderSrv interface {
	//根据订单号获取订单
	GetOrders(c *gin.Context)
	//根据用户id获取订单
	GetUserOrders(c *gin.Context)
	//支付唤醒 不需关注
	AliPayNotify(c *gin.Context)
	//支付回调 不需关注
	AliPayReturn(c *gin.Context)
	//创建订单
	CreateOrder(c *gin.Context)
	//创建支付
	Alipay(c *gin.Context)
	//取消订单 更新订单
	CancelOrder(c *gin.Context)
	//支付订单 更新订单
	FinishOrder(c *gin.Context)
}

type OrderService struct {
	Repo data.OrderRepoInterface
}

var temp string

// 根据订单号获取订单
func (srv *OrderService) GetOrders(c *gin.Context) {
	orderid := c.PostForm("orderId")

	orders := srv.Repo.GetOrdersById(orderid)
	if orders == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg":        "根据订单号获取订单失败",
			"entitylist": orders,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":        "根据订单号获取订单成功",
		"entitylist": orders,
	})
}

// 根据用户id获取订单
func (srv *OrderService) GetUserOrders(c *gin.Context) {
	var o model.Order
	o.UserId = c.PostForm("userID")

	orders := srv.Repo.GetUserOrders(o.UserId)
	if orders == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg":        "获取用户订单失败",
			"entitylist": orders,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":        "获取用户订单成功",
		"entitylist": orders,
	})
}

// 支付宝支付页面使用 支付状态识别和Todo业务 需要post orderId用于订单检索和更新 trade_status是支付宝环境自带的
func (srv *OrderService) AliPayNotify(c *gin.Context) {
	//var o model.Order
	orderid := c.PostForm("out_trade_no")
	tradeStatus := c.PostForm("status")
	orders := srv.Repo.GetOrders(orderid)
	if orders == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "找不到该订单",
		})
		return
	}

	// 支付页面关闭
	if tradeStatus == "TRADE_CLOSED" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "交易已关闭",
		})
	}

	// 支付成功
	if tradeStatus == "TRADE_SUCCESS" {
		//做自己的业务
		//更新订单支付状态
		for _, ors := range orders {
			ors.Status = 2
			_, err := srv.Repo.Edit(*ors)
			if err != nil {
				xlog.Error(err)
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"msg": "订单更新出错！",
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "订单支付成功！",
		})
	}
}

// 支付完成后回调 解析和验证 暂无post需求
func (srv *OrderService) AliPayReturn(c *gin.Context) {
	// 解析网址参数
	notifyReq, err := alipay.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		xlog.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "网址参数错误",
		})
		return
	}

	// 验签
	ok, err := alipay.VerifySign(config.AliPublicKey, notifyReq)
	if err != nil {
		xlog.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验签时参数错误",
		})
		return
	}
	msg := ""
	if ok {
		msg = "验签成功"
	} else {
		msg = "验签失败"
	}
	//做自己的业务
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

type entity struct {
	entitylist []model.OrderTemp `json:"entitylist" binding:"required"`
}

// 创建订单 需要post整个model.order中OrderTemp的属性 按数组形式获取后切片
func (srv *OrderService) CreateOrder(c *gin.Context) {
	var entitylist []model.OrderTemp
	//if err := c.ShouldBind(&entitylist); err != nil {=
	if err := c.ShouldBindJSON(&entitylist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "获取实体失败，不能生成订单！",
			"len": len(entitylist),
		})
		return
	}

	// 生成order_id、create_time, state设为1
	// 订单号获取 时间戳表示
	ts := time.Now().UnixMilli()
	outTradeNo := fmt.Sprintf("%d", ts)
	// 创建时间获取
	timeStr := time.Now().Format("2006-01-02 15:04:05") //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法

	// 创建订单
	println(len(entitylist))
	for _, ot := range entitylist {
		var on model.Order
		on.OrderId = outTradeNo
		on.UserId = ot.UserId
		on.ProductID = ot.ProductID
		on.ProductName = ot.ProductName
		on.ProductCount = ot.ProductCount
		pricetemp, _ := strconv.ParseFloat(ot.TotalPrice, 64)
		on.TotalPrice = pricetemp
		on.Status = 1
		on.CreateTime = timeStr
		on.UserAddress = ot.UserAddress
		on.Mobile = ot.Mobile
		on.Image_Id = ot.Image_Id
		on.Url = ot.Url
		on.Value1 = ot.Value1
		on.Value2 = ot.Value2

		_, err := srv.Repo.Add(on)
		if err != nil {
			xlog.Error(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg":     "订单创建失败！",
				"orderId": outTradeNo,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":     "订单创建完成！",
		"orderId": outTradeNo,
		"strr":    timeStr,
		"entity":  entitylist,
	})
}

type payentity struct {
	UserId      string `json:"userId" binding:"required"`
	OrderId     string `json:"orderId" binding:"required"`
	TotalPrice  string `json:"totalPrice" binding:"required"`
	Mobile      string `json:"mobile" binding:"required"`
	UserAddress string `json:"useraddress" binding:"required"`
}

// 支付宝页面设置 支付登录界面url获取  post：用户id，总金额，订单id
func (srv *OrderService) Alipay(c *gin.Context) {
	var pe payentity
	if err := c.ShouldBindJSON(&pe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "获取订单信息失败，不能生成订单！",
			"len": pe,
		})
		return
	}

	//获取url进行支付
	client, err := alipay.NewClient(config.AppId, config.PrivateKey, config.IsProduction)
	if err != nil {
		xlog.Error(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg":    "client初始化失败",
			"payUrl": "",
		})
		return
	}
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetNotifyUrl(config.NotifyURL).
		SetReturnUrl(config.ReturnURL)

	bm := make(gopay.BodyMap)
	// 支付标题 用户名
	bm.Set("subject", pe.UserId)
	// 支付账单号
	temp = pe.OrderId
	bm.Set("out_trade_no", pe.OrderId)
	// 支付金额
	bm.Set("total_amount", pe.TotalPrice)
	bm.Set("product_code", config.ProductCode)
	//设置自定义参数
	passback_params := "" + pe.Mobile + ";" + pe.UserAddress
	//将参数encode
	bm.Set("passback_params", passback_params)

	payUrl, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		xlog.Error(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg":    "payurl创建失败",
			"payUrl": "",
		})
		return
	}
	//println("temp:" + payUrl)
	//xlog.Debugf("payUrl", payUrl)
	// 以下为payurl的调整信息 如果payurl为https://openapi.alipaydev.com/gateway.do则将以下内容注释掉 否则endpayurl才是正确的url
	//editpayurl := strings.SplitAfter(payUrl, "/")
	//editpayurl[2] = "openapi.alipaydev.com/"
	//endpayurl := strings.Join(editpayurl, "")
	//fmt.Println("payurl: " + endpayurl)

	orders := srv.Repo.GetOrders(pe.OrderId)
	if orders == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "找不到该订单",
		})
		return
	}
	//更新订单状态
	for _, ors := range orders {
		ors.UserAddress = pe.UserAddress
		ors.Mobile = pe.Mobile
		_, err := srv.Repo.Edit(*ors)
		if err != nil {
			xlog.Error(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": "订单更新出错！",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    "成功生成支付订单",
		"payUrl": payUrl,
	})
}

// 用户取消订单 根据订单号修改订单信息
func (srv *OrderService) CancelOrder(c *gin.Context) {
	var o model.Order
	o.OrderId = c.PostForm("orderId")

	orders := srv.Repo.GetOrdersById(o.OrderId)
	for _, ors := range orders {
		ors.Status = 6
		_, err := srv.Repo.Edit(*ors)
		if err != nil {
			xlog.Error(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": "订单取消失败！",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "订单取消完成",
	})
}

// 用户支付订单后更新订单信息 需要post orderid mobile useraddress
func (srv *OrderService) FinishOrder(c *gin.Context) {
	orderId := c.PostForm("orderId")
	orders := srv.Repo.GetOrders(orderId)
	if orders == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "找不到该订单",
		})
		return
	}

	for _, ors := range orders {
		ors.Status = 2
		_, err := srv.Repo.Edit(*ors)
		if err != nil {
			xlog.Error(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": "订单更新出错！",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "订单支付完成",
	})
}
