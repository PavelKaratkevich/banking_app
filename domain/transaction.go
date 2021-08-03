package domain

import (
	"banking/dto"
	"banking/errs"
)

type Transaction struct {
	TransactionId int
	AccountId     int
	Amount        float64
	Type          string
	Date          string
}

func (a Transaction) ToNewTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		a.TransactionId,
		a.Amount,
	}
}

// TransactionRepository is a secondary port connecting Transaction with a Data Base
type TransactionRepository interface {
	Update(Transaction) (*Transaction, *errs.AppErr)
}
