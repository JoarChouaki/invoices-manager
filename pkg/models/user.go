package models

type User struct {
	ID        uint64 `db:"id" json:"user_id"`
	FirstName string `db:"first_name" json:"first_name" validate:"required"`
	LastName  string `db:"last_name" json:"last_name" validate:"required"`
	Balance   int    `db:"balance" json:"balance"`
}
