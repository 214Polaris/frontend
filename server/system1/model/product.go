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
	//参数
	Options []Option `gorm:"many2many:product_options"`
	//参数图片
	Images []Image `gorm:"many2many:product_images"`
}

// 参数
type Option struct {
	ID string `gorm:"primaryKey"`
	//类型
	Type string
	//类型名字
	Value string
	//关联
	//Products []Product `gorm:"many2many:product_options"`
}

// 参数图片
type Image struct {
	ID string `gorm:"primaryKey"`
	//图片链接
	URL string
	//关联
	//Products []Product `gorm:"many2many:product_images"`
}
