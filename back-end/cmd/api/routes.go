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
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)

	// private routes for admin
	router.HandlerFunc(http.MethodPost, "/v1/admin/movie/add", app.AddNewMovie)
	router.HandlerFunc(http.MethodPut, "/v1/admin/movie/edit", app.AddNewMovie)
	router.HandlerFunc(http.MethodDelete, "/v1/admin/movie/delete/:id", app.deleteMovie)

	return app.enableCORS(router)
}
