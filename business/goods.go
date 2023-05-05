package business

import (
	"mall/global"
	"mall/model"
)

type GoodsBusiness struct {
	ID         int64
	IDs        []int64
	CategoryID int64
	OnSale     *bool
	Page       int64
	PerPage    int64
}

func (b *GoodsBusiness) GoodsList() ([]*model.Goods, int64) {
	query := global.DB.Model(&model.Goods{})

	var goods []*model.Goods
	condition := model.Goods{}
	cb := CategoryBusiness{}
	var categoryIds []interface{}

	// 分类查询
	if b.CategoryID != 0 {
		cb.ID = b.CategoryID
		categoryIds, _ = cb.GetMultistageCategoryIds()

		query = query.Where("category_id in ?", categoryIds)
	}

	// ids 查询
	if b.IDs != nil {
		query = query.Where("id in ?", b.IDs)
	}

	// 上架状态查询
	if b.OnSale != nil {
		condition.OnSale = *b.OnSale
	}

	// 分页查询
	if b.Page != 0 {
		query = query.Scopes(model.Paginate(int(b.Page), int(b.PerPage)))
	}

	// 结果
	query = query.Where(&condition)
	query.Preload("Products.Product").Find(&goods)

	var count int64
	query.Count(&count)
	return goods, count
}
