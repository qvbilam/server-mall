package business

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mall/global"
	"mall/model"
)

type CategoryBusiness struct {
	ID   int64
	PID  *int64
	Name string
}

func (b *CategoryBusiness) Get() []*model.Category {
	var cs []*model.Category
	condition := &model.Category{
		IsOpen: true,
	}
	if b.PID != nil {
		condition.PID = *b.PID
	}

	global.DB.Model(&model.Category{}).Where(condition).Order("sort desc").Find(&cs)
	return cs
}

func (b *CategoryBusiness) GetAllCategoryIds() []interface{} {
	cs := b.Get()
	var categoryIds []interface{}
	for _, c := range cs {
		categoryIds = append(categoryIds, c.ID)
	}
	return categoryIds
}

// GetMultistageCategoryIds 获取多级分类
func (b *CategoryBusiness) GetMultistageCategoryIds() ([]interface{}, error) {
	entity := model.Category{}
	if res := global.DB.First(&entity, b.ID); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "分类不存在")
	}

	var subQuery string
	if entity.Level == 1 {
		subQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_id IN (SELECT id FROM category WHERE pid = %d)", b.ID)
	} else if entity.Level == 2 { // 二级分类
		subQuery = fmt.Sprintf("SELECT id FROM category WHERE pid = %d", b.ID)
	} else { // 三级分类
		subQuery = fmt.Sprintf("SELECT id FROM category WHERE id = %d", b.ID)
	}

	var categoryIds []interface{}
	var categories []model.Category
	if res := global.DB.Model(model.Category{}).Raw(subQuery).Scan(&categories); res.RowsAffected == 0 {
		return nil, nil
	}
	for _, c := range categories {
		categoryIds = append(categoryIds, c.ID)
	}

	return categoryIds, nil
}
