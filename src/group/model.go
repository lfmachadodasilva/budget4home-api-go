package group

type Group struct {
	ID uint `gorm:"primary_key"`
	// gorm.Model
	Name string

	Users []GroupUser `gorm:"foreignKey:GroupId"`
}

type GroupUser struct {
	GroupId uint   `gorm:"primary_key"`
	UserId  string `gorm:"primary_key"`
}
