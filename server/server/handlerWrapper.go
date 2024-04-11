package server

import "net/http"

type handlerWrapper interface {
	wrapHandler(http.Handler) http.Handler
}

func wrapHandler(h http.Handler, wrappers ...handlerWrapper) http.Handler {
	for _, wrapper := range wrappers {
		h = wrapper.wrapHandler(h)
	}
	return h
}
