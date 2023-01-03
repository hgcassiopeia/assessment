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

	result, err := h.UseCase.CreateExpense(&expenses)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, &result)
}

func (h *HttpHandler) GetExpenseDetail(c echo.Context) error {
	id := c.Param("id")

	result, err := h.UseCase.GetExpense(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) UpdateExpense(c echo.Context) error {
	id := c.Param("id")
	var expenses entities.Expenses
	err := c.Bind(&expenses)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}

	result, err := h.UseCase.UpdateExpense(id, &expenses)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
