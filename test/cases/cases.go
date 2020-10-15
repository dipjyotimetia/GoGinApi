package cases

import "github.com/GoGinApi/v2/internal/entity"

type Suite []Case

type Case struct {
	name    string
	Login   entity.Login
	Expense entity.Expense
}

var TestCases = Suite{
	{
		name: "Verify valid login request",
		Login: entity.Login{
			Password: "Password1",
			Email:    "test@gmail.com",
		},
		Expense: entity.Expense{
			ExpenseType:   "",
			ExpenseAmount: 0,
			ExpenseDate:   "",
			ClientID:      0,
		},
	},
}
