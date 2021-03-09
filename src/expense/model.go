package expense

import (
	"budget4home/src/group"
	"budget4home/src/label"
)

type Expense struct {
	ID uint `gorm:"primary_key"`
	// gorm.Model
	Name string

	LabelId uint
	GroupId uint

	Label label.Label `gorm:"foreignkey:LabelId"`
	Group group.Group `gorm:"foreignkey:GroupId"`
}
