package model

// 订单
type Order struct {
	//id自增
	Id int `json:"id" gorm:"column:id"`
	//订单idc
	OrderId string `json:"orderId" gorm:"column:order_id"`
	//客户id
	UserId string `json:"userId" gorm:"column:user_id"`
	//商品名称 包含类别
	ProductName string `json:"productName" gorm:"column:product_name"`
	//商品num
	ProductCount int `json:"productCount" gorm:"column:product_count"`
	//总价格
	TotalPrice float64 `json:"totalPrice" gorm:"column:total_price"`
	//订单状态
	Status int `json:"status" gorm:"column:status"`
	//创建时间
	CreateTime string `json:"createTime" gorm:"column:create_time"`
	//订单地址
	UserAddress string `json:"userAddress" gorm:"column:user_address"`
	//联系电话
	Mobile string `json:"mobile" gorm:"column:mobile"`
	//图片编号
	Image_Id string `json:"imageId" gorm:"column:image_id"`
	//图片url
	Url string `json:"url" gorm:"column:url"`
	//类别1
	Value1 string `json:"value1" gorm:"column:value1"`
	//类别2
	Value2 string `json:"value2" gorm:"column:value2"`
}

// 订单创建的中间状态
type OrderTemp struct {
	CartID       int     `json:"cartID" binding:"required"`
	ProductID    string  `json:"productID" binding:"required"`
	ProductLink  string  `json:"productLink" binding:"required"`
	ProductName  string  `json:"productName" binding:"required"`
	ProductCount int     `json:"productCount" binding:"required"`
	ProductPrice float64 `json:"productPrice" binding:"required"`
	Option1_id   string  `json:"option1_id" binding:"required"`
	Type1        string  `json:"type1" binding:"required"`
	Value1       string  `json:"value1" binding:"required"`
	Image_Id     string  `json:"image_id" binding:"required"`
	Url          string  `json:"url" binding:"required"`
	Option2_id   string  `json:"option2_id" binding:"required"`
	Type2        string  `json:"type2" binding:"required"`
	Value2       string  `json:"value2" binding:"required"`
	Selected     bool    `json:"selected" binding:"required"`
	UserId       string  `json:"userID" binding:"required"`
	TotalPrice   string  `json:"TotalPrice" binding:"required"`
	UserAddress  string  `json:"userAddress"`
	Mobile       string  `json:"mobile"`
}
