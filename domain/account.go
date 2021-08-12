package domain

import (
	"banking/dto"
	"banking/errs"
)

// домен
type Account struct {
	AccountID   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		a.AccountID,
	}
}

// senodary port для связи с базой данных
//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain banking/domain AccountRepository
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppErr)
	Update(Transaction) (*Transaction, *errs.AppErr)
	FindBy(accountId int) (*Account, *errs.AppErr)
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}
