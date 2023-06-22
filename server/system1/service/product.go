package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"system/data"
)

type ProductSrv interface {
	//List(req *model.ListPage) (Products []*model.Product, err error)
	//GetTotal(req *model.ListPage) (total int, err error)
	//Get(Product model.Product) (*model.Product, error)
	//Exist(Product model.Product) *model.Product
	//ExistByProductID(id string) *model.Product
	//Add(Product model.Product) (*model.Product, error)
	//Edit(Product model.Product) (bool, error)
	//Delete(id string) (bool, error)
	SendAll(ctx *gin.Context)
	Search(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Paging(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
}

type ProductService struct {
	Repo data.ProductRepoInterface
}

//func (srv *ProductService) List(req *model.ListPage) (products []*model.Product, err error) {
//	return srv.Repo.List(req)
//}
//
////	func (srv *ProductService) GetTotal(req *model.ListPage) (total int64, err error) {
////		return srv.Repo.GetTotal(req)
////	}
//func (srv *ProductService) Get(product model.Product) (*model.Product, error) {
//	return srv.Repo.Get(product)
//}
//func (srv *ProductService) Exist(product model.Product) *model.Product {
//	return srv.Repo.Exist(product)
//}
//
//func (srv *ProductService) ExistByProductID(id string) *model.Product {
//	return srv.Repo.ExistByProductID(id)
//}
//
//func (srv *ProductService) Add(product model.Product) (*model.Product, error) {
//	// 如果id为空 给一个uuid
//	if product.ProductId == "" {
//		product.ProductId = uuid.NewV4().String()
//	}
//	return srv.Repo.Add(product)
//}
//func (srv *ProductService) Edit(product model.Product) (bool, error) {
//	p := srv.ExistByProductID(product.ProductId)
//	if p == nil || p.ProductName == "" {
//		return false, errors.New("参数传递错误")
//	}
//	return srv.Repo.Edit(product)
//}
//func (srv *ProductService) Delete(id string) (bool, error) {
//	if id == "" {
//		return false, errors.New("参数错误")
//	}
//
//	p := srv.ExistByProductID(id)
//	if p == nil {
//		return false, errors.New("参数错误")
//	}
//	return srv.Repo.Delete(*p)
//}

func (srv *ProductService) SendAll(ctx *gin.Context) {
	products := srv.Repo.GetTotal()
	//获取所有物品失败，返回状态码422
	if products == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "获取失败"})
		return
	}
	ctx.JSON(200, gin.H{
		"products": products,
	})
}

func (srv *ProductService) Search(ctx *gin.Context) {
	name := ctx.PostForm("name")
	page := ctx.PostForm("pageNum")
	//名字为空不做操作
	if len(name) == 0 {
		return
	}
	products := srv.Repo.Get(name)
	//没有找到结果，返回状态码422
	if len(products) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "没有找到该商品"})
	}
	//以下为分页操作
	//确定要发送的商品的下标
	productsNum := len(products)
	//总页数
	totalPage := (productsNum-1)/10 + 1
	//确定商品下标范围
	pageNum, _ := strconv.Atoi(page)
	startIndex := (pageNum - 1) * 10
	endIndex := startIndex + 9
	//防止越界
	if endIndex+1 > productsNum {
		endIndex = productsNum - 1
	}
	selectedProducts := products[startIndex : endIndex+1]
	ctx.JSON(200, gin.H{
		"totalPage": totalPage,
		"products":  selectedProducts,
	})
}

func (srv *ProductService) GetByID(ctx *gin.Context) {
	id := ctx.PostForm("productID")
	//没有找到结果，返回状态码422
	product := srv.Repo.GetByID(id)
	if product == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "没有找到该商品"})
	}
	ctx.JSON(200, product)
}

func (srv *ProductService) GetDetail(ctx *gin.Context) {
	id := ctx.PostForm("productID")
	//没有找到结果，返回状态码401
	product := srv.Repo.GetOptions(id)
	fmt.Print(id)
	if product == nil {
		ctx.JSON(401, gin.H{"code": 401, "msg": "没有找到该商品"})
	}
	fmt.Print(len(product.Options))
	dataWithPic := []map[string]interface{}{}
	dataWithoutPic := []map[string]interface{}{}
	imageLen := len(product.Images)
	for i := 0; i < len(product.Options); i++ {
		if i < imageLen {
			item := map[string]interface{}{
				"option": product.Options[i],
				"image":  product.Images[i],
			}
			dataWithPic = append(dataWithPic, item)
		} else {
			item := map[string]interface{}{
				"option": product.Options[i],
			}
			dataWithoutPic = append(dataWithoutPic, item)
		}
	}
	ctx.JSON(200, gin.H{"dataWithPic": dataWithPic, "dataWithoutPic": dataWithoutPic})
}
