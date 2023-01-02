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

	given := &entities.Expenses{
		Title:  "Isakaya Bangna",
		Amount: 899,
		Note:   "central bangna",
		Tags:   []string{"food", "beverage"},
	}

	expected := entities.Expenses{
		Id:     1,
		Title:  "Isakaya Bangna",
		Amount: 899,
		Note:   "central bangna",
		Tags:   []string{"food", "beverage"},
	}

	mock.ExpectQuery("INSERT INTO expenses").
		WithArgs(given.Title, given.Amount, given.Note, pq.Array(given.Tags)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expected.Id))

	// Act
	result, err := repo.CreateExpense(given)

	// Assert
	if assert.NoError(t, err) {
		assert.Equal(t, expected.Id, result.Id)
		assert.Equal(t, expected.Title, result.Title)
		assert.Equal(t, expected.Amount, result.Amount)
		assert.Equal(t, expected.Note, result.Note)
		assert.Equal(t, expected.Tags, result.Tags)
	}
}
