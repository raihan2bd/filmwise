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

	if len(payload.ID) > 0 {
		editMovieID, err := strconv.Atoi(payload.ID)
		if err != nil {
			app.logger.Println(err)
			app.errorJSON(w, errors.New("invalid movie id"))
			return
		}

		m, err := app.models.DB.Get(editMovieID)
		if err != nil {
			app.logger.Println(err)
			app.errorJSON(w, errors.New("invalid movie id"))
			return
		}

		movie.ID = m.ID
	}

	var movieID int
	var moviesGenres map[int]string
	respMsg := "Movie is inserted successfully!"

	if movie.ID > 0 {
		respMsg = "Movie is successfully updated"
		movieID, moviesGenres, err = app.models.DB.UpdateMovie(&movie)
	} else {
		movieID, moviesGenres, err = app.models.DB.InsertMovie(&movie)
	}

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
	resp.Message = respMsg

	err = app.writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.errorJSON(w, err)
	}
}

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, errors.New("invalid id parameter"))
		return
	}

	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.errorJSON(w, errors.New("failed to fetch the movie"))
		return
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, errors.New("invalid id"))
		return
	}

	err = app.models.DB.DeleteMovie(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var resp struct {
		Error   bool   `json:"error"`
		ID      int    `json:"id"`
		Message string `json:"message"`
	}

	resp.Error = false
	resp.ID = id
	resp.Message = "movie is successfully deleted!"

	err = app.writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// Add or Update genre
func (app *application) addOrUpdateGenre(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID        string `json:"genre_id"`
		GenreName string `json:"genre_name"`
	}

	// read json from the body
	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequest(w, r, errors.New("invalid json"))
		return
	}

	id := 0

	if payload.ID != "" {
		id, err = strconv.Atoi(payload.ID)
		if err != nil {
			app.badRequest(w, r, errors.New("invalid genre id"))
			return
		}
		_, err = app.models.DB.CheckGenre(id)
		if err != nil {
			app.badRequest(w, r, errors.New("invalid genre id"))
			return
		}
	}

	// validate genre name
	validator := validator.New()
	validator.IsLength(payload.GenreName, "genre_name", 3, 50)

	// if len(payload.GenreName) < 3 || len(payload.GenreName) > 50 {
	// 	app.badRequest(w, r, errors.New("genre name must be between 3 and 50 characters"))
	// 	return
	// }

	if !validator.Valid() {
		err := app.writeJSON(w, http.StatusBadRequest, validator)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}
		return
	}

	var genreID int
	respMsg := "Genre is inserted successfully!"

	if id > 0 {
		genreID, err = app.models.DB.UpdateGenre(id, payload.GenreName)
		respMsg = "Genre is successfully updated"
	} else {
		genreID, err = app.models.DB.InsertGenre(payload.GenreName)
	}

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var resp struct {
		OK      bool   `json:"ok"`
		Message string `json:"message"`
		ID      int    `json:"id"`
	}

	resp.OK = true
	resp.Message = respMsg
	resp.ID = genreID

	err = app.writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.errorJSON(w, err)
	}

}

// deleteGenre deletes a genre
func (app *application) deleteGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, errors.New("invalid id"))
		return
	}

	err = app.models.DB.DeleteGenre(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var resp struct {
		OK      bool   `json:"ok"`
		ID      int    `json:"id"`
		Message string `json:"message"`
	}

	resp.OK = true
	resp.ID = id
	resp.Message = "genre is successfully deleted!"

	err = app.writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

// Add or Update Rating
func (app *application) addOrUpdateRating(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID      string `json:"id"`
		MovieID string `json:"movie_id"`
		Rating  string `json:"rating"`
	}

	// read json from the body
	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, errors.New("invalid json request"))
		return
	}

	ratingID := 0

	if payload.ID != "" {
		ratingID, err = strconv.Atoi(payload.ID)

		if err != nil {
			app.errorJSON(w, errors.New("invalid rating id"))
			return
		}

		err = app.models.DB.CheckRating(ratingID)
		if err != nil {
			app.errorJSON(w, errors.New("invalid rating id"))
			return
		}
	}

	validator := validator.New()
	validator.Required(payload.MovieID, "movie_id", "movie id is required")
	validator.Required(payload.Rating, "rating", "rating is required")
	movieID, err := strconv.Atoi(payload.MovieID)
	if err != nil {
		validator.AddError("movie_id", "invalid movie_id")
	}

	if movieID > 0 {
		_, err = app.models.DB.Get(movieID)
		if err != nil {
			validator.AddError("movie_id", "invalid movie id")
		}
	}

	movieRating, err := strconv.ParseFloat(payload.Rating, 32)
	if err != nil {
		validator.AddError("rating", "invalid rating")
	}

	if movieRating < 1.0 || movieRating > 10.0 {
		validator.AddError("rating", "movie rating shoulbe between 1.0 to 10.0")
	}

	if !validator.Valid() {
		err := app.writeJSON(w, http.StatusBadRequest, validator)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}
		return
	}

	// user id
	userID := 1

	rating := models.Rating{
		ID:        ratingID,
		MovieID:   movieID,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var id int
	respMsg := "Rating is added successfully!"

	if ratingID > 0 {
		id, err = app.models.DB.UpdateRating(&rating)
		respMsg = "Rating is updated successfully!"
	} else {
		id, err = app.models.DB.InsertRating(&rating)
	}

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var resp struct {
		OK      bool   `json:"ok"`
		ID      int    `json:"id"`
		Message string `json:"message"`
	}

	resp.OK = true
	resp.ID = id
	resp.Message = respMsg

	err = app.writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.errorJSON(w, err)
	}

}

