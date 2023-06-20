package model

// 用于首页推荐 轮播图
type Banner struct {
	//推荐位id
	BannerId string `json:"bannerId" gorm:"column:banner_id"`
	//图片url
	Url string `json:"url" gorm:"column:url"`
	//商家url
	RedirectUrl string `json:"redirectUrl" gorm:"column:redirect_url"`
	//排序 序号
	Order int `json:"order" gorm:"column:order"`
	//创建者
	CreaateUser string `json:"creaateUser" gorm:"column:creaate_user"`
	//更新者
	UpdataUser string `json:"updataUser" gorm:"column:updata_user"`
}
