package domain

import (
	"banking/dto"
	"banking/errs"
)

// домен
type Account struct {
	AccountID   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		a.AccountID,
	}
}

// senodary port для связи с базой данных
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppErr)
}
