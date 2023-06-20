package model

// 分类
type Category struct {
	//种类id
	CategoryId string `json:"catagoryId" gorm:"column:catagory_id"`
	//种类名
	Name string `json:"name" gorm:"column:name"`
	//位置
	Desc string `json:"desc" gorm:"column:desc"`
	//种类序号
	Order int `json:"order" gorm:"column:order"`
	//父种类（分级）
	ParentId string `json:"parentId" gorm:"column:parent_id"`
	//注销
	IsDeleted bool `json:"isDeleted" gorm:"column:is_deleted"`
}

// 三级分类 一级设为0
type CategoryResult struct {
	C1CategoryID string `gorm:"c1_category_id"`
	C1Name       string `gorm:"column:c1_name"`
	C1Desc       string `gorm:"column:c1_desc"`
	C1Order      int    `gorm:"column:c1_order"`
	C1ParentId   string `gorm:"column:c1_parent_id"`

	C2CategoryID string `gorm:"c2_category_id"`
	C2Name       string `gorm:"column:c2_name"`
	C2Order      int    `gorm:"column:c2_order"`
	C2ParentId   string `gorm:"column:c2_parent_id"`

	C3CategoryID string `gorm:"c3_category_id"`
	C3Name       string `gorm:"column:c3_name"`
	C3Order      int    `gorm:"column:c3_order"`
	C3ParentId   string `gorm:"column:c3_parent_id"`
	IsDeleted    bool   `gorm:"column:c3_is_deleted" json:"isDeleted"`
}
