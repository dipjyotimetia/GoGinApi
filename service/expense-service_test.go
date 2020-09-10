package service

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

var expense = entity.Expense{
	ExpenseID:     1,
	ClientID:      1,
	ExpenseType:   "Test",
	ExpenseAmount: 11.5,
	ExpenseDate:   "12/12/2019",
}

var expenses = []entity.Expense{
	{
		ExpenseID:     1,
		ClientID:      1,
		ExpenseType:   "Test",
		ExpenseAmount: 11.5,
		ExpenseDate:   "12/12/2019",
	},
	{
		ExpenseID:     2,
		ClientID:      1,
		ExpenseType:   "Test2",
		ExpenseAmount: 12.5,
		ExpenseDate:   "15/12/2019",
	},
}

func TestExpensesServices(t *testing.T) {
	for _, c := range []TestCase{
		{"Successful add expense details", testAddExpense},
		{"Successful get expense details", testGetExpense},
		{"Successful get all expense details", testGetAllExpense},
		{"Successful update expense details", testUpdateExpense},
		{"Successful delete expense details", testDeleteExpense},
	} {
		t.Run(c.name, c.test)
	}
}

func testAddExpense(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	mockRepo.On("AddExpense").Return(1)

	testService := NewExpense(mockRepo)
	res := testService.AddExpense(expense)
	assert.Equal(t, res, int64(1))
}

func testUpdateExpense(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	mockRepo.On("UpdateExpense").Return(nil)

	testService := NewExpense(mockRepo)
	err := testService.UpdateExpense(1, expense)
	assert.Equal(t, err, nil)
}

func testGetExpense(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	mockRepo.On("GetExpense").Return(expense, nil)

	testService := NewExpense(mockRepo)
	result, _ := testService.GetExpense(1)
	assert.Equal(t, result, expense)
}

func testGetAllExpense(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	mockRepo.On("GetAllExpense").Return(expenses, nil)

	testService := NewExpense(mockRepo)
	res := testService.GetAllExpense()
	assert.Equal(t, res, expenses)
}

func testDeleteExpense(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	mockRepo.On("DeleteExpense").Return(nil)

	testService := NewExpense(mockRepo)
	err := testService.DeleteExpense(1)
	assert.Equal(t, err, nil)
}
