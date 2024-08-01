package service

import (
	"github.com/GoGinApi/v2/internal/entity"
	"github.com/GoGinApi/v2/internal/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/pact-foundation/pact-go/v2/dsl"
)

type TestCase struct {
	name string
	test func(t *testing.T)
}

var account = entity.Account{
	AccountID:    1,
	CurrencyCode: "AUD",
	StatusCode:   "active",
	Balance:      200,
	ClientID:     1,
}

func TestAccountServices(t *testing.T) {
	for _, c := range []TestCase{
		{"Successful add account details", testAddAccountDetails},
		{"Successful get account details", testGetAccountDetails},
		{"Successful update account details", testUpdateAccountDetails},
		{"Successful Pact test for account service", testPactAccountService},
	} {
		t.Run(c.name, c.test)
	}
}

func testAddAccountDetails(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	mockRepo.On("AddAccountDetails").Return(nil)

	testService := NewAccountService(mockRepo)
	err := testService.AddAccountDetails(account)
	assert.Equal(t, err, nil)
}

func testUpdateAccountDetails(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	mockRepo.On("UpdateAccountDetails").Return(nil)

	testService := NewAccountService(mockRepo)
	err := testService.UpdateAccountDetails(10, account)
	assert.Equal(t, err, nil)
}

func testGetAccountDetails(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	mockRepo.On("GetAccountDetails").Return(account, nil)

	testService := NewAccountService(mockRepo)
	result, _ := testService.GetAccountDetails(10)
	assert.Equal(t, result, account)
}

func testPactAccountService(t *testing.T) {
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
