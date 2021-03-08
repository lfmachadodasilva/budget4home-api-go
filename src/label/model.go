package label

import (
	"budget4home/src/group"
)

type Label struct {
	Id   int `gorm:"primaryKey"`
	Name string

	GroupId int
	Group   group.Group `gorm:"foreignKey:GroupId"`
}
