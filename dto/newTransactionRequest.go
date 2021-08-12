package dto

import (
	"banking/errs"
	"strings"
)

type NewTransactionRequest struct {
	AccountId  int     `json:"account_id"`
	CustomerId int     `json:"customer_id"`
	Type       string  `json:"transaction_type"`
	Amount     float64 `json:"amount"`
}

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

func (r NewTransactionRequest) Validate() *errs.AppErr {
	if r.Amount < 1 {
		return errs.NewValidationError("To make a transaction you need to deposit an amount equal to at least 1.0")
	}
	if strings.ToLower(r.Type) != "withdrawal" && strings.ToLower(r.Type) != "deposit" {
		return errs.NewValidationError("Transaction type should be 'withdrawal' or 'deposit'")
	}
	return nil
}

func (r NewTransactionRequest) IsTransactionTypeWithdrawal() bool {
	if strings.ToLower(r.Type) == WITHDRAWAL {
		return true
	} else {
		return false
	}
}

func (r NewTransactionRequest) IsTransactionTypeDeposit() bool {
	if r.Type == DEPOSIT {
		return true
	} else {
		return false
	}
}
