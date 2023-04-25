package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.GetStatus)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMoviesByFilter)
	router.HandlerFunc(http.MethodGet, "/v1/movies/all", app.getAllMovies)
	router.HandlerFunc(http.MethodGet, "/v1/movies/genre/:genre_id", app.getAllMoviesByGenre)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/v1/movie/get_one/:id", app.getOneMovie)

	// private routes for user
	router.HandlerFunc(http.MethodPost, "/v1/movie/rating/add", app.addOrUpdateRating)
	router.HandlerFunc(http.MethodPost, "/v1/movie/rating/update", app.addOrUpdateRating)

	// routes for comments
	router.HandlerFunc(http.MethodPost, "/v1/movie/comments/add", app.addOrUpdateComment)
	router.HandlerFunc(http.MethodPut, "/v1/movie/comments/update", app.addOrUpdateComment)
	router.HandlerFunc(http.MethodGet, "/v1/movie/comments/delete/:id", app.deleteComment)

	// routes for favorites
	router.HandlerFunc(http.MethodGet, "/v1/movie/favorites/add/:id", app.addFavorite)
	router.HandlerFunc(http.MethodGet, "/v1/movie/favorites/delete/:id", app.removeFavorite)

	// private routes for admin
	router.HandlerFunc(http.MethodPost, "/v1/images/upload", app.uploadImage)
	router.HandlerFunc(http.MethodPost, "/v1/admin/genre/add", app.addOrUpdateGenre)
	router.HandlerFunc(http.MethodPut, "/v1/admin/genre/edit", app.addOrUpdateGenre)
	router.HandlerFunc(http.MethodGet, "/v1/admin/genre/delete/:id", app.deleteGenre)
	router.HandlerFunc(http.MethodPost, "/v1/admin/movie/add", app.AddNewMovie)
	router.HandlerFunc(http.MethodPut, "/v1/admin/movie/edit", app.AddNewMovie)
	router.HandlerFunc(http.MethodGet, "/v1/admin/movie/delete/:id", app.deleteMovie)

	return app.enableCORS(router)
}
