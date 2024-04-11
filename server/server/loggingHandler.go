package server

import (
	"log"
	"net/http"
)

type loggingHandler struct{}

func (l *loggingHandler) wrapHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Default().Printf("%s %s", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
