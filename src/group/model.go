package group

// Group group db table
type Group struct {
	ID   uint `gorm:"primary_key"`
	Name string

	Users []UserGroup `gorm:"foreignKey:GroupID"`
}

// UserGroup group user many-to-many table
type UserGroup struct {
	GroupID uint   `gorm:"primary_key"`
	UserID  string `gorm:"primary_key"`
}
