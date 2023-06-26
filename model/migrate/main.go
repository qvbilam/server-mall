package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"mall/enum"
	"mall/model"
)

func main() {
	user := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	database := "qvbilam_mall"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //不带表名
		},
	})
	if err != nil {
		panic(any(err))
	}

	_ = db.AutoMigrate(
		&model.Category{},
		&model.Goods{},
		&model.GoodsProduct{},
		&model.Product{},
	)

	initCategory(db)
	initGoods(db)
	initProduct(db)
	initGoodsProduct(db)
}

// 初始化分类
func initCategory(tx *gorm.DB) {
	categories := []model.Category{
		{
			PID:    0,
			Name:   "金币商城",
			Sort:   0,
			IsOpen: true,
		},
		{
			PID:    0,
			Name:   "月卡",
			Sort:   0,
			IsOpen: true,
		},
	}

	tx.Create(&categories)
}

func initGoods(tx *gorm.DB) {
	payTypeMoney := "money"
	//payTypeCoin := "coin"
	categoryCoin := int64(1)
	categoryMonthCard := int64(2)

	goods := []model.Goods{
		{
			CategoryID:    categoryCoin,
			Name:          "6元金币",
			PayType:       payTypeMoney,
			Price:         6,
			OriginalPrice: 6,
			IsUnlimited:   true,
			OnSale:        true,
		},
		{
			CategoryID:    categoryCoin,
			Name:          "30元金币",
			PayType:       payTypeMoney,
			Price:         30,
			OriginalPrice: 30,
			IsUnlimited:   true,
			OnSale:        true,
		},
		{
			CategoryID:    categoryCoin,
			Name:          "60元金币",
			PayType:       payTypeMoney,
			Price:         60,
			OriginalPrice: 60,
			IsUnlimited:   true,
			OnSale:        true,
		},
		{
			CategoryID:    categoryCoin,
			Name:          "98元金币",
			PayType:       payTypeMoney,
			Price:         98,
			OriginalPrice: 98,
			IsUnlimited:   true,
			OnSale:        true,
		},
		{
			CategoryID:    categoryCoin,
			Name:          "198元金币",
			PayType:       payTypeMoney,
			Price:         198,
			OriginalPrice: 198,
			IsUnlimited:   true,
			OnSale:        true,
		},
		{
			CategoryID:    categoryCoin,
			Name:          "328元金币",
			PayType:       payTypeMoney,
			Price:         328,
			OriginalPrice: 328,
			IsUnlimited:   true,
			OnSale:        true,
		},
		{
			CategoryID:    categoryCoin,
			Name:          "648元金币",
			PayType:       payTypeMoney,
			Price:         648,
			OriginalPrice: 648,
			IsUnlimited:   true,
			OnSale:        true,
		},
		{
			CategoryID:    categoryMonthCard,
			Name:          "普通月卡",
			PayType:       payTypeMoney,
			Price:         30,
			OriginalPrice: 30,
			IsUnlimited:   true,
			OnSale:        true,
		},
		{
			CategoryID:    categoryMonthCard,
			Name:          "超级月卡",
			PayType:       payTypeMoney,
			Price:         98,
			OriginalPrice: 98,
			IsUnlimited:   true,
			OnSale:        true,
		},
	}

	tx.Create(&goods)
}

func initProduct(tx *gorm.DB) {
	product := []model.Product{
		{
			Name:  "金币",
			Type:  enum.ProductTypeCoin,
			Tag:   enum.ProductTypeCoin,
			Price: 1,
		},
		{
			Name:  "月卡",
			Type:  enum.ProductTypeTicket,
			Tag:   enum.ProductTagMonthTicket,
			Price: 1,
		},
		{
			Name:  "超级月卡",
			Type:  enum.ProductTypeTicket,
			Tag:   enum.ProductTagSuperMonthTicket,
			Price: 1,
		},
	}

	tx.Create(&product)
}

func initGoodsProduct(tx *gorm.DB) {
	goodsProduct := []model.GoodsProduct{
		{
			GoodsID:   1,
			ProductID: 1,
			Count:     600,
		},
		{
			GoodsID:   2,
			ProductID: 1,
			Count:     3000,
		},
		{
			GoodsID:   3,
			ProductID: 1,
			Count:     6000,
		},
		{
			GoodsID:   4,
			ProductID: 1,
			Count:     9800,
		},
		{
			GoodsID:   5,
			ProductID: 1,
			Count:     19800,
		},
		{
			GoodsID:   6,
			ProductID: 1,
			Count:     32800,
		},
		{
			GoodsID:   7,
			ProductID: 1,
			Count:     64800,
		},
		{
			GoodsID:   8,
			ProductID: 2,
			Count:     1,
		},
		{
			GoodsID:   8,
			ProductID: 1,
			Count:     3000,
		},
		{
			GoodsID:   9,
			ProductID: 3,
			Count:     1,
		},
		{
			GoodsID:   9,
			ProductID: 1,
			Count:     9800,
		},
	}

	tx.Create(&goodsProduct)
}
