package label

import (
	"budget4home/src/group"
)

// Label table
type Label struct {
	ID uint `gorm:"primaryKey"`
	// gorm.Model
	Name string

	GroupID uint
	Group   group.Group `gorm:"foreignKey:GroupID"`
}
