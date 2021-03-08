package label

import (
	"gorm.io/gorm"
)

type ILabelRepository interface {
	Fetch() (res []Label, err error)
}

type LabelRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) ILabelRepository {
	return &LabelRepository{
		db: db,
	}
}

func (this *LabelRepository) Fetch() ([]Label, error) {
	var labels []Label
	// var groups []group.Group
	this.db.Preload("Group").Find(&labels)
	return labels, nil
}
