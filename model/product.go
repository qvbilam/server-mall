package model

// Product 产品
type Product struct {
	IDModel
	Name      string `gorm:"type:varchar(30) NOT NULL DEFAULT '';comment:名称"`
	Icon      string `gorm:"type:varchar(30) NOT NULL DEFAULT '';comment:图标"`
	Introduce string `gorm:"type:varchar(30) NOT NULL DEFAULT '';comment:简介"`
	Type      string `gorm:"type:varchar(30) NOT NULL DEFAULT '';comment:类型"`
	Tag       string `gorm:"type:varchar(30) NOT NULL DEFAULT '';comment:标签"`
	Price     int64  `gorm:"type:int(30) NOT NULL DEFAULT 0;comment:价值(显示用) 1分/1金币"`
	DateModel
}
