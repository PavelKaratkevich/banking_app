package dto

import (
	"testing"
)

func Test_if_validate_returns_an_error_while_passing_a_wrong_transaction_type(t *testing.T) {

	transaction := NewTransactionRequest{
		Type:   "invalid_type",
		Amount: -2000,
	}
	err := transaction.Validate()
	if err.Message != "Transaction type should be 'withdrawal' or 'deposit'" && err.Message != "To make a transaction you need to deposit an amount equal to at least 1.0" {
		t.Error("Error while checking transaction type")
	}
}
