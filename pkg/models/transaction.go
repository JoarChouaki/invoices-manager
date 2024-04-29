package models

type TransactionInput struct {
	InvoiceID uint64 `json:"invoice_id" validate:"required"`
	Amount    int    `json:"amount" validate:"required,gte=0"`
	Reference string `json:"reference"`
}
