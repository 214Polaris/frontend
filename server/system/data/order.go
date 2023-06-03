package data

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"system/model"
	"system/utils"
)

type OrderData struct {
	DB *gorm.DB
}

type OrderRepoInterface interface {
	List(req *model.ListPage) (Orders []*model.Order, err error)
	GetTotal(req *model.ListPage) (total int64, err error)
	Get(Order model.Order) (*model.Order, error)
	Exist(Order model.Order) *model.Order
	ExistByOrderID(id string) *model.Order
	Add(Order model.Order) (*model.Order, error)
	Edit(Order model.Order) (bool, error)
	Delete(u model.Order) (bool, error)
}

func (o *OrderData) List(req *model.ListPage) (Orders []*model.Order, err error) {
	//fmt.Println(req)
	db := o.DB
	limit, offset := utils.Page(req.PageSize, req.Page) // 分页
	sort := utils.Sort(req.Sort)
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&Orders).Error; err != nil {
		return nil, err
	}
	return Orders, nil
}

func (o *OrderData) GetTotal(req *model.ListPage) (total int64, err error) {
	var orders []model.Order
	db := o.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&orders).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (o *OrderData) Get(order model.Order) (*model.Order, error) {
	if err := o.DB.Where(&order).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderData) Exist(order model.Order) *model.Order {
	if order.OrderId != "" {
		o.DB.Model(&order).Where("order_id= ?", order.OrderId)
		return &order
	}
	return nil
}

func (o *OrderData) ExistByOrderID(id string) *model.Order {
	var order model.Order
	o.DB.Where("order_id = ?", id).First(&order)
	return &order
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
	o1 := &model.Order{}
	err := o.DB.Model(o1).Where("order_id=?", order.OrderId).Updates(map[string]interface{}{
		"nick_name":    order.NickName,
		"mobile":       order.Mobile,
		"pay_status":   order.PayStatus,
		"order_status": order.OrderStatus,
		"extra_info":   order.ExtraInfo,
		"user_address": order.UserAddress,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (o *OrderData) Delete(order model.Order) (bool, error) {
	err := o.DB.Model(&order).
		Where("order_id=?", order.OrderId).
		Update("is_deleted", order.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
