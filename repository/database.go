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
	ResetPassword(resetPassword entity.ResetPassword) error
	Create(user entity.Register) error
	Login(name, email, password, createdAt, updatedAt string, user entity.Login) error
	CheckUserExist(user entity.Register) bool
	CheckAndRetrieveUserIDViaEmail(createReset entity.CreateReset) (int, bool)
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
