package expenseso

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func AddNewExpenseHandler(c echo.Context) error {
	var expenses Expenses
	err := c.Bind(&expenses)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) VALUES ($1, $2, $3, $4) RETURNING id", expenses.Title, expenses.Amount, expenses.Note, pq.Array(&expenses.Tags))
	err = row.Scan(&expenses.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, expenses)
}
