package expense

type IExpenseService interface {
	Fetch() (res []Expense, err error)
}

type ExpenseService struct {
	repo IExpenseRepository
}

func NewService(r IExpenseRepository) IExpenseService {
	return &ExpenseService{
		repo: r,
	}
}

func (this *ExpenseService) Fetch() (res []Expense, err error) {
	res, err = this.repo.Fetch()
	if err != nil {
		return nil, err
	}
	return
}
