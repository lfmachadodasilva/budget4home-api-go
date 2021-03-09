package group

import (
	"gorm.io/gorm"
)

type IGroupRepository interface {
	Fetch() (res []Group, err error)
}

type GroupRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IGroupRepository {
	return &GroupRepository{
		db: db,
	}
}

func (this *GroupRepository) Fetch() ([]Group, error) {
	var groups []Group
	this.db.Preload("Users").Find(&groups)
	return groups, nil
}
