package main

import (
	"database/sql"
	http1 "invoices-manager/pkg/http"
	"invoices-manager/pkg/models"

	"net/http"

	"golang.org/x/sync/errgroup"
)

func (s *Server) ExecuteTransaction(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var input models.TransactionInput
	if err := http1.Parse(r, &input); err != nil {
		http1.RespondBadRequest(w, r, "unable to parse input")
		return
	}

	var invoice models.Invoice
	if err := s.GetContext(
		ctx,
		&invoice,
		"SELECT * FROM invoices WHERE id = $1",
		input.InvoiceID,
	); err != nil {
		if err == sql.ErrNoRows {
			http1.RespondNotFound(w, r)
			return
		}
		http1.RespondInternalServerError(w, r, err)
		return
	}

	if invoice.Amount != input.Amount {
		http1.RespondBadRequest(w, r, "transaction amount should be equal to invoice amount")
		return
	}

	if invoice.Status == models.Paid {
		http1.RespondUnprocessable(w, r)
		return
	}

	group, ctx1 := errgroup.WithContext(ctx)
	group.Go(func() error {
		_, err := s.ExecContext(
			ctx1,
			`UPDATE invoices SET status = $1 WHERE id = $2 `,
			models.Paid,
			input.InvoiceID,
		)
		return err
	})

	group.Go(func() error {
		_, err := s.ExecContext(
			ctx1,
			`UPDATE users SET balance = balance + $1 WHERE id = $2 `,
			input.Amount,
			invoice.UserID,
		)
		return err
	})

	if err := group.Wait(); err != nil {
		http1.RespondInternalServerError(w, r, err)
		return
	}

	http1.RespondNoContent(w)
}
