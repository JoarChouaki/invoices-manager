package http

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/sirupsen/logrus"
)

func All() []alice.Constructor {
	return []alice.Constructor{
		LogRequest,
	}
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &StatusRecorder{
			ResponseWriter: w,
			Status:         200,
		}
		next.ServeHTTP(recorder, r)
		path := r.URL.Path
		logrus.Infof("%s %s %d", r.Method, path, recorder.Status)
	})
}
