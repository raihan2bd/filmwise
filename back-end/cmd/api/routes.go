package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.GetStatus)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)

	return app.enableCORS(router)
}
