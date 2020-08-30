package repository

import (
	"database/sql"
	"fmt"

	"github.com/GoGinApi/v2/entity"
	_ "github.com/lib/pq" //nolint:golint
)

const (
	host string = "db"
	db   string = "goland"
	user string = "goland"
	pass string = "goland"
)

//DataStore having all repository interface
type DataStore interface {
	AddExpense(expense entity.Expense) int64
	GetAllExpense() []entity.Expense
	InsertUser(user entity.User) int64
	GetAllUsers() []entity.User
	GetUser(id int64) entity.User
	UpdateUser(id int64, user entity.User) int64
	DeleteUser(id int64) int64
	CloseDB()
}

//Database initialization
type Database struct {
	connection *sql.DB
}

//DatabaseConnection establish database connection
func DatabaseConnection() DataStore {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, db, pass)
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

	return &Database{
		connection: db,
	}
}
