package repository

import (
	"log"

	"github.com/GoGinApi/v2/entity"
)

const (
	addExpenseStatement    = `INSERT INTO expense (username, expenseType, expenseAmount, expenseDate) VALUES ($1,$2,$3,$4) RETURNING eid`
	getAllExpenseStatement = `SELECT * FROM expense`
	//getExpenseStatement = `SELECT * FROM users WHERE uid=$1`
	//deleteExpenseStatement = `DELETE FROM users WHERE uid=$1`
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

	err := db.QueryRow(addExpenseStatement, expense.Username, expense.ExpenseType, expense.ExpenseAmount, expense.ExpenseDate).Scan(&id)
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

//// get one user from the DB by its userid
//func (db UserDatabase) GetExpense(id int64) entity.User { // create a user of models.User type
//	var user entity.User
//
//	// create the select sql query
//	sqlStatement := `SELECT * FROM users WHERE uid=$1`
//
//	// execute the sql statement
//	row := db.connection.QueryRow(sqlStatement, id)
//
//	// unmarshal the row object to user
//	err := row.Scan(&user.ID, &user.Name, &user.Location, &user.Age)
//
//	switch err {
//	case sql.ErrNoRows:
//		fmt.Println("No rows were returned!")
//		return user
//	case nil:
//		return user
//	default:
//		log.Fatalf("Unable to scan the row. %v", err)
//	}
//
//	// return empty user on error
//	return user
//}

//// delete user in the DB
//func (db UserDatabase) DeleteExpense(id int64) int64 { // create the delete sql query
//	sqlStatement := `DELETE FROM users WHERE uid=$1`
//
//	// execute the sql statement
//	res, err := db.connection.Exec(sqlStatement, id)
//
//	if err != nil {
//		log.Fatalf("Unable to execute the query. %v", err)
//	}
//
//	// check how many rows affected
//	rowsAffected, err := res.RowsAffected()
//
//	if err != nil {
//		log.Fatalf("Error while checking the affected rows. %v", err)
//	}
//
//	fmt.Printf("Total rows/record affected %v", rowsAffected)
//
//	return rowsAffected
//}
