package dto

type NewTransactionResponse struct {
	TransactionId int     `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}
