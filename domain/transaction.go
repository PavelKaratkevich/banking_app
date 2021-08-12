package domain

import (
	"banking/dto"
)

type Transaction struct {
	TransactionId int     `db:"transaction_id"`
	AccountId     int     `db:"account_id"`
	Amount        float64 `db:"amount"`
	Type          string  `db:"transaction_type"`
	Date          string  `db:"transaction_date"`
}

func (a Transaction) ToNewTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		a.TransactionId,
		a.Amount,
	}
}
