package server

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"pjm.dev/sfs/app"
	"pjm.dev/sfs/config"
	"pjm.dev/sfs/graph"
)

func New(config config.Server, app *app.App) http.Handler {
	mux := http.NewServeMux()

	// initialize graph handler
	graphPattern := fmt.Sprintf("/%s", config.GraphEndpoint)
	graphHandler := graph.NewHandler(app, app)
	graphHandler = wrapHandler(
		graphHandler,
		&loggingHandler{},
		&corsHandler{},
	)
	mux.Handle(graphPattern, graphHandler)

	// initialize playground handler
	playgroundPattern := "/"
	playgroundHandler := playground.Handler("SFS Playground", graphPattern)
	mux.Handle(playgroundPattern, playgroundHandler)

	return mux
}
