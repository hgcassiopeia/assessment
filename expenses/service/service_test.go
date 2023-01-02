package service

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hgcassiopeia/assessment/expenses/entities"
	mock "github.com/hgcassiopeia/assessment/expenses/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddNewExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)

	t.Run("Success - TestAddNewExpense", func(t *testing.T) {
		// Arrange
		given := &entities.Expenses{
			Id:     0,
			Title:  "Isakaya Bangna",
			Amount: 899,
			Note:   "central bangna",
			Tags:   []string{"food", "beverage"},
		}

		expected := &entities.Expenses{
			Id:     1,
			Title:  "Isakaya Bangna",
			Amount: 899,
			Note:   "central bangna",
			Tags:   []string{"food", "beverage"},
		}

		mockRepo.EXPECT().CreateExpense(given).Return(expected, nil)

		// Act
		service := Init(mockRepo)
		result, err := service.CreateExpense(given)

		// Assert
		if assert.NoError(t, err) {
			assert.Equal(t, expected.Id, result.Id)
			assert.Equal(t, expected.Title, result.Title)
			assert.Equal(t, expected.Amount, result.Amount)
			assert.Equal(t, expected.Note, result.Note)
			assert.Equal(t, expected.Tags, result.Tags)
		}
	})

	t.Run("Fail - TestAddNewExpense", func(t *testing.T) {
		// Arrange
		given := &entities.Expenses{
			Id:     0,
			Title:  "Isakaya Bangna",
			Amount: 899,
			Note:   "central bangna",
			Tags:   []string{"food", "beverage"},
		}

		expected := fmt.Errorf("error service create expenses")
		mockRepo.EXPECT().CreateExpense(given).Return(nil, expected)

		// Act
		service := Init(mockRepo)
		_, err := service.CreateExpense(given)

		// Assert
		if err != nil {
			assert.Equal(t, expected, err)
		}
	})
}

func TestGetExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)

	t.Run("Success - TestGetExpense", func(t *testing.T) {
		// Arrange
		given := "1"
		expected := &entities.Expenses{
			Id:     0,
			Title:  "",
			Amount: 0,
			Note:   "",
			Tags:   []string{},
		}

		mockRepo.EXPECT().GetExpense(given).Return(expected, nil)

		// Act
		service := Init(mockRepo)
		result, err := service.GetExpense(given)

		// Assert
		if assert.NoError(t, err) {
			assert.Equal(t, expected.Id, result.Id)
			assert.Equal(t, expected.Title, result.Title)
			assert.Equal(t, expected.Amount, result.Amount)
			assert.Equal(t, expected.Note, result.Note)
			assert.Equal(t, expected.Tags, result.Tags)
		}
	})

	t.Run("Fail - TestGetExpense", func(t *testing.T) {
		// Arrange
		given := "1"

		expected := fmt.Errorf("error service get expenses detail")
		mockRepo.EXPECT().GetExpense(given).Return(nil, expected)

		// Act
		service := Init(mockRepo)
		_, err := service.GetExpense(given)

		// Assert
		if err != nil {
			assert.Equal(t, expected, err)
		}
	})
}
