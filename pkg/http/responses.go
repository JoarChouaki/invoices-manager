package http

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

// Used to retrieve the response status in a middleware.
func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func Respond(w http.ResponseWriter, out interface{}) {
	outByte, _ := json.Marshal(out)
	w.WriteHeader(http.StatusOK)
	w.Write(outByte)
}

func RespondNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func RespondInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	logrus.WithError(err).Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func RespondNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}

func RespondBadRequest(w http.ResponseWriter, r *http.Request, m string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(m))
}

func RespondUnprocessable(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte("422 Unprocessable Entity"))
}
