package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/pkg/utils"
	"log"
)

//ResetPassword reset password database
func (db Database) ResetPassword(resetPassword entity.ResetPassword) error {
	sqlStatement := `UPDATE users SET password = $2 WHERE id = $1`
	if ok, _ := utils.ValidatePasswordReset(resetPassword); ok {
		password := entity.CreateHashedPassword(resetPassword.Password)
		_, err := db.Query(sqlStatement, resetPassword.ID, password)
		return err
	}
	return nil
}

//Create create new user in database
func (db Database) Create(user entity.Register) error {
	sqlStatement := `INSERT INTO users(id,name,password,email) VALUES (DEFAULT, $1 , $2, $3)`

	entity.HashPassword(&user)
	_, err := db.Query(sqlStatement, user.Name, user.Password, user.Email)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record")
	return nil
}

//Login verify login in database
func (db Database) Login(name, email, password, createdAt, updatedAt string, user entity.Login) error {
	sqlStatement := `SELECT * from users WHERE email = $1`

	row := db.QueryRow(sqlStatement, user.Email)

	var id int

	err := row.Scan(&id, &name, &password, &email, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		return fmt.Errorf("user does not exist")
	}

	err = entity.CheckPasswordHash(user.Password, password)
	if err != nil {
		return errors.New("incorrect password")
	}

	return nil
}

//CheckUserExist in the database
func (db Database) CheckUserExist(user entity.Register) bool {
	sqlStatement := `SELECT id from users WHERE email = $1`
	rows, err := db.Query(sqlStatement, user.Email)
	if err != nil {
		return false
	}
	if !rows.Next() {
		return false
	}
	return true
}

//CheckAndRetrieveUserIDViaEmail in user database
func (db Database) CheckAndRetrieveUserIDViaEmail(createReset entity.CreateReset) (int, bool) {
	sqlStatement := `SELECT id from users WHERE email = $1`
	rows, err := db.Query(sqlStatement, createReset.Email)
	if err != nil {
		return -1, false
	}
	if !rows.Next() {
		return -1, false
	}
	var id int
	err = rows.Scan(&id)
	if err != nil {
		return -1, false
	}
	return id, true
}
