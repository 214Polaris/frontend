package model

// 订单
type Order struct {
	//订单id
	OrderId string `json:"orderId" gorm:"column:order_id"`
	//客户id
	UserId string `json:"userId" gorm:"column:user_id"`
	//客服联系电话
	Mobile string `json:"mobile" gorm:"column:mobile"`
	//总价格
	TotalPrice int64 `json:"totalPrice" gorm:"column:total_price"`
	//支付状态
	PayStatus string `json:"payStatus" gorm:"column:pay_status"`
	//订单地址
	UserAddress string `json:"userAddress" gorm:"column:user_address"`
	//注销
	IsDeleted bool `json:"isDeleted" gorm:"column:is_deleted"`
	//商品列表 用id检索
	//ProductsItem []string `json:"productsItem" gorm:"column:products_item"`
	//支付url
	PayUrl string `json:"payUrl" gorm:"column:pay_url"`
	//订单物流状态
	//OrderStatus int `json:"orderStatus" gorm:"column:order_status"`

}
