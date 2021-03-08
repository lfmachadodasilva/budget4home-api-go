package expense

import (
	"budget4home/src/group"
	"budget4home/src/label"
)

type Expense struct {
	Id   int `gorm:"primary_key"`
	Name string

	LabelId int
	GroupId int

	Label label.Label `gorm:"foreignkey:LabelId"`
	Group group.Group `gorm:"foreignkey:GroupId"`
}
