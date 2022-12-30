package expenses

import "github.com/hgcassiopeia/assessment/expenses/entities"

type Repository interface {
	CreateExpense(expense *entities.Expenses) error
}
