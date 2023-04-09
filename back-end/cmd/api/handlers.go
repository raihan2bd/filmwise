package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/raihan2bd/filmwise/models"
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

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

	movies, err := app.models.DB.GetAllMovies()
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

// -----------------------------

// func (app *application) getAllMoviesByFilter(w http.ResponseWriter, r *http.Request) {
// 	// get query params from request
// 	queryValues := r.URL.Query()

// 	// create a slice of queryParam structs to store the query parameters and their values
// 	queryParams := make([]models.QueryParam, 0)

// 	// iterate over the query values and validate them using the helper function
// 	for key, values := range queryValues {
// 		if len(values) > 0 {
// 			value, err := validateQueryParam(key, values[0])
// 			if err != nil {
// 				app.errorJSON(w, err)
// 				return
// 			}
// 			queryParams = append(queryParams, models.QueryParam{Key: key, Value: value})
// 		}
// 	}
// 	// if queryValues.Get("s") == "" {
// 	// 	queryParams = append(queryParams, QueryParam{Key: "s", Value: ""})
// 	// }
// 	// if queryValues.Get("order_by") == "" {
// 	// 	queryParams = append(queryParams, QueryParam{Key: "order_by", Value: "new"})
// 	// }

// 	fmt.Println("hi")

// 	// set up default values for page, perPage and orderBy if not provided in the query
// 	page := defaultPage
// 	perPage := defaultPerPage
// 	orderBy := defaultOrderBy

// 	// set up filters for name, genre and year based on the query parameters
// 	// var filter models.MovieFilter
// 	findByName := ""
// 	var filterByGenre int
// 	var filterByYear int

// 	// iterate over the queryParams slice and assign values to the variables based on the key
// 	for _, param := range queryParams {
// 		switch param.Key {
// 		case "s":
// 			// findByName = fmt.Sprintf("(title ILIKE '%%%s%%' OR description ILIKE '%%%s%%')", param.Value, param.Value)
// 			findByName = param.Value.(string)
// 		case "page":
// 			page = param.Value.(int)
// 		case "limit":
// 			perPage = param.Value.(int)
// 		case "genre":
// 			filterByGenre = param.Value.(int)
// 		case "year":
// 			filterByYear = param.Value.(int)
// 		case "order_by":
// 			switch param.Value {
// 			case "rating", "runtime":
// 				orderBy = fmt.Sprintf("%s desc", param.Value)
// 			case "old":
// 				orderBy = "release_date asc"
// 			case "name":
// 				orderBy = "title asc"
// 			default:
// 				orderBy = "release_date desc"
// 			}
// 		}
// 	}

// 	// filter.FindByName = findByName
// 	// filter.FilterByGenre = filterByGenre
// 	// filter.FilterByYear = filterByYear
// 	// filter.OrderBy = orderBy

// 	// app.logger.Println(page, perPage, findByName, filterByYear, filterByGenre, orderBy)

// 	movies, err := app.models.DB.GetAllMoviesByFilter(page, perPage, filterByGenre, filterByYear, findByName, orderBy)
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	err = app.writeJSON(w, http.StatusOK, movies, "movies")
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}
// }
