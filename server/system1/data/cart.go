package data

import (
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
	repo.DB.Where("userID = ?", userID).Find(&carts).Count(&count)
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
			Where("productID = ? and userID = ? and type1 = ? and type2 = ?", cart.ProductID, cart.UserID, cart.Type1, cart.Type2).
			Update("productCount", gorm.Expr("productCount + ?", cart.ProductCount))
	}
	var count int64
	repo.DB.Model(&cart).Count(&count)
	cart.ID = int(count) + 1
	err := repo.DB.Model(&cart).Create(cart).Error
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
