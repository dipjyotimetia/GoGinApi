package repository

import (
	"fmt"
	"log"

	"github.com/GoGinApi/v2/entity"
)

const (
	addExpenseStatement    = `INSERT INTO expense (clientID, expenseType, expenseAmount, expenseDate) VALUES ($1,$2,$3,$4) RETURNING eid`
	getAllExpenseStatement = `SELECT * FROM expense`
	getExpenseStatement    = `SELECT * FROM expense WHERE eid=$1`
	updateExpenseStatement = `UPDATE expense SET clientID=$2, expenseType=$3, expenseAmount=$4,expenseDate=$5 WHERE eid=$1`
	deleteExpenseStatement = `DELETE FROM expense WHERE eid=$1`
)

// CloseDB closing db connection
func (db Database) CloseDB() {
	err := db.Close()
	if err != nil {
		panic("failed to close connection")
	}
}

// AddExpense add expense in expense table
func (db Database) AddExpense(expense entity.Expense) int64 {
	var id int64

	err := db.QueryRow(addExpenseStatement, expense.ClientID, expense.ExpenseType, expense.ExpenseAmount, expense.ExpenseDate).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// return the inserted id
	return id
}

// GetAllExpense query from expense table
func (db Database) GetAllExpense() []entity.Expense {
	var expenses []entity.Expense

	rows, err := db.Query(getAllExpenseStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var expense entity.Expense

		// unmarshal the row object to user
		err = rows.Scan(&expense.ExpenseID, &expense.ClientID, &expense.ExpenseType, &expense.ExpenseAmount, &expense.ExpenseDate)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		// append the user in the users slice
		expenses = append(expenses, expense)
	}
	// return empty user on error
	return expenses
}

// get one expense from the DB by its expense id
func (db Database) GetExpense(id int64) (entity.Expense, error) {
	var expense entity.Expense

	// execute the sql statement
	row := db.QueryRow(getExpenseStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&expense.ExpenseID, &expense.ClientID, &expense.ExpenseType, &expense.ExpenseAmount, &expense.ExpenseDate)

	if err != nil {
		return expense, fmt.Errorf("no rows found")
	}

	return expense, nil
}

func (db Database) UpdateExpense(id int64, expense entity.Expense) error {
	res, err := db.Exec(updateExpenseStatement, id, expense.ClientID, expense.ExpenseType, expense.ExpenseAmount, expense.ExpenseDate)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error while checking the affected rows. %v", err)
	}

	return nil
}

// delete expense in the DB
func (db Database) DeleteExpense(id int64) error {
	res, err := db.Exec(deleteExpenseStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return fmt.Errorf("error while checking the affected rows. %v", err)
	}

	return nil
}
