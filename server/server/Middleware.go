package server

import (
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func ApplyMiddleware(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

var CORSMiddleware Middleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// allow all headers
		w.Header().Set("Access-Control-Allow-Headers", "*")
		// allow all methods
		w.Header().Set("Access-Control-Allow-Methods", "*")
		// allow credentials
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// call next handler
		next.ServeHTTP(w, r)
	})
}

var LogMiddleware Middleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log request
		log.Printf("%s %s", r.Method, r.URL.Path)
		// call next handler
		next.ServeHTTP(w, r)
	})
}
