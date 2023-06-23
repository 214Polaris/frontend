package data

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"system/model"
)

type OrderData struct {
	DB *gorm.DB
}

type OrderRepoInterface interface {
	Exist(Order model.Order) *model.Order
	GetOrders(id string) (pro []*model.Order)
	Add(Order model.Order) (*model.Order, error)
	Edit(Order model.Order) (bool, error)
	GetUserOrders(id string) (pro []*model.Order)
	GetOrdersById(id string) (pro []*model.Order)
}

func (o *OrderData) Exist(order model.Order) *model.Order {
	if order.OrderId != "" {
		o.DB.Model(&order).Where("order_id= ?", order.OrderId)
		return &order
	}
	return nil
}

func (o *OrderData) GetOrders(id string) (pro []*model.Order) {
	var orders []*model.Order
	db := o.DB
	if err := db.Where("order_id = ?", id).Find(&orders).Error; err != nil {
		return nil
	}
	return orders
}

func (o *OrderData) Add(order model.Order) (*model.Order, error) {
	err := o.DB.Create(order).Error
	if err != nil {
		return &order, fmt.Errorf("订单添加失败")
	}
	return &order, nil
}

func (o *OrderData) Edit(order model.Order) (bool, error) {
	if order.OrderId == "" {
		return false, errors.New("请传入参数")
	}
	err := o.DB.Model(&order).Where("order_id=?", order.OrderId).Updates(map[string]interface{}{
		"status":       order.Status,
		"user_address": order.UserAddress,
		"mobile":       order.Mobile,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (o *OrderData) GetUserOrders(id string) (pro []*model.Order) {
	var orders []*model.Order
	db := o.DB
	if err := db.Where("user_id = ?", id).Find(&orders).Error; err != nil {
		return nil
	}
	return orders
}

func (o *OrderData) GetOrdersById(id string) (pro []*model.Order) {
	var orders []*model.Order
	db := o.DB
	if err := db.Where("order_id = ?", id).Find(&orders).Error; err != nil {
		return nil
	}
	return orders
}