// comment payload
type commentPayload struct {
	Comment   string `json:"comment"`
	MovieID   string `json:"movie_id"`
	CommentID string `json:"comment_id"`
}

// Add or Update a comment
func (app *application) addOrUpdateComment(w http.ResponseWriter, r *http.Request) {
	var payload commentPayload

	// read json from the body
	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequest(w, r, errors.New("invalid json"))
		return
	}

	id := 0

	if payload.CommentID != "" {
		id, err = strconv.Atoi(payload.CommentID)
		if err != nil {
			app.badRequest(w, r, errors.New("invalid comment id"))
			return
		}
		_, err = app.models.DB.CheckComment(id)
		if err != nil {
			app.badRequest(w, r, errors.New("invalid comment id"))
			return
		}
	}

	validator := validator.New()
	validator.IsLength(payload.Comment, "comment", 10, 500)
	validator.Required(payload.MovieID, "movie_id", "movie_id is required")

	movieID, err := strconv.Atoi(payload.MovieID)
	if err != nil {
		validator.AddError("movie_id", "invalid movie_id!")
	}

	if !validator.Valid() {
		err := app.writeJSON(w, http.StatusBadRequest, validator)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}
		return
	}

	var comment models.Comment
	comment.Comment = strings.Trim(payload.Comment, "")
	comment.MovieID = movieID
	comment.UserID = 1
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	var resp struct {
		OK      bool   `json:"ok"`
		ID      int    `json:"id"`
		Message string `json:"message"`
	}

	var commentID int
	respMsg := "Comment is inserted successfully!"
	if comment.ID > 0 {
		commentID, err = app.models.DB.UpdateComment(&comment)
		respMsg = "Comment is updated successfully!"
	} else {
		commentID, err = app.models.DB.InsertComment(&comment)
	}

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp.OK = false
	resp.ID = commentID
	resp.Message = respMsg

	err = app.writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.errorJSON(w, err)
	}
}

// Delete a comment
func (app *application) deleteComment(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, errors.New("invalid id"))
		return
	}

	// check if the comment exists
	_, err = app.models.DB.GetComment(id)
	if err != nil {
		app.errorJSON(w, errors.New("invalid comment id"))
		return
	}

	// user authentication will be added later
	err = app.models.DB.DeleteComment(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var resp struct {
		OK      bool   `json:"ok"`
		ID      int    `json:"id"`
		Message string `json:"message"`
	}

	resp.OK = false
	resp.ID = id
	resp.Message = "comment is successfully deleted!"

	err = app.writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// Add Favorite
func (app *application) addFavorite(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, errors.New("invalid id"))
		return
	}

	// check if the movie exists
	_, err = app.models.DB.Get(id)
	if err != nil {
		app.errorJSON(w, errors.New("invalid movie id"))
		return
	}

	// user authentication will be added later
	favorite := models.Favorite{
		UserID:    1,
		MovieID:   id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = app.models.DB.AddFavorite(&favorite)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var resp struct {
		OK      bool   `json:"ok"`
		ID      int    `json:"id"`
		Message string `json:"message"`
	}

	resp.OK = false
	resp.ID = id
	resp.Message = "movie is successfully added to favorites!"

	err = app.writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// Remove Favorite
func (app *application) removeFavorite(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, errors.New("invalid id"))
		return
	}

	// check if the movie exists
	_, err = app.models.DB.Get(id)
	if err != nil {
		app.errorJSON(w, errors.New("invalid movie id"))
		return
	}

	// user authentication will be added later
	err = app.models.DB.RemoveFavorite(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var resp struct {
		OK      bool   `json:"ok"`
		ID      int    `json:"id"`
		Message string `json:"message"`
	}

	resp.OK = false
	resp.ID = id
	resp.Message = "movie is successfully removed from favorites!"

	err = app.writeJSON(w, http.StatusOK, resp)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
