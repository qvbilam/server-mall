package Business

import (
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
