package handler

import (
	"net/http"

	"github.com/hgcassiopeia/assessment/expenses"
	"github.com/hgcassiopeia/assessment/expenses/entities"
	"github.com/labstack/echo/v4"
)

type Error struct {
	Message string `json:"message"`
}

type HttpHandler struct {
	UseCase expenses.UseCase
}

func (h *HttpHandler) AddNewExpense(c echo.Context) error {
	var expenses entities.Expenses
	err := c.Bind(&expenses)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}

	h.UseCase.CreateExpense(&expenses)

	return c.JSON(http.StatusCreated, expenses)
}
