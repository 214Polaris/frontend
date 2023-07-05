package model

// 用于首页推荐 轮播图
type Banner struct {
	//推荐位id
	ProductId string `json:"productId" gorm:"column:product_id"`
	//图片url
	ProductLink string `json:"productLink" gorm:"column:product_link"`
}
