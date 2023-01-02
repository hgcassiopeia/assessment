package repo

import (
	"database/sql"
	"fmt"

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
		return nil, fmt.Errorf("can't Scan row into variables : %v", err.Error())
	}

	return expense, nil
}

func (r *RepoImpl) GetExpense(id string) (*entities.Expenses, error) {
	stmt, err := r.DB.Prepare("SELECT * FROM expenses WHERE id=$1")
	if err != nil {
		return nil, fmt.Errorf("can't prepare query one row statment : %v", err.Error())
	}

	row := stmt.QueryRow(id)
	var result entities.Expenses

	err = row.Scan(&result.Id, &result.Title, &result.Amount, &result.Note, pq.Array(&result.Tags))
	if err != nil {
		return nil, fmt.Errorf("can't Scan row into variables : %v", err.Error())
	}

	return &result, nil
}
