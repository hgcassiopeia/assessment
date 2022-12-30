package repo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hgcassiopeia/assessment/expenses/entities"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := InitRepository(db)

	expense := &entities.Expenses{
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}
	expectedId := 1

	mock.ExpectQuery("INSERT INTO expenses").
		WithArgs(expense.Title, expense.Amount, expense.Note, pq.Array(expense.Tags)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))

	// Act
	err = repo.CreateExpense(expense)

	// Assert
	if assert.NoError(t, err) {
		assert.Equal(t, expectedId, expense.Id)
	}
}
