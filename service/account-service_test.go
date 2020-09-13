package service

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
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
