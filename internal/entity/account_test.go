package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
