package model

// 订单
type Order struct {
	//订单id
	OrderId string `json:"orderId" gorm:"column:order_id"`
	//客户id
	UserId string `json:"userId" gorm:"column:user_Id"`
	//订单名称
	NickName string `json:"nickName" gorm:"column:nick_name"`
	//客服联系电话
	Mobile string `json:"mobile" gorm:"column:mobile"`
	//总价格
	TotalPrice int64 `json:"totalPrice" gorm:"column:total_price"`
	//支付状态
	PayStatus int `json:"payStatus" gorm:"column:pay_status"`
	//支付方式
	PayType int `json:"payType" gorm:"column:pay_type"`
	//支付时间
	PayTime string `json:"payTime" gorm:"column:pay_time"`
	//订单物流状态
	OrderStatus int `json:"orderStatus" gorm:"column:order_status"`
	//备注 额外信息
	ExtraInfo string `json:"extraInfo" gorm:"column:extra_info"`
	//订单地址
	UserAddress string `json:"userAddress" gorm:"column:user_address"`
	//注销
	IsDeleted bool `json:"isDeleted" gorm:"column:is_deleted"`
}
