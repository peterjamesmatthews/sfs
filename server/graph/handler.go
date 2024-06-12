package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
)

func New(resolver Resolver) http.Handler {
	executableSchema := NewExecutableSchema(Config{
		Resolvers:  &resolver,
		Directives: DirectiveRoot{Authenticated: authenticated},
	})
	var h http.Handler = handler.NewDefaultServer(executableSchema)
	h = applyMiddleware(h, resolver, putAuthorizationInContext, putResolverInContext)
	return h
}

type middleware func(http.Handler, Resolver) http.Handler

func applyMiddleware(h http.Handler, resolver Resolver, middlewares ...middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h, resolver)
	}
	return h
}

var putAuthorizationInContext middleware = func(next http.Handler, _ Resolver) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		r = r.WithContext(getContextWithAuthorization(r.Context(), authorization))
		next.ServeHTTP(w, r)
	})
}

var putResolverInContext middleware = func(next http.Handler, resolver Resolver) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(getContextWithResolver(r.Context(), resolver))
		next.ServeHTTP(w, r)
	})
}
