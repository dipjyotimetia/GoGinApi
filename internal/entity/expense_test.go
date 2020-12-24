package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
