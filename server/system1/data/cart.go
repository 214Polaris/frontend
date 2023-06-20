package data

import (
	"fmt"
	"gorm.io/gorm"
	"system/model"
)

type CartData struct {
	DB *gorm.DB
}

type CartRepoInterface interface {
	// 获取一个用户的全部记录
	GetTotal(userID string) []*model.Cart
	// 查询一个用户的一个商品记录
	Query(cart *model.Cart) bool
	// 增加数量
	Plus(id int, productCount int) bool
	// 减少数量
	Diminish(id int, productCount int) bool
	// 增加商品
	Add(cart model.Cart) bool
	// 删除商品
	Delete(id int) bool
}

// 获取一个用户的全部记录
func (repo *CartData) GetTotal(userID string) []*model.Cart {
	var carts []*model.Cart
	var count int64
	repo.DB.Preload("Options").Where("userID = ?", userID).Find(&carts).Count(&count)
	fmt.Print(carts[0].Options)
	if count > 0 {
		return carts
	}
	return nil
}

// 查询一个用户的一个商品记录
func (repo *CartData) Query(cart *model.Cart) bool {
	var count int64
	repo.DB.Where("id = ?", cart.ID).Find(&cart).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// 增加数量
func (repo *CartData) Plus(id int, productCount int) bool {
	var cart model.Cart
	err := repo.DB.Model(&cart).
		Where("id = ?", id).
		Update("productCount", productCount+1).Error
	if err != nil {
		return false
	}
	return true
}

// 减少数量
func (repo *CartData) Diminish(id int, productCount int) bool {
	var cart model.Cart
	err := repo.DB.Model(&cart).
		Where("id = ?", id).
		Update("productCount", productCount-1).
		Error
	if err != nil {
		return false
	}
	return true
}

// 增加商品
func (repo *CartData) Add(cart model.Cart) bool {
	//先判断有没有该商品在购物车
	if repo.Query(&cart) == true {
		repo.DB.Model(&cart).
			Where("id = ?", cart.ID).
			Update("productCount", gorm.Expr("productCount + ?", cart.ProductCount))
	}
	// 创建或更新 CartItem 的选项关联
	var options []model.Option
	for _, option := range cart.Options {
		options = append(options, option)
	}
	repo.DB.Model(&cart).Association("Options").Replace(&options)
	err := repo.DB.Create(cart)
	if err != nil {
		return false
	}
	return true
}

// 删除商品
func (repo *CartData) Delete(id int) bool {
	var cart model.Cart
	err := repo.DB.Where("id = ?", id).Delete(&cart)
	if err != nil {
		return false
	}
	return true
}
