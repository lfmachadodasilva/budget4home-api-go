package expense

import (
	"budget4home/src/group"
	"budget4home/src/label"
)

// Expense table
type Expense struct {
	ID uint `gorm:"primary_key"`
	// gorm.Model
	Name string

	LabelID uint
	GroupID uint

	Label label.Label `gorm:"foreignkey:LabelID"`
	Group group.Group `gorm:"foreignkey:GroupID"`
}
