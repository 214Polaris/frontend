package data

import (
	"gorm.io/gorm"
	"system/model"
)

type ProductData struct {
	DB *gorm.DB
}

type ProductRepoInterface interface {
	//// 配合gettotal分页
	//List(req *model.ListPage) (Products []*model.Product, err error)
	// 获取所有商品记录
	GetTotal() (pro []*model.Product)
	//// 获取单个记录
	//Get(Product model.Product) (*model.Product, error)
	//// 根据整个product查询某记录是否存在
	//Exist(Product model.Product) *model.Product
	//// 根据product id查询某记录是否存在
	//ExistByProductID(id string) *model.Product
	//// 添加记录
	//Add(Product model.Product) (*model.Product, error)
	//// 编辑记录
	//Edit(Product model.Product) (bool, error)
	//// 软删除
	//Delete(u model.Product) (bool, error)
}

//func (repo *ProductData) List(req *model.ListPage) (products []*model.Product, err error) {
//	//fmt.Println(req)
//	db := repo.DB
//	limit, offset := utils.Page(req.PageSize, req.Page) // 分页
//	if req.Where != "" {
//		db = db.Where(req.Where)
//	}
//	if err := db.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
//		return nil, err
//	}
//	return products, nil
//}

// have been edited
func (repo *ProductData) GetTotal() (pro []*model.Product) {
	var products []*model.Product
	db := repo.DB
	if err := db.Find(&products).Error; err != nil {
		return nil
	}
	return products
}

//func (repo *ProductData) Get(product model.Product) (*model.Product, error) {
//	if err := repo.DB.Where(&product).Find(&product).Error; err != nil {
//		return nil, err
//	}
//	return &product, nil
//}
//
//func (repo *ProductData) Exist(product model.Product) *model.Product {
//	if product.ProductName != "" {
//		var temp model.Product
//		repo.DB.Where("product_name= ?", product.ProductName).First(&temp)
//		return &temp
//	}
//	return nil
//}
//
//func (repo *ProductData) ExistByProductID(id string) *model.Product {
//	var p model.Product
//	repo.DB.Where("product_id = ?", id).First(&p)
//	return &p
//}
//
//func (repo *ProductData) Add(product model.Product) (*model.Product, error) {
//	exist := repo.Exist(product)
//	if exist != nil && exist.ProductName != "" {
//		return &product, fmt.Errorf("商品已存在")
//	}
//	err := repo.DB.Create(product).Error
//	if err != nil {
//		return nil, fmt.Errorf("商品添加失败")
//	}
//	return &product, nil
//}

//func (repo *ProductData) Edit(product model.Product) (bool, error) {
//	if product.ProductId == "" {
//		return false, fmt.Errorf("请传入更新 ID")
//	}
//	p := &model.Product{}
//	err := repo.DB.Model(p).Where("product_id=?", product.ProductId).Updates(map[string]interface{}{
//		"product_name": product.ProductName, "product_intro": product.ProductIntro,
//		"category_id": product.ProductCatagory, "product_cover_img": product.ProductCoverImg,
//		"product_banner": product.ProductBanner, "original_price": product.OriginalPrice,
//		"selling_price": product.SellingPrice, "stock_num": product.StockNum,
//		"tag": product.Tag, "sell_status": product.SellStatus,
//		"product_detail_content": product.ProductDetailContent,
//	}).Error
//	if err != nil {
//		return false, err
//	}
//	return true, nil
//}

//func (repo *ProductData) Delete(product model.Product) (bool, error) {
//
//	//err := repo.DB.Model(&product).Where("product_id = ?", product.ProductId).Update("is_deleted", product.IsDeleted).Error
//	err := repo.DB.Delete(&product).Error
//	if err != nil {
//		return false, err
//	}
//	return true, nil
//}
