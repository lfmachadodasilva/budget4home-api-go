package label

import (
	"budget4home/src/group"
)

type Label struct {
	ID uint `gorm:"primaryKey"`
	// gorm.Model
	Name string

	GroupId uint
	Group   group.Group `gorm:"foreignKey:GroupId"`
}
