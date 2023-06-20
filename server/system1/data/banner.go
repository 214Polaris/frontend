package data

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"system/model"
	"system/utils"
)

type BannerData struct {
	DB *gorm.DB
}

type BannerRepoInterface interface {
	List(req *model.ListPage) (Banners []*model.Banner, err error)
	GetTotal(req *model.ListPage) (total int64, err error)
	Get(Banner model.Banner) (*model.Banner, error)
	Exist(Banner model.Banner) *model.Banner
	ExistByBannerID(id string) *model.Banner
	Add(Banner model.Banner) (*model.Banner, error)
	Edit(Banner model.Banner) (bool, error)
	Delete(id string) (bool, error)
}

func (b *BannerData) List(req *model.ListPage) (banners []*model.Banner, err error) {
	//fmt.Println(req)
	db := b.DB
	limit, offset := utils.Page(req.PageSize, req.Page) // 分页
	sort := utils.Sort(req.Sort)
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&banners).Error; err != nil {
		return nil, err
	}
	return banners, nil
}

func (b *BannerData) GetTotal(req *model.ListPage) (total int64, err error) {
	var banners []model.Banner
	db := b.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&banners).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (b *BannerData) Get(banner model.Banner) (*model.Banner, error) {
	if err := b.DB.Where(&banner).Find(&banner).Error; err != nil {
		return nil, err
	}
	return &banner, nil
}

func (b *BannerData) Exist(banner model.Banner) *model.Banner {

	if banner.Url != "" && banner.RedirectUrl != "" {
		b.DB.Model(&banner).Where("url= ? and redirect_url", banner.Url, banner.RedirectUrl).First(&b)
		return &banner
	}
	return nil
}

func (b *BannerData) ExistByBannerID(id string) *model.Banner {
	var banner model.Banner
	b.DB.Where("order_id = ?", id).First(&banner)
	return &banner
}

func (b *BannerData) Add(banner model.Banner) (*model.Banner, error) {
	exist := b.Exist(banner)
	if exist != nil && exist.Url == banner.Url && exist.RedirectUrl == banner.RedirectUrl {
		return nil, errors.New("轮播已存在")
	}
	err := b.DB.Create(banner).Error
	if err != nil {
		return nil, errors.New("轮播添加失败")
	}
	return &banner, nil
}

func (b *BannerData) Edit(banner model.Banner) (bool, error) {
	if banner.BannerId == "" {
		return false, fmt.Errorf("请传入参数")
	}
	banner1 := &model.Banner{}
	err := b.DB.Model(banner1).Where("banner_id=?", banner.BannerId).Updates(map[string]interface{}{
		"Url":         banner.Url,
		"RedirectUrl": banner.RedirectUrl,
		"OrderBy":     banner.Order,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (b *BannerData) Delete(id string) (bool, error) {
	err := b.DB.Where("banner_id = ?", id).Delete(&model.Banner{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
