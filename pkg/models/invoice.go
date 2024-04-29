package models

type InvoiceStatus string

const (
	Pending InvoiceStatus = "pending"
	Paid    InvoiceStatus = "paid"
)

type Invoice struct {
	ID        uint64        `db:"id" json:"id"`
	UserID    uint64        `db:"user_id" json:"user_id"`
	Status    InvoiceStatus `db:"status" json:"status" validate:"oneof=paid pending"`
	Label     string        `db:"label" json:"label"`
	Amount    int           `db:"amount" json:"amount" validate:"gte=0"`
	Reference string        `json:"reference"`
}

type InvoiceInput struct {
	UserID uint64 `json:"user_id" validate:"required"`
	Label  string `json:"label" validate:"required"`
	Amount int    `json:"amount" validate:"required,gte=0"`
}
