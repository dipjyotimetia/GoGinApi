package service

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/repository"
)

type UserService interface {
	InsertUser(user entity.User) int64
	GetAllUsers() []entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUser(repo repository.UserRepository) UserService {
	return &userService{userRepository: repo}
}

func (u userService) InsertUser(user entity.User) int64 {
	u.userRepository.InsertUser(user)
	return user.ID
}

func (u userService) GetAllUsers() []entity.User {
	return u.userRepository.GetAllUsers()
}

//
//// get one user from the DB by its userid
//func GetUser(id int64) (entity.User, error) {
//	// create the postgres db connection
//	db := createConnection()
//
//	// close the db connection
//	defer db.Close()
//
//	// create a user of models.User type
//	var user entity.User
//
//	// create the select sql query
//	sqlStatement := `SELECT * FROM users WHERE userid=$1`
//
//	// execute the sql statement
//	row := db.QueryRow(sqlStatement, id)
//
//	// unmarshal the row object to user
//	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)
//
//	switch err {
//	case sql.ErrNoRows:
//		fmt.Println("No rows were returned!")
//		return user, nil
//	case nil:
//		return user, nil
//	default:
//		log.Fatalf("Unable to scan the row. %v", err)
//	}
//
//	// return empty user on error
//	return user, err
//}

//// update user in the DB
//func updateUser(id int64, user entity.User) int64 {
//
//	// create the postgres db connection
//	db := createConnection()
//
//	// close the db connection
//	defer db.Close()
//
//	// create the update sql query
//	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`
//
//	// execute the sql statement
//	res, err := db.Exec(sqlStatement, id, user.Name, user.Location, user.Age)
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
//
//// delete user in the DB
//func deleteUser(id int64) int64 {
//
//	// create the postgres db connection
//	db := createConnection()
//
//	// close the db connection
//	defer db.Close()
//
//	// create the delete sql query
//	sqlStatement := `DELETE FROM users WHERE userid=$1`
//
//	// execute the sql statement
//	res, err := db.Exec(sqlStatement, id)
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
