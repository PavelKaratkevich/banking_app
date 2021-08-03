package domain

import (
	"banking/errs"
	"banking/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"strings"
	"time"
)

// Адаптер для связи с Secondary port
type TransactionRepositoryDb struct {
	client *sqlx.DB
}

// реализация метода, предусмотренного Secondary port
func (d TransactionRepositoryDb) Update(t Transaction) (*Transaction, *errs.AppErr) {
	var accountSql string
	var transactionSql string

	if strings.ToLower(t.Type) == "deposit" {
		accountSql = "UPDATE accounts SET amount = amount + ? where account_id = ?"
		transactionSql = "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)"
	} else {
		accountSql = "UPDATE accounts SET amount = amount - ? where account_id = ?"
		transactionSql = "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES  (?, ?, ?, ?)"
	}

	_, err := d.client.Exec(accountSql, t.Amount, t.AccountId)
	if err != nil {
		logger.Error("Error while updating accounts table")
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	result, err := d.client.Exec(transactionSql, t.AccountId, t.Amount, t.Type, t.Date)

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted transaction ID from database")
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	t.TransactionId = int(transactionId)
	return &t, nil
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	dbclient, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbclient.SetMaxOpenConns(10)
	dbclient.SetMaxIdleConns(10)
	return TransactionRepositoryDb{dbClient}
}
