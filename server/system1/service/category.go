package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"system/data"
	"system/model"
)

type CategorySrv interface {
	List(req *model.ListPage) (Categories []*model.CategoryResult, err error)
	GetTotal(req *model.ListPage) (total int64, err error)
	Get(id string) ([]*model.CategoryResult, error)
	Exist(Category model.Category) *model.Category
	ExistByCategoryID(id string) *model.Category
	Add(Category model.CategoryResult) (bool, error)
	Edit(Category model.Category) (bool, error)
	Delete(Category model.Category) (bool, error)
}

type CategoryService struct {
	Repo data.CategoryRepoInterface
}

func (srv *CategoryService) List(req *model.ListPage) (categories []*model.CategoryResult, err error) {
	return srv.Repo.List(req)
}
func (srv *CategoryService) GetTotal(req *model.ListPage) (total int64, err error) {
	return srv.Repo.GetTotal(req)
}
func (srv *CategoryService) Get(id string) ([]*model.CategoryResult, error) {
	return srv.Repo.Get(id)
}
func (srv *CategoryService) Exist(category model.Category) *model.Category {
	return srv.Repo.Exist(category)
}
func (srv *CategoryService) ExistByCategoryID(id string) *model.Category {
	return srv.Repo.ExistByCategoryID(id)
}

func (srv *CategoryService) Add(category model.CategoryResult) (bool, error) {

	var c1CategoryId string
	var c2CategoryId string
	if category.C1CategoryID == "" {
		c1CategoryId = uuid.NewV4().String()
		category.C1CategoryID = c1CategoryId
	}
	if category.C2CategoryID == "" {
		c2CategoryId = uuid.NewV4().String()
		category.C2CategoryID = c2CategoryId
		category.C2ParentId = c1CategoryId
	}
	if category.C3CategoryID == "" {
		category.C3CategoryID = uuid.NewV4().String()
		category.C3ParentId = c2CategoryId
	}
	//判断3个category是否都存在，就重复，有任何一个不重复，都可以添加
	c1 := model.Category{
		CategoryId: category.C1CategoryID,
		Name:       category.C1Name,
		Desc:       category.C1Desc,
		Order:      category.C1Order,
		ParentId:   "0",
		IsDeleted:  false,
	}
	r1 := srv.Exist(c1)

	c2 := model.Category{
		CategoryId: category.C2CategoryID,
		Name:       category.C2Name,
		Desc:       "",
		Order:      category.C2Order,
		ParentId:   category.C2ParentId,
		IsDeleted:  false,
	}
	r2 := srv.Exist(c2)

	c3 := model.Category{
		CategoryId: category.C3CategoryID,
		Name:       category.C3Name,
		Desc:       "",
		Order:      category.C3Order,
		ParentId:   category.C3ParentId,
		IsDeleted:  false,
	}
	r3 := srv.Exist(c3)

	if r1.Name != "" && r2.Name != "" && r3.Name != "" {
		return false, errors.New("分类已存在")
	}

	if r1.Name == "" {
		srv.Repo.Add(c1)
	}

	if r2.Name == "" {
		srv.Repo.Add(c2)
	}

	if r3.Name == "" {
		srv.Repo.Add(c3)
	}

	return true, nil
}
func (srv *CategoryService) Edit(category model.Category) (bool, error) {
	return srv.Repo.Edit(category)
}
func (srv *CategoryService) Delete(Category model.Category) (bool, error) {
	if Category.CategoryId == "" {
		return false, errors.New("参数错误")
	}
	category := srv.ExistByCategoryID(Category.CategoryId)
	category.IsDeleted = !category.IsDeleted
	return srv.Repo.Delete(*category)
}
