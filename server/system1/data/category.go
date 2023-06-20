package data

import (
	"errors"
	"gorm.io/gorm"
	"system/model"
)

type CategoryData struct {
	DB *gorm.DB
}

type CategoryRepoInterface interface {
	List(req *model.ListPage) (Categorys []*model.CategoryResult, err error)
	GetTotal(req *model.ListPage) (total int64, err error)
	Get(id string) ([]*model.CategoryResult, error)
	Exist(Category model.Category) *model.Category
	ExistByCategoryID(id string) *model.Category
	Add(Category model.Category) (*model.Category, error)
	Edit(Category model.Category) (bool, error)
	Delete(c model.Category) (bool, error)
}

func (c *CategoryData) List(req *model.ListPage) (categories []*model.CategoryResult, err error) {
	var list []*model.CategoryResult
	err = c.DB.Raw("SELECT c1.category_id as c1_category_id," +
		"c1.name as c1_name,c1.desc as c1_desc," +
		"c1.order as c1_order," +
		"c1.parent_id as c1_parent_id, " +
		"c2.category_id as c2_category_id," +
		"c2.name as c2_name," +
		"c2.order as c2_order," +
		"c2.parent_id as c2_parent_id," +
		"c3.category_id as c3_category_id," +
		"c3.name as c3_name," +
		"c3.order as c3_order," +
		"c3.parent_id as c3_parent_id," +
		"c3.is_deleted as c3_is_deleted " +
		"FROM category c1 join category c2 on c1.category_id = c2.parent_id " +
		"join category c3 on c2.category_id=c3.parent_id").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *CategoryData) GetTotal(req *model.ListPage) (total int64, err error) {
	err = c.DB.Raw("SELECT count(c3.category_id) FROM" +
		"category c1 join category c2 on c1.category_id = c2.parent_id " +
		"join category c3 on c2.category_id=c3.parent_id").Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (c *CategoryData) Get(id string) ([]*model.CategoryResult, error) {
	var list []*model.CategoryResult
	err := c.DB.Raw("SELECT c1.category_id as c1_category_id,c1.name as c1_name,"+
		"c1.desc as c1_desc,"+
		"c1.order as c1_order,"+
		"c2.category_id as c2_category_id,"+
		"c2.name as c2_name,"+
		"c2.order as c2_order,"+
		"c3.category_id as c3_category_id,"+
		"c3.name as c3_name,"+
		"c3.order as c3_order "+
		"FROM category c1 join category c2 on c1.category_id = c2.parent_id "+
		"join category c3 on c2.category_id=c3.parent_id "+
		"where c3.category_id = ?", id).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *CategoryData) Exist(Category model.Category) *model.Category {
	var category model.Category
	if category.Name != "" {
		c.DB.Model(&category).Where("name= ?", Category.Name)
		return &category
	}
	return nil
}

func (c *CategoryData) ExistByCategoryID(id string) *model.Category {
	var category model.Category
	c.DB.Where("category_id = ?", id).First(&category)
	return &category

}

func (c *CategoryData) Add(category model.Category) (*model.Category, error) {
	err := c.DB.Create(category).Error
	if err != nil {
		return nil, errors.New("分类添加失败")
	}
	return &category, nil
}

func (c *CategoryData) Edit(category model.Category) (bool, error) {
	if category.CategoryId == "" {
		return false, errors.New("请传入更新数据")
	}
	temp := &model.Category{CategoryId: category.CategoryId}
	err := c.DB.Model(temp).Where("category_id=?", category.CategoryId).Updates(map[string]interface{}{
		"name":      category.Name,
		"order":     category.Order,
		"parent_id": category.ParentId,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *CategoryData) Delete(category model.Category) (bool, error) {
	err := c.DB.Model(&category).
		Where("category_id=?", category.CategoryId).
		Update("is_deleted", category.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
