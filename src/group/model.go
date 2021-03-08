package group

type Group struct {
	Id   int `gorm:"primary_key"`
	Name string

	Users []GroupUser `gorm:"foreignKey:GroupId;references:id"`
}

type GroupUser struct {
	GroupId int    `gorm:"primary_key"`
	UserId  string `gorm:"primary_key"`
}
