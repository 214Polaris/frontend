package model

type ListPage struct {
	PageSize int    `json:"pageSize"`
	Page     int    `json:"page"`
	Where    string `json:"where"`
	Sort     string `json:"sort"`
}
