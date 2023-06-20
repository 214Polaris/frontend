package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"net/http"
	"strings"
	"system/config"
	"system/data"
	"system/model"
	"time"
)

type OrderSrv interface {
	GetUserOrders(ctx *gin.Context)
	AliPayNotify(c *gin.Context)
	AliPayReturn(c *gin.Context)
	Alipay(c *gin.Context)
}

type OrderService struct {
	Repo data.OrderRepoInterface
}

func (srv *OrderService) GetUserOrders(c *gin.Context) {
	var o model.Order
	o.OrderId = c.PostForm("orderId")

	orders := srv.Repo.GetUserOrders(o.OrderId)
	if orders == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "获取用户订单失败"})
		return
	}
	c.JSON(200, orders)
}

// 支付宝支付状态识别和Todo业务 需要post orderId用于订单检索和更新
func (srv *OrderService) AliPayNotify(c *gin.Context) {
	var o model.Order
	o.OrderId = c.PostForm("orderId")
	tradeStatus := c.PostForm("trade_status")

	order := srv.Repo.ExistByOrderID(o.OrderId)
	if order == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "找不到该订单",
		})
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
		//更新订单支付
		order.PayStatus = "已支付"
		_, err := srv.Repo.Edit(*order)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "订单状态更新失败",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "交易成功！",
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
			"msg": "参数错误",
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

// 支付宝页面设置、url获取、订单创建 post：订单id，用户id，客服电话，总金额，收件地址
func (srv *OrderService) Alipay(c *gin.Context) {
	var o model.Order
	// 订单号获取 时间戳表示
	ts := time.Now().UnixMilli()
	outTradeNo := fmt.Sprintf("%d", ts)
	o.OrderId = outTradeNo
	o.UserId = c.PostForm("userId")
	o.Mobile = c.PostForm("mobile")
	o.TotalPrice = c.GetInt64("totalPrice")
	o.PayStatus = "未支付"
	o.UserAddress = c.PostForm("userAddress")
	//获取url进行支付
	client, err := alipay.NewClient(config.AppId, config.PrivateKey, config.IsProduction)
	if err != nil {
		xlog.Error(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "client初始化失败",
		})
		return
	}
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetNotifyUrl(config.NotifyURL).
		SetReturnUrl(config.ReturnURL)

	bm := make(gopay.BodyMap)
	// 支付标题 可以改成用户名
	bm.Set("subject", o.UserId)
	// 支付账单号 这里用时间戳
	bm.Set("out_trade_no", outTradeNo)
	// 支付金额
	bm.Set("total_amount", o.TotalPrice)
	bm.Set("product_code", config.ProductCode)

	payUrl, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		xlog.Error(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "payurl初始化失败",
		})
		return
	}
	//println("temp:" + payUrl)
	//xlog.Debugf("payUrl", payUrl)

	// 以下为payurl的调整信息 如果payurl为https://openapi.alipaydev.com/gateway.do则将以下内容注释掉 否则endpayurl才是正确的url
	editpayurl := strings.SplitAfter(payUrl, "/")
	editpayurl[2] = "openapi.alipaydev.com/"

	endpayurl := strings.Join(editpayurl, "")
	fmt.Println("payurl: " + endpayurl)
	o.PayUrl = endpayurl

	// 创建订单
	srv.Repo.Add(o)

	c.JSON(http.StatusOK, gin.H{
		"msg": "成功生成订单",
	})
}
