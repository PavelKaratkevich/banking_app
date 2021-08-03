package dto

type NewTransactionResponse struct {
	TransactionId int     `json:"transaction_id"`
	Balance       float64 `json:"balance"`
}
