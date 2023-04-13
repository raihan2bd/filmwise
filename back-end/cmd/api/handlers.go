package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/raihan2bd/filmwise/models"
	"github.com/raihan2bd/filmwise/validator"
)

// constants for default values
const (
	defaultPage    = 1
	defaultPerPage = 3
)

type MoviePayload struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Year        string         `json:"year"`
	ReleaseDate string         `json:"release_date"`
	Runtime     string         `json:"runtime"`
	Rating      string         `json:"rating"`
	MovieGenre  map[int]string `json:"genres"`
}

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
	var payload MoviePayload

	// read json from the body
	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	validator := validator.New()
	validator.IsLength(payload.Title, "title", 3, 255)
	validator.IsLength(payload.Description, "description", 20, 500)
	validator.Required(payload.Year, "year", "year is required")
	validator.Required(payload.ReleaseDate, "release_date", "release_date is required")

	year, err := strconv.Atoi(payload.Year)
	if err != nil {
		validator.AddError("year", "invalid year!")
	}

	releaseDate, err := time.Parse("2006-01-02", payload.ReleaseDate)
	if err != nil {
		validator.AddError("release_date", "invalid release_date!")
	}

	runtime, err := strconv.Atoi(payload.Runtime)
	if err != nil {
		validator.AddError("runtime", "invalid runtime!")
	}

	rating, err := strconv.ParseFloat(payload.Rating, 64)
	if err != nil {
		validator.AddError("rating", "invalid rating!")
	}

	if len(payload.MovieGenre) <= 0 {
		validator.AddError("genres", "movie genre is required")
	}

	if len(payload.MovieGenre) > 5 {
		validator.AddError("genres", "maximum 5 genres are allowed")
	}

	if !validator.Valid() {
		err := app.writeJSON(w, http.StatusBadRequest, validator)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}
		return
	}

	var movie models.Movie
	movie.Title = strings.Trim(payload.Title, "")
	movie.Description = strings.Trim(payload.Description, "")
	movie.Year = year
	movie.ReleaseDate = releaseDate
	movie.Runtime = runtime
	movie.Rating = rating
	movie.MovieGenre = payload.MovieGenre

	movieID, moviesGenres, err := app.models.DB.InsertMovie(&movie)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var resp struct {
		Error        bool           `json:"error"`
		ID           int            `json:"id"`
		MoviesGenres map[int]string `json:"movies_genres"`
		Message      string         `json:"message"`
	}

	resp.Error = false
	resp.ID = movieID
	resp.MoviesGenres = moviesGenres
	resp.Message = "Movie is inserted successfully!"

	err = app.writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.errorJSON(w, err)
	}
}
