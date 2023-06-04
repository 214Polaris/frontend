package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	ctx.JSON(200, products)
}
