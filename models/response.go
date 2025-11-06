package models

type Response struct {
	TransactionStatus string `json:"transaction_status"`
	ErrorMessage      string `json:"error_message"`
}
