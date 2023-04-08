package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/raihan2bd/filmwise/models"
)

// constants for default values
const (
	defaultPage    = 1
	defaultPerPage = 10
	defaultOrderBy = "order by release_date desc"
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

// // get all movies by filter
// func (app *application) getAllMoviesByFilter(w http.ResponseWriter, r *http.Request) {
// 	// get query params from request
// 	queryValues := r.URL.Query()

// 	// find by search query
// 	searchInput := strings.ToLower(queryValues.Get("s"))
// 	findByName := fmt.Sprintf("(title ILIKE '%%%s%%' OR description ILIKE '%%%s%%')", searchInput, searchInput)

// 	var page int
// 	var perPage int
// 	// set up current page
// 	if queryValues.Get("page") == "" {
// 		page = 1
// 	} else {
// 		p, err := strconv.Atoi(queryValues.Get("page"))
// 		if err != nil {
// 			app.errorJSON(w, errors.New("current page should be a number"))
// 			return
// 		}
// 		page = p
// 	}

// 	// set up per page limit
// 	if queryValues.Get("limit") == "" {
// 		perPage = 10
// 	} else {
// 		pp, err := strconv.Atoi(queryValues.Get("limit"))
// 		if err != nil {
// 			app.errorJSON(w, errors.New("per page limit should be a number"))
// 			return
// 		}
// 		perPage = pp
// 	}

// 	filterByGenre := ""
// 	gID, err := strconv.Atoi(queryValues.Get("genre"))
// 	if err == nil {
// 		filterByGenre = fmt.Sprintf("and id in (select movie_id from movies_genres where genre_id = %d)", gID)
// 	}

// 	filterByYear := ""
// 	if queryValues.Get("year") != "" {
// 		year, err := strconv.Atoi(queryValues.Get("year"))
// 		if err == nil {
// 			filterByYear = fmt.Sprintf("and year = %d", year)
// 		}
// 	}

// 	// get sort value
// 	var orderBy string
// 	sort := queryValues.Get("order_by")
// 	switch sort {
// 	case "rating", "runtime":
// 		orderBy = fmt.Sprintf("order by %s desc", sort)
// 	case "old":
// 		orderBy = "order by release_date asc"
// 	case "name":
// 		orderBy = "order by title asc"
// 	default:
// 		orderBy = "order by release_date desc"
// 	}
// 	app.logger.Println(page, perPage, findByName, filterByYear, filterByGenre, orderBy)

// 	movies, err := app.models.DB.GetAllMoviesByFilter(page, perPage, findByName, filterByGenre, filterByYear, orderBy)
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

func (app *application) getAllMoviesByFilter(w http.ResponseWriter, r *http.Request) {
	// get query params from request
	queryValues := r.URL.Query()

	// create a slice of queryParam structs to store the query parameters and their values
	queryParams := make([]queryParam, 0)

	// iterate over the query values and validate them using the helper function
	for key, values := range queryValues {
		if len(values) > 0 {
			value, err := validateQueryParam(key, values[0])
			if err != nil {
				app.errorJSON(w, err)
				return
			}
			queryParams = append(queryParams, queryParam{key: key, value: value})
		}
	}
	if queryValues.Get("s") == "" {
		queryParams = append(queryParams, queryParam{key: "s", value: ""})
	}
	if queryValues.Get("order_by") == "" {
		queryParams = append(queryParams, queryParam{key: "order_by", value: "new"})
	}

	// set up default values for page, perPage and orderBy if not provided in the query
	page := defaultPage
	perPage := defaultPerPage
	orderBy := defaultOrderBy

	// set up filters for name, genre and year based on the query parameters
	var filter models.MovieFilter
	findByName := ""
	filterByGenre := ""
	filterByYear := ""

	// iterate over the queryParams slice and assign values to the variables based on the key
	for _, param := range queryParams {
		switch param.key {
		case "s":
			findByName = fmt.Sprintf("(title ILIKE '%%%s%%' OR description ILIKE '%%%s%%')", param.value, param.value)
		case "page":
			page = param.value.(int)
		case "limit":
			perPage = param.value.(int)
		case "genre":
			filterByGenre = fmt.Sprintf("and id in (select movie_id from movies_genres where genre_id = %d)", param.value)
		case "year":
			filterByYear = fmt.Sprintf("and year = %d", param.value)
		case "order_by":
			switch param.value {
			case "rating", "runtime":
				orderBy = fmt.Sprintf("%s desc", param.value)
			case "old":
				orderBy = "release_date asc"
			case "name":
				orderBy = "title asc"
			default:
				orderBy = "release_date desc"
			}
		}
	}

	filter.FindByName = findByName
	filter.FilterByGenre = filterByGenre
	filter.FilterByYear = filterByYear
	filter.OrderBy = orderBy

	// app.logger.Println(page, perPage, findByName, filterByYear, filterByGenre, orderBy)

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
