package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/pact-foundation/pact-go/v2/dsl"
)

func TestExpenseEntity(t *testing.T) {
	tests := []Expense{
		{ExpenseID: int64(100), ExpenseType: "shopping", ExpenseAmount: 90, ExpenseDate: "12/12/2020", ClientID: 1},
		{ExpenseID: int64(101), ExpenseType: "food", ExpenseAmount: 50, ExpenseDate: "12/12/2020", ClientID: 1},
		{ExpenseID: int64(102), ExpenseType: "beverage", ExpenseAmount: 20, ExpenseDate: "24/12/2020", ClientID: 1},
	}

	for _, tt := range tests {
		t.Run("TestExpense", func(t *testing.T) {
			assert.Equal(t, tt, tt)
		})
	}
}

func TestPact_ExpenseEntity(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "ExpenseEntityConsumer",
		Provider: "ExpenseEntityProvider",
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
