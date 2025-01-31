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
		assert.Error(t, err)
		assert.Equal(t, expected, err)
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
		assert.Error(t, err)
		assert.Equal(t, expected, err)
	})
}

func TestUpdateExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)

	t.Run("Success - TestUpdateExpense", func(t *testing.T) {
		// Arrange
		id := "1"
		newExpense := &entities.Expenses{
			Id:     1,
			Title:  "Isakaya Bangna",
			Amount: 990,
			Note:   "Central bangna near by BigC",
			Tags:   []string{"food", "beverage"},
		}
		expected := &entities.Expenses{
			Id:     1,
			Title:  "Isakaya Bangna",
			Amount: 990,
			Note:   "Central bangna near by BigC",
			Tags:   []string{"food", "beverage"},
		}

		mockRepo.EXPECT().UpdateExpense(id, newExpense).Return(expected, nil)

		// Act
		service := Init(mockRepo)
		result, err := service.UpdateExpense(id, newExpense)

		// Assert
		if assert.NoError(t, err) {
			assert.Equal(t, expected.Id, result.Id)
			assert.Equal(t, expected.Title, result.Title)
			assert.Equal(t, expected.Amount, result.Amount)
			assert.Equal(t, expected.Note, result.Note)
			assert.Equal(t, expected.Tags, result.Tags)
		}
	})

	t.Run("Fail - TestUpdateExpense", func(t *testing.T) {
		// Arrange
		id := "1"
		newExpense := entities.Expenses{}

		expected := fmt.Errorf("error service update expenses")
		mockRepo.EXPECT().UpdateExpense(id, &newExpense).Return(nil, expected)

		// Act
		service := Init(mockRepo)
		_, err := service.UpdateExpense(id, &newExpense)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expected, err)
	})
}

func TestGetExpenseList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)

	t.Run("Success - TestGetExpenseList", func(t *testing.T) {
		// Arrange
		expected := []entities.Expenses{
			{
				Id:     1,
				Title:  "Isakaya Bangna",
				Amount: 899,
				Note:   "central bangna",
				Tags:   []string{"food", "beverage"},
			},
		}

		mockRepo.EXPECT().GetExpenseList().Return(expected, nil)

		// Act
		service := Init(mockRepo)
		result, _ := service.GetExpenseList()

		// Assert
		assert.Equal(t, expected, result)
		assert.Len(t, expected, 1)
	})

	t.Run("Fail - TestGetExpenseList", func(t *testing.T) {
		// Arrange
		expected := fmt.Errorf("error service get expenses list")
		mockRepo.EXPECT().GetExpenseList().Return(nil, expected)

		// Act
		service := Init(mockRepo)
		_, err := service.GetExpenseList()

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expected, err)
	})
}
