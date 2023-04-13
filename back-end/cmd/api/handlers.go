package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/raihan2bd/filmwise/models"
	"github.com/raihan2bd/filmwise/validator"
)

// constants for default values
const (
	defaultPage    = 1
	defaultPerPage = 3
)

func (app *application) GetStatus(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "available",
		Environment: app.config.env,
		Version:     version,
	}
	err := app.writeJSON(w, http.StatusOK, currentStatus, "app_status")
	if err != nil {
		app.logger.Println(err)
	}
}

// get all movies by filter
func (app *application) getAllMoviesByFilter(w http.ResponseWriter, r *http.Request) {
	// get query params from request
	queryValues := r.URL.Query()

	// find by search query
	searchInput := strings.ToLower(queryValues.Get("s"))
	var filter models.MovieFilter
	filter.FindByName = searchInput

	page := defaultPage
	perPage := defaultPerPage

	// set up current page
	if queryValues.Get("page") != "" {
		p, err := strconv.Atoi(queryValues.Get("page"))
		if err != nil {
			app.errorJSON(w, errors.New("current page should be a number"))
			return
		}
		page = p
	}

	// set up per page limit
	if queryValues.Get("limit") != "" {
		pp, err := strconv.Atoi(queryValues.Get("limit"))
		if err != nil {
			app.errorJSON(w, errors.New("per page limit should be a number"))
			return
		}
		perPage = pp
	}

	gID, err := strconv.Atoi(queryValues.Get("genre"))
	if err == nil {
		filter.FilterByGenre = gID
	}

	if queryValues.Get("year") != "" {
		year, err := strconv.Atoi(queryValues.Get("year"))
		if err == nil {
			filter.FilterByYear = year
		}
	}

	filter.OrderBy = queryValues.Get("order_by")

	movies, err := app.models.DB.GetAllMoviesByFilter(page, perPage, &filter)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// Get all movies by genre
func (app *application) getAllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	genreID, err := strconv.Atoi(params.ByName("genre_id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	movies, err := app.models.DB.GetAllMovies(genreID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// Get all genres
func (app *application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.models.DB.GenresAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// AddNewMovie will insert a new movie
func (app *application) AddNewMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	// read json from the body
	err := app.readJSON(w, r, &movie)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	validation := validator.New()
	validation.IsLength(movie.Title, "title", "title should be minimum 3 to 255 characters long!", 3, 255)

	if validation.Valid() {
		app.logger.Println("form is validated")
	} else {
		// read json from the body
		err := app.writeJSON(w, http.StatusBadRequest, validation)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}
		app.logger.Println("for isn't validated")
		return
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
	}
}
