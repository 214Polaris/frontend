package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"system/data"
	"system/service"
)

var (
	db          *gorm.DB = nil
	BannerSrv   service.BannerService
	CategorySrv service.CategoryService
	OrderSrv    service.OrderSrv
	ProductSrc  service.ProductService
	UserSrv     service.UserService
	CartSrv     service.CartService
)

// 数据库配置
const (
	userName = "root"
	//password = "Hzm13602985871"
	//password = "chen8574jun"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "shoppingmall"
)

func InitDB() *gorm.DB {
	// 如果已经建立数据库，直接返回
	if db != nil {
		return db
	}
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	//验证连接
	if err != nil {
		fmt.Println("open database fail")
		return DB
	}
	sqlDB, _ := DB.DB()
	//设置数据库最大连接数
	sqlDB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	sqlDB.SetMaxIdleConns(10)
	fmt.Println("connect success")
	db = DB
	return db
}

func initService() {
	BannerSrv = service.BannerService{
		Repo: &data.BannerData{
			DB: db,
		},
	}

	ProductSrc = service.ProductService{
		Repo: &data.ProductData{
			DB: db,
		},
	}

	CategorySrv = service.CategoryService{
		Repo: &data.CategoryData{
			DB: db,
		},
	}

	OrderSrv = &service.OrderService{
		Repo: &data.OrderData{
			DB: db,
		},
	}

	UserSrv = service.UserService{
		Repo: &data.UserData{
			DB: db,
		},
	}

	CartSrv = service.CartService{
		Repo: &data.CartData{
			DB: db,
		},
	}
}

func init() {
	InitDB()      // 配置数据库
	initService() // 配置service
}
