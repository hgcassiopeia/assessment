package repo

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hgcassiopeia/assessment/expenses/entities"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	db, mock := newSqlMock(t)
	defer db.Close()
	repo := InitRepository(db)

	statement := "INSERT INTO expenses"

	t.Run("Success - TestCreateExpense", func(t *testing.T) {
		// Arrange
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

		mock.ExpectQuery(statement).
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
	})

	t.Run("Fail - TestCreateExpense", func(t *testing.T) {
		// Arrange
		given := &entities.Expenses{
			Title:  "Isakaya Bangna",
			Amount: 899,
			Note:   "central bangna",
			Tags:   []string{"food", "beverage"},
		}

		mockErr := fmt.Errorf("something went wrong")
		expected := fmt.Errorf("can't Scan row into variables : something went wrong")
		mock.ExpectQuery(statement).
			WithArgs(given.Title, given.Amount, given.Note, pq.Array(given.Tags)).
			WillReturnError(mockErr)

		// Act
		_, err := repo.CreateExpense(given)

		// Assert
		if err != nil {
			assert.Equal(t, expected, err)
		}
	})
}

func TestGetExpense(t *testing.T) {
	db, mock := newSqlMock(t)
	defer db.Close()
	repo := InitRepository(db)

	statement := "SELECT (.+) FROM expenses WHERE id=(.+)"

	t.Run("Success - TestGetExpense", func(t *testing.T) {
		// Arrange
		given := "1"

		columns := []string{"id", "title", "amount", "note", "tags"}
		expected := entities.Expenses{
			Id:     1,
			Title:  "Isakaya Bangna",
			Amount: 899,
			Note:   "central bangna",
			Tags:   []string{"food", "beverage"},
		}
		expectedRow := sqlmock.NewRows(columns).AddRow(expected.Id, expected.Title, expected.Amount, expected.Note, pq.Array(expected.Tags))

		mock.ExpectPrepare(statement).ExpectQuery().WithArgs(given).WillReturnRows(expectedRow)

		// Act
		result, err := repo.GetExpense(given)

		// Assert
		if assert.NoError(t, err) {
			assert.Equal(t, expected.Id, result.Id)
			assert.Equal(t, expected.Title, result.Title)
			assert.Equal(t, expected.Amount, result.Amount)
			assert.Equal(t, expected.Note, result.Note)
			assert.Equal(t, expected.Tags, result.Tags)
		}
	})

	t.Run("Fail - TestGetExpense prepare query failed", func(t *testing.T) {
		// Arrange
		given := "1"

		mockErr := fmt.Errorf("something went wrong")
		expected := fmt.Errorf("can't prepare statment : something went wrong")

		mock.ExpectPrepare(statement).WillReturnError(mockErr)

		// Act
		_, err := repo.GetExpense(given)

		// Assert
		if err != nil {
			assert.Equal(t, expected, err)
		}
	})

	t.Run("Fail - TestGetExpense scan row into variable failed", func(t *testing.T) {
		// Arrange
		given := "1"

		mockErr := fmt.Errorf("something went wrong")
		expected := fmt.Errorf("can't Scan row into variables : something went wrong")

		mock.ExpectPrepare(statement).ExpectQuery().WithArgs(given).WillReturnError(mockErr)

		// Act
		_, err := repo.GetExpense(given)

		// Assert
		if err != nil {
			assert.Equal(t, expected, err)
		}
	})
}

func TestUpdateExpense(t *testing.T) {
	db, mock := newSqlMock(t)
	defer db.Close()
	repo := InitRepository(db)

	statement := "UPDATE expenses SET (.+) WHERE id=(.+) RETURNING (.+)"

	t.Run("Success - TestUpdateExpense", func(t *testing.T) {
		// Arrange
		id := "1"
		newExpense := entities.Expenses{
			Id:     1,
			Title:  "Isakaya Bangna",
			Amount: 990,
			Note:   "Central bangna near by BigC",
			Tags:   []string{"food", "beverage"},
		}

		columns := []string{"id", "title", "amount", "note", "tags"}
		expected := entities.Expenses{
			Id:     1,
			Title:  "Isakaya Bangna",
			Amount: 990,
			Note:   "Central bangna near by BigC",
			Tags:   []string{"food", "beverage"},
		}
		expectedRow := sqlmock.NewRows(columns).AddRow(expected.Id, expected.Title, expected.Amount, expected.Note, pq.Array(expected.Tags))
		mock.ExpectPrepare(statement).ExpectQuery().WithArgs(id, newExpense.Title, newExpense.Amount, newExpense.Note, pq.Array(newExpense.Tags)).WillReturnRows(expectedRow)

		// Act
		result, err := repo.UpdateExpense(id, &expected)

		// Assert
		if assert.NoError(t, err) {
			assert.Equal(t, expected.Id, result.Id)
			assert.Equal(t, expected.Title, result.Title)
			assert.Equal(t, expected.Amount, result.Amount)
			assert.Equal(t, expected.Note, result.Note)
			assert.Equal(t, expected.Tags, result.Tags)
		}
	})

	t.Run("Fail - TestUpdateExpense prepare query failed", func(t *testing.T) {
		// Arrange
		id := "1"
		newExpense := entities.Expenses{}

		mockErr := fmt.Errorf("something went wrong")
		expected := fmt.Errorf("can't prepare statment : something went wrong")

		mock.ExpectPrepare(statement).WillReturnError(mockErr)

		// Act
		_, err := repo.UpdateExpense(id, &newExpense)

		// Assert
		if err != nil {
			assert.Equal(t, expected, err)
		}
	})

	t.Run("Fail - TestUpdateExpense scan row into variable failed", func(t *testing.T) {
		// Arrange
		id := "1"
		newExpense := entities.Expenses{
			Id:     1,
			Title:  "Isakaya Bangna",
			Amount: 990,
			Note:   "Central bangna near by BigC",
			Tags:   []string{"food", "beverage"},
		}

		mockErr := fmt.Errorf("something went wrong")
		expected := fmt.Errorf("can't Scan row into variables : something went wrong")

		mock.ExpectPrepare(statement).ExpectQuery().WithArgs(id, newExpense.Title, newExpense.Amount, newExpense.Note, pq.Array(newExpense.Tags)).WillReturnError(mockErr)

		// Act
		_, err := repo.UpdateExpense(id, &newExpense)

		// Assert
		if err != nil {
			assert.Equal(t, expected, err)
		}
	})
}

func newSqlMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
