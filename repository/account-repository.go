package repository

import (
	"database/sql"
	"fmt"
	"github.com/GoGinApi/v2/entity"
	"log"
)

const (
	addAccountStatement           = `INSERT INTO accounts(currencyCode,statusCode,balance,clientId) VALUES ( $1 , $2, $3, $4)`
	getAccountDetailsStatement    = `SELECT * FROM accounts WHERE accountID = $1`
	updateAccountDetailsStatement = `UPDATE accounts SET currencyCode=$2, statusCode=$3, balance=$4,clientId=$5 WHERE accountID=$1`
)

func (db *Database) AddAccountDetails(account entity.Account) error {
	_, err := db.Query(addAccountStatement, account.CurrencyCode, account.StatusCode, account.Balance, account.ClientID)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record")
	return nil
}

func (db *Database) GetAccountDetails(clientID int64) (entity.Account, error) {
	var account entity.Account

	row := db.QueryRow(getAccountDetailsStatement, clientID)

	err := row.Scan(&account.AccountID, &account.CurrencyCode, &account.StatusCode, &account.Balance, &account.ClientID)
	if err == sql.ErrNoRows {
		return entity.Account{}, fmt.Errorf("account details not exist")
	}
	return account, nil
}

func (db *Database) UpdateAccountDetails(accountID int64, account entity.Account) error {
	res, err := db.Exec(updateAccountDetailsStatement, accountID, account.CurrencyCode, account.StatusCode, account.Balance, account.ClientID)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error while checking the affected rows. %v", err)
	}

	return nil
}
