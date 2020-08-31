package repository

import (
	"log"

	"github.com/GoGinApi/v2/entity"
)

//CloseDB closing db connection
func (db Database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("failed to close connection")
	}
}

//AddExpense add expense in expense table
func (db Database) AddExpense(expense entity.Expense) int64 {
	sqlStatement := `INSERT INTO expense (username, expenseType, expenseAmount, expenseDate) VALUES ($1,$2,$3,$4) RETURNING eid`

	var id int64

	err := db.connection.QueryRow(sqlStatement, expense.Username, expense.ExpenseType, expense.ExpenseAmount, expense.ExpenseDate).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// return the inserted id
	return id
}

//GetAllExpense query from expense table
func (db Database) GetAllExpense() []entity.Expense {
	var expenses []entity.Expense

	sqlStatement := `SELECT * FROM expense`

	rows, err := db.connection.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var expense entity.Expense

		// unmarshal the row object to user
		err = rows.Scan(&expense.ExpenseID, &expense.Username, &expense.ExpenseType, &expense.ExpenseAmount, &expense.ExpenseDate)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		// append the user in the users slice
		expenses = append(expenses, expense)
	}
	// return empty user on error
	return expenses
}
