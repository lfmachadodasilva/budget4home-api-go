package expense

import (
	"gorm.io/gorm"
)

type IExpenseRepository interface {
	Fetch() (res []Expense, err error)
}

type ExpenseRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IExpenseRepository {
	return &ExpenseRepository{
		db: db,
	}
}

func (this *ExpenseRepository) Fetch() ([]Expense, error) {
	var expenses []Expense
	this.db.Joins("Group").Joins("Label").Find(&expenses)
	return expenses, nil
}
