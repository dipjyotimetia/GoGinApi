package repository

import (
	"database/sql"
	"fmt"
	"github.com/GoGinApi/v2/entity"
	"log"
)

const (
	addAccountStatement           = `INSERT INTO accounts(clientId,currencyCode,statusCode,balance) VALUES ( $1 , $2, $3, $4)`
	getAccountDetailsStatement    = `SELECT * FROM accounts WHERE clientId = $1`
	updateAccountDetailsStatement = `UPDATE accounts SET currencyCode=$2, statusCode=$3, balance=$4 WHERE clientId=$1`
)

func (db *Database) AddAccountDetails(account entity.Account) error {
	_, err := db.Query(addAccountStatement, account.ClientID, account.CurrencyCode, account.StatusCode, account.Balance)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record")
	return nil
}

func (db *Database) GetAccountDetails(clientID int64) (entity.Account, error) {
	var account entity.Account

	row := db.QueryRow(getAccountDetailsStatement, clientID)

	err := row.Scan(&account.ClientID, &account.CurrencyCode, &account.StatusCode, &account.Balance)
	if err == sql.ErrNoRows {
		return entity.Account{}, fmt.Errorf("account details not exist")
	}
	return account, nil
}

func (db *Database) UpdateAccountDetails(clientID int64, account entity.Account) error {
	res, err := db.Exec(updateAccountDetailsStatement, clientID, account.CurrencyCode, account.StatusCode, account.Balance)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error while checking the affected rows. %v", err)
	}

	return nil
}
