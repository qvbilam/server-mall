package model

// Category 商品分类
type Category struct {
	IDModel
	PID    int64  `gorm:"type:int not null default 0;comment:父类id;index"`
	Name   string `gorm:"type:varchar(30) NOT NULL DEFAULT '';comment:名称"`
	Sort   int64  `gorm:"type:int not null default 0;comment:排序(大到小);index"`
	IsOpen bool   `gorm:"type:tinyint(1) NOT NULL DEFAULT 0;comment:是否开启"`
	DateModel
}
