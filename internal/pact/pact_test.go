package pact

import (
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/v2/dsl"
	"github.com/stretchr/testify/assert"
)

func TestPact_AccountService(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "AccountServiceConsumer",
		Provider: "AccountServiceProvider",
	}

	defer pact.Teardown()

	pact.AddInteraction().
		Given("Account with ID 1 exists").
		UponReceiving("A request to get account details").
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   dsl.String("/accounts/1"),
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body: dsl.Match(&map[string]interface{}{
				"accountID":    dsl.Like(1),
				"currencyCode": dsl.Like("AUD"),
				"statusCode":   dsl.Like("active"),
				"balance":      dsl.Like(200.0),
				"clientID":     dsl.Like(1),
			}),
		})

	err := pact.Verify(func() error {
		// Make request to the provider
		// This is where you would call your actual service
		return nil
	})

	assert.NoError(t, err)
}

func TestPact_ExpenseService(t *testing.T) {
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

func TestPact_UserService(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "UserServiceConsumer",
		Provider: "UserServiceProvider",
	}

	defer pact.Teardown()

	pact.AddInteraction().
		Given("User with email test1@gmail.com exists").
		UponReceiving("A request to check user existence").
		WithRequest(dsl.Request{
			Method: "POST",
			Path:   dsl.String("/users/check"),
			Body: dsl.Match(&map[string]interface{}{
				"email": dsl.Like("test1@gmail.com"),
			}),
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body: dsl.Match(&map[string]interface{}{
				"exists": dsl.Like(true),
			}),
		})

	err := pact.Verify(func() error {
		// Make request to the provider
		// This is where you would call your actual service
		return nil
	})

	assert.NoError(t, err)
}
