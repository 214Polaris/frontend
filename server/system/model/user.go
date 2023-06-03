package model

// 用户
type User struct {
	//用户名
	Username string `json:"username" gorm:"column:username"`
	//电话
	Mobile string `json:"mobile" gorm:"column:mobile"`
	//密码
	Password string `json:"password" gorm:"column:password"`
	//地址
	Address string `json:"address" gorm:"column:address"`
}
