package domain

import (
	"banking/dto"
	"banking/errs"
	"banking/logger"
	//"banking/domain"
	//"banking/service"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
	//"strings"
	"time"
)

// Адаптер для связи с Secondary port
type AccountRepositoryDb struct {
	client *sqlx.DB
}

// реализация метода, предусмотренного Secondary port
func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppErr) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while inserting account data")
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted account ID from database")
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountID = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDb) Update(t Transaction) (*Transaction, *errs.AppErr) {
	var accountSql string
	var transactionSql string

	if t.Type == dto.DEPOSIT {
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
	account, _ := d.FindBy(t.AccountId)
	t.Amount = account.Amount

	return &t, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
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
	return AccountRepositoryDb{dbClient}
}

func (d AccountRepositoryDb) FindBy(accountId int) (*Account, *errs.AppErr) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}
