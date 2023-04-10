package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.GetStatus)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMoviesByFilter)
	router.HandlerFunc(http.MethodGet, "/v1/movies/genre/:genre_id", app.getAllMoviesByGenre)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)

	// private routes for admin
	router.HandlerFunc(http.MethodGet, "/v1/admin/movie/add", app.AddNewMovie)

	return app.enableCORS(router)
}
