package repo

import (
	"database/sql"

	"github.com/hgcassiopeia/assessment/expenses"
	"github.com/hgcassiopeia/assessment/expenses/entities"
	"github.com/lib/pq"
)

type RepoImpl struct {
	DB *sql.DB
}

func InitRepository(Conn *sql.DB) expenses.Repository {
	return &RepoImpl{Conn}
}

func (r *RepoImpl) CreateExpense(expense *entities.Expenses) (*entities.Expenses, error) {
	row := r.DB.QueryRow("INSERT INTO expenses (title, amount, note, tags) VALUES ($1, $2, $3, $4) RETURNING id", expense.Title, expense.Amount, expense.Note, pq.Array(&expense.Tags))
	err := row.Scan(&expense.Id)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (r *RepoImpl) GetExpense(id string) (entities.Expenses, error) {
	return entities.Expenses{
		Id:     1,
		Title:  "Isakaya Bangna",
		Amount: 899,
		Note:   "central bangna",
		Tags:   []string{"food", "beverage"},
	}, nil
}
