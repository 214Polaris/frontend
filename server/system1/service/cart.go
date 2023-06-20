package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"system/data"
	"system/model"
)

type CartSrv interface {
	// 获取一个用户的全部记录
	GetTotal(ctx *gin.Context)
	// 查询一个用户的一个商品记录
	Query(ctx *gin.Context)
	// 增加数量
	Plus(ctx *gin.Context)
	// 减少数量
	Diminish(ctx *gin.Context)
	// 增加商品
	Add(ctx *gin.Context)
	// 删除商品
	Delete(ctx *gin.Context)
}

type CartService struct {
	Repo data.CartRepoInterface
}

// 定义报文格式
type CartRequest struct {
	CartID       int            `json:"cartID" binding:"required"`
	ProductID    string         `json:"productID" binding:"required"`
	UserID       string         `json:"userID" binding:"required"`
	ProductLink  string         `json:"productLink" binding:"required"`
	ProductName  string         `json:"productName" binding:"required"`
	ProductCount int            `json:"productCount" binding:"required"`
	ProductPrice float32        `json:"productPrice" binding:"required"`
	Options      []model.Option `json:"options" binding:"required"`
}

// 获取一个用户的全部记录
func (srv *CartService) GetTotal(ctx *gin.Context) {
	userID := ctx.PostForm("userID")
	carts := srv.Repo.GetTotal(userID)
	if carts == nil {
		ctx.JSON(421, gin.H{"msg": "获取购物车数据失败"})
		return
	}
	ctx.JSON(200, carts)
}

// 查询一个用户的一个商品记录
func (srv *CartService) Query(ctx *gin.Context) {
	var cart *model.Cart
	cart.ProductID = ctx.PostForm("productID")
	cart.UserID = ctx.PostForm("userID")
	if srv.Repo.Query(cart) == false {
		ctx.JSON(401, gin.H{"msg": "没有该商品"})
		return
	}
	ctx.JSON(200, cart)
}

// 增加数量
func (srv *CartService) Plus(ctx *gin.Context) {
	cartID := ctx.PostForm("cartID")
	id, _ := strconv.Atoi(cartID)
	ProductCountStr := ctx.PostForm("productCount")
	ProductCount, _ := strconv.Atoi(ProductCountStr)
	if srv.Repo.Plus(id, ProductCount) == false {
		ctx.JSON(401, gin.H{"msg": "添加失败"})
	}
	ctx.JSON(200, gin.H{"msg": "添加成功"})
}

// 减少数量
func (srv *CartService) Diminish(ctx *gin.Context) {
	cartID := ctx.PostForm("cartID")
	id, _ := strconv.Atoi(cartID)
	ProductCountStr := ctx.PostForm("productCount")
	//string 转 int
	ProductCount, _ := strconv.Atoi(ProductCountStr)
	if srv.Repo.Diminish(id, ProductCount) == false {
		ctx.JSON(401, gin.H{"msg": "减少失败"})
	}
	ctx.JSON(200, gin.H{"msg": "减少成功"})
}

// 增加商品
func (srv *CartService) Add(ctx *gin.Context) {
	var cartReq CartRequest
	if err := ctx.ShouldBindJSON(&cartReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 创建或更新 CartItem
	cartItem := model.Cart{
		ID:           cartReq.CartID,
		ProductID:    cartReq.ProductID,
		UserID:       cartReq.UserID,
		ProductLink:  cartReq.ProductLink,
		ProductName:  cartReq.ProductName,
		ProductCount: cartReq.ProductCount,
		ProductPrice: cartReq.ProductPrice,
		Options:      cartReq.Options,
	}
	if srv.Repo.Add(cartItem) == false {
		ctx.JSON(401, gin.H{"msg": "添加失败"})
	}
	ctx.JSON(200, gin.H{"msg": "添加成功"})
}

// 删除商品
func (srv *CartService) Delete(ctx gin.Context) {
	cartID := ctx.PostForm("cartID")
	id, _ := strconv.Atoi(cartID)
	if srv.Repo.Delete(id) == false {
		ctx.JSON(401, gin.H{"msg": "删除失败"})
		return
	}
	ctx.JSON(200, gin.H{"msg": "删除成功"})
}
