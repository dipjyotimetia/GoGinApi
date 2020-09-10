package mocks

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/stretchr/testify/mock"
)

type MyMockedObject struct {
	mock.Mock
}

func (mock *MyMockedObject) AddExpense(expense entity.Expense) int64 {
	args := mock.Called()
	return int64(args.Int(0))
}

func (mock *MyMockedObject) GetAllExpense() []entity.Expense {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Expense)
}

func (mock *MyMockedObject) GetExpense(id int64) (entity.Expense, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.Expense), args.Error(1)
}

func (mock *MyMockedObject) UpdateExpense(id int64, expense entity.Expense) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MyMockedObject) DeleteExpense(id int64) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MyMockedObject) ResetPassword(resetPassword entity.ResetPassword) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MyMockedObject) Create(user entity.Register) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MyMockedObject) Login(name, email, password, createdAt, updatedAt string, user entity.Login) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MyMockedObject) CheckUserExist(user entity.Register) bool {
	args := mock.Called()
	return args.Bool(0)
}

func (mock *MyMockedObject) CheckAndRetrieveUserIDViaEmail(createReset entity.CreateReset) (int, bool) {
	args := mock.Called()
	return 1, args.Bool(1)
}

func (mock *MyMockedObject) UpdateAccountDetails(clientID int64, account entity.Account) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MyMockedObject) CloseDB() {
	panic("implement me")
}

func (mock *MyMockedObject) GetAccountDetails(clientID int64) (entity.Account, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.Account), args.Error(1)
}

func (mock *MyMockedObject) AddAccountDetails(account entity.Account) error {
	args := mock.Called()
	return args.Error(0)
}
