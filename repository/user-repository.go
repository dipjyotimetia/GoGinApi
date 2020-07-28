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
	GetUser(id int64) entity.User
	UpdateUser(id int64, user entity.User) int64
	DeleteUser(id int64) int64
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

// get one user from the DB by its userid
func (db UserDatabase) GetUser(id int64) entity.User {

	// create a user of models.User type
	var user entity.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE uid=$1`

	// execute the sql statement
	row := db.connection.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&user.ID, &user.Name, &user.Location, &user.Age)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user
	case nil:
		return user
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return user
}

// update user in the DB
func (db UserDatabase) UpdateUser(id int64, user entity.User) int64 {

	// close the db connection
	//defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE uid=$1`

	// execute the sql statement
	res, err := db.connection.Exec(sqlStatement, id, user.Name, user.Location, user.Age)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete user in the DB
func (db UserDatabase) DeleteUser(id int64) int64 {

	// create the delete sql query
	sqlStatement := `DELETE FROM users WHERE uid=$1`

	// execute the sql statement
	res, err := db.connection.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func NewUserRepository() UserRepository {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "database", "goland", "goland", "goland")
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