package main

import "net/http"

type HandlerWrapper interface {
	WrapHandler(http.Handler) http.Handler
}

func WrapHandler(h http.Handler, wrappers ...HandlerWrapper) http.Handler {
	for _, wrapper := range wrappers {
		h = wrapper.WrapHandler(h)
	}
	return h
}
