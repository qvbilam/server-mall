package model

type GoodsProduct struct {
	IDModel
	GoodsID   int64    `gorm:"type:int(30) NOT NULL DEFAULT 0;comment:商品id;index"`
	ProductID int64    `gorm:"type:int(30) NOT NULL DEFAULT 0;comment:产品id"`
	Count     int64    `gorm:"type:int(30) NOT NULL DEFAULT 0;comment:数量"`
	Product   *Product `gorm:"foreignKey:id;references:product_id"`
	DateModel
}
