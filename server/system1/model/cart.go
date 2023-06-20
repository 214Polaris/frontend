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
	//商品种类1
	OptionID1 string `json:"option1_id" gorm:"column:option1_id"`
	//商品选项1
	Type1 string `json:"type1" gorm:"column:type1"`
	//商品类型1
	Value1 string `json:"value1" gorm:"column:value1"`
	//图片id
	ImageID string `json:"image_id" gorm:"column:image_id"`
	//图片url
	Url string `json:"url" gorm:"column:url"`
	//商品种类2
	OptionID2 string `json:"option2_id" gorm:"column:option2_id"`
	//商品选项2
	Type2 string `json:"type2" gorm:"column:type2"`
	//商品类型2
	Value2 string `json:"value2" gorm:"column:value2"`
}
