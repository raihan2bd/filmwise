package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// wrap function will help to chain between multiple middlewares
func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// pass httprouter.Params to request context
		ctx := context.WithValue(r.Context(), "params", ps)
		// call next middleware with new context
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// routes funtion helps to create back-end routes
func (app *application) routes() http.Handler {
	// initialize the router
	router := httprouter.New()

	// initialize secure middleware
	secure := alice.New(app.authenticate)
	// public routes
	router.HandlerFunc(http.MethodGet, "/status", app.GetStatus)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMoviesByFilter)
	router.HandlerFunc(http.MethodGet, "/v1/movies/all", app.getAllMovies)
	router.HandlerFunc(http.MethodGet, "/v1/movies/genre/:genre_id", app.getAllMoviesByGenre)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/v1/movie/get_one/:id", app.getOneMovie)

	// router.HandlerFunc(http.MethodPost, "/v1/user/signup/", app.signUp)
	router.HandlerFunc(http.MethodPost, "/v1/user/login/", app.loginUser)
	// router.HandlerFunc(http.MethodGet, "/v1/user/logout/", app.logout)

	/* private routes for user */

	// routes for ratings
	router.POST("/v1/rating/add", app.wrap(secure.ThenFunc(app.addOrUpdateRating)))

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
