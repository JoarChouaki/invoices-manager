package main

import (
	"database/sql"
	"fmt"
	http1 "invoices-manager/pkg/http"
	"invoices-manager/pkg/models"

	"net/http"
)

func (s *Server) CreateInvoice(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var invoice models.InvoiceInput
	if err := http1.Parse(r, &invoice); err != nil {
		http1.RespondBadRequest(w, r, "unable to parse input")
		return
	}

	fmt.Println("invoice = ", invoice)

	var user models.User
	if err := s.GetContext(
		ctx,
		&user,
		"SELECT * FROM users WHERE id = $1",
		invoice.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			http1.RespondNotFound(w, r)
			return
		}
		http1.RespondInternalServerError(w, r, err)
		return
	}

	if _, err := s.ExecContext(
		ctx,
		`INSERT INTO invoices (user_id, status, label, amount) 
		VALUES ($1, $2, $3, $4) `,
		invoice.UserID, models.Pending, invoice.Label, invoice.Amount,
	); err != nil {
		http1.RespondInternalServerError(w, r, err)
		return
	}

	http1.RespondNoContent(w)
}
