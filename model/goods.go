package model

// Goods Commodity 商品
type Goods struct {
	IDModel
	CategoryID    int64           `gorm:"type:varchar(30) NOT NULL DEFAULT '';comment:分类id;index"`
	Name          string          `gorm:"type:varchar(30) NOT NULL DEFAULT '';comment:名称"`
	Introduce     string          `gorm:"type:varchar(30) NOT NULL DEFAULT '';comment:简介"`
	Icon          string          `gorm:"type:varchar(30) NOT NULL DEFAULT '';comment:图标"`
	PayType       string          `gorm:"type:varchar(30) NOT NULL DEFAULT '';comment:支付类型"`
	Price         int64           `gorm:"type:int(30) NOT NULL DEFAULT 0;comment:价格 1分/1金币"`
	OriginalPrice int64           `gorm:"type:int(30) NOT NULL DEFAULT 0;comment:原价 1分/1金币"`
	Stocks        int64           `gorm:"type:int(30) NOT NULL DEFAULT 0;comment:库存"`
	SoldCount     int64           `gorm:"type:int(30) NOT NULL DEFAULT 0;comment:销量"`
	IsHot         bool            `gorm:"type:tinyint(1) NOT NULL DEFAULT 0;comment:是否热销"`
	IsUnlimited   bool            `gorm:"type:tinyint(1) NOT NULL DEFAULT 0;comment:是否无限"`
	OnSale        bool            `gorm:"type:tinyint(1) NOT NULL DEFAULT 0;comment:是否上架"`
	Products      []*GoodsProduct `gorm:"foreignKey:goods_id;references:id"`
	DateModel
	DeletedModel
}
