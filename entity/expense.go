package entity

type Expense struct {
	ExpenseID     int64   `json:"expenseID"`
	ClientID      int64   `json:"clientID"`
	ExpenseType   string  `json:"expenseType"`
	ExpenseAmount float64 `json:"expenseAmount"`
	ExpenseDate   string  `json:"expenseDate"`
}
