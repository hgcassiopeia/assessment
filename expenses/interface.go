package expenses

import "github.com/hgcassiopeia/assessment/expenses/entities"

type Repository interface {
	CreateExpense(expense *entities.Expenses) (*entities.Expenses, error)
	GetExpense(id string) (*entities.Expenses, error)
}

type UseCase interface {
	CreateExpense(expense *entities.Expenses) (*entities.Expenses, error)
	GetExpense(id string) (*entities.Expenses, error)
}
