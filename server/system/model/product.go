package model

// 产品
type Product struct {
	//商品id
	ProductId string `json:"productId" gorm:"column:id"`
	//商品名称
	ProductName string `json:"productName" gorm:"column:name"`
	//简介
	ProductIntro string `json:"productIntro" gorm:"column:introduction"`
	//分类
	ProductCatagory string `json:"productCategory" gorm:"column:category"`
	//进价
	ProductPrice float64 `json:"productPrice" gorm:"column:price"`
	//图片链接
	ProductLink string `json:"productLink" gorm:"column:link"`
}
