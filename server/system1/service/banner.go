package service

import (
	"github.com/gin-gonic/gin"
	"system/data"
)

type BannerSrv interface {
	//List(req *model.ListPage) (Banners []*model.Banner, err error)
	//GetTotal(req *model.ListPage) (total int, err error)
	//Get(Banner model.Banner) (*model.Banner, error)
	//Exist(Banner model.Banner) *model.Banner
	//ExistByBannerID(id string) *model.Banner
	//Add(Banner model.Banner) (*model.Banner, error)
	//Edit(Banner model.Banner) (bool, error)
	//Delete(id string) (bool, error)
	// 获取所有信息
	GetAll(ctx *gin.Context)
}

type BannerService struct {
	Repo data.BannerRepoInterface
}

//func (srv *BannerService) List(req *model.ListPage) (banners []*model.Banner, err error) {
//	return srv.Repo.List(req)
//}
//
//func (srv *BannerService) GetTotal(req *model.ListPage) (total int64, err error) {
//	return srv.Repo.GetTotal(req)
//}
//
//func (srv *BannerService) Get(banner model.Banner) (*model.Banner, error) {
//	return srv.Repo.Get(banner)
//}
//
//func (srv *BannerService) Exist(banner model.Banner) *model.Banner {
//	return srv.Repo.Exist(banner)
//}
//
//func (srv *BannerService) ExistByBannerID(id string) *model.Banner {
//	return srv.Repo.ExistByBannerID(id)
//}
//
//func (srv *BannerService) Add(banner model.Banner) (*model.Banner, error) {
//	if banner.BannerId == "" {
//		banner.BannerId = uuid.NewV4().String()
//	}
//	return srv.Repo.Add(banner)
//}
//
//func (srv *BannerService) Edit(banner model.Banner) (bool, error) {
//	return srv.Repo.Edit(banner)
//}
//
//func (srv *BannerService) Delete(id string) (bool, error) {
//	return srv.Repo.Delete(id)
//}

// 获取所有信息
func (srv *BannerService) GetAll(ctx *gin.Context) {
	banners := srv.Repo.GetAll()
	if banners == nil {
		ctx.JSON(401, gin.H{"meg": "查询失败"})
		return
	}
	ctx.JSON(200, banners)
}
