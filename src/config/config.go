package config

import (
	"budget4home/src/label"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PrepareDb() *gorm.DB {
	if len(os.Getenv("DATABASE")) == 0 {
		// set default connection string
		os.Setenv("DATABASE", "host=localhost port=5432 database=postgres user=postgres password=postgres sslmode=disable")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE")), &gorm.Config{})
	if err != nil {
		fmt.Println("Error starting db")
		return nil
	}

	db.AutoMigrate(&label.Label{})
	// TODO add expense db.AutoMigrate(&expense.Expense{})
	// TODO add group db.AutoMigrate(&group.Group{})
	return db
}
