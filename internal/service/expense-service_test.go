package service

import (
	"github.com/GoGinApi/v2/internal/entity"
	"github.com/GoGinApi/v2/internal/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/pact-foundation/pact-go/v2/dsl"
)

var expense = entity.Expense{
	ExpenseID:     1,
	ExpenseType:   "Test",
	ExpenseAmount: 11.5,
	ExpenseDate:   "12/12/2019",
	ClientID:      1,
}

var expenses = []entity.Expense{
	{
		ExpenseID:     1,
		ExpenseType:   "Test",
		ExpenseAmount: 11.5,
		ExpenseDate:   "12/12/2019",
		ClientID:      1,
	},
	{
		ExpenseID:     2,
		ExpenseType:   "Test2",
		ExpenseAmount: 12.5,
		ExpenseDate:   "15/12/2019",
		ClientID:      1,
	},
}

func TestExpensesServices(t *testing.T) {
	for _, c := range []TestCase{
		{"Successful add expense details", testAddExpense},
		{"Successful get expense details", testGetExpense},
		{"Successful get all expense details", testGetAllExpense},
		{"Successful update expense details", testUpdateExpense},
		{"Successful delete expense details", testDeleteExpense},
		{"Successful Pact test for expense service", testPactExpenseService},
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

func testPactExpenseService(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "ExpenseServiceConsumer",
		Provider: "ExpenseServiceProvider",
	}

	defer pact.Teardown()

	pact.AddInteraction().
		Given("Expense with ID 1 exists").
		UponReceiving("A request to get expense details").
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   dsl.String("/expenses/1"),
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body: dsl.Match(&map[string]interface{}{
				"expenseID":     dsl.Like(1),
				"expenseType":   dsl.Like("Test"),
				"expenseAmount": dsl.Like(11.5),
				"expenseDate":   dsl.Like("12/12/2019"),
				"clientID":      dsl.Like(1),
			}),
		})

	err := pact.Verify(func() error {
		// Make request to the provider
		// This is where you would call your actual service
		return nil
	})

	assert.NoError(t, err)
}
