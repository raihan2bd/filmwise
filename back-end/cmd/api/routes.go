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

	// initialize admin middleware
	secureAdmin := alice.New(app.adminAuth)
	// public routes
	router.HandlerFunc(http.MethodGet, "/status", app.GetStatus)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMoviesByFilter)
	router.HandlerFunc(http.MethodGet, "/v1/movies/all", app.getAllMovies)
	router.HandlerFunc(http.MethodGet, "/v1/movies/genre/:genre_id", app.getAllMoviesByGenre)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/v1/movie/get_one/:id", app.getOneMovie)

	// router.HandlerFunc(http.MethodPost, "/v1/user/signup/", app.signUp)
	router.HandlerFunc(http.MethodPost, "/v1/user/login/", app.loginUser)

	/* private routes for user */

	// routes for ratings
	router.POST("/v1/rating/add", app.wrap(secure.ThenFunc(app.addOrUpdateRating)))

	// user private routes to manage comments
	router.POST("/v1/movie/comments/add", app.wrap(secure.ThenFunc(app.addOrUpdateComment)))
	router.PUT("/v1/movie/comments/update", app.wrap(secure.ThenFunc(app.addOrUpdateComment)))
	router.GET("/v1/movie/comments/delete/:id", app.wrap(secure.ThenFunc(app.deleteComment)))

	// routes for favorites
	router.HandlerFunc(http.MethodGet, "/v1/movie/favorites/add/:id", app.addFavorite)
	router.HandlerFunc(http.MethodGet, "/v1/movie/favorites/delete/:id", app.removeFavorite)

	// private routes for admin
	router.POST("/v1/images/upload", app.wrap(secureAdmin.ThenFunc(app.uploadImage)))

	// admin routes for manage movie genre
	router.POST("/v1/admin/genre/add", app.wrap(secureAdmin.ThenFunc(app.addOrUpdateGenre)))
	router.PUT("/v1/admin/genre/edit", app.wrap(secureAdmin.ThenFunc(app.addOrUpdateGenre)))
	router.GET("/v1/admin/genre/delete/:id", app.wrap(secureAdmin.ThenFunc(app.deleteGenre)))

	// admin routes to manage movie
	router.POST("/v1/admin/movie/add", app.wrap(secureAdmin.ThenFunc(app.AddNewMovie)))
	router.PUT("/v1/admin/movie/edit", app.wrap(secureAdmin.ThenFunc(app.AddNewMovie)))
	router.GET("/v1/admin/movie/delete/:id", app.wrap(secureAdmin.ThenFunc(app.deleteMovie)))

	return app.enableCORS(router)
}
