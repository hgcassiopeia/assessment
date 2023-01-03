package service

import (
	"github.com/hgcassiopeia/assessment/expenses"
	"github.com/hgcassiopeia/assessment/expenses/entities"
)

type UseCaseImpl struct {
	Repository expenses.Repository
}

func Init(databaseRepo expenses.Repository) expenses.UseCase {
	return &UseCaseImpl{Repository: databaseRepo}
}

func (u *UseCaseImpl) CreateExpense(expense *entities.Expenses) (*entities.Expenses, error) {
	result, err := u.Repository.CreateExpense(expense)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UseCaseImpl) GetExpense(id string) (*entities.Expenses, error) {
	result, err := u.Repository.GetExpense(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UseCaseImpl) UpdateExpense(id string) (*entities.Expenses, error) {
	return nil, nil
}
