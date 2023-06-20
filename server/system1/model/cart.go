package model

type Cart struct {
	//id
	ID int `json:"cartID" gorm:"column:id"`
	//商品ID
	ProductID string `json:"productID" gorm:"column:productID"`
	//用户ID
	UserID string `json:"userID" gorm:"column:userID"`
	//图片链接
	ProductLink string `json:"productLink" gorm:"column:productLink"`
	//商品名称
	ProductName string `json:"productName" gorm:"column:productName"`
	//商品数量
	ProductCount int `json:"productCount" gorm:"column:productCount"`
	//商品价格
	ProductPrice float32 `json:"productPrice" gorm:"column:productPrice"`
	//商品种类
	Options []Option `gorm:"many2many:cart_options"`
}
