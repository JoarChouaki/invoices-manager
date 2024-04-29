package main

import (
	http1 "invoices-manager/pkg/http"
	"invoices-manager/pkg/models"

	"net/http"
)

func (s *Server) GetUserInvoices(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var id uint64
	if err := http1.Parse(r, &id); err != nil {
		http1.RespondBadRequest(w, r, "unable to parse input")
		return
	}

	var invoices []*models.Invoice
	if err := s.SelectContext(
		ctx,
		&invoices,
		"SELECT * FROM invoices WHERE user_id = $1",
		id,
	); err != nil {
		http1.RespondInternalServerError(w, r, err)
		return
	}

	http1.Respond(w, invoices)
}
