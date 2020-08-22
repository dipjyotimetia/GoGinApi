package repository

import (
	"database/sql"
	"fmt"
	"github.com/GoGinApi/v2/entity"
	"log"
)

type ExpenseRepository interface {
	AddExpense(expense entity.Expense) int64
	GetAllExpense() []entity.Expense
	CloseDB()
}

type ExpenseDatabase struct {
	connection *sql.DB
}

func (db ExpenseDatabase) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("failed to close connection")
	}
}

func (db ExpenseDatabase) AddExpense(expense entity.Expense) int64 {
	sqlStatement := `INSERT INTO expense (username, expenseType, expenseAmount, expenseDate) VALUES ($1,$2,$3,$4) RETURNING eid`

	var id int64

	err := db.connection.QueryRow(sqlStatement, expense.Username, expense.ExpenseType, expense.ExpenseAmount, expense.ExpenseDate).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", username)

	// return the inserted id
	return id
}

func (db ExpenseDatabase) GetAllExpense() []entity.Expense {
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

func NewExpenseRepository() ExpenseRepository {
	//dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "db", "goland", "goland", "goland")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "10.102.176.3", "postgres", "postgres", "postgres")
	fmt.Println(dbURI)

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection

	return &ExpenseDatabase{
		connection: db,
	}
}
