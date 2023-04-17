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

	// private routes for user

	// routes for comments
	router.HandlerFunc(http.MethodPost, "/v1/movie/comments/add", app.addComment)
	router.HandlerFunc(http.MethodPut, "/v1/movie/comments/update", app.updateComment)
	router.HandlerFunc(http.MethodGet, "/v1/movie/comments/delete", app.deleteComment)

	// routes for favorites
	router.HandlerFunc(http.MethodPost, "/v1/movie/favorites/add", app.addFavorite)
	router.HandlerFunc(http.MethodGet, "/v1/movie/favorites/delete/:id", app.deleteFavorite)

	// private routes for admin
	router.HandlerFunc(http.MethodPost, "/v1/images/upload", app.uploadImage)
	router.HandlerFunc(http.MethodPost, "/v1/admin/genre/add", app.addGenre)
	router.HandlerFunc(http.MethodPut, "/v1/admin/genre/edit", app.editGenre)
	router.HandlerFunc(http.MethodGet, "/v1/admin/genre/delete/:id", app.deleteGenre)
	router.HandlerFunc(http.MethodPost, "/v1/admin/movie/add", app.AddNewMovie)
	router.HandlerFunc(http.MethodPut, "/v1/admin/movie/edit", app.AddNewMovie)
	router.HandlerFunc(http.MethodGet, "/v1/admin/movie/delete/:id", app.deleteMovie)

	return app.enableCORS(router)
}
