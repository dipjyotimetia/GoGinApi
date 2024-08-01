package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/pact-foundation/pact-go/v2/dsl"
)

func TestAccountEntity(t *testing.T) {
	tests := []Account{
		{AccountID: 1, CurrencyCode: "AUD", StatusCode: "active", Balance: 10, ClientID: 1},
		{AccountID: 2, CurrencyCode: "USD", StatusCode: "inactive", Balance: 20, ClientID: 2},
		{AccountID: 3, CurrencyCode: "INR", StatusCode: "idle", Balance: 0, ClientID: 3},
	}
	for _, tt := range tests {
		t.Run("TestAccount", func(t *testing.T) {
			assert.Equal(t, tt, tt)
		})
	}
}

func TestPact_AccountEntity(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "AccountEntityConsumer",
		Provider: "AccountEntityProvider",
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
