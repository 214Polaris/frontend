package model

// 订单
type Order struct {
	//id自增
	Id string `json:"id" gorm:"column:id"`
	//订单id
	OrderId string `json:"orderId" gorm:"column:order_id"`
	//客户id
	UserId string `json:"userId" gorm:"column:user_id"`
	//商品id
	ProductId string `json:"productId" gorm:"column:product_id"`
	//商品num
	ProductNum string `json:"productNum" gorm:"column:product_num"`
	//总价格
	TotalPrice int64 `json:"totalPrice" gorm:"column:total_price"`
	//订单状态
	Status int64 `json:"status" gorm:"column:status"`
	//创建时间
	CreateTime string `json:"createTime" gorm:"column:create_time"`
	//订单地址
	UserAddress string `json:"userAddress" gorm:"column:user_address"`
	//联系电话
	Mobile string `json:"mobile" gorm:"column:mobile"`
}

// 订单创建的中间状态
type OrderTemp struct {
	//客户id
	UserId string `json:"userId" gorm:"column:user_id"`
	//商品id
	ProductId string `json:"productId" gorm:"column:product_id"`
	//商品num
	ProductNum string `json:"productNum" gorm:"column:product_num"`
	//总价格
	TotalPrice int64 `json:"totalPrice" gorm:"column:total_price"`
	//订单地址
	UserAddress string `json:"userAddress" gorm:"column:user_address"`
	//联系电话
	Mobile string `json:"mobile" gorm:"column:mobile"`
}
