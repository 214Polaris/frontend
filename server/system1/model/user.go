package model

// 用户
type User struct {
	//用户账号
	UserID int `json:"userID" gorm:"column:id"`
	//电话
	Mobile string `json:"mobile" gorm:"column:mobile"`
	//密码
	Password string `json:"password" gorm:"column:password"`
	//地址
	Address string `json:"address" gorm:"column:address"`
	//用户名
	Username string `json:"userName" gorm:"column:name"`
	//邮箱
	UserEmail string `json:"userEmail" gorm:"column:email"`
	//出生日期
	UserBornDate string `json:"userBornDate" gorm:"column:bornDate"`
	//姓名
	UserRealName string `json:"userRealName" gorm:"column:realName"`
}
