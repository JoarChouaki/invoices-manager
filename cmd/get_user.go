package main

import (
	"database/sql"
	http1 "invoices-manager/pkg/http"
	"invoices-manager/pkg/models"

	"net/http"
)

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var id uint64
	if err := http1.Parse(r, &id); err != nil {
		http1.RespondBadRequest(w, r, "unable to parse input")
		return
	}

	var user models.User
	if err := s.GetContext(
		ctx,
		&user,
		"SELECT * FROM users WHERE id = $1",
		id,
	); err != nil {
		if err == sql.ErrNoRows {
			http1.RespondNotFound(w, r)
			return
		}
		http1.RespondInternalServerError(w, r, err)
		return
	}

	http1.Respond(w, user)
}
