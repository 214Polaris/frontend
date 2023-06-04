package data

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"system/model"
	"system/utils"
)

type UserData struct {
	DB *gorm.DB
}

func (data *UserData) ExistByuserID(id string) *model.User {
	//TODO implement me
	panic("implement me")
}

type UserRepoInterface interface {
	// MyQuery 查询
	Query(u model.User) bool
	// MyInsert 添加
	Insert(u model.User)
	// MyDelete 删除
	Delete(u model.User)
	// 检查密码
	CheckPassword(u model.User) bool
	// 配合gettotal分页
	List(req *model.ListPage) (users []*model.User, err error)
	// 获取数据库中的所有user记录 给前端分页的时候提供数据
	GetTotal(req *model.ListPage) (total int64, err error)
	// 获取单个user记录
	Get(user model.User) (*model.User, error)
	// 检测是否存在某记录
	Exist(user model.User) *model.User
	// 通过id检测是否存在某记录
	ExistByuserID(id string) *model.User
	// 通过电话检测是否存在某记录
	ExistByMobile(mobile string) *model.User
	// 编辑记录
	Edit(user model.User) (bool, error)
}

// Query 查询
func (data *UserData) Query(u model.User) bool {
	var count int64
	data.DB.Where("username = ?", u.Username).Find(&u).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// Insert 添加
func (data *UserData) Insert(u model.User) {
	if exist := data.Exist(u); exist != nil {
		fmt.Println("添加失败，数据库中已有该账号")
		return
	}
	err := data.DB.Create(&u).Error
	if err != nil {
		fmt.Println("添加失败", err)
		return
	}
	fmt.Println("添加成功")
}

// Delete 删除
func (data *UserData) Delete(u model.User) {
	if data.Query(u) == false {
		fmt.Println("删除失败，数据库中无该账号")
		return
	}
	err := data.DB.Delete(&u).Error
	if err != nil {
		fmt.Println("删除失败，数据库中无该账号")
	}
	fmt.Println("删除成功")
}

// 密码检查
func (data *UserData) CheckPassword(u model.User) bool {
	DbUser, _ := data.Get(u)
	//密码不相同
	//fmt.Println(DbUser.Password, u.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(DbUser.Password), []byte(u.Password)); err != nil {
		return false
	}
	//相同返回true
	return true
}

func (data *UserData) List(req *model.ListPage) (users []*model.User, err error) {
	db := data.DB
	limit, offset := utils.Page(req.PageSize, req.Page) // 分页

	if err := db.Order("id desc").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (data *UserData) GetTotal(req *model.ListPage) (total int64, err error) {
	var users []model.User
	db := data.DB

	if err := db.Find(&users).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (data *UserData) Get(user model.User) (*model.User, error) {
	if err := data.DB.Where("username = ?", user.Username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (data *UserData) Exist(user model.User) *model.User {
	var count int
	data.DB.Find(&user).Where("user_name = ?", user.Username)
	if count > 0 {
		return &user
	}
	return nil
}

func (data *UserData) ExistByMobile(mobile string) *model.User {
	var count int
	var user model.User
	data.DB.Find(&user).Where("mobile = ?", mobile)
	if count > 0 {
		return &user
	}
	return nil
}

func (data *UserData) Edit(user model.User) (bool, error) {
	err := data.DB.Model(&user).Where("username = ?").Updates(map[string]interface{}{"mobile": user.Mobile, "address": user.Address}).Error
	//err := repo.DB.Save(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
