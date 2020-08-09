package entity

type Expense struct {
	ExpenseID     int64   `json:"expenseID"`
	Username      string  `json:"username"`
	ExpenseType   string  `json:"expenseType"`
	ExpenseAmount float64 `json:"expenseAmount"`
	ExpenseDate   string  `json:"expenseDate"`
}
