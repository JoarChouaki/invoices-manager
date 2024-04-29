package main

import (
	http1 "invoices-manager/pkg/http"
	"invoices-manager/pkg/models"

	"net/http"
)

func (s *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var users []*models.User
	if err := s.SelectContext(
		ctx,
		&users,
		"SELECT * FROM users ORDER BY first_name ASC",
	); err != nil {
		http1.RespondInternalServerError(w, r, err)
		return
	}

	http1.Respond(w, users)
}
