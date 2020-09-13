package entity

type Expense struct {
	ExpenseID     int64   `json:"expenseID"`
	ExpenseType   string  `json:"expenseType"`
	ExpenseAmount float64 `json:"expenseAmount"`
	ExpenseDate   string  `json:"expenseDate"`
	ClientID      int64   `json:"clientID"`
}
