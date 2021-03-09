package group

import (
	"gorm.io/gorm"
)

// IGroupRepository group repository interface
type IGroupRepository interface {
	Fetch() (res []Group, err error)
	Add(toAdd *Group) error
	AddUserGroup(toAdd []UserGroup) error
}

// Repository group repository
type Repository struct {
	db *gorm.DB
}

// NewRepository constructor
func NewRepository(db *gorm.DB) IGroupRepository {
	return &Repository{
		db: db,
	}
}

// Fetch get all groups
func (r *Repository) Fetch() ([]Group, error) {
	var groups []Group
	r.db.Preload("Users").Find(&groups)
	return groups, nil
}

// Add group
func (r *Repository) Add(toAdd *Group) error {
	r.db.Create(toAdd)
	return nil
}

// AddUserGroup group
func (r *Repository) AddUserGroup(toAdd []UserGroup) error {
	r.db.Create(toAdd)
	return nil
}
