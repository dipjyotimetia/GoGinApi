package repository

import (
	"database/sql"
	"fmt"
	"github.com/GoGinApi/v2/entity"
	"github.com/joho/godotenv"
	"log"
)

type UserRepository interface {
	InsertUser(user entity.User) int64
	GetAllUsers() []entity.User
	CloseDB()
}

type UserDatabase struct {
	connection *sql.DB
}

func (db UserDatabase) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("failed to close connection")
	}
}

func (db UserDatabase) InsertUser(user entity.User) int64 {

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING uid`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.connection.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

func (db UserDatabase) GetAllUsers() []entity.User {

	var users []entity.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users`

	// execute the sql statement
	rows, err := db.connection.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user entity.User

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Name, &user.Location, &user.Age)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		// append the user in the users slice
		users = append(users, user)

	}
	// return empty user on error
	return users
}

func NewUserRepository() UserRepository {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "localhost", "goland", "goland", "goland")
	fmt.Println(dbUri)

	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection

	return &UserDatabase{
		connection: db,
	}
}
