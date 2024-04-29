package main

import (
	http1 "invoices-manager/pkg/http"
	"invoices-manager/pkg/models"

	"net/http"
)

func (s *Server) GetAllInvoices(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var invoices []*models.Invoice
	if err := s.SelectContext(
		ctx,
		&invoices,
		"SELECT * FROM invoices ORDER BY id ASC",
	); err != nil {
		http1.RespondInternalServerError(w, r, err)
		return
	}

	http1.Respond(w, invoices)
}
